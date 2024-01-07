package gateway

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/jsn4ke/chat/internal/inter"
	ltvcodec "github.com/jsn4ke/chat/pkg/ltv-codec"
	"github.com/jsn4ke/chat/pkg/pb/message_obj"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

func NewGateway(gatewayid uint32, addr, rpcAddr string, done <-chan struct{},
	authCluster inter.RpcAuthCliCluster,
	logicCluster inter.RpcLogicCliCluster,
	gatewayRpcWrap inter.GateWayWrapFunc,
) error {
	listener, err := net.Listen(`tcp`, addr)
	if nil != err {
		return err
	}

	g := new(gateway)
	g.gatewayId = gatewayid
	g.rpcAddr = rpcAddr
	g.done = done
	g.authCluster = authCluster
	g.logicCluster = logicCluster
	// 首次同步全部
	g.syncDiff = false

	g.cliAcceptor = jsn_net.NewTcpAcceptor(listener.(*net.TCPListener), 1, 30000, func() jsn_net.Pipe {
		c := new(clientHandler)
		c.in = make(chan any, 1)
		c.done = make(chan struct{})
		c.gateway = g
		return c
	}, new(ltvcodec.LTVCodec), 0, 64)
	g.cliAcceptor.Start()

	gatewayRpcSvr := jsn_rpc.NewServer(rpcAddr, 128, 4)

	jsn_net.WaitGo(&g.wg, gatewayRpcWrap(gatewayRpcSvr, 10, done, g).Run)

	// logicCluster.Range(func(rlc rpcinter.RpcLogicCli) {
	// 	jsn_net.WaitGo(&g.wg, func() {
	// 		cli := api.NewRpcUnionGatewayCli(rlc.Cli())
	// 		req := &message_rpc.RpcUnionGatewayUploadGatewayInfoAsk{
	// 			GatewaId: gatewayid,
	// 			Address:  rpcAddr,
	// 		}
	// 		err := cli.RpcUnionGatewayUploadGatewayInfoAsk(req, done, time.Second)
	// 		if nil == err {
	// 			return
	// 		}
	// 		tk := time.NewTicker(time.Second)
	// 		defer tk.Stop()
	// 		for {
	// 			select {
	// 			case <-done:
	// 				return
	// 			case <-tk.C:
	// 				err := cli.RpcUnionGatewayUploadGatewayInfoAsk(req, done, time.Second)
	// 				if nil == err {
	// 					return
	// 				}
	// 			}
	// 		}
	// 	})
	// })

	jsn_net.WaitGo(&g.wg, g.upload)
	return nil
}

type gateway struct {
	gatewayId uint32

	rpcAddr string

	authCluster  inter.RpcAuthCliCluster
	logicCluster inter.RpcLogicCliCluster

	cliAcceptor *jsn_net.TcpAcceptor

	////////////////
	mtx   sync.Mutex
	users map[uint64]uint64
	// key - uid - ssid
	channels    [message_obj.ChannelMessageType_ChannelMessageType_Max]map[uint64]map[uint64]uint64
	channelSync [message_obj.ChannelMessageType_ChannelMessageType_Max]map[uint64]bool
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

func (g *gateway) ChatTimeout() time.Duration {
	return time.Millisecond * 50
}

func (g *gateway) bindUser(uid uint64, ssid uint64, guildId uint64, worlId uint64) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
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
	g.bindChanel(message_obj.ChannelMessageType_ChannelMessageType_Guild, guildId, uid, ssid)
	g.bindChanel(message_obj.ChannelMessageType_ChannelMessageType_World, worlId, uid, ssid)
}

func (g *gateway) unbindUser(uid uint64, ssid uint64, guild uint64, world uint64) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	curssid := g.users[uid]
	if curssid == ssid {
		delete(g.users, uid)
	}
	g.unbindChannel(message_obj.ChannelMessageType_ChannelMessageType_Guild, guild, uid, ssid)
	g.unbindChannel(message_obj.ChannelMessageType_ChannelMessageType_World, world, uid, ssid)
}

func (g *gateway) bindChanel(ct message_obj.ChannelMessageType, key uint64, uid uint64, ssid uint64) {
	if 0 == key {
		return
	}
	sub := g.channels[ct]
	mp, ok := sub[key]
	if !ok {
		if nil == sub {
			sub = map[uint64]map[uint64]uint64{}
			g.channels[ct] = sub
		}
		mp = map[uint64]uint64{}
		sub[key] = mp
	}
	_, replace := mp[uid]
	mp[uid] = ssid
	// 0-1 new sub or 1-1
	if !replace && 1 != len(mp) {
		return
	}
	sm := g.channelSync[ct]
	if _, ok := sm[key]; ok {
		delete(sm, key)
	} else {
		if nil == sm {
			sm = map[uint64]bool{}
			g.channelSync[ct] = sm
		}
		sm[key] = true
	}
}

func (g *gateway) unbindChannel(ct message_obj.ChannelMessageType, key uint64, uid uint64, ssid uint64) {
	if 0 == key {
		return
	}
	ch := g.channels[ct]
	sub := ch[key]
	curssid := sub[uid]
	if curssid == ssid {
		delete(sub, uid)
		if 0 == len(sub) {
			cs := g.channelSync[ct]
			if _, ok := cs[key]; ok {
				delete(cs, key)
			} else {
				if nil == cs {
					cs = map[uint64]bool{}
					g.channelSync[ct] = cs
				}
				cs[key] = false
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
		g.channelSync = [message_obj.ChannelMessageType_ChannelMessageType_Max]map[uint64]bool{}
		g.mtx.Unlock()

		gen := func(ct message_obj.ChannelMessageType) (add []uint64, rem []uint64) {
			for k, v := range sync[ct] {
				if v {
					add = append(add, k)
				} else {
					rem = append(rem, k)
				}
			}
			return
		}
		msg := &message_rpc.RpcLogicSubscribeOrUnsubscribAsk{
			GatewayId: g.gatewayId,
		}
		msg.GuildsAdd, msg.GuildsRem = gen(message_obj.ChannelMessageType_ChannelMessageType_Guild)
		msg.WorldAdd, msg.WorldRem = gen(message_obj.ChannelMessageType_ChannelMessageType_World)
		if 0 != len(msg.WorldAdd) || 0 != len(msg.WorldRem) ||
			0 != len(msg.GuildsAdd) || 0 != len(msg.GuildsRem) {
			err := g.logicCluster.Get().RpcLogicSubscribeOrUnsubscribAsk(msg, g.done, g.SubscribTimeout())
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
		msg := &message_rpc.RpcLogicReSubscribeAsk{
			GatewayId: g.gatewayId,
		}
		g.mtx.Lock()
		gen := func(ct message_obj.ChannelMessageType) (ret []uint64) {
			sub := g.channels[ct]
			for key := range sub {
				ret = append(ret, key)
			}
			return
		}
		msg.Guilds = gen(message_obj.ChannelMessageType_ChannelMessageType_Guild)
		msg.Worlds = gen(message_obj.ChannelMessageType_ChannelMessageType_World)
		g.channelSync = [message_obj.ChannelMessageType_ChannelMessageType_Max]map[uint64]bool{}
		g.mtx.Unlock()
		err := g.logicCluster.Get().RpcLogicReSubscribeAsk(msg, g.done, g.SubscribTimeout())
		if nil == err {
			g.syncDiff = true
		} else {
			fmt.Println("uploadReSubscrib", err)
		}
	}
}
