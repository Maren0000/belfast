// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: DETAIL_INFO.proto

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

type DETAIL_INFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 *uint32      `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name               *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Title              *uint32      `protobuf:"varint,3,req,name=title" json:"title,omitempty"`
	Lv                 *uint32      `protobuf:"varint,4,req,name=lv" json:"lv,omitempty"`
	ShipCount          *uint32      `protobuf:"varint,5,req,name=ship_count,json=shipCount" json:"ship_count,omitempty"`
	CollectionCount    *uint32      `protobuf:"varint,6,req,name=collection_count,json=collectionCount" json:"collection_count,omitempty"`
	PvpAttackCount     *uint32      `protobuf:"varint,7,req,name=pvp_attack_count,json=pvpAttackCount" json:"pvp_attack_count,omitempty"`
	PvpWinCount        *uint32      `protobuf:"varint,8,req,name=pvp_win_count,json=pvpWinCount" json:"pvp_win_count,omitempty"`
	CollectAttackCount *uint32      `protobuf:"varint,9,req,name=collect_attack_count,json=collectAttackCount" json:"collect_attack_count,omitempty"`
	AttackCount        *uint32      `protobuf:"varint,10,req,name=attack_count,json=attackCount" json:"attack_count,omitempty"`
	WinCount           *uint32      `protobuf:"varint,11,req,name=win_count,json=winCount" json:"win_count,omitempty"`
	Adv                *string      `protobuf:"bytes,12,req,name=adv" json:"adv,omitempty"`
	Online             *uint32      `protobuf:"varint,13,req,name=online" json:"online,omitempty"`
	PreOnlineTime      *uint32      `protobuf:"varint,14,req,name=pre_online_time,json=preOnlineTime" json:"pre_online_time,omitempty"`
	Score              *uint32      `protobuf:"varint,15,req,name=score" json:"score,omitempty"`
	MedalId            []uint32     `protobuf:"varint,16,rep,name=medal_id,json=medalId" json:"medal_id,omitempty"`
	Display            *DISPLAYINFO `protobuf:"bytes,17,opt,name=display" json:"display,omitempty"`
}

func (x *DETAIL_INFO) Reset() {
	*x = DETAIL_INFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DETAIL_INFO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DETAIL_INFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DETAIL_INFO) ProtoMessage() {}

func (x *DETAIL_INFO) ProtoReflect() protoreflect.Message {
	mi := &file_DETAIL_INFO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DETAIL_INFO.ProtoReflect.Descriptor instead.
func (*DETAIL_INFO) Descriptor() ([]byte, []int) {
	return file_DETAIL_INFO_proto_rawDescGZIP(), []int{0}
}

func (x *DETAIL_INFO) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *DETAIL_INFO) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *DETAIL_INFO) GetTitle() uint32 {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return 0
}

func (x *DETAIL_INFO) GetLv() uint32 {
	if x != nil && x.Lv != nil {
		return *x.Lv
	}
	return 0
}

func (x *DETAIL_INFO) GetShipCount() uint32 {
	if x != nil && x.ShipCount != nil {
		return *x.ShipCount
	}
	return 0
}

func (x *DETAIL_INFO) GetCollectionCount() uint32 {
	if x != nil && x.CollectionCount != nil {
		return *x.CollectionCount
	}
	return 0
}

func (x *DETAIL_INFO) GetPvpAttackCount() uint32 {
	if x != nil && x.PvpAttackCount != nil {
		return *x.PvpAttackCount
	}
	return 0
}

func (x *DETAIL_INFO) GetPvpWinCount() uint32 {
	if x != nil && x.PvpWinCount != nil {
		return *x.PvpWinCount
	}
	return 0
}

func (x *DETAIL_INFO) GetCollectAttackCount() uint32 {
	if x != nil && x.CollectAttackCount != nil {
		return *x.CollectAttackCount
	}
	return 0
}

func (x *DETAIL_INFO) GetAttackCount() uint32 {
	if x != nil && x.AttackCount != nil {
		return *x.AttackCount
	}
	return 0
}

func (x *DETAIL_INFO) GetWinCount() uint32 {
	if x != nil && x.WinCount != nil {
		return *x.WinCount
	}
	return 0
}

func (x *DETAIL_INFO) GetAdv() string {
	if x != nil && x.Adv != nil {
		return *x.Adv
	}
	return ""
}

func (x *DETAIL_INFO) GetOnline() uint32 {
	if x != nil && x.Online != nil {
		return *x.Online
	}
	return 0
}

func (x *DETAIL_INFO) GetPreOnlineTime() uint32 {
	if x != nil && x.PreOnlineTime != nil {
		return *x.PreOnlineTime
	}
	return 0
}

func (x *DETAIL_INFO) GetScore() uint32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

func (x *DETAIL_INFO) GetMedalId() []uint32 {
	if x != nil {
		return x.MedalId
	}
	return nil
}

func (x *DETAIL_INFO) GetDisplay() *DISPLAYINFO {
	if x != nil {
		return x.Display
	}
	return nil
}

var File_DETAIL_INFO_proto protoreflect.FileDescriptor

var file_DETAIL_INFO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x11, 0x44, 0x49,
	0x53, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x04, 0x0a, 0x0b, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x76, 0x18,
	0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x6c, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73,
	0x68, 0x69, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x76, 0x70, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0e, 0x70,
	0x76, 0x70, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0d, 0x70, 0x76, 0x70, 0x5f, 0x77, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x76, 0x70, 0x57, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x6e, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x64, 0x76, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x64, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x0d, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x70, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0e, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0f,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x65, 0x64, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x6d,
	0x65, 0x64, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73,
	0x74, 0x2e, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x46, 0x4f, 0x52, 0x07, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66,
}

var (
	file_DETAIL_INFO_proto_rawDescOnce sync.Once
	file_DETAIL_INFO_proto_rawDescData = file_DETAIL_INFO_proto_rawDesc
)

func file_DETAIL_INFO_proto_rawDescGZIP() []byte {
	file_DETAIL_INFO_proto_rawDescOnce.Do(func() {
		file_DETAIL_INFO_proto_rawDescData = protoimpl.X.CompressGZIP(file_DETAIL_INFO_proto_rawDescData)
	})
	return file_DETAIL_INFO_proto_rawDescData
}

var file_DETAIL_INFO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DETAIL_INFO_proto_goTypes = []interface{}{
	(*DETAIL_INFO)(nil), // 0: belfast.DETAIL_INFO
	(*DISPLAYINFO)(nil), // 1: belfast.DISPLAYINFO
}
var file_DETAIL_INFO_proto_depIdxs = []int32{
	1, // 0: belfast.DETAIL_INFO.display:type_name -> belfast.DISPLAYINFO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DETAIL_INFO_proto_init() }
func file_DETAIL_INFO_proto_init() {
	if File_DETAIL_INFO_proto != nil {
		return
	}
	file_DISPLAYINFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DETAIL_INFO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DETAIL_INFO); i {
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
			RawDescriptor: file_DETAIL_INFO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DETAIL_INFO_proto_goTypes,
		DependencyIndexes: file_DETAIL_INFO_proto_depIdxs,
		MessageInfos:      file_DETAIL_INFO_proto_msgTypes,
	}.Build()
	File_DETAIL_INFO_proto = out.File
	file_DETAIL_INFO_proto_rawDesc = nil
	file_DETAIL_INFO_proto_goTypes = nil
	file_DETAIL_INFO_proto_depIdxs = nil
}