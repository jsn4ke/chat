// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: message_rpc/logic.proto

package message_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RpcLogicUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RpcLogicUnit) Reset() {
	*x = RpcLogicUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicUnit) ProtoMessage() {}

func (x *RpcLogicUnit) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicUnit.ProtoReflect.Descriptor instead.
func (*RpcLogicUnit) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{0}
}

type RpcLogicSigninRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      uint64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ClientId *ClientId `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *RpcLogicSigninRequest) Reset() {
	*x = RpcLogicSigninRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicSigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicSigninRequest) ProtoMessage() {}

func (x *RpcLogicSigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicSigninRequest.ProtoReflect.Descriptor instead.
func (*RpcLogicSigninRequest) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{1}
}

func (x *RpcLogicSigninRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RpcLogicSigninRequest) GetClientId() *ClientId {
	if x != nil {
		return x.ClientId
	}
	return nil
}

type RpcLogicSigninResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId uint64 `protobuf:"varint,1,opt,name=guildId,proto3" json:"guildId,omitempty"`
	WorldId uint64 `protobuf:"varint,2,opt,name=worldId,proto3" json:"worldId,omitempty"`
}

func (x *RpcLogicSigninResponse) Reset() {
	*x = RpcLogicSigninResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicSigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicSigninResponse) ProtoMessage() {}

func (x *RpcLogicSigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicSigninResponse.ProtoReflect.Descriptor instead.
func (*RpcLogicSigninResponse) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{2}
}

func (x *RpcLogicSigninResponse) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *RpcLogicSigninResponse) GetWorldId() uint64 {
	if x != nil {
		return x.WorldId
	}
	return 0
}

type RpcLogicSubscribeOrUnsubscribAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId uint32   `protobuf:"varint,1,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	GuildsAdd []uint64 `protobuf:"varint,2,rep,packed,name=guildsAdd,proto3" json:"guildsAdd,omitempty"`
	GuildsRem []uint64 `protobuf:"varint,3,rep,packed,name=guildsRem,proto3" json:"guildsRem,omitempty"`
	WorldAdd  []uint64 `protobuf:"varint,4,rep,packed,name=worldAdd,proto3" json:"worldAdd,omitempty"`
	WorldRem  []uint64 `protobuf:"varint,5,rep,packed,name=worldRem,proto3" json:"worldRem,omitempty"`
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) Reset() {
	*x = RpcLogicSubscribeOrUnsubscribAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicSubscribeOrUnsubscribAsk) ProtoMessage() {}

func (x *RpcLogicSubscribeOrUnsubscribAsk) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicSubscribeOrUnsubscribAsk.ProtoReflect.Descriptor instead.
func (*RpcLogicSubscribeOrUnsubscribAsk) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{3}
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) GetGatewayId() uint32 {
	if x != nil {
		return x.GatewayId
	}
	return 0
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) GetGuildsAdd() []uint64 {
	if x != nil {
		return x.GuildsAdd
	}
	return nil
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) GetGuildsRem() []uint64 {
	if x != nil {
		return x.GuildsRem
	}
	return nil
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) GetWorldAdd() []uint64 {
	if x != nil {
		return x.WorldAdd
	}
	return nil
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) GetWorldRem() []uint64 {
	if x != nil {
		return x.WorldRem
	}
	return nil
}

type RpcLogicReSubscribeAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId uint32   `protobuf:"varint,1,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	Guilds    []uint64 `protobuf:"varint,2,rep,packed,name=guilds,proto3" json:"guilds,omitempty"`
	Worlds    []uint64 `protobuf:"varint,3,rep,packed,name=worlds,proto3" json:"worlds,omitempty"`
}

func (x *RpcLogicReSubscribeAsk) Reset() {
	*x = RpcLogicReSubscribeAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicReSubscribeAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicReSubscribeAsk) ProtoMessage() {}

func (x *RpcLogicReSubscribeAsk) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicReSubscribeAsk.ProtoReflect.Descriptor instead.
func (*RpcLogicReSubscribeAsk) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{4}
}

func (x *RpcLogicReSubscribeAsk) GetGatewayId() uint32 {
	if x != nil {
		return x.GatewayId
	}
	return 0
}

func (x *RpcLogicReSubscribeAsk) GetGuilds() []uint64 {
	if x != nil {
		return x.Guilds
	}
	return nil
}

func (x *RpcLogicReSubscribeAsk) GetWorlds() []uint64 {
	if x != nil {
		return x.Worlds
	}
	return nil
}

type RpcLogicChat2GuildAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *RpcLogicChat2GuildAsk) Reset() {
	*x = RpcLogicChat2GuildAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicChat2GuildAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicChat2GuildAsk) ProtoMessage() {}

func (x *RpcLogicChat2GuildAsk) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicChat2GuildAsk.ProtoReflect.Descriptor instead.
func (*RpcLogicChat2GuildAsk) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{5}
}

func (x *RpcLogicChat2GuildAsk) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RpcLogicChat2GuildAsk) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type RpcLogicChat2WorldAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *RpcLogicChat2WorldAsk) Reset() {
	*x = RpcLogicChat2WorldAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicChat2WorldAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicChat2WorldAsk) ProtoMessage() {}

func (x *RpcLogicChat2WorldAsk) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicChat2WorldAsk.ProtoReflect.Descriptor instead.
func (*RpcLogicChat2WorldAsk) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{6}
}

func (x *RpcLogicChat2WorldAsk) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RpcLogicChat2WorldAsk) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type RpcLogicChat2DirectAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ReceiveId uint64 `protobuf:"varint,2,opt,name=receiveId,proto3" json:"receiveId,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *RpcLogicChat2DirectAsk) Reset() {
	*x = RpcLogicChat2DirectAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_logic_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicChat2DirectAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicChat2DirectAsk) ProtoMessage() {}

func (x *RpcLogicChat2DirectAsk) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_logic_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcLogicChat2DirectAsk.ProtoReflect.Descriptor instead.
func (*RpcLogicChat2DirectAsk) Descriptor() ([]byte, []int) {
	return file_message_rpc_logic_proto_rawDescGZIP(), []int{7}
}

func (x *RpcLogicChat2DirectAsk) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RpcLogicChat2DirectAsk) GetReceiveId() uint64 {
	if x != nil {
		return x.ReceiveId
	}
	return 0
}

func (x *RpcLogicChat2DirectAsk) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_message_rpc_logic_proto protoreflect.FileDescriptor

var file_message_rpc_logic_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x1a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x55, 0x6e, 0x69, 0x74,
	0x22, 0x5c, 0x0a, 0x15, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4c,
	0x0a, 0x16, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x22, 0xb4, 0x01, 0x0a,
	0x20, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4f, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x41, 0x73,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x41, 0x64, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x41, 0x64, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x41, 0x64, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x41, 0x64, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x52, 0x65, 0x6d, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x52, 0x65, 0x6d, 0x22, 0x66, 0x0a, 0x16, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x73, 0x6b, 0x12, 0x1c, 0x0a,
	0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x73, 0x22, 0x3d, 0x0a, 0x15, 0x52,
	0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x43, 0x68, 0x61, 0x74, 0x32, 0x47, 0x75, 0x69, 0x6c,
	0x64, 0x41, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3d, 0x0a, 0x15, 0x52, 0x70,
	0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x43, 0x68, 0x61, 0x74, 0x32, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x41, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x5c, 0x0a, 0x16, 0x52, 0x70, 0x63,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x43, 0x68, 0x61, 0x74, 0x32, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x41, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x6e, 0x34, 0x6b, 0x65, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_rpc_logic_proto_rawDescOnce sync.Once
	file_message_rpc_logic_proto_rawDescData = file_message_rpc_logic_proto_rawDesc
)

func file_message_rpc_logic_proto_rawDescGZIP() []byte {
	file_message_rpc_logic_proto_rawDescOnce.Do(func() {
		file_message_rpc_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_rpc_logic_proto_rawDescData)
	})
	return file_message_rpc_logic_proto_rawDescData
}

var file_message_rpc_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_rpc_logic_proto_goTypes = []interface{}{
	(*RpcLogicUnit)(nil),                     // 0: message_rpc.RpcLogicUnit
	(*RpcLogicSigninRequest)(nil),            // 1: message_rpc.RpcLogicSigninRequest
	(*RpcLogicSigninResponse)(nil),           // 2: message_rpc.RpcLogicSigninResponse
	(*RpcLogicSubscribeOrUnsubscribAsk)(nil), // 3: message_rpc.RpcLogicSubscribeOrUnsubscribAsk
	(*RpcLogicReSubscribeAsk)(nil),           // 4: message_rpc.RpcLogicReSubscribeAsk
	(*RpcLogicChat2GuildAsk)(nil),            // 5: message_rpc.RpcLogicChat2GuildAsk
	(*RpcLogicChat2WorldAsk)(nil),            // 6: message_rpc.RpcLogicChat2WorldAsk
	(*RpcLogicChat2DirectAsk)(nil),           // 7: message_rpc.RpcLogicChat2DirectAsk
	(*ClientId)(nil),                         // 8: message_rpc.ClientId
}
var file_message_rpc_logic_proto_depIdxs = []int32{
	8, // 0: message_rpc.RpcLogicSigninRequest.clientId:type_name -> message_rpc.ClientId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_rpc_logic_proto_init() }
func file_message_rpc_logic_proto_init() {
	if File_message_rpc_logic_proto != nil {
		return
	}
	file_message_rpc_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_rpc_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicUnit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_rpc_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicSigninRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_rpc_logic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicSigninResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_rpc_logic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicSubscribeOrUnsubscribAsk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_rpc_logic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicReSubscribeAsk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_rpc_logic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicChat2GuildAsk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_rpc_logic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicChat2WorldAsk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_rpc_logic_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcLogicChat2DirectAsk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_rpc_logic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_rpc_logic_proto_goTypes,
		DependencyIndexes: file_message_rpc_logic_proto_depIdxs,
		MessageInfos:      file_message_rpc_logic_proto_msgTypes,
	}.Build()
	File_message_rpc_logic_proto = out.File
	file_message_rpc_logic_proto_rawDesc = nil
	file_message_rpc_logic_proto_goTypes = nil
	file_message_rpc_logic_proto_depIdxs = nil
}
