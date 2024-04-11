// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_17004.proto

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

type SC_17004 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipInfo *SHIP_STATISTICS_INFO `protobuf:"bytes,1,req,name=ship_info,json=shipInfo" json:"ship_info,omitempty"`
}

func (x *SC_17004) Reset() {
	*x = SC_17004{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_17004_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_17004) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_17004) ProtoMessage() {}

func (x *SC_17004) ProtoReflect() protoreflect.Message {
	mi := &file_SC_17004_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_17004.ProtoReflect.Descriptor instead.
func (*SC_17004) Descriptor() ([]byte, []int) {
	return file_SC_17004_proto_rawDescGZIP(), []int{0}
}

func (x *SC_17004) GetShipInfo() *SHIP_STATISTICS_INFO {
	if x != nil {
		return x.ShipInfo
	}
	return nil
}

var File_SC_17004_proto protoreflect.FileDescriptor

var file_SC_17004_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x31, 0x37, 0x30, 0x30, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x1a, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x49, 0x53, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x31, 0x37, 0x30, 0x30,
	0x34, 0x12, 0x3a, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x53,
	0x48, 0x49, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x49, 0x53, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_17004_proto_rawDescOnce sync.Once
	file_SC_17004_proto_rawDescData = file_SC_17004_proto_rawDesc
)

func file_SC_17004_proto_rawDescGZIP() []byte {
	file_SC_17004_proto_rawDescOnce.Do(func() {
		file_SC_17004_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_17004_proto_rawDescData)
	})
	return file_SC_17004_proto_rawDescData
}

var file_SC_17004_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_17004_proto_goTypes = []interface{}{
	(*SC_17004)(nil),             // 0: belfast.SC_17004
	(*SHIP_STATISTICS_INFO)(nil), // 1: belfast.SHIP_STATISTICS_INFO
}
var file_SC_17004_proto_depIdxs = []int32{
	1, // 0: belfast.SC_17004.ship_info:type_name -> belfast.SHIP_STATISTICS_INFO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SC_17004_proto_init() }
func file_SC_17004_proto_init() {
	if File_SC_17004_proto != nil {
		return
	}
	file_SHIP_STATISTICS_INFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_17004_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_17004); i {
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
			RawDescriptor: file_SC_17004_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_17004_proto_goTypes,
		DependencyIndexes: file_SC_17004_proto_depIdxs,
		MessageInfos:      file_SC_17004_proto_msgTypes,
	}.Build()
	File_SC_17004_proto = out.File
	file_SC_17004_proto_rawDesc = nil
	file_SC_17004_proto_goTypes = nil
	file_SC_17004_proto_depIdxs = nil
}
