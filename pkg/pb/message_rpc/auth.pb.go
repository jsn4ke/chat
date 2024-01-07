// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: message_rpc/auth.proto

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

type RpcAuthUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RpcAuthUnit) Reset() {
	*x = RpcAuthUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcAuthUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcAuthUnit) ProtoMessage() {}

func (x *RpcAuthUnit) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcAuthUnit.ProtoReflect.Descriptor instead.
func (*RpcAuthUnit) Descriptor() ([]byte, []int) {
	return file_message_rpc_auth_proto_rawDescGZIP(), []int{0}
}

type RpcAuthCheckAsk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RpcAuthCheckAsk) Reset() {
	*x = RpcAuthCheckAsk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_rpc_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcAuthCheckAsk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcAuthCheckAsk) ProtoMessage() {}

func (x *RpcAuthCheckAsk) ProtoReflect() protoreflect.Message {
	mi := &file_message_rpc_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcAuthCheckAsk.ProtoReflect.Descriptor instead.
func (*RpcAuthCheckAsk) Descriptor() ([]byte, []int) {
	return file_message_rpc_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RpcAuthCheckAsk) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RpcAuthCheckAsk) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_message_rpc_auth_proto protoreflect.FileDescriptor

var file_message_rpc_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x70, 0x63, 0x41, 0x75, 0x74, 0x68,
	0x55, 0x6e, 0x69, 0x74, 0x22, 0x39, 0x0a, 0x0f, 0x52, 0x70, 0x63, 0x41, 0x75, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73,
	0x6e, 0x34, 0x6b, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_rpc_auth_proto_rawDescOnce sync.Once
	file_message_rpc_auth_proto_rawDescData = file_message_rpc_auth_proto_rawDesc
)

func file_message_rpc_auth_proto_rawDescGZIP() []byte {
	file_message_rpc_auth_proto_rawDescOnce.Do(func() {
		file_message_rpc_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_rpc_auth_proto_rawDescData)
	})
	return file_message_rpc_auth_proto_rawDescData
}

var file_message_rpc_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_rpc_auth_proto_goTypes = []interface{}{
	(*RpcAuthUnit)(nil),     // 0: message_rpc.RpcAuthUnit
	(*RpcAuthCheckAsk)(nil), // 1: message_rpc.RpcAuthCheckAsk
}
var file_message_rpc_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_rpc_auth_proto_init() }
func file_message_rpc_auth_proto_init() {
	if File_message_rpc_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_rpc_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcAuthUnit); i {
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
		file_message_rpc_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcAuthCheckAsk); i {
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
			RawDescriptor: file_message_rpc_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_rpc_auth_proto_goTypes,
		DependencyIndexes: file_message_rpc_auth_proto_depIdxs,
		MessageInfos:      file_message_rpc_auth_proto_msgTypes,
	}.Build()
	File_message_rpc_auth_proto = out.File
	file_message_rpc_auth_proto_rawDesc = nil
	file_message_rpc_auth_proto_goTypes = nil
	file_message_rpc_auth_proto_depIdxs = nil
}
