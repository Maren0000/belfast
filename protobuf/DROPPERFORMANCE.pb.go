// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: DROPPERFORMANCE.proto

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

type DROPPERFORMANCE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnemyId     *uint32 `protobuf:"varint,1,req,name=enemy_id,json=enemyId" json:"enemy_id,omitempty"`
	ResourceNum *uint32 `protobuf:"varint,2,req,name=resource_num,json=resourceNum" json:"resource_num,omitempty"`
	OtherNum    *uint32 `protobuf:"varint,3,req,name=other_num,json=otherNum" json:"other_num,omitempty"`
}

func (x *DROPPERFORMANCE) Reset() {
	*x = DROPPERFORMANCE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DROPPERFORMANCE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DROPPERFORMANCE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DROPPERFORMANCE) ProtoMessage() {}

func (x *DROPPERFORMANCE) ProtoReflect() protoreflect.Message {
	mi := &file_DROPPERFORMANCE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DROPPERFORMANCE.ProtoReflect.Descriptor instead.
func (*DROPPERFORMANCE) Descriptor() ([]byte, []int) {
	return file_DROPPERFORMANCE_proto_rawDescGZIP(), []int{0}
}

func (x *DROPPERFORMANCE) GetEnemyId() uint32 {
	if x != nil && x.EnemyId != nil {
		return *x.EnemyId
	}
	return 0
}

func (x *DROPPERFORMANCE) GetResourceNum() uint32 {
	if x != nil && x.ResourceNum != nil {
		return *x.ResourceNum
	}
	return 0
}

func (x *DROPPERFORMANCE) GetOtherNum() uint32 {
	if x != nil && x.OtherNum != nil {
		return *x.OtherNum
	}
	return 0
}

var File_DROPPERFORMANCE_proto protoreflect.FileDescriptor

var file_DROPPERFORMANCE_proto_rawDesc = []byte{
	0x0a, 0x15, 0x44, 0x52, 0x4f, 0x50, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43,
	0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x22, 0x6c, 0x0a, 0x0f, 0x44, 0x52, 0x4f, 0x50, 0x50, 0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x4e, 0x43, 0x45, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x65, 0x6d, 0x79, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x75,
	0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_DROPPERFORMANCE_proto_rawDescOnce sync.Once
	file_DROPPERFORMANCE_proto_rawDescData = file_DROPPERFORMANCE_proto_rawDesc
)

func file_DROPPERFORMANCE_proto_rawDescGZIP() []byte {
	file_DROPPERFORMANCE_proto_rawDescOnce.Do(func() {
		file_DROPPERFORMANCE_proto_rawDescData = protoimpl.X.CompressGZIP(file_DROPPERFORMANCE_proto_rawDescData)
	})
	return file_DROPPERFORMANCE_proto_rawDescData
}

var file_DROPPERFORMANCE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DROPPERFORMANCE_proto_goTypes = []interface{}{
	(*DROPPERFORMANCE)(nil), // 0: belfast.DROPPERFORMANCE
}
var file_DROPPERFORMANCE_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DROPPERFORMANCE_proto_init() }
func file_DROPPERFORMANCE_proto_init() {
	if File_DROPPERFORMANCE_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DROPPERFORMANCE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DROPPERFORMANCE); i {
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
			RawDescriptor: file_DROPPERFORMANCE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DROPPERFORMANCE_proto_goTypes,
		DependencyIndexes: file_DROPPERFORMANCE_proto_depIdxs,
		MessageInfos:      file_DROPPERFORMANCE_proto_msgTypes,
	}.Build()
	File_DROPPERFORMANCE_proto = out.File
	file_DROPPERFORMANCE_proto_rawDesc = nil
	file_DROPPERFORMANCE_proto_goTypes = nil
	file_DROPPERFORMANCE_proto_depIdxs = nil
}
