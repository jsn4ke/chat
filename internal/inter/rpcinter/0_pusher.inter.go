// Code generated by protoc-gen-chat. DO NOT EDIT.
// rpcinter option

package rpcinter

import (
	"github.com/jsn4ke/chat/pkg/pb/message_rpc"
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
)

type RpcPusherCli interface {
	Cli() *jsn_rpc.Client
	RpcPusherChat2GuildSync(in *message_rpc.RpcPusherChat2GuildSync)
	RpcPusherChat2WorldSync(in *message_rpc.RpcPusherChat2WorldSync)
	RpcPusherChat2DirectSync(in *message_rpc.RpcPusherChat2DirectSync)
}
type RpcPusherSvrWrap interface {
	RpcPusherChat2GuildSync(in *message_rpc.RpcPusherChat2GuildSync)
	RpcPusherChat2WorldSync(in *message_rpc.RpcPusherChat2WorldSync)
	RpcPusherChat2DirectSync(in *message_rpc.RpcPusherChat2DirectSync)
}
type RpcPusherSvr interface {
	Run()
	RpcPusherSvrWrap
}
