package gateway

import (
	"fmt"

	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	ltvcodec "github.com/jsn4ke/chat/pkg/ltv-codec"
	"github.com/jsn4ke/chat/pkg/pb/message_cli"
	"github.com/jsn4ke/chat/pkg/pb/message_obj"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	"github.com/jsn4ke/jsn_net"
)

var (
	_ rpcinter.RpcGatewaySvrWrap = (*gateway)(nil)
)

// RpcGatewayInfoRequest implements rpcinter.RpcGatewaySvr.
func (g *gateway) RpcGatewayInfoRequest(in *message_rpc.RpcGatewayInfoRequest) (*message_rpc.RpcGatewayInfoResponse, error) {
	return &message_rpc.RpcGatewayInfoResponse{
		GatewayId: g.gatewayId,
	}, nil
}

// RpcGatewayKickUserSync implements rpcinter.RpcGatewaySvr.
func (g *gateway) RpcGatewayKickUserSync(in *message_rpc.RpcGatewayKickUserSync) {
	var (
		uid      = in.GetUid()
		gatewaId = in.GetClientId().GetGatewayId()
		ssid     = in.GetClientId().GetSsid()
	)
	if g.gatewayId != gatewaId {
		return
	}

	g.mtx.Lock()

	currentSsid := g.users[uid]
	g.mtx.Unlock()

	if currentSsid != ssid {
		return
	}
	sess := g.cliAcceptor.Session(currentSsid)
	if nil == sess {
		fmt.Println("no session while receive uid in user map")
		return
	}
	resp := &message_cli.ResponseCliFailure{}
	resp.ErrorString = "signin in other location"
	resp.Kick = true
	body, _ := resp.Marshal()
	sess.Send(&ltvcodec.LTVWritePacket{
		CmdId: resp.CmdId(),
		Body:  body,
	})
	sess.Close()
}

// RpcGatewayChannelPushSync implements rpcinter.RpcGatewaySvr.
func (g *gateway) RpcGatewayChannelPushSync(in *message_rpc.RpcGatewayChannelPushSync) {
	var (
		ssids = map[uint64]*message_cli.ResponseChatInfo{}
	)

	g.mtx.Lock()
	for _, v := range in.GetCm() {
		v := v
		switch v.GetCt() {
		case message_obj.ChannelMessageType_ChannelMessageType_Drirect:
			info := ssids[v.GetChannelId()]
			if nil == info {
				info = &message_cli.ResponseChatInfo{}
				ssids[v.GetChannelId()] = info
			}
			v.ChannelId = 0
			info.Cms = append(info.Cms, v)
		default:
			for _, vv := range g.channels[v.GetCt()][v.GetChannelId()] {
				info := ssids[vv]
				if nil == info {
					info = &message_cli.ResponseChatInfo{}
					ssids[vv] = info
				}
				info.Cms = append(info.Cms, v)
			}
		}
	}
	g.mtx.Unlock()

	for ssid, v := range ssids {
		v := v
		session := g.cliAcceptor.Session(ssid)
		if nil == session {
			continue
		}
		jsn_net.WaitGo(&g.wg, func() {
			body, err := v.Marshal()
			if nil != err {
				return
			}
			session.Send(&ltvcodec.LTVWritePacket{
				CmdId: v.CmdId(),
				Body:  body,
			})
		})
	}
}
