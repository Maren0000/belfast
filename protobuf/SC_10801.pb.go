// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_10801.proto

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

type SC_10801 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayIp               *string  `protobuf:"bytes,1,req,name=gateway_ip,json=gatewayIp" json:"gateway_ip,omitempty"`
	GatewayPort             *uint32  `protobuf:"varint,2,req,name=gateway_port,json=gatewayPort" json:"gateway_port,omitempty"`
	Url                     *string  `protobuf:"bytes,3,req,name=url" json:"url,omitempty"`
	Version                 []string `protobuf:"bytes,4,rep,name=version" json:"version,omitempty"`
	ProxyIp                 *string  `protobuf:"bytes,5,opt,name=proxy_ip,json=proxyIp" json:"proxy_ip,omitempty"`
	ProxyPort               *uint32  `protobuf:"varint,6,opt,name=proxy_port,json=proxyPort" json:"proxy_port,omitempty"`
	IsTs                    *uint32  `protobuf:"varint,7,req,name=is_ts,json=isTs" json:"is_ts,omitempty"`
	Timestamp               *uint32  `protobuf:"varint,8,req,name=timestamp" json:"timestamp,omitempty"`
	Monday_0OclockTimestamp *uint32  `protobuf:"varint,9,req,name=monday_0oclock_timestamp,json=monday0oclockTimestamp" json:"monday_0oclock_timestamp,omitempty"`
}

func (x *SC_10801) Reset() {
	*x = SC_10801{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_10801_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_10801) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_10801) ProtoMessage() {}

func (x *SC_10801) ProtoReflect() protoreflect.Message {
	mi := &file_SC_10801_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_10801.ProtoReflect.Descriptor instead.
func (*SC_10801) Descriptor() ([]byte, []int) {
	return file_SC_10801_proto_rawDescGZIP(), []int{0}
}

func (x *SC_10801) GetGatewayIp() string {
	if x != nil && x.GatewayIp != nil {
		return *x.GatewayIp
	}
	return ""
}

func (x *SC_10801) GetGatewayPort() uint32 {
	if x != nil && x.GatewayPort != nil {
		return *x.GatewayPort
	}
	return 0
}

func (x *SC_10801) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *SC_10801) GetVersion() []string {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *SC_10801) GetProxyIp() string {
	if x != nil && x.ProxyIp != nil {
		return *x.ProxyIp
	}
	return ""
}

func (x *SC_10801) GetProxyPort() uint32 {
	if x != nil && x.ProxyPort != nil {
		return *x.ProxyPort
	}
	return 0
}

func (x *SC_10801) GetIsTs() uint32 {
	if x != nil && x.IsTs != nil {
		return *x.IsTs
	}
	return 0
}

func (x *SC_10801) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *SC_10801) GetMonday_0OclockTimestamp() uint32 {
	if x != nil && x.Monday_0OclockTimestamp != nil {
		return *x.Monday_0OclockTimestamp
	}
	return 0
}

var File_SC_10801_proto protoreflect.FileDescriptor

var file_SC_10801_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x31, 0x30, 0x38, 0x30, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x9f, 0x02, 0x0a, 0x08, 0x53, 0x43,
	0x5f, 0x31, 0x30, 0x38, 0x30, 0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x49, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x69, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x70, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x13,
	0x0a, 0x05, 0x69, 0x73, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x69,
	0x73, 0x54, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x5f, 0x30, 0x6f, 0x63, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20,
	0x02, 0x28, 0x0d, 0x52, 0x16, 0x6d, 0x6f, 0x6e, 0x64, 0x61, 0x79, 0x30, 0x6f, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_10801_proto_rawDescOnce sync.Once
	file_SC_10801_proto_rawDescData = file_SC_10801_proto_rawDesc
)

func file_SC_10801_proto_rawDescGZIP() []byte {
	file_SC_10801_proto_rawDescOnce.Do(func() {
		file_SC_10801_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_10801_proto_rawDescData)
	})
	return file_SC_10801_proto_rawDescData
}

var file_SC_10801_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_10801_proto_goTypes = []interface{}{
	(*SC_10801)(nil), // 0: belfast.SC_10801
}
var file_SC_10801_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SC_10801_proto_init() }
func file_SC_10801_proto_init() {
	if File_SC_10801_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SC_10801_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_10801); i {
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
			RawDescriptor: file_SC_10801_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_10801_proto_goTypes,
		DependencyIndexes: file_SC_10801_proto_depIdxs,
		MessageInfos:      file_SC_10801_proto_msgTypes,
	}.Build()
	File_SC_10801_proto = out.File
	file_SC_10801_proto_rawDesc = nil
	file_SC_10801_proto_goTypes = nil
	file_SC_10801_proto_depIdxs = nil
}
