// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: CS_26105.proto

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

type CS_26105 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Combine []*CS_26103 `protobuf:"bytes,1,rep,name=combine" json:"combine,omitempty"`
}

func (x *CS_26105) Reset() {
	*x = CS_26105{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_26105_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_26105) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_26105) ProtoMessage() {}

func (x *CS_26105) ProtoReflect() protoreflect.Message {
	mi := &file_CS_26105_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_26105.ProtoReflect.Descriptor instead.
func (*CS_26105) Descriptor() ([]byte, []int) {
	return file_CS_26105_proto_rawDescGZIP(), []int{0}
}

func (x *CS_26105) GetCombine() []*CS_26103 {
	if x != nil {
		return x.Combine
	}
	return nil
}

var File_CS_26105_proto protoreflect.FileDescriptor

var file_CS_26105_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x32, 0x36, 0x31, 0x30, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0e, 0x43, 0x53, 0x5f, 0x32, 0x36,
	0x31, 0x30, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x08, 0x43, 0x53, 0x5f,
	0x32, 0x36, 0x31, 0x30, 0x35, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x2e, 0x43, 0x53, 0x5f, 0x32, 0x36, 0x31, 0x30, 0x33, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CS_26105_proto_rawDescOnce sync.Once
	file_CS_26105_proto_rawDescData = file_CS_26105_proto_rawDesc
)

func file_CS_26105_proto_rawDescGZIP() []byte {
	file_CS_26105_proto_rawDescOnce.Do(func() {
		file_CS_26105_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_26105_proto_rawDescData)
	})
	return file_CS_26105_proto_rawDescData
}

var file_CS_26105_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_26105_proto_goTypes = []interface{}{
	(*CS_26105)(nil), // 0: belfast.CS_26105
	(*CS_26103)(nil), // 1: belfast.CS_26103
}
var file_CS_26105_proto_depIdxs = []int32{
	1, // 0: belfast.CS_26105.combine:type_name -> belfast.CS_26103
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CS_26105_proto_init() }
func file_CS_26105_proto_init() {
	if File_CS_26105_proto != nil {
		return
	}
	file_CS_26103_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CS_26105_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_26105); i {
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
			RawDescriptor: file_CS_26105_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_26105_proto_goTypes,
		DependencyIndexes: file_CS_26105_proto_depIdxs,
		MessageInfos:      file_CS_26105_proto_msgTypes,
	}.Build()
	File_CS_26105_proto = out.File
	file_CS_26105_proto_rawDesc = nil
	file_CS_26105_proto_goTypes = nil
	file_CS_26105_proto_depIdxs = nil
}
