// Code generated by protoc-gen-chat. DO NOT EDIT.
// rpccli option

package api

import (
	"github.com/jsn4ke/chat/internal/inter/rpcinter"
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

func NewRpcPusherCli(cli *jsn_rpc.Client) rpcinter.RpcPusherCli {
	c := new(rpcPusherCli)
	c.cli = cli
	return c
}

type (
	rpcPusherCli struct {
		cli *jsn_rpc.Client
	}
	EmptyRpcPusherCli struct{}
)

func (c *rpcPusherCli) Cli() *jsn_rpc.Client {
	return c.cli
}
func (c *EmptyRpcPusherCli) Cli() *jsn_rpc.Client {
	return nil
}
func (c *rpcPusherCli) RpcPusherChat2GuildSync(in *message_rpc.RpcPusherChat2GuildSync) {
	c.cli.Sync(in)
}
func (c *EmptyRpcPusherCli) RpcPusherChat2GuildSync(in *message_rpc.RpcPusherChat2GuildSync) {
}
func (c *rpcPusherCli) RpcPusherChat2WorldSync(in *message_rpc.RpcPusherChat2WorldSync) {
	c.cli.Sync(in)
}
func (c *EmptyRpcPusherCli) RpcPusherChat2WorldSync(in *message_rpc.RpcPusherChat2WorldSync) {
}
func (c *rpcPusherCli) RpcPusherChat2DirectSync(in *message_rpc.RpcPusherChat2DirectSync) {
	c.cli.Sync(in)
}
func (c *EmptyRpcPusherCli) RpcPusherChat2DirectSync(in *message_rpc.RpcPusherChat2DirectSync) {
}
