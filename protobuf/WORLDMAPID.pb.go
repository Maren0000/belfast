// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: WORLDMAPID.proto

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

type WORLDMAPID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RandomId   *uint32 `protobuf:"varint,1,req,name=random_id,json=randomId" json:"random_id,omitempty"`
	TemplateId *uint32 `protobuf:"varint,2,req,name=template_id,json=templateId" json:"template_id,omitempty"`
}

func (x *WORLDMAPID) Reset() {
	*x = WORLDMAPID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WORLDMAPID_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WORLDMAPID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WORLDMAPID) ProtoMessage() {}

func (x *WORLDMAPID) ProtoReflect() protoreflect.Message {
	mi := &file_WORLDMAPID_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WORLDMAPID.ProtoReflect.Descriptor instead.
func (*WORLDMAPID) Descriptor() ([]byte, []int) {
	return file_WORLDMAPID_proto_rawDescGZIP(), []int{0}
}

func (x *WORLDMAPID) GetRandomId() uint32 {
	if x != nil && x.RandomId != nil {
		return *x.RandomId
	}
	return 0
}

func (x *WORLDMAPID) GetTemplateId() uint32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

var File_WORLDMAPID_proto protoreflect.FileDescriptor

var file_WORLDMAPID_proto_rawDesc = []byte{
	0x0a, 0x10, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x4d, 0x41, 0x50, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x0a, 0x57,
	0x4f, 0x52, 0x4c, 0x44, 0x4d, 0x41, 0x50, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_WORLDMAPID_proto_rawDescOnce sync.Once
	file_WORLDMAPID_proto_rawDescData = file_WORLDMAPID_proto_rawDesc
)

func file_WORLDMAPID_proto_rawDescGZIP() []byte {
	file_WORLDMAPID_proto_rawDescOnce.Do(func() {
		file_WORLDMAPID_proto_rawDescData = protoimpl.X.CompressGZIP(file_WORLDMAPID_proto_rawDescData)
	})
	return file_WORLDMAPID_proto_rawDescData
}

var file_WORLDMAPID_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WORLDMAPID_proto_goTypes = []interface{}{
	(*WORLDMAPID)(nil), // 0: belfast.WORLDMAPID
}
var file_WORLDMAPID_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WORLDMAPID_proto_init() }
func file_WORLDMAPID_proto_init() {
	if File_WORLDMAPID_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WORLDMAPID_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WORLDMAPID); i {
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
			RawDescriptor: file_WORLDMAPID_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WORLDMAPID_proto_goTypes,
		DependencyIndexes: file_WORLDMAPID_proto_depIdxs,
		MessageInfos:      file_WORLDMAPID_proto_msgTypes,
	}.Build()
	File_WORLDMAPID_proto = out.File
	file_WORLDMAPID_proto_rawDesc = nil
	file_WORLDMAPID_proto_goTypes = nil
	file_WORLDMAPID_proto_depIdxs = nil
}
