// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: CS_11204.proto

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

type CS_11204 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId *uint32      `protobuf:"varint,1,req,name=activity_id,json=activityId" json:"activity_id,omitempty"`
	GroupList  []*GROUPINFO `protobuf:"bytes,2,rep,name=group_list,json=groupList" json:"group_list,omitempty"`
}

func (x *CS_11204) Reset() {
	*x = CS_11204{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_11204_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_11204) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_11204) ProtoMessage() {}

func (x *CS_11204) ProtoReflect() protoreflect.Message {
	mi := &file_CS_11204_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_11204.ProtoReflect.Descriptor instead.
func (*CS_11204) Descriptor() ([]byte, []int) {
	return file_CS_11204_proto_rawDescGZIP(), []int{0}
}

func (x *CS_11204) GetActivityId() uint32 {
	if x != nil && x.ActivityId != nil {
		return *x.ActivityId
	}
	return 0
}

func (x *CS_11204) GetGroupList() []*GROUPINFO {
	if x != nil {
		return x.GroupList
	}
	return nil
}

var File_CS_11204_proto protoreflect.FileDescriptor

var file_CS_11204_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x31, 0x31, 0x32, 0x30, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x08, 0x43, 0x53,
	0x5f, 0x31, 0x31, 0x32, 0x30, 0x34, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x65,
	0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x49, 0x4e, 0x46, 0x4f, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CS_11204_proto_rawDescOnce sync.Once
	file_CS_11204_proto_rawDescData = file_CS_11204_proto_rawDesc
)

func file_CS_11204_proto_rawDescGZIP() []byte {
	file_CS_11204_proto_rawDescOnce.Do(func() {
		file_CS_11204_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_11204_proto_rawDescData)
	})
	return file_CS_11204_proto_rawDescData
}

var file_CS_11204_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_11204_proto_goTypes = []interface{}{
	(*CS_11204)(nil),  // 0: belfast.CS_11204
	(*GROUPINFO)(nil), // 1: belfast.GROUPINFO
}
var file_CS_11204_proto_depIdxs = []int32{
	1, // 0: belfast.CS_11204.group_list:type_name -> belfast.GROUPINFO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CS_11204_proto_init() }
func file_CS_11204_proto_init() {
	if File_CS_11204_proto != nil {
		return
	}
	file_GROUPINFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CS_11204_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_11204); i {
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
			RawDescriptor: file_CS_11204_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_11204_proto_goTypes,
		DependencyIndexes: file_CS_11204_proto_depIdxs,
		MessageInfos:      file_CS_11204_proto_msgTypes,
	}.Build()
	File_CS_11204_proto = out.File
	file_CS_11204_proto_rawDesc = nil
	file_CS_11204_proto_goTypes = nil
	file_CS_11204_proto_depIdxs = nil
}
