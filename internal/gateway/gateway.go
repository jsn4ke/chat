package gateway

import (
	"sync"
	"time"

	"github.com/jsn4ke/chat/pkg/inter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
)

const (
	channelTypeNone = iota
	channelTypeGuild
	channelTypeMax
)

func NewGateway(gatewayid uint32, done <-chan struct{}) {
	g := new(gateway)
	g.gatewayId = gatewayid

	g.done = done

	jsn_net.WaitGo(&g.wg, g.upload)
}

type gateway struct {
	gatewayId uint32

	authCluster  inter.RpcAuthCliCluster
	logicCluster inter.RpcLogicCliCluster

	cliAcceptor *jsn_net.TcpAcceptor

	mtx         sync.Mutex
	channels    [channelTypeMax]map[uint64]map[uint64]struct{}
	channelSync [channelTypeMax]map[uint64]bool
	syncDiff    bool

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

func (g *gateway) Subscrib(ct int, key uint64, uid uint64) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	keyMap := g.channels[ct]
	if nil == keyMap {
		keyMap = map[uint64]map[uint64]struct{}{}
		g.channels[ct] = keyMap
	}
	mp, ok := keyMap[key]
	if !ok {
		mp = map[uint64]struct{}{}
		keyMap[key] = mp
	}
	mp[uid] = struct{}{}
	// 0-1 new sub
	if 1 != len(mp) {
		return
	}
	sm := g.channelSync[ct]
	if nil == sm {
		sm = map[uint64]bool{}
		g.channelSync[ct] = sm
	}
	if _, ok := sm[key]; ok {
		delete(sm, key)
	} else {
		sm[key] = true
	}
}

func (g *gateway) Unsubscrib(ct int, key uint64, uid uint64) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	delete(g.channels[ct][key], uid)

	if 0 != len(g.channels[ct][key]) {
		return
	}
	sm := g.channelSync[ct]
	if nil == sm {
		sm = map[uint64]bool{}
		g.channelSync[ct] = sm
	}
	if _, ok := sm[key]; ok {
		delete(sm, key)
	} else {
		sm[key] = false
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
