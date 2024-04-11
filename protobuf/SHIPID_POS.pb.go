// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SHIPID_POS.proto

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

type SHIPID_POS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos    *uint32 `protobuf:"varint,1,req,name=pos" json:"pos,omitempty"`
	ShipId *uint32 `protobuf:"varint,2,req,name=shipId" json:"shipId,omitempty"`
}

func (x *SHIPID_POS) Reset() {
	*x = SHIPID_POS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SHIPID_POS_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SHIPID_POS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SHIPID_POS) ProtoMessage() {}

func (x *SHIPID_POS) ProtoReflect() protoreflect.Message {
	mi := &file_SHIPID_POS_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SHIPID_POS.ProtoReflect.Descriptor instead.
func (*SHIPID_POS) Descriptor() ([]byte, []int) {
	return file_SHIPID_POS_proto_rawDescGZIP(), []int{0}
}

func (x *SHIPID_POS) GetPos() uint32 {
	if x != nil && x.Pos != nil {
		return *x.Pos
	}
	return 0
}

func (x *SHIPID_POS) GetShipId() uint32 {
	if x != nil && x.ShipId != nil {
		return *x.ShipId
	}
	return 0
}

var File_SHIPID_POS_proto protoreflect.FileDescriptor

var file_SHIPID_POS_proto_rawDesc = []byte{
	0x0a, 0x10, 0x53, 0x48, 0x49, 0x50, 0x49, 0x44, 0x5f, 0x50, 0x4f, 0x53, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x0a, 0x53,
	0x48, 0x49, 0x50, 0x49, 0x44, 0x5f, 0x50, 0x4f, 0x53, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x69, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x69,
	0x70, 0x49, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66,
}

var (
	file_SHIPID_POS_proto_rawDescOnce sync.Once
	file_SHIPID_POS_proto_rawDescData = file_SHIPID_POS_proto_rawDesc
)

func file_SHIPID_POS_proto_rawDescGZIP() []byte {
	file_SHIPID_POS_proto_rawDescOnce.Do(func() {
		file_SHIPID_POS_proto_rawDescData = protoimpl.X.CompressGZIP(file_SHIPID_POS_proto_rawDescData)
	})
	return file_SHIPID_POS_proto_rawDescData
}

var file_SHIPID_POS_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SHIPID_POS_proto_goTypes = []interface{}{
	(*SHIPID_POS)(nil), // 0: belfast.SHIPID_POS
}
var file_SHIPID_POS_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SHIPID_POS_proto_init() }
func file_SHIPID_POS_proto_init() {
	if File_SHIPID_POS_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SHIPID_POS_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SHIPID_POS); i {
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
			RawDescriptor: file_SHIPID_POS_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SHIPID_POS_proto_goTypes,
		DependencyIndexes: file_SHIPID_POS_proto_depIdxs,
		MessageInfos:      file_SHIPID_POS_proto_msgTypes,
	}.Build()
	File_SHIPID_POS_proto = out.File
	file_SHIPID_POS_proto_rawDesc = nil
	file_SHIPID_POS_proto_goTypes = nil
	file_SHIPID_POS_proto_depIdxs = nil
}
