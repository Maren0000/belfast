// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_19009.proto

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

type SC_19009 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp         *uint32 `protobuf:"varint,1,req,name=exp" json:"exp,omitempty"`
	FoodConsume *uint32 `protobuf:"varint,2,req,name=food_consume,json=foodConsume" json:"food_consume,omitempty"`
}

func (x *SC_19009) Reset() {
	*x = SC_19009{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_19009_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_19009) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_19009) ProtoMessage() {}

func (x *SC_19009) ProtoReflect() protoreflect.Message {
	mi := &file_SC_19009_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_19009.ProtoReflect.Descriptor instead.
func (*SC_19009) Descriptor() ([]byte, []int) {
	return file_SC_19009_proto_rawDescGZIP(), []int{0}
}

func (x *SC_19009) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *SC_19009) GetFoodConsume() uint32 {
	if x != nil && x.FoodConsume != nil {
		return *x.FoodConsume
	}
	return 0
}

var File_SC_19009_proto protoreflect.FileDescriptor

var file_SC_19009_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x31, 0x39, 0x30, 0x30, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x08, 0x53, 0x43, 0x5f,
	0x31, 0x39, 0x30, 0x30, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6f, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x66,
	0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_19009_proto_rawDescOnce sync.Once
	file_SC_19009_proto_rawDescData = file_SC_19009_proto_rawDesc
)

func file_SC_19009_proto_rawDescGZIP() []byte {
	file_SC_19009_proto_rawDescOnce.Do(func() {
		file_SC_19009_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_19009_proto_rawDescData)
	})
	return file_SC_19009_proto_rawDescData
}

var file_SC_19009_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_19009_proto_goTypes = []interface{}{
	(*SC_19009)(nil), // 0: belfast.SC_19009
}
var file_SC_19009_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SC_19009_proto_init() }
func file_SC_19009_proto_init() {
	if File_SC_19009_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SC_19009_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_19009); i {
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
			RawDescriptor: file_SC_19009_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_19009_proto_goTypes,
		DependencyIndexes: file_SC_19009_proto_depIdxs,
		MessageInfos:      file_SC_19009_proto_msgTypes,
	}.Build()
	File_SC_19009_proto = out.File
	file_SC_19009_proto_rawDesc = nil
	file_SC_19009_proto_goTypes = nil
	file_SC_19009_proto_depIdxs = nil
}
