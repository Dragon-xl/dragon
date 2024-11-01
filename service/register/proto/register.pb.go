// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/register.proto

package register

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

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	CapCode string `protobuf:"bytes,2,opt,name=capCode,proto3" json:"capCode,omitempty"`
	Uuid    string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	mi := &file_proto_register_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_register_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_proto_register_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CallRequest) GetCapCode() string {
	if x != nil {
		return x.CapCode
	}
	return ""
}

func (x *CallRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	mi := &file_proto_register_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_register_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_proto_register_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetErrno() string {
	if x != nil {
		return x.Errno
	}
	return ""
}

func (x *CallResponse) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

type RegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	SmsCode  string `protobuf:"bytes,3,opt,name=smsCode,proto3" json:"smsCode,omitempty"`
}

func (x *RegRequest) Reset() {
	*x = RegRequest{}
	mi := &file_proto_register_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegRequest) ProtoMessage() {}

func (x *RegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_register_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegRequest.ProtoReflect.Descriptor instead.
func (*RegRequest) Descriptor() ([]byte, []int) {
	return file_proto_register_proto_rawDescGZIP(), []int{2}
}

func (x *RegRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegRequest) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

var File_proto_register_proto protoreflect.FileDescriptor

var file_proto_register_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x51, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73,
	0x67, 0x22, 0x5a, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x82, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x6d, 0x73, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_register_proto_rawDescOnce sync.Once
	file_proto_register_proto_rawDescData = file_proto_register_proto_rawDesc
)

func file_proto_register_proto_rawDescGZIP() []byte {
	file_proto_register_proto_rawDescOnce.Do(func() {
		file_proto_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_register_proto_rawDescData)
	})
	return file_proto_register_proto_rawDescData
}

var file_proto_register_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_register_proto_goTypes = []any{
	(*CallRequest)(nil),  // 0: register.CallRequest
	(*CallResponse)(nil), // 1: register.CallResponse
	(*RegRequest)(nil),   // 2: register.RegRequest
}
var file_proto_register_proto_depIdxs = []int32{
	0, // 0: register.Register.SendSms:input_type -> register.CallRequest
	2, // 1: register.Register.Register:input_type -> register.RegRequest
	1, // 2: register.Register.SendSms:output_type -> register.CallResponse
	1, // 3: register.Register.Register:output_type -> register.CallResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_register_proto_init() }
func file_proto_register_proto_init() {
	if File_proto_register_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_register_proto_goTypes,
		DependencyIndexes: file_proto_register_proto_depIdxs,
		MessageInfos:      file_proto_register_proto_msgTypes,
	}.Build()
	File_proto_register_proto = out.File
	file_proto_register_proto_rawDesc = nil
	file_proto_register_proto_goTypes = nil
	file_proto_register_proto_depIdxs = nil
}
