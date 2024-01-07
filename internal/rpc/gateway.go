package rpc

import (
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	"github.com/jsn4ke/jsn_net"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

func NewRpcGatewaySvrWrap(svr *jsn_rpc.Server, runNum int, done <-chan struct{}, core rpcinter.RpcGatewaySvrWrap) rpcinter.RpcGatewaySvr {
	s := new(rpcGatewaySvr)
	c := new(rpcCore)
	c.svr = svr
	c.in = make(chan *jsn_rpc.AsyncRpc, 128)
	c.done = done
	c.runNum = runNum

	s.rpcGatewayCore.rpcCore = c
	s.rpcGatewayCore.wrap = core

	return s
}

type rpcGatewaySvr struct {
	rpcGatewayCore
	rpcinter.RpcGatewaySvrWrap
}

func (s *rpcGatewaySvr) Run() {
	s.registerRpc()

	s.svr.Start()
	for i := 0; i < s.runNum; i++ {
		jsn_net.WaitGo(&s.wg, s.run)
	}
	s.wg.Wait()
}
