// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_26001.proto

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

type SC_26001 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *uint32      `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	CellList  []*CELLSINFO `protobuf:"bytes,2,rep,name=cell_list,json=cellList" json:"cell_list,omitempty"`
	ColorList []*COLORINFO `protobuf:"bytes,3,rep,name=color_list,json=colorList" json:"color_list,omitempty"`
	AwardList []*AWARDINFO `protobuf:"bytes,4,rep,name=award_list,json=awardList" json:"award_list,omitempty"`
	StartTime *uint32      `protobuf:"varint,5,req,name=start_time,json=startTime" json:"start_time,omitempty"`
}

func (x *SC_26001) Reset() {
	*x = SC_26001{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_26001_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_26001) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_26001) ProtoMessage() {}

func (x *SC_26001) ProtoReflect() protoreflect.Message {
	mi := &file_SC_26001_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_26001.ProtoReflect.Descriptor instead.
func (*SC_26001) Descriptor() ([]byte, []int) {
	return file_SC_26001_proto_rawDescGZIP(), []int{0}
}

func (x *SC_26001) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SC_26001) GetCellList() []*CELLSINFO {
	if x != nil {
		return x.CellList
	}
	return nil
}

func (x *SC_26001) GetColorList() []*COLORINFO {
	if x != nil {
		return x.ColorList
	}
	return nil
}

func (x *SC_26001) GetAwardList() []*AWARDINFO {
	if x != nil {
		return x.AwardList
	}
	return nil
}

func (x *SC_26001) GetStartTime() uint32 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

var File_SC_26001_proto protoreflect.FileDescriptor

var file_SC_26001_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x32, 0x36, 0x30, 0x30, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0f, 0x43, 0x45, 0x4c, 0x4c, 0x53,
	0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x43, 0x4f, 0x4c, 0x4f,
	0x52, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x41, 0x57, 0x41,
	0x52, 0x44, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a,
	0x08, 0x53, 0x43, 0x5f, 0x32, 0x36, 0x30, 0x30, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x65, 0x6c,
	0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62,
	0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x45, 0x4c, 0x4c, 0x53, 0x49, 0x4e, 0x46, 0x4f,
	0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x49, 0x4e,
	0x46, 0x4f, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x0a, 0x61, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x41, 0x57, 0x41, 0x52,
	0x44, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x09, 0x61, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_26001_proto_rawDescOnce sync.Once
	file_SC_26001_proto_rawDescData = file_SC_26001_proto_rawDesc
)

func file_SC_26001_proto_rawDescGZIP() []byte {
	file_SC_26001_proto_rawDescOnce.Do(func() {
		file_SC_26001_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_26001_proto_rawDescData)
	})
	return file_SC_26001_proto_rawDescData
}

var file_SC_26001_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_26001_proto_goTypes = []interface{}{
	(*SC_26001)(nil),  // 0: belfast.SC_26001
	(*CELLSINFO)(nil), // 1: belfast.CELLSINFO
	(*COLORINFO)(nil), // 2: belfast.COLORINFO
	(*AWARDINFO)(nil), // 3: belfast.AWARDINFO
}
var file_SC_26001_proto_depIdxs = []int32{
	1, // 0: belfast.SC_26001.cell_list:type_name -> belfast.CELLSINFO
	2, // 1: belfast.SC_26001.color_list:type_name -> belfast.COLORINFO
	3, // 2: belfast.SC_26001.award_list:type_name -> belfast.AWARDINFO
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SC_26001_proto_init() }
func file_SC_26001_proto_init() {
	if File_SC_26001_proto != nil {
		return
	}
	file_CELLSINFO_proto_init()
	file_COLORINFO_proto_init()
	file_AWARDINFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_26001_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_26001); i {
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
			RawDescriptor: file_SC_26001_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_26001_proto_goTypes,
		DependencyIndexes: file_SC_26001_proto_depIdxs,
		MessageInfos:      file_SC_26001_proto_msgTypes,
	}.Build()
	File_SC_26001_proto = out.File
	file_SC_26001_proto_rawDesc = nil
	file_SC_26001_proto_goTypes = nil
	file_SC_26001_proto_depIdxs = nil
}
