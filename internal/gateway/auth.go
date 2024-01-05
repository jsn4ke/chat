package gateway

import (
	"fmt"
	"time"

	ltvcodec "github.com/jsn4ke/chat/pkg/ltv-codec"
	"github.com/jsn4ke/chat/pkg/pb/message_cli"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
)

var (
	_ jsn_net.Pipe = (*clientHandler)(nil)

	handlers = map[uint32]func(c *clientHandler, session *jsn_net.TcpSession, in message_cli.CliMessage) error{}
)

func init() {
	handlers[uint32(message_cli.CliCmd_CliCmd_RequestSignIn)] = func(c *clientHandler, session *jsn_net.TcpSession, in message_cli.CliMessage) error {
		return auth(c, session, in.(*message_cli.RequestSignIn))
	}
}

type clientHandler struct {
	in       chan any
	inflight []any
	done     chan struct{}

	session *jsn_net.TcpSession

	gateway *gateway

	authed bool

	uid     uint64
	guildId uint64
}

// Close implements jsn_net.Pipe.
func (c *clientHandler) Close() {
	close(c.done)
}

// Post implements jsn_net.Pipe.
func (c *clientHandler) Post(in any) bool {
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
	switch tp := in.(type) {
	case *jsn_net.TcpEventAdd:
		c.session = tp.S
		c.session.SetReadTimeoutOnce(time.Second)
	case *jsn_net.TcpEventClose:
		if 0 == c.uid {
			return
		}
		c.gateway.unbindUser(c.uid, c.session.Sid(), c.guildId)
	case *jsn_net.TcpEventPacket:
		var (
			body  []byte
			err   error
			cmdId uint32
		)
		packet := tp.Packet.(*ltvcodec.LTVReadPacket)
		h, ok := handlers[packet.CmdId]
		if !ok {
			err = fmt.Errorf("%v in coming no handler", packet.CmdId)
		} else {
			err = h(c, tp.S, packet.Body)
		}
		if nil != err {
			resp := &message_cli.ResponseCliFailure{
				InCmdId:     packet.CmdId,
				ErrorString: err.Error(),
			}
			body, err = resp.Marshal()
			if nil != err {
				fmt.Println("handlerIn marshal error")
			}
			cmdId = resp.CmdId()

		} else {
			resp := &message_cli.ResponseCliSucces{
				InCmdId: packet.CmdId,
			}
			body, err = resp.Marshal()
			if nil != err {
				fmt.Println("handlerIn marshal error")
			}
			cmdId = resp.CmdId()
		}
		tp.S.Send(&ltvcodec.LTVWritePacket{
			CmdId: cmdId,
			Body:  body,
		})
		for _, v := range c.inflight {
			v := v
			c.session.Send(v)
		}
	case *jsn_net.TcpEventOutSide:
	}
}

func auth(c *clientHandler, session *jsn_net.TcpSession, in *message_cli.RequestSignIn) error {
	var (
		uid     uint64 = in.GetUid()
		token   []byte = in.GetToken()
		gateway        = c.gateway
		ssid           = session.Sid()
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
			Ssid:      ssid,
			GatewayId: gateway.gatewayId,
		},
	}, c.done, gateway.SigninTimeout())
	if nil != err {
		return err
	}
	gateway.bindUser(uid, ssid, reply.GetGuildId())
	c.uid = uid
	c.guildId = reply.GetGuildId()
	return nil
}
