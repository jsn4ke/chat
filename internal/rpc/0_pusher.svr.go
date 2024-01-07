// Code generated by protoc-gen-chat. DO NOT EDIT.
// rpcsvr option

package rpc

import (
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
)

/*
	func NewRpcPusherSvr(svr *jsn_rpc.Server, runNum int, done <-chan struct{}) rpcinter.RpcPusherSvr {
		s := new(rpcPusherSvr)
		c := new(rpcCore)
		c.svr = svr
		c.in = make(chan *jsn_rpc.AsyncRpc, 128)
		c.done = done
		c.runNum = runNum

		s.rpcPusherCore.rpcCore = c
		s.rpcPusherCore.wrap = s
		return s
	}

	type rpcPusherSvr struct {
		rpcPusherCore
	}

	func (s *rpcPusherSvr) Run() {
		s.registerRpc()

		s.svr.Start()
		for i := 0; i < s.runNum; i++ {
			jsn_net.WaitGo(&s.wg, s.run)
		}
		s.wg.Wait()
	}

func(s *rpcPusherSvr)RpcPusherChat2GuildSync(in *message_rpc.RpcPusherChat2GuildSync) {}
func(s *rpcPusherSvr)RpcPusherChat2WorldSync(in *message_rpc.RpcPusherChat2WorldSync) {}
func(s *rpcPusherSvr)RpcPusherChat2DirectSync(in *message_rpc.RpcPusherChat2DirectSync) {}
*/
type rpcPusherCore struct {
	*rpcCore
	wrap rpcinter.RpcPusherSvrWrap
}

func (s *rpcPusherCore) registerRpc() {
	s.svr.RegisterExecutor(new(message_rpc.RpcPusherChat2GuildSync), s.syncRpc)
	s.svr.RegisterExecutor(new(message_rpc.RpcPusherChat2WorldSync), s.syncRpc)
	s.svr.RegisterExecutor(new(message_rpc.RpcPusherChat2DirectSync), s.syncRpc)
}
func (s *rpcPusherCore) syncRpc(in jsn_rpc.RpcUnit) (jsn_rpc.RpcUnit, error) {
	wrap := jsn_rpc.AsyncRpcPool.Get()
	wrap.In = in
	wrap.Error = make(chan error, 1)
	s.in <- wrap
	err := <-wrap.Error
	return wrap.Reply, err
}
func (s *rpcPusherCore) run() {
	for {
		select {
		case <-s.done:
			select {
			case in := <-s.in:
				in.Error <- RpcDownError
			default:
				return
			}
		case in := <-s.in:
			s.handleIn(in)
		}
	}
}
func (s *rpcPusherCore) handleIn(wrap *jsn_rpc.AsyncRpc) {
	var err error
	defer func() {
		wrap.Error <- err
	}()
	switch in := wrap.In.(type) {
	case *message_rpc.RpcPusherChat2GuildSync:
		// func(s *rpcPusherSvr)RpcPusherChat2GuildSync(in *message_rpc.RpcPusherChat2GuildSync)
		s.wrap.RpcPusherChat2GuildSync(in)
	case *message_rpc.RpcPusherChat2WorldSync:
		// func(s *rpcPusherSvr)RpcPusherChat2WorldSync(in *message_rpc.RpcPusherChat2WorldSync)
		s.wrap.RpcPusherChat2WorldSync(in)
	case *message_rpc.RpcPusherChat2DirectSync:
		// func(s *rpcPusherSvr)RpcPusherChat2DirectSync(in *message_rpc.RpcPusherChat2DirectSync)
		s.wrap.RpcPusherChat2DirectSync(in)
	default:
		err = InvalidRpcInputError
	}
}