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

	// return error, kickuser==true
	handlers = map[uint32]func(c *clientHandler, session *jsn_net.TcpSession, in message_cli.CliMessage) (error, bool){}
)

func init() {
	handlers[uint32(message_cli.CliCmd_CliCmd_RequestSignIn)] = func(c *clientHandler, session *jsn_net.TcpSession, in message_cli.CliMessage) (error, bool) {
		return auth(c, session, in.(*message_cli.RequestSignIn))
	}
	handlers[uint32(message_cli.CliCmd_CliCmd_RequestChat2Guild)] = func(c *clientHandler, session *jsn_net.TcpSession, in message_cli.CliMessage) (error, bool) {
		return chat2Guild(c, session, in.(*message_cli.RequestChat2Guild))
	}
	handlers[uint32(message_cli.CliCmd_CliCmd_RequestChat2World)] = func(c *clientHandler, session *jsn_net.TcpSession, in message_cli.CliMessage) (error, bool) {
		return chat2World(c, session, in.(*message_cli.RequestChat2World))
	}
	handlers[uint32(message_cli.CliCmd_CliCmd_RequestChat2Direct)] = func(c *clientHandler, session *jsn_net.TcpSession, in message_cli.CliMessage) (error, bool) {
		return chat2Direct(c, session, in.(*message_cli.RequestChat2Direct))
	}
}

type clientHandler struct {
	in       chan any
	inflight []any
	done     chan struct{}

	session *jsn_net.TcpSession

	gateway *gateway

	uid     uint64
	guildId uint64
	world   uint64
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
		// fmt.Println(tp.S.Sid(), "close", tp.ManualClose, tp.Err)
		uid := c.uid
		if 0 == uid {
			return
		}
		c.uid = 0
		c.gateway.unbindUser(uid, c.session.Sid(), c.guildId, c.world)
	case *jsn_net.TcpEventPacket:
		c.session.CancelLastTimeout()
		var (
			body  []byte
			err   error
			cmdId uint32
			kick  bool
		)
		packet := tp.Packet.(*ltvcodec.LTVReadPacket)
		h, ok := handlers[packet.CmdId]
		if !ok {
			err = fmt.Errorf("%v in coming no handler", packet.CmdId)
		} else {
			err, kick = h(c, tp.S, packet.Body)
		}
		outErr := err
		if nil != outErr {
			resp := &message_cli.ResponseCliFailure{
				InCmdId:     packet.CmdId,
				ErrorString: outErr.Error(),
				Kick:        kick,
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
		if kick {
			c.session.Close()
		}
	case *jsn_net.TcpEventOutSide:
	}
}

func auth(c *clientHandler, session *jsn_net.TcpSession, in *message_cli.RequestSignIn) (error, bool) {
	var (
		uid     uint64 = in.GetUid()
		token   []byte = in.GetToken()
		gateway        = c.gateway
		ssid           = session.Sid()
	)
	if 0 != c.uid {
		return fmt.Errorf("session authed"), true
	}
	err := gateway.authCluster.Get().RpcAuthCheckAsk(&message_rpc.RpcAuthCheckAsk{
		Uid:   uid,
		Token: token,
	}, c.done, gateway.AuthRpcTimeout())
	if nil != err {
		return err, true
	}
	reply, err := gateway.logicCluster.Get().RpcLogicSigninRequest(&message_rpc.RpcLogicSigninRequest{
		Uid: uid,
		ClientId: &message_rpc.ClientId{
			Ssid:      ssid,
			GatewayId: gateway.gatewayId,
		},
	}, c.done, gateway.SigninTimeout())
	if nil != err {
		return err, true
	}
	gateway.bindUser(uid, ssid, reply.GetGuildId(), reply.GetWorldId())
	c.uid = uid
	c.guildId = reply.GetGuildId()
	c.world = reply.GetWorldId()
	return nil, false
}

func chat2Guild(c *clientHandler, session *jsn_net.TcpSession, in *message_cli.RequestChat2Guild) (error, bool) {
	if 0 == c.uid {
		return fmt.Errorf("session unauthed"), true
	}
	var (
		text    = in.GetText()
		gateway = c.gateway
	)
	if 0 == len(text) {
		return fmt.Errorf("chat2Guild no text"), false
	}
	if 0 == c.guildId {
		return fmt.Errorf("not in guild"), false
	}
	err := gateway.logicCluster.Get().RpcLogicChat2GuildAsk(&message_rpc.RpcLogicChat2GuildAsk{
		Uid:  c.uid,
		Text: text,
	}, c.done, gateway.ChatTimeout())
	if nil != err {
		return err, false
	}
	return nil, false
}

func chat2World(c *clientHandler, session *jsn_net.TcpSession, in *message_cli.RequestChat2World) (error, bool) {
	if 0 == c.uid {
		return fmt.Errorf("session unauthed"), true
	}
	var (
		text    = in.GetText()
		gateway = c.gateway
	)
	if 0 == len(text) {
		return fmt.Errorf("chat2Guild no text"), false
	}
	if 0 == c.world {
		return fmt.Errorf("not in guild"), false
	}
	err := gateway.logicCluster.Get().RpcLogicChat2WorldAsk(&message_rpc.RpcLogicChat2WorldAsk{
		Uid:  c.uid,
		Text: text,
	}, c.done, gateway.ChatTimeout())
	if nil != err {
		return err, false
	}
	return nil, false
}

func chat2Direct(c *clientHandler, session *jsn_net.TcpSession, in *message_cli.RequestChat2Direct) (error, bool) {
	if 0 == c.uid {
		return fmt.Errorf("session unauthed"), true
	}
	var (
		rid     = in.GetReceiveId()
		text    = in.GetText()
		gateway = c.gateway
	)
	if 0 == len(text) {
		return fmt.Errorf("chat2Guild no text"), false
	}
	if 0 == rid {
		return fmt.Errorf("not in guild"), false
	}
	err := gateway.logicCluster.Get().RpcLogicChat2DirectAsk(&message_rpc.RpcLogicChat2DirectAsk{
		Uid:       c.uid,
		ReceiveId: rid,
		Text:      text,
	}, c.done, gateway.ChatTimeout())
	if nil != err {
		return err, false
	}
	return nil, false
}
