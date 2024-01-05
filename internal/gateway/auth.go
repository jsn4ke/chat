package gateway

import (
	"fmt"
	"sync/atomic"

	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
)

var (
	_ jsn_net.Pipe = (*clientHandler)(nil)
)

type clientHandler struct {
	in   chan any
	done chan struct{}

	closed int32

	authed bool
}

// Close implements jsn_net.Pipe.
func (c *clientHandler) Close() {
	atomic.StoreInt32(&c.closed, 1)

	select {
	case c.done <- struct{}{}:
	default:
	}
}

// Post implements jsn_net.Pipe.
func (c *clientHandler) Post(in any) bool {
	if 1 == atomic.LoadInt32(&c.closed) {
		return false
	}
	c.in <- in
	return true
}

// Run implements jsn_net.Pipe.
func (c *clientHandler) Run() {
	defer func() {
		for {
			select {
			case <-c.in:
			default:
				return
			}
		}
	}()
	for {
		select {
		case in := <-c.in:
			c.handlerIn(in)
		case <-c.done:
			return
		}
	}
}

func (c *clientHandler) handlerIn(in any) {
}

func auth(c *clientHandler, session *jsn_net.TcpSession, gateway *gateway) error {
	var (
		uid   uint64
		token []byte
	)
	if c.authed {
		return fmt.Errorf("session authed")
	}
	err := gateway.authCluster.Get().RpcAuthCheckAsk(&message_rpc.RpcAuthCheckAsk{
		Uid:   uid,
		Token: token,
	}, c.done, gateway.AuthRpcTimeout())
	if nil != err {
		return err
	}
	reply, err := gateway.logicCluster.Get().RpcLogicSigninRequest(&message_rpc.RpcLogicSigninRequest{
		ClientId: &message_rpc.ClientId{
			Ssid:      session.Sid(),
			GatewayId: gateway.gatewayId,
		},
	}, c.done, gateway.SigninTimeout())
	if nil != err {
		return err
	}
	if 0 != reply.GetGuildId() {
		gateway.Subscrib(channelTypeGuild, reply.GetGuildId(), uid)
	}
	return nil
}
