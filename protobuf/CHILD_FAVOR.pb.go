// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: CHILD_FAVOR.proto

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

type CHILD_FAVOR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lv  *uint32 `protobuf:"varint,1,req,name=lv" json:"lv,omitempty"`
	Exp *uint32 `protobuf:"varint,2,req,name=exp" json:"exp,omitempty"`
}

func (x *CHILD_FAVOR) Reset() {
	*x = CHILD_FAVOR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CHILD_FAVOR_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CHILD_FAVOR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CHILD_FAVOR) ProtoMessage() {}

func (x *CHILD_FAVOR) ProtoReflect() protoreflect.Message {
	mi := &file_CHILD_FAVOR_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CHILD_FAVOR.ProtoReflect.Descriptor instead.
func (*CHILD_FAVOR) Descriptor() ([]byte, []int) {
	return file_CHILD_FAVOR_proto_rawDescGZIP(), []int{0}
}

func (x *CHILD_FAVOR) GetLv() uint32 {
	if x != nil && x.Lv != nil {
		return *x.Lv
	}
	return 0
}

func (x *CHILD_FAVOR) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

var File_CHILD_FAVOR_proto protoreflect.FileDescriptor

var file_CHILD_FAVOR_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x46, 0x41, 0x56, 0x4f, 0x52, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x0b,
	0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x46, 0x41, 0x56, 0x4f, 0x52, 0x12, 0x0e, 0x0a, 0x02, 0x6c,
	0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x6c, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x78, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CHILD_FAVOR_proto_rawDescOnce sync.Once
	file_CHILD_FAVOR_proto_rawDescData = file_CHILD_FAVOR_proto_rawDesc
)

func file_CHILD_FAVOR_proto_rawDescGZIP() []byte {
	file_CHILD_FAVOR_proto_rawDescOnce.Do(func() {
		file_CHILD_FAVOR_proto_rawDescData = protoimpl.X.CompressGZIP(file_CHILD_FAVOR_proto_rawDescData)
	})
	return file_CHILD_FAVOR_proto_rawDescData
}

var file_CHILD_FAVOR_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CHILD_FAVOR_proto_goTypes = []interface{}{
	(*CHILD_FAVOR)(nil), // 0: belfast.CHILD_FAVOR
}
var file_CHILD_FAVOR_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CHILD_FAVOR_proto_init() }
func file_CHILD_FAVOR_proto_init() {
	if File_CHILD_FAVOR_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CHILD_FAVOR_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CHILD_FAVOR); i {
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
			RawDescriptor: file_CHILD_FAVOR_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CHILD_FAVOR_proto_goTypes,
		DependencyIndexes: file_CHILD_FAVOR_proto_depIdxs,
		MessageInfos:      file_CHILD_FAVOR_proto_msgTypes,
	}.Build()
	File_CHILD_FAVOR_proto = out.File
	file_CHILD_FAVOR_proto_rawDesc = nil
	file_CHILD_FAVOR_proto_goTypes = nil
	file_CHILD_FAVOR_proto_depIdxs = nil
}
