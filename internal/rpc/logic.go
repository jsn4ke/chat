package rpc

import (
	"sync"

	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

type rpcLogicServer struct {
	svr    *jsn_rpc.Server
	in     chan *jsn_rpc.AsyncRpc
	done   <-chan struct{}
	wg     sync.WaitGroup
	runNum int
}

func (s *rpcLogicServer) RpcLogicSigninRequest(in *message_rpc.RpcLogicSigninRequest) (*message_rpc.RpcLogicSigninResponse, error) {
	return nil, nil
}

func (s *rpcLogicServer) RpcLogicSubscribeOrUnsubscribAsk(in *message_rpc.RpcLogicSubscribeOrUnsubscribAsk) error {
	return nil
}

func (s *rpcLogicServer) RpcLogicReSubscribeAsk(in *message_rpc.RpcLogicReSubscribeAsk) error {
	return nil
}
