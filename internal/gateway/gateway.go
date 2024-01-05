package gateway

import (
	"net"
	"sync"
	"time"

	"github.com/jsn4ke/chat/pkg/inter"
	ltvcodec "github.com/jsn4ke/chat/pkg/ltv-codec"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
)

const (
	channelTypeNone = iota
	channelTypeGuild
	channelTypeMax
)

func NewGateway(gatewayid uint32, addr string, done <-chan struct{},
	authCluster inter.RpcAuthCliCluster,
	logicCluster inter.RpcLogicCliCluster,
) error {
	listener, err := net.Listen(`tcp`, addr)
	if nil != err {
		return err
	}

	g := new(gateway)
	g.gatewayId = gatewayid
	g.done = done
	g.authCluster = authCluster
	g.logicCluster = logicCluster
	g.syncDiff = true

	g.cliAcceptor = jsn_net.NewTcpAcceptor(listener.(*net.TCPListener), 1, 30000, func() jsn_net.Pipe {
		c := new(clientHandler)
		c.in = make(chan any, 1)
		c.done = make(chan struct{})
		c.gateway = g
		return c
	}, new(ltvcodec.LTVCodec), 0, 64)
	g.cliAcceptor.Start()

	jsn_net.WaitGo(&g.wg, g.upload)
	return nil
}

type gateway struct {
	gatewayId uint32

	authCluster  inter.RpcAuthCliCluster
	logicCluster inter.RpcLogicCliCluster

	cliAcceptor *jsn_net.TcpAcceptor

	////////////////
	mtx         sync.Mutex
	users       map[uint64]uint64
	channels    [channelTypeMax]map[uint64]map[uint64]uint64
	channelSync [channelTypeMax]map[uint64]bool
	syncDiff    bool
	///////////////

	done <-chan struct{}

	wg sync.WaitGroup
}

func (g *gateway) AuthRpcTimeout() time.Duration {
	return time.Second
}

func (g *gateway) SigninTimeout() time.Duration {
	return time.Second
}

func (g *gateway) SubscribTimeout() time.Duration {
	return time.Millisecond * 50
}

func (g *gateway) SubscribHeatbeat() time.Duration {
	return time.Millisecond * 100
}

func (g *gateway) bindUser(uid uint64, ssid uint64, guildId uint64) {
	g.mtx.Lock()
	g.mtx.Unlock()
	oldSess := g.users[uid]
	if nil == g.users {
		g.users = map[uint64]uint64{}
	}
	g.users[uid] = ssid
	if 0 != oldSess {
		sess := g.cliAcceptor.Session(oldSess)
		if nil != sess {
			sess.Close()
		}
	}
	if 0 != guildId {
		ct := channelTypeGuild
		keyMap := g.channels[ct]
		mp, ok := keyMap[guildId]
		if !ok {
			if nil == keyMap {
				keyMap = map[uint64]map[uint64]uint64{}
				g.channels[ct] = keyMap
			}
			mp = map[uint64]uint64{}
			keyMap[guildId] = mp
		}
		_, replace := mp[uid]
		mp[uid] = ssid
		// 0-1 new sub or 1-1
		if !replace && 1 != len(mp) {
			return
		}
		sm := g.channelSync[ct]
		if _, ok := sm[guildId]; ok {
			delete(sm, guildId)
		} else {
			if nil == sm {
				sm = map[uint64]bool{}
				g.channelSync[ct] = sm
			}
			sm[guildId] = true
		}
	}
}

func (g *gateway) unbindUser(uid uint64, ssid uint64, guild uint64) {
	g.mtx.Lock()
	g.mtx.Unlock()
	curssid := g.users[uid]
	if curssid == ssid {
		delete(g.users, uid)
	}
	if 0 != guild {
		ch := g.channels[channelTypeGuild]
		guildSub := ch[guild]
		curssid = guildSub[uid]
		if curssid == ssid {
			delete(guildSub, uid)
			if 0 == len(guildSub) {
				cs := g.channelSync[channelTypeGuild]
				if _, ok := cs[guild]; ok {
					delete(cs, guild)
				} else {
					if nil == cs {
						cs = map[uint64]bool{}
						g.channelSync[channelTypeGuild] = cs
					}
					cs[guild] = false
				}
			}
		}
	}
}

func (g *gateway) upload() {
	for {
		if g.syncDiff {
			g.uploadSubOrUnsub()
		} else {
			g.uploadReSubscrib()
		}
	}
}

func (g *gateway) uploadSubOrUnsub() {
	tk := time.NewTicker(g.SubscribHeatbeat())
	defer tk.Stop()
	for g.syncDiff {
		select {
		case <-tk.C:
		}
		g.mtx.Lock()
		sync := g.channelSync
		g.channelSync = [channelTypeMax]map[uint64]bool{}
		g.mtx.Unlock()
		if 0 == len(sync) {
			return
		}
		if 0 != len(sync[channelTypeGuild]) {
			err := g.logicCluster.Get().RpcLogicSubscribeOrUnsubscribAsk(&message_rpc.RpcLogicSubscribeOrUnsubscribAsk{
				GatewayId: g.gatewayId,
				Guild:     sync[channelTypeGuild],
			}, g.done, g.SubscribTimeout())
			if nil != err {
				g.syncDiff = false
			}
		}
	}
}

func (g *gateway) uploadReSubscrib() {
	tk := time.NewTicker(g.SubscribHeatbeat())
	defer tk.Stop()
	for !g.syncDiff {
		select {
		case <-tk.C:
		}
		var sync [channelTypeMax]map[uint64]bool
		g.mtx.Lock()
		g.channelSync = [channelTypeMax]map[uint64]bool{}
		for ct, v := range g.channels {
			sync[ct] = map[uint64]bool{}
			for key := range v {
				sync[ct][key] = true
			}
		}
		g.mtx.Unlock()
		err := g.logicCluster.Get().RpcLogicReSubscribeAsk(&message_rpc.RpcLogicReSubscribeAsk{
			GatewayId: g.gatewayId,
			Guild:     sync[channelTypeGuild],
		}, g.done, g.SubscribTimeout())
		if nil == err {
			g.syncDiff = true
		}
	}
}
