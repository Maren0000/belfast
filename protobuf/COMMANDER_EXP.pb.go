// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: COMMANDER_EXP.proto

package protobuf

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

type COMMANDER_EXP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommanderId *uint32 `protobuf:"varint,1,req,name=commander_id,json=commanderId" json:"commander_id,omitempty"`
	Exp         *uint32 `protobuf:"varint,2,req,name=exp" json:"exp,omitempty"`
}

func (x *COMMANDER_EXP) Reset() {
	*x = COMMANDER_EXP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_COMMANDER_EXP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COMMANDER_EXP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COMMANDER_EXP) ProtoMessage() {}

func (x *COMMANDER_EXP) ProtoReflect() protoreflect.Message {
	mi := &file_COMMANDER_EXP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COMMANDER_EXP.ProtoReflect.Descriptor instead.
func (*COMMANDER_EXP) Descriptor() ([]byte, []int) {
	return file_COMMANDER_EXP_proto_rawDescGZIP(), []int{0}
}

func (x *COMMANDER_EXP) GetCommanderId() uint32 {
	if x != nil && x.CommanderId != nil {
		return *x.CommanderId
	}
	return 0
}

func (x *COMMANDER_EXP) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

var File_COMMANDER_EXP_proto protoreflect.FileDescriptor

var file_COMMANDER_EXP_proto_rawDesc = []byte{
	0x0a, 0x13, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x44,
	0x0a, 0x0d, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x03, 0x65, 0x78, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66,
}

var (
	file_COMMANDER_EXP_proto_rawDescOnce sync.Once
	file_COMMANDER_EXP_proto_rawDescData = file_COMMANDER_EXP_proto_rawDesc
)

func file_COMMANDER_EXP_proto_rawDescGZIP() []byte {
	file_COMMANDER_EXP_proto_rawDescOnce.Do(func() {
		file_COMMANDER_EXP_proto_rawDescData = protoimpl.X.CompressGZIP(file_COMMANDER_EXP_proto_rawDescData)
	})
	return file_COMMANDER_EXP_proto_rawDescData
}

var file_COMMANDER_EXP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_COMMANDER_EXP_proto_goTypes = []interface{}{
	(*COMMANDER_EXP)(nil), // 0: belfast.COMMANDER_EXP
}
var file_COMMANDER_EXP_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_COMMANDER_EXP_proto_init() }
func file_COMMANDER_EXP_proto_init() {
	if File_COMMANDER_EXP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_COMMANDER_EXP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*COMMANDER_EXP); i {
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
			RawDescriptor: file_COMMANDER_EXP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_COMMANDER_EXP_proto_goTypes,
		DependencyIndexes: file_COMMANDER_EXP_proto_depIdxs,
		MessageInfos:      file_COMMANDER_EXP_proto_msgTypes,
	}.Build()
	File_COMMANDER_EXP_proto = out.File
	file_COMMANDER_EXP_proto_rawDesc = nil
	file_COMMANDER_EXP_proto_goTypes = nil
	file_COMMANDER_EXP_proto_depIdxs = nil
}
