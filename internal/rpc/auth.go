package rpc

import (
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

type rpcAuthServer struct {
	svr *jsn_rpc.Server
	in  chan *jsn_rpc.AsyncRpc
}

func (s *rpcAuthServer) RpcAuthCheckAsk(in *message_rpc.RpcAuthCheckAsk) error {
	return nil
}
