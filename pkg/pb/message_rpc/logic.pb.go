// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: rpc/logic.proto

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
		mi := &file_rpc_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicUnit) ProtoMessage() {}

func (x *RpcLogicUnit) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_logic_proto_msgTypes[0]
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
	return file_rpc_logic_proto_rawDescGZIP(), []int{0}
}

type RpcLogicSigninRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId *ClientId `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *RpcLogicSigninRequest) Reset() {
	*x = RpcLogicSigninRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicSigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicSigninRequest) ProtoMessage() {}

func (x *RpcLogicSigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_logic_proto_msgTypes[1]
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
	return file_rpc_logic_proto_rawDescGZIP(), []int{1}
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
}

func (x *RpcLogicSigninResponse) Reset() {
	*x = RpcLogicSigninResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_logic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicSigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicSigninResponse) ProtoMessage() {}

func (x *RpcLogicSigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_logic_proto_msgTypes[2]
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
	return file_rpc_logic_proto_rawDescGZIP(), []int{2}
}

func (x *RpcLogicSigninResponse) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

type RpcLogicSubscribeOrUnsubscribAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId uint32 `protobuf:"varint,1,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	// key-bool true-sub false-unsub
	Guild map[uint64]bool `protobuf:"bytes,2,rep,name=guild,proto3" json:"guild,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) Reset() {
	*x = RpcLogicSubscribeOrUnsubscribAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_logic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicSubscribeOrUnsubscribAsk) ProtoMessage() {}

func (x *RpcLogicSubscribeOrUnsubscribAsk) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_logic_proto_msgTypes[3]
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
	return file_rpc_logic_proto_rawDescGZIP(), []int{3}
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) GetGatewayId() uint32 {
	if x != nil {
		return x.GatewayId
	}
	return 0
}

func (x *RpcLogicSubscribeOrUnsubscribAsk) GetGuild() map[uint64]bool {
	if x != nil {
		return x.Guild
	}
	return nil
}

type RpcLogicReSubscribeAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId uint32 `protobuf:"varint,1,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	// key-sub
	Guild map[uint64]bool `protobuf:"bytes,2,rep,name=guild,proto3" json:"guild,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *RpcLogicReSubscribeAsk) Reset() {
	*x = RpcLogicReSubscribeAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_logic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcLogicReSubscribeAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcLogicReSubscribeAsk) ProtoMessage() {}

func (x *RpcLogicReSubscribeAsk) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_logic_proto_msgTypes[4]
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
	return file_rpc_logic_proto_rawDescGZIP(), []int{4}
}

func (x *RpcLogicReSubscribeAsk) GetGatewayId() uint32 {
	if x != nil {
		return x.GatewayId
	}
	return 0
}

func (x *RpcLogicReSubscribeAsk) GetGuild() map[uint64]bool {
	if x != nil {
		return x.Guild
	}
	return nil
}

var File_rpc_logic_proto protoreflect.FileDescriptor

var file_rpc_logic_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x1a, 0x10,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x55, 0x6e, 0x69, 0x74,
	0x22, 0x4a, 0x0a, 0x15, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x16,
	0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64,
	0x22, 0xca, 0x01, 0x0a, 0x20, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x41, 0x73, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4f, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x41, 0x73,
	0x6b, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x1a, 0x38, 0x0a, 0x0a, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb6, 0x01,
	0x0a, 0x16, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x73, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x73, 0x6b, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x1a, 0x38, 0x0a, 0x0a,
	0x47, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_logic_proto_rawDescOnce sync.Once
	file_rpc_logic_proto_rawDescData = file_rpc_logic_proto_rawDesc
)

func file_rpc_logic_proto_rawDescGZIP() []byte {
	file_rpc_logic_proto_rawDescOnce.Do(func() {
		file_rpc_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_logic_proto_rawDescData)
	})
	return file_rpc_logic_proto_rawDescData
}

var file_rpc_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_logic_proto_goTypes = []interface{}{
	(*RpcLogicUnit)(nil),                     // 0: message_rpc.RpcLogicUnit
	(*RpcLogicSigninRequest)(nil),            // 1: message_rpc.RpcLogicSigninRequest
	(*RpcLogicSigninResponse)(nil),           // 2: message_rpc.RpcLogicSigninResponse
	(*RpcLogicSubscribeOrUnsubscribAsk)(nil), // 3: message_rpc.RpcLogicSubscribeOrUnsubscribAsk
	(*RpcLogicReSubscribeAsk)(nil),           // 4: message_rpc.RpcLogicReSubscribeAsk
	nil,                                      // 5: message_rpc.RpcLogicSubscribeOrUnsubscribAsk.GuildEntry
	nil,                                      // 6: message_rpc.RpcLogicReSubscribeAsk.GuildEntry
	(*ClientId)(nil),                         // 7: message_rpc.ClientId
}
var file_rpc_logic_proto_depIdxs = []int32{
	7, // 0: message_rpc.RpcLogicSigninRequest.clientId:type_name -> message_rpc.ClientId
	5, // 1: message_rpc.RpcLogicSubscribeOrUnsubscribAsk.guild:type_name -> message_rpc.RpcLogicSubscribeOrUnsubscribAsk.GuildEntry
	6, // 2: message_rpc.RpcLogicReSubscribeAsk.guild:type_name -> message_rpc.RpcLogicReSubscribeAsk.GuildEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_logic_proto_init() }
func file_rpc_logic_proto_init() {
	if File_rpc_logic_proto != nil {
		return
	}
	file_rpc_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_logic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_logic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_logic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_logic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_logic_proto_goTypes,
		DependencyIndexes: file_rpc_logic_proto_depIdxs,
		MessageInfos:      file_rpc_logic_proto_msgTypes,
	}.Build()
	File_rpc_logic_proto = out.File
	file_rpc_logic_proto_rawDesc = nil
	file_rpc_logic_proto_goTypes = nil
	file_rpc_logic_proto_depIdxs = nil
}