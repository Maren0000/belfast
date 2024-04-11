// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_63314.proto

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

type SC_63314 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipId    *uint32        `protobuf:"varint,1,req,name=ship_id,json=shipId" json:"ship_id,omitempty"`
	DoubleExp *uint32        `protobuf:"varint,2,req,name=double_exp,json=doubleExp" json:"double_exp,omitempty"`
	Exp       *uint32        `protobuf:"varint,3,req,name=exp" json:"exp,omitempty"`
	SkillId   *uint32        `protobuf:"varint,4,req,name=skill_id,json=skillId" json:"skill_id,omitempty"`
	Tasks     []*FINISH_TASK `protobuf:"bytes,5,rep,name=tasks" json:"tasks,omitempty"`
	SwitchCnt *uint32        `protobuf:"varint,6,req,name=switch_cnt,json=switchCnt" json:"switch_cnt,omitempty"`
	SkillExp  []*SKILL_EXP   `protobuf:"bytes,7,rep,name=skill_exp,json=skillExp" json:"skill_exp,omitempty"`
}

func (x *SC_63314) Reset() {
	*x = SC_63314{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_63314_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_63314) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_63314) ProtoMessage() {}

func (x *SC_63314) ProtoReflect() protoreflect.Message {
	mi := &file_SC_63314_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_63314.ProtoReflect.Descriptor instead.
func (*SC_63314) Descriptor() ([]byte, []int) {
	return file_SC_63314_proto_rawDescGZIP(), []int{0}
}

func (x *SC_63314) GetShipId() uint32 {
	if x != nil && x.ShipId != nil {
		return *x.ShipId
	}
	return 0
}

func (x *SC_63314) GetDoubleExp() uint32 {
	if x != nil && x.DoubleExp != nil {
		return *x.DoubleExp
	}
	return 0
}

func (x *SC_63314) GetExp() uint32 {
	if x != nil && x.Exp != nil {
		return *x.Exp
	}
	return 0
}

func (x *SC_63314) GetSkillId() uint32 {
	if x != nil && x.SkillId != nil {
		return *x.SkillId
	}
	return 0
}

func (x *SC_63314) GetTasks() []*FINISH_TASK {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *SC_63314) GetSwitchCnt() uint32 {
	if x != nil && x.SwitchCnt != nil {
		return *x.SwitchCnt
	}
	return 0
}

func (x *SC_63314) GetSkillExp() []*SKILL_EXP {
	if x != nil {
		return x.SkillExp
	}
	return nil
}

var File_SC_63314_proto protoreflect.FileDescriptor

var file_SC_63314_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x36, 0x33, 0x33, 0x31, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x11, 0x46, 0x49, 0x4e, 0x49, 0x53,
	0x48, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x4b,
	0x49, 0x4c, 0x4c, 0x5f, 0x45, 0x58, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01,
	0x0a, 0x08, 0x53, 0x43, 0x5f, 0x36, 0x33, 0x33, 0x31, 0x34, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x69,
	0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x78,
	0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45,
	0x78, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x03, 0x65, 0x78, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x5f,
	0x54, 0x41, 0x53, 0x4b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x09, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x43, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x45, 0x58,
	0x50, 0x52, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x78, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_63314_proto_rawDescOnce sync.Once
	file_SC_63314_proto_rawDescData = file_SC_63314_proto_rawDesc
)

func file_SC_63314_proto_rawDescGZIP() []byte {
	file_SC_63314_proto_rawDescOnce.Do(func() {
		file_SC_63314_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_63314_proto_rawDescData)
	})
	return file_SC_63314_proto_rawDescData
}

var file_SC_63314_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_63314_proto_goTypes = []interface{}{
	(*SC_63314)(nil),    // 0: belfast.SC_63314
	(*FINISH_TASK)(nil), // 1: belfast.FINISH_TASK
	(*SKILL_EXP)(nil),   // 2: belfast.SKILL_EXP
}
var file_SC_63314_proto_depIdxs = []int32{
	1, // 0: belfast.SC_63314.tasks:type_name -> belfast.FINISH_TASK
	2, // 1: belfast.SC_63314.skill_exp:type_name -> belfast.SKILL_EXP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SC_63314_proto_init() }
func file_SC_63314_proto_init() {
	if File_SC_63314_proto != nil {
		return
	}
	file_FINISH_TASK_proto_init()
	file_SKILL_EXP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_63314_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_63314); i {
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
			RawDescriptor: file_SC_63314_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_63314_proto_goTypes,
		DependencyIndexes: file_SC_63314_proto_depIdxs,
		MessageInfos:      file_SC_63314_proto_msgTypes,
	}.Build()
	File_SC_63314_proto = out.File
	file_SC_63314_proto_rawDesc = nil
	file_SC_63314_proto_goTypes = nil
	file_SC_63314_proto_depIdxs = nil
}
