// Code generated by protoc-gen-chat. DO NOT EDIT.
// rpccli option

package api

import (
	"time"

	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

func NewRpcGatewayCli(cli *jsn_rpc.Client) rpcinter.RpcGatewayCli {
	c := new(rpcGatewayCli)
	c.cli = cli
	return c
}
func NewRpcUnionGatewayCli(cli *jsn_rpc.Client) rpcinter.RpcUnionGatewayCli {
	c := new(rpcUnionGatewayCli)
	c.cli = cli
	return c
}

type (
	rpcGatewayCli struct {
		cli *jsn_rpc.Client
	}
	EmptyRpcGatewayCli struct{}
	rpcUnionGatewayCli struct {
		cli *jsn_rpc.Client
	}
	EmptyRpcUnionGatewayCli struct{}
)

func (c *rpcGatewayCli) Cli() *jsn_rpc.Client {
	return c.cli
}
func (c *EmptyRpcGatewayCli) Cli() *jsn_rpc.Client {
	return nil
}
func (c *rpcUnionGatewayCli) Cli() *jsn_rpc.Client {
	return c.cli
}
func (c *EmptyRpcUnionGatewayCli) Cli() *jsn_rpc.Client {
	return nil
}
func (c *rpcGatewayCli) RpcGatewayKickUserSync(in *message_rpc.RpcGatewayKickUserSync) {
	c.cli.Sync(in)
}
func (c *EmptyRpcGatewayCli) RpcGatewayKickUserSync(in *message_rpc.RpcGatewayKickUserSync) {
}
func (c *rpcGatewayCli) RpcGatewayInfoRequest(in *message_rpc.RpcGatewayInfoRequest, cancel <-chan struct{}, timeout time.Duration) (*message_rpc.RpcGatewayInfoResponse, error) {
	reply := new(message_rpc.RpcGatewayInfoResponse)
	err := c.cli.Call(in, reply, cancel, timeout)
	return reply, err
}
func (c *EmptyRpcGatewayCli) RpcGatewayInfoRequest(in *message_rpc.RpcGatewayInfoRequest, cancel <-chan struct{}, timeout time.Duration) (*message_rpc.RpcGatewayInfoResponse, error) {
	return nil, EmptyRpcCliError
}
func (c *rpcGatewayCli) RpcGatewayChannelPushSync(in *message_rpc.RpcGatewayChannelPushSync) {
	c.cli.Sync(in)
}
func (c *EmptyRpcGatewayCli) RpcGatewayChannelPushSync(in *message_rpc.RpcGatewayChannelPushSync) {
}
func (c *rpcUnionGatewayCli) RpcUnionGatewayUploadGatewayInfoAsk(in *message_rpc.RpcUnionGatewayUploadGatewayInfoAsk, cancel <-chan struct{}, timeout time.Duration) error {
	return c.cli.Ask(in, cancel, timeout)
}
func (c *EmptyRpcUnionGatewayCli) RpcUnionGatewayUploadGatewayInfoAsk(in *message_rpc.RpcUnionGatewayUploadGatewayInfoAsk, cancel <-chan struct{}, timeout time.Duration) error {
	return EmptyRpcCliError
}
