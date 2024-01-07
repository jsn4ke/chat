package rpc

import (
	"sync"

	"github.com/jsn4ke/chat/internal/api"
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

type rpcUnionGatewaySvr struct {
	rpcUnionGatewayCore

	mtx      sync.Mutex
	gateways map[uint32]rpcinter.RpcGatewayCli
}

func (s *rpcUnionGatewaySvr) RpcUnionGatewayUploadGatewayInfoAsk(in *message_rpc.RpcUnionGatewayUploadGatewayInfoAsk) error {
	var (
		gatewaId = in.GetGatewaId()
		address  = in.GetAddress()
	)
	newCli := jsn_rpc.NewClient(address, 4, 4)
	s.mtx.Lock()
	cli := s.gateways[gatewaId]
	if nil != cli {
		cli.Cli().Close()
	}
	if nil == s.gateways {
		s.gateways = make(map[uint32]rpcinter.RpcGatewayCli)
	}
	s.gateways[gatewaId] = api.NewRpcGatewayCli(newCli)
	return nil
}

func (s *rpcUnionGatewaySvr) GetGateway(gatewayId uint32) rpcinter.RpcGatewayCli {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	cli := s.gateways[gatewayId]
	if nil == cli {
		return &api.EmptyRpcGatewayCli{}
	}
	return cli
}
