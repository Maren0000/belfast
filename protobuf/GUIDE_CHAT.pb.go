// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: GUIDE_CHAT.proto

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

type GUIDE_CHAT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player  *PLAYER_INFO `protobuf:"bytes,1,req,name=player" json:"player,omitempty"`
	Content *string      `protobuf:"bytes,2,req,name=content" json:"content,omitempty"`
	Time    *uint32      `protobuf:"varint,3,req,name=time" json:"time,omitempty"`
}

func (x *GUIDE_CHAT) Reset() {
	*x = GUIDE_CHAT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GUIDE_CHAT_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GUIDE_CHAT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GUIDE_CHAT) ProtoMessage() {}

func (x *GUIDE_CHAT) ProtoReflect() protoreflect.Message {
	mi := &file_GUIDE_CHAT_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GUIDE_CHAT.ProtoReflect.Descriptor instead.
func (*GUIDE_CHAT) Descriptor() ([]byte, []int) {
	return file_GUIDE_CHAT_proto_rawDescGZIP(), []int{0}
}

func (x *GUIDE_CHAT) GetPlayer() *PLAYER_INFO {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *GUIDE_CHAT) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *GUIDE_CHAT) GetTime() uint32 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

var File_GUIDE_CHAT_proto protoreflect.FileDescriptor

var file_GUIDE_CHAT_proto_rawDesc = []byte{
	0x0a, 0x10, 0x47, 0x55, 0x49, 0x44, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x11, 0x50, 0x4c, 0x41,
	0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68,
	0x0a, 0x0a, 0x47, 0x55, 0x49, 0x44, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x12, 0x2c, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62,
	0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e,
	0x46, 0x4f, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_GUIDE_CHAT_proto_rawDescOnce sync.Once
	file_GUIDE_CHAT_proto_rawDescData = file_GUIDE_CHAT_proto_rawDesc
)

func file_GUIDE_CHAT_proto_rawDescGZIP() []byte {
	file_GUIDE_CHAT_proto_rawDescOnce.Do(func() {
		file_GUIDE_CHAT_proto_rawDescData = protoimpl.X.CompressGZIP(file_GUIDE_CHAT_proto_rawDescData)
	})
	return file_GUIDE_CHAT_proto_rawDescData
}

var file_GUIDE_CHAT_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GUIDE_CHAT_proto_goTypes = []interface{}{
	(*GUIDE_CHAT)(nil),  // 0: belfast.GUIDE_CHAT
	(*PLAYER_INFO)(nil), // 1: belfast.PLAYER_INFO
}
var file_GUIDE_CHAT_proto_depIdxs = []int32{
	1, // 0: belfast.GUIDE_CHAT.player:type_name -> belfast.PLAYER_INFO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GUIDE_CHAT_proto_init() }
func file_GUIDE_CHAT_proto_init() {
	if File_GUIDE_CHAT_proto != nil {
		return
	}
	file_PLAYER_INFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GUIDE_CHAT_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GUIDE_CHAT); i {
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
			RawDescriptor: file_GUIDE_CHAT_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GUIDE_CHAT_proto_goTypes,
		DependencyIndexes: file_GUIDE_CHAT_proto_depIdxs,
		MessageInfos:      file_GUIDE_CHAT_proto_msgTypes,
	}.Build()
	File_GUIDE_CHAT_proto = out.File
	file_GUIDE_CHAT_proto_rawDesc = nil
	file_GUIDE_CHAT_proto_goTypes = nil
	file_GUIDE_CHAT_proto_depIdxs = nil
}
