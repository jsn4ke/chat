// Code generated by protoc-gen-chat. DO NOT EDIT.
// rpcext option
package message_rpc

import (
	jsn_rpc "github.com/jsn4ke/jsn_net/rpc"
	"google.golang.org/protobuf/proto"
)

var (
	_ jsn_rpc.RpcUnit = (*RpcGatewayKickUserSync)(nil)
	_ jsn_rpc.RpcUnit = (*RpcGatewayInfoRequest)(nil)
	_ jsn_rpc.RpcUnit = (*RpcGatewayInfoResponse)(nil)
	_ jsn_rpc.RpcUnit = (*RpcGatewayChannelPushSync)(nil)
	_ jsn_rpc.RpcUnit = (*RpcUnionGatewayUploadGatewayInfoAsk)(nil)
)

func (*RpcGatewayKickUserSync) CmdId() uint32 {
	return uint32(RpcCmd_RpcCmd_RpcGatewayKickUserSync)
}
func (x *RpcGatewayKickUserSync) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}
func (x *RpcGatewayKickUserSync) Unmarshal(in []byte) error {
	return proto.Unmarshal(in, x)
}
func (*RpcGatewayKickUserSync) New() jsn_rpc.RpcUnit {
	return new(RpcGatewayKickUserSync)
}
func (*RpcGatewayInfoRequest) CmdId() uint32 {
	return uint32(RpcCmd_RpcCmd_RpcGatewayInfoRequest)
}
func (x *RpcGatewayInfoRequest) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}
func (x *RpcGatewayInfoRequest) Unmarshal(in []byte) error {
	return proto.Unmarshal(in, x)
}
func (*RpcGatewayInfoRequest) New() jsn_rpc.RpcUnit {
	return new(RpcGatewayInfoRequest)
}
func (*RpcGatewayInfoResponse) CmdId() uint32 {
	return uint32(RpcCmd_RpcCmd_RpcGatewayInfoResponse)
}
func (x *RpcGatewayInfoResponse) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}
func (x *RpcGatewayInfoResponse) Unmarshal(in []byte) error {
	return proto.Unmarshal(in, x)
}
func (*RpcGatewayInfoResponse) New() jsn_rpc.RpcUnit {
	return new(RpcGatewayInfoResponse)
}
func (*RpcGatewayChannelPushSync) CmdId() uint32 {
	return uint32(RpcCmd_RpcCmd_RpcGatewayChannelPushSync)
}
func (x *RpcGatewayChannelPushSync) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}
func (x *RpcGatewayChannelPushSync) Unmarshal(in []byte) error {
	return proto.Unmarshal(in, x)
}
func (*RpcGatewayChannelPushSync) New() jsn_rpc.RpcUnit {
	return new(RpcGatewayChannelPushSync)
}
func (*RpcUnionGatewayUploadGatewayInfoAsk) CmdId() uint32 {
	return uint32(RpcCmd_RpcCmd_RpcUnionGatewayUploadGatewayInfoAsk)
}
func (x *RpcUnionGatewayUploadGatewayInfoAsk) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}
func (x *RpcUnionGatewayUploadGatewayInfoAsk) Unmarshal(in []byte) error {
	return proto.Unmarshal(in, x)
}
func (*RpcUnionGatewayUploadGatewayInfoAsk) New() jsn_rpc.RpcUnit {
	return new(RpcUnionGatewayUploadGatewayInfoAsk)
}
