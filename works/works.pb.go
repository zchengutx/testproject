// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: works.proto

package works

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

type CreatedWorksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title          string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Desc           string `protobuf:"bytes,2,opt,name=Desc,proto3" json:"Desc,omitempty"`
	MusicId        int64  `protobuf:"varint,3,opt,name=MusicId,proto3" json:"MusicId,omitempty"`
	WorkType       string `protobuf:"bytes,4,opt,name=WorkType,proto3" json:"WorkType,omitempty"`
	IpAddress      string `protobuf:"bytes,5,opt,name=IpAddress,proto3" json:"IpAddress,omitempty"`
	WorkPermission string `protobuf:"bytes,6,opt,name=WorkPermission,proto3" json:"WorkPermission,omitempty"`
}

func (x *CreatedWorksReq) Reset() {
	*x = CreatedWorksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_works_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedWorksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedWorksReq) ProtoMessage() {}

func (x *CreatedWorksReq) ProtoReflect() protoreflect.Message {
	mi := &file_works_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedWorksReq.ProtoReflect.Descriptor instead.
func (*CreatedWorksReq) Descriptor() ([]byte, []int) {
	return file_works_proto_rawDescGZIP(), []int{0}
}

func (x *CreatedWorksReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatedWorksReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreatedWorksReq) GetMusicId() int64 {
	if x != nil {
		return x.MusicId
	}
	return 0
}

func (x *CreatedWorksReq) GetWorkType() string {
	if x != nil {
		return x.WorkType
	}
	return ""
}

func (x *CreatedWorksReq) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *CreatedWorksReq) GetWorkPermission() string {
	if x != nil {
		return x.WorkPermission
	}
	return ""
}

type CreatedWorksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreatedWorksResp) Reset() {
	*x = CreatedWorksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_works_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatedWorksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatedWorksResp) ProtoMessage() {}

func (x *CreatedWorksResp) ProtoReflect() protoreflect.Message {
	mi := &file_works_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatedWorksResp.ProtoReflect.Descriptor instead.
func (*CreatedWorksResp) Descriptor() ([]byte, []int) {
	return file_works_proto_rawDescGZIP(), []int{1}
}

func (x *CreatedWorksResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreatedWorksResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_works_proto protoreflect.FileDescriptor

var file_works_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0xb7, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x57, 0x6f, 0x72, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x38, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x4a, 0x0a, 0x05, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_works_proto_rawDescOnce sync.Once
	file_works_proto_rawDescData = file_works_proto_rawDesc
)

func file_works_proto_rawDescGZIP() []byte {
	file_works_proto_rawDescOnce.Do(func() {
		file_works_proto_rawDescData = protoimpl.X.CompressGZIP(file_works_proto_rawDescData)
	})
	return file_works_proto_rawDescData
}

var file_works_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_works_proto_goTypes = []any{
	(*CreatedWorksReq)(nil),  // 0: stream.CreatedWorksReq
	(*CreatedWorksResp)(nil), // 1: stream.CreatedWorksResp
}
var file_works_proto_depIdxs = []int32{
	0, // 0: stream.Works.CreatedWorks:input_type -> stream.CreatedWorksReq
	1, // 1: stream.Works.CreatedWorks:output_type -> stream.CreatedWorksResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_works_proto_init() }
func file_works_proto_init() {
	if File_works_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_works_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreatedWorksReq); i {
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
		file_works_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatedWorksResp); i {
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
			RawDescriptor: file_works_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_works_proto_goTypes,
		DependencyIndexes: file_works_proto_depIdxs,
		MessageInfos:      file_works_proto_msgTypes,
	}.Build()
	File_works_proto = out.File
	file_works_proto_rawDesc = nil
	file_works_proto_goTypes = nil
	file_works_proto_depIdxs = nil
}
