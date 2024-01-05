// Code generated by protoc-gen-chat. DO NOT EDIT.
// rpcsvr option

package rpc

import (
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

// type rpcLogicServer struct  {
// svr *jsn_rpc.Server
// in  chan *jsn_rpc.AsyncRpc
// }
func (s *rpcLogicServer) RegisterRpc() {
	s.svr.RegisterExecutor(new(message_rpc.RpcLogicSigninRequest), s.SyncRpc)
	s.svr.RegisterExecutor(new(message_rpc.RpcLogicSubscribeOrUnsubscribAsk), s.SyncRpc)
	s.svr.RegisterExecutor(new(message_rpc.RpcLogicReSubscribeAsk), s.SyncRpc)
}
func (s *rpcLogicServer) SyncRpc(in jsn_rpc.RpcUnit) (jsn_rpc.RpcUnit, error) {
	wrap := jsn_rpc.AsyncRpcPool.Get()
	wrap.Error = make(chan error, 1)
	s.in <- wrap
	err := <-wrap.Error
	return wrap.Reply, err
}
func (s *rpcLogicServer) handleIn(wrap *jsn_rpc.AsyncRpc) {
	var err error
	defer func() {
		wrap.Error <- err
	}()
	switch in := wrap.In.(type) {
	case *message_rpc.RpcLogicSigninRequest:
		// func(s *rpcLogicServer)RpcLogicSigninRequest(in *message_rpc.RpcLogicSigninRequest)(*message_rpc.RpcLogicSigninResponse, error)
		wrap.Reply, err = s.RpcLogicSigninRequest(in)
	case *message_rpc.RpcLogicSubscribeOrUnsubscribAsk:
		// func(s *rpcLogicServer)RpcLogicSubscribeOrUnsubscribAsk(in *message_rpc.RpcLogicSubscribeOrUnsubscribAsk) error
		err = s.RpcLogicSubscribeOrUnsubscribAsk(in)
	case *message_rpc.RpcLogicReSubscribeAsk:
		// func(s *rpcLogicServer)RpcLogicReSubscribeAsk(in *message_rpc.RpcLogicReSubscribeAsk) error
		err = s.RpcLogicReSubscribeAsk(in)
	default:
		err = InvalidRpcInputError
	}
}