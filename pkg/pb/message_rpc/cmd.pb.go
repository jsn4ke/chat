// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: message_rpc/cmd.proto

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

type RpcCmd int32

const (
	RpcCmd_RpcCmd_None                                RpcCmd = 0
	RpcCmd_RpcCmd_RpcAuthCheckAsk                     RpcCmd = 10001
	RpcCmd_RpcCmd_RpcLogicSigninRequest               RpcCmd = 20001
	RpcCmd_RpcCmd_RpcLogicSigninResponse              RpcCmd = 20002
	RpcCmd_RpcCmd_RpcLogicSubscribeOrUnsubscribAsk    RpcCmd = 20003
	RpcCmd_RpcCmd_RpcLogicReSubscribeAsk              RpcCmd = 20004
	RpcCmd_RpcCmd_RpcLogicChat2GuildAsk               RpcCmd = 20005
	RpcCmd_RpcCmd_RpcLogicChat2WorldAsk               RpcCmd = 20006
	RpcCmd_RpcCmd_RpcLogicChat2DirectAsk              RpcCmd = 20007
	RpcCmd_RpcCmd_RpcGatewayKickUserSync              RpcCmd = 30001
	RpcCmd_RpcCmd_RpcGatewayChannelPushSync           RpcCmd = 30002
	RpcCmd_RpcCmd_RpcGatewayInfoRequest               RpcCmd = 30003
	RpcCmd_RpcCmd_RpcGatewayInfoResponse              RpcCmd = 30004
	RpcCmd_RpcCmd_RpcUnionGatewayUploadGatewayInfoAsk RpcCmd = 40001
	RpcCmd_RpcCmd_RpcPusherChat2GuildSync             RpcCmd = 50001
	RpcCmd_RpcCmd_RpcPusherChat2WorldSync             RpcCmd = 50002
	RpcCmd_RpcCmd_RpcPusherChat2DirectSync            RpcCmd = 50003
)

// Enum value maps for RpcCmd.
var (
	RpcCmd_name = map[int32]string{
		0:     "RpcCmd_None",
		10001: "RpcCmd_RpcAuthCheckAsk",
		20001: "RpcCmd_RpcLogicSigninRequest",
		20002: "RpcCmd_RpcLogicSigninResponse",
		20003: "RpcCmd_RpcLogicSubscribeOrUnsubscribAsk",
		20004: "RpcCmd_RpcLogicReSubscribeAsk",
		20005: "RpcCmd_RpcLogicChat2GuildAsk",
		20006: "RpcCmd_RpcLogicChat2WorldAsk",
		20007: "RpcCmd_RpcLogicChat2DirectAsk",
		30001: "RpcCmd_RpcGatewayKickUserSync",
		30002: "RpcCmd_RpcGatewayChannelPushSync",
		30003: "RpcCmd_RpcGatewayInfoRequest",
		30004: "RpcCmd_RpcGatewayInfoResponse",
		40001: "RpcCmd_RpcUnionGatewayUploadGatewayInfoAsk",
		50001: "RpcCmd_RpcPusherChat2GuildSync",
		50002: "RpcCmd_RpcPusherChat2WorldSync",
		50003: "RpcCmd_RpcPusherChat2DirectSync",
	}
	RpcCmd_value = map[string]int32{
		"RpcCmd_None":                                0,
		"RpcCmd_RpcAuthCheckAsk":                     10001,
		"RpcCmd_RpcLogicSigninRequest":               20001,
		"RpcCmd_RpcLogicSigninResponse":              20002,
		"RpcCmd_RpcLogicSubscribeOrUnsubscribAsk":    20003,
		"RpcCmd_RpcLogicReSubscribeAsk":              20004,
		"RpcCmd_RpcLogicChat2GuildAsk":               20005,
		"RpcCmd_RpcLogicChat2WorldAsk":               20006,
		"RpcCmd_RpcLogicChat2DirectAsk":              20007,
		"RpcCmd_RpcGatewayKickUserSync":              30001,
		"RpcCmd_RpcGatewayChannelPushSync":           30002,
		"RpcCmd_RpcGatewayInfoRequest":               30003,
		"RpcCmd_RpcGatewayInfoResponse":              30004,
		"RpcCmd_RpcUnionGatewayUploadGatewayInfoAsk": 40001,
		"RpcCmd_RpcPusherChat2GuildSync":             50001,
		"RpcCmd_RpcPusherChat2WorldSync":             50002,
		"RpcCmd_RpcPusherChat2DirectSync":            50003,
	}
)

func (x RpcCmd) Enum() *RpcCmd {
	p := new(RpcCmd)
	*p = x
	return p
}

func (x RpcCmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RpcCmd) Descriptor() protoreflect.EnumDescriptor {
	return file_message_rpc_cmd_proto_enumTypes[0].Descriptor()
}

func (RpcCmd) Type() protoreflect.EnumType {
	return &file_message_rpc_cmd_proto_enumTypes[0]
}

func (x RpcCmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RpcCmd.Descriptor instead.
func (RpcCmd) EnumDescriptor() ([]byte, []int) {
	return file_message_rpc_cmd_proto_rawDescGZIP(), []int{0}
}

var File_message_rpc_cmd_proto protoreflect.FileDescriptor

var file_message_rpc_cmd_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6d,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x70, 0x63, 0x2a, 0xfb, 0x04, 0x0a, 0x06, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x12,
	0x0f, 0x0a, 0x0b, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x16, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x41, 0x75,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x73, 0x6b, 0x10, 0x91, 0x4e, 0x12, 0x22, 0x0a,
	0x1c, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0xa1, 0x9c,
	0x01, 0x12, 0x23, 0x0a, 0x1d, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x10, 0xa2, 0x9c, 0x01, 0x12, 0x2d, 0x0a, 0x27, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64,
	0x5f, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4f, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x41, 0x73,
	0x6b, 0x10, 0xa3, 0x9c, 0x01, 0x12, 0x23, 0x0a, 0x1d, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f,
	0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x41, 0x73, 0x6b, 0x10, 0xa4, 0x9c, 0x01, 0x12, 0x22, 0x0a, 0x1c, 0x52, 0x70,
	0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x43, 0x68, 0x61,
	0x74, 0x32, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x73, 0x6b, 0x10, 0xa5, 0x9c, 0x01, 0x12, 0x22,
	0x0a, 0x1c, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x4c, 0x6f, 0x67, 0x69,
	0x63, 0x43, 0x68, 0x61, 0x74, 0x32, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x41, 0x73, 0x6b, 0x10, 0xa6,
	0x9c, 0x01, 0x12, 0x23, 0x0a, 0x1d, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63,
	0x4c, 0x6f, 0x67, 0x69, 0x63, 0x43, 0x68, 0x61, 0x74, 0x32, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x41, 0x73, 0x6b, 0x10, 0xa7, 0x9c, 0x01, 0x12, 0x23, 0x0a, 0x1d, 0x52, 0x70, 0x63, 0x43, 0x6d,
	0x64, 0x5f, 0x52, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4b, 0x69, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x10, 0xb1, 0xea, 0x01, 0x12, 0x26, 0x0a, 0x20,
	0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x75, 0x73, 0x68, 0x53, 0x79, 0x6e, 0x63,
	0x10, 0xb2, 0xea, 0x01, 0x12, 0x22, 0x0a, 0x1c, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52,
	0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x10, 0xb3, 0xea, 0x01, 0x12, 0x23, 0x0a, 0x1d, 0x52, 0x70, 0x63, 0x43,
	0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0xb4, 0xea, 0x01, 0x12, 0x30, 0x0a,
	0x2a, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x55, 0x6e, 0x69, 0x6f, 0x6e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x73, 0x6b, 0x10, 0xc1, 0xb8, 0x02, 0x12,
	0x24, 0x0a, 0x1e, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x50, 0x75, 0x73,
	0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x32, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x79, 0x6e,
	0x63, 0x10, 0xd1, 0x86, 0x03, 0x12, 0x24, 0x0a, 0x1e, 0x52, 0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f,
	0x52, 0x70, 0x63, 0x50, 0x75, 0x73, 0x68, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x32, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x53, 0x79, 0x6e, 0x63, 0x10, 0xd2, 0x86, 0x03, 0x12, 0x25, 0x0a, 0x1f, 0x52,
	0x70, 0x63, 0x43, 0x6d, 0x64, 0x5f, 0x52, 0x70, 0x63, 0x50, 0x75, 0x73, 0x68, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x74, 0x32, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x10, 0xd3,
	0x86, 0x03, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x73, 0x6e, 0x34, 0x6b, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_rpc_cmd_proto_rawDescOnce sync.Once
	file_message_rpc_cmd_proto_rawDescData = file_message_rpc_cmd_proto_rawDesc
)

func file_message_rpc_cmd_proto_rawDescGZIP() []byte {
	file_message_rpc_cmd_proto_rawDescOnce.Do(func() {
		file_message_rpc_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_rpc_cmd_proto_rawDescData)
	})
	return file_message_rpc_cmd_proto_rawDescData
}

var file_message_rpc_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_rpc_cmd_proto_goTypes = []interface{}{
	(RpcCmd)(0), // 0: message_rpc.RpcCmd
}
var file_message_rpc_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_rpc_cmd_proto_init() }
func file_message_rpc_cmd_proto_init() {
	if File_message_rpc_cmd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_rpc_cmd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_rpc_cmd_proto_goTypes,
		DependencyIndexes: file_message_rpc_cmd_proto_depIdxs,
		EnumInfos:         file_message_rpc_cmd_proto_enumTypes,
	}.Build()
	File_message_rpc_cmd_proto = out.File
	file_message_rpc_cmd_proto_rawDesc = nil
	file_message_rpc_cmd_proto_goTypes = nil
	file_message_rpc_cmd_proto_depIdxs = nil
}
