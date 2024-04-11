// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_18101.proto

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

type SC_18101 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlashCount    *uint32      `protobuf:"varint,1,req,name=flash_count,json=flashCount" json:"flash_count,omitempty"`
	ArenaShopList []*ARENASHOP `protobuf:"bytes,2,rep,name=arena_shop_list,json=arenaShopList" json:"arena_shop_list,omitempty"`
	NextFlashTime *uint32      `protobuf:"varint,3,req,name=next_flash_time,json=nextFlashTime" json:"next_flash_time,omitempty"`
}

func (x *SC_18101) Reset() {
	*x = SC_18101{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_18101_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_18101) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_18101) ProtoMessage() {}

func (x *SC_18101) ProtoReflect() protoreflect.Message {
	mi := &file_SC_18101_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_18101.ProtoReflect.Descriptor instead.
func (*SC_18101) Descriptor() ([]byte, []int) {
	return file_SC_18101_proto_rawDescGZIP(), []int{0}
}

func (x *SC_18101) GetFlashCount() uint32 {
	if x != nil && x.FlashCount != nil {
		return *x.FlashCount
	}
	return 0
}

func (x *SC_18101) GetArenaShopList() []*ARENASHOP {
	if x != nil {
		return x.ArenaShopList
	}
	return nil
}

func (x *SC_18101) GetNextFlashTime() uint32 {
	if x != nil && x.NextFlashTime != nil {
		return *x.NextFlashTime
	}
	return 0
}

var File_SC_18101_proto protoreflect.FileDescriptor

var file_SC_18101_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x31, 0x38, 0x31, 0x30, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0f, 0x41, 0x52, 0x45, 0x4e, 0x41,
	0x53, 0x48, 0x4f, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x08, 0x53,
	0x43, 0x5f, 0x31, 0x38, 0x31, 0x30, 0x31, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x61, 0x73, 0x68,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a, 0x66, 0x6c,
	0x61, 0x73, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0f, 0x61, 0x72, 0x65, 0x6e,
	0x61, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x41, 0x52, 0x45, 0x4e,
	0x41, 0x53, 0x48, 0x4f, 0x50, 0x52, 0x0d, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x53, 0x68, 0x6f, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x6c, 0x61,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_18101_proto_rawDescOnce sync.Once
	file_SC_18101_proto_rawDescData = file_SC_18101_proto_rawDesc
)

func file_SC_18101_proto_rawDescGZIP() []byte {
	file_SC_18101_proto_rawDescOnce.Do(func() {
		file_SC_18101_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_18101_proto_rawDescData)
	})
	return file_SC_18101_proto_rawDescData
}

var file_SC_18101_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_18101_proto_goTypes = []interface{}{
	(*SC_18101)(nil),  // 0: belfast.SC_18101
	(*ARENASHOP)(nil), // 1: belfast.ARENASHOP
}
var file_SC_18101_proto_depIdxs = []int32{
	1, // 0: belfast.SC_18101.arena_shop_list:type_name -> belfast.ARENASHOP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SC_18101_proto_init() }
func file_SC_18101_proto_init() {
	if File_SC_18101_proto != nil {
		return
	}
	file_ARENASHOP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_18101_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_18101); i {
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
			RawDescriptor: file_SC_18101_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_18101_proto_goTypes,
		DependencyIndexes: file_SC_18101_proto_depIdxs,
		MessageInfos:      file_SC_18101_proto_msgTypes,
	}.Build()
	File_SC_18101_proto = out.File
	file_SC_18101_proto_rawDesc = nil
	file_SC_18101_proto_goTypes = nil
	file_SC_18101_proto_depIdxs = nil
}
