// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/common/base_response.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StatusCode    int64                  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	IsError       bool                   `protobuf:"varint,3,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_proto_common_base_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_base_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_common_base_response_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaseResponse) GetIsError() bool {
	if x != nil {
		return x.IsError
	}
	return false
}

var File_proto_common_base_response_proto protoreflect.FileDescriptor

const file_proto_common_base_response_proto_rawDesc = "" +
	"\n" +
	" proto/common/base_response.proto\x12\x06common\"d\n" +
	"\fBaseResponse\x12\x1f\n" +
	"\vstatus_code\x18\x01 \x01(\x03R\n" +
	"statusCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x19\n" +
	"\bis_error\x18\x03 \x01(\bR\aisErrorB8Z6github.com/Ilhamkawe/commerce-proto/protogen/go/commonb\x06proto3"

var (
	file_proto_common_base_response_proto_rawDescOnce sync.Once
	file_proto_common_base_response_proto_rawDescData []byte
)

func file_proto_common_base_response_proto_rawDescGZIP() []byte {
	file_proto_common_base_response_proto_rawDescOnce.Do(func() {
		file_proto_common_base_response_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_common_base_response_proto_rawDesc), len(file_proto_common_base_response_proto_rawDesc)))
	})
	return file_proto_common_base_response_proto_rawDescData
}

var file_proto_common_base_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_common_base_response_proto_goTypes = []any{
	(*BaseResponse)(nil), // 0: common.BaseResponse
}
var file_proto_common_base_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_common_base_response_proto_init() }
func file_proto_common_base_response_proto_init() {
	if File_proto_common_base_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_common_base_response_proto_rawDesc), len(file_proto_common_base_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_base_response_proto_goTypes,
		DependencyIndexes: file_proto_common_base_response_proto_depIdxs,
		MessageInfos:      file_proto_common_base_response_proto_msgTypes,
	}.Build()
	File_proto_common_base_response_proto = out.File
	file_proto_common_base_response_proto_goTypes = nil
	file_proto_common_base_response_proto_depIdxs = nil
}
