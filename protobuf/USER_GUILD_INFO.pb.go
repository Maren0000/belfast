// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: USER_GUILD_INFO.proto

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

type USER_GUILD_INFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DonateCount    *uint32  `protobuf:"varint,1,req,name=donate_count,json=donateCount" json:"donate_count,omitempty"`
	DonateTasks    []uint32 `protobuf:"varint,2,rep,name=donate_tasks,json=donateTasks" json:"donate_tasks,omitempty"`
	BenefitTime    *uint32  `protobuf:"varint,3,req,name=benefit_time,json=benefitTime" json:"benefit_time,omitempty"`
	TechId         []uint32 `protobuf:"varint,4,rep,name=tech_id,json=techId" json:"tech_id,omitempty"`
	WeeklyTaskFlag *uint32  `protobuf:"varint,5,req,name=weekly_task_flag,json=weeklyTaskFlag" json:"weekly_task_flag,omitempty"`
	ExtraDonate    *uint32  `protobuf:"varint,6,req,name=extra_donate,json=extraDonate" json:"extra_donate,omitempty"`
	ExtraOperation *uint32  `protobuf:"varint,7,req,name=extra_operation,json=extraOperation" json:"extra_operation,omitempty"`
}

func (x *USER_GUILD_INFO) Reset() {
	*x = USER_GUILD_INFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_USER_GUILD_INFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *USER_GUILD_INFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*USER_GUILD_INFO) ProtoMessage() {}

func (x *USER_GUILD_INFO) ProtoReflect() protoreflect.Message {
	mi := &file_USER_GUILD_INFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use USER_GUILD_INFO.ProtoReflect.Descriptor instead.
func (*USER_GUILD_INFO) Descriptor() ([]byte, []int) {
	return file_USER_GUILD_INFO_proto_rawDescGZIP(), []int{0}
}

func (x *USER_GUILD_INFO) GetDonateCount() uint32 {
	if x != nil && x.DonateCount != nil {
		return *x.DonateCount
	}
	return 0
}

func (x *USER_GUILD_INFO) GetDonateTasks() []uint32 {
	if x != nil {
		return x.DonateTasks
	}
	return nil
}

func (x *USER_GUILD_INFO) GetBenefitTime() uint32 {
	if x != nil && x.BenefitTime != nil {
		return *x.BenefitTime
	}
	return 0
}

func (x *USER_GUILD_INFO) GetTechId() []uint32 {
	if x != nil {
		return x.TechId
	}
	return nil
}

func (x *USER_GUILD_INFO) GetWeeklyTaskFlag() uint32 {
	if x != nil && x.WeeklyTaskFlag != nil {
		return *x.WeeklyTaskFlag
	}
	return 0
}

func (x *USER_GUILD_INFO) GetExtraDonate() uint32 {
	if x != nil && x.ExtraDonate != nil {
		return *x.ExtraDonate
	}
	return 0
}

func (x *USER_GUILD_INFO) GetExtraOperation() uint32 {
	if x != nil && x.ExtraOperation != nil {
		return *x.ExtraOperation
	}
	return 0
}

var File_USER_GUILD_INFO_proto protoreflect.FileDescriptor

var file_USER_GUILD_INFO_proto_rawDesc = []byte{
	0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x49, 0x4e, 0x46,
	0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74,
	0x22, 0x89, 0x02, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x6f, 0x6e, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x64,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0b, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06,
	0x74, 0x65, 0x63, 0x68, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0e, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x6c, 0x61, 0x67,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x6f, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0e, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_USER_GUILD_INFO_proto_rawDescOnce sync.Once
	file_USER_GUILD_INFO_proto_rawDescData = file_USER_GUILD_INFO_proto_rawDesc
)

func file_USER_GUILD_INFO_proto_rawDescGZIP() []byte {
	file_USER_GUILD_INFO_proto_rawDescOnce.Do(func() {
		file_USER_GUILD_INFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_USER_GUILD_INFO_proto_rawDescData)
	})
	return file_USER_GUILD_INFO_proto_rawDescData
}

var file_USER_GUILD_INFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_USER_GUILD_INFO_proto_goTypes = []interface{}{
	(*USER_GUILD_INFO)(nil), // 0: belfast.USER_GUILD_INFO
}
var file_USER_GUILD_INFO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_USER_GUILD_INFO_proto_init() }
func file_USER_GUILD_INFO_proto_init() {
	if File_USER_GUILD_INFO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_USER_GUILD_INFO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*USER_GUILD_INFO); i {
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
			RawDescriptor: file_USER_GUILD_INFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_USER_GUILD_INFO_proto_goTypes,
		DependencyIndexes: file_USER_GUILD_INFO_proto_depIdxs,
		MessageInfos:      file_USER_GUILD_INFO_proto_msgTypes,
	}.Build()
	File_USER_GUILD_INFO_proto = out.File
	file_USER_GUILD_INFO_proto_rawDesc = nil
	file_USER_GUILD_INFO_proto_goTypes = nil
	file_USER_GUILD_INFO_proto_depIdxs = nil
}
