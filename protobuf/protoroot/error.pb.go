// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: error.proto

package protoroot

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_UNSPECIFIED                ErrorCode = 0
	ErrorCode_ERROR_CODE_INDEX_NOT_PRESENT_ON_FIELD ErrorCode = 1
	ErrorCode_ERROR_CODE_API_ERROR_RETRY            ErrorCode = 2
	ErrorCode_ERROR_CODE_INVALID_TOKEN              ErrorCode = 3
	ErrorCode_ERROR_CODE_DB_QUERY_TIMEOUT           ErrorCode = 4
	ErrorCode_ERROR_CODE_DB_CAPACITY_EXCEEDED       ErrorCode = 5
	ErrorCode_ERROR_CODE_DB_SESSION_EXPIRED         ErrorCode = 6
	ErrorCode_ERROR_CODE_DB_SESSION_LIMIT_EXCEEDED  ErrorCode = 7
	ErrorCode_ERROR_CODE_DB_RATE_LIMIT_EXCEEDED     ErrorCode = 8
	ErrorCode_ERROR_CODE_GENERATE_ACCOUNT_NUMBER    ErrorCode = 14
	ErrorCode_ERROR_CODE_APIKEY_ID_ERROR            ErrorCode = 9
	ErrorCode_ERROR_CODE_APIKEY_ID_RETRY_EXCEEDED   ErrorCode = 10
	ErrorCode_ERROR_CODE_AUTH_ERROR                 ErrorCode = 12
	ErrorCode_ERROR_CODE_INVALID_SHOW_TOKEN         ErrorCode = 13
	ErrorCode_ERROR_CODE_INTERNAL_ERROR             ErrorCode = 16
	ErrorCode_ERROR_CODE_PROTO_VALIDATION_FAILED    ErrorCode = 15
	ErrorCode_ERROR_CODE_UNAUTHORIZED               ErrorCode = 17
	ErrorCode_ERROR_CODE_REDIS_ERROR                ErrorCode = 18
	// DB Errors
	ErrorCode_ERROR_CODE_DB_ERROR                    ErrorCode = 21
	ErrorCode_ERROR_CODE_DB_TYPE_ERROR               ErrorCode = 22
	ErrorCode_ERROR_CODE_DB_VERSION_MISMATCH         ErrorCode = 23 // error generated when update is triggered but version is incorrect
	ErrorCode_ERROR_CODE_AUTH_FORBIDDEN              ErrorCode = 100
	ErrorCode_ERROR_CODE_AUTH_EXPIRED                ErrorCode = 101
	ErrorCode_ERROR_CODE_INVALID_FIELD               ErrorCode = 400
	ErrorCode_ERROR_CODE_RESOURCE_ALREADY_EXIST      ErrorCode = 401
	ErrorCode_ERROR_CODE_RESOURCE_NOT_FOUND          ErrorCode = 402
	ErrorCode_ERROR_CODE_RESOURCE_ERROR              ErrorCode = 404
	ErrorCode_ERROR_CODE_INVALID_REQUEST_BODY        ErrorCode = 405
	ErrorCode_ERROR_CODE_RESOURCE_UPDATE_NOT_ALLOWED ErrorCode = 406
	ErrorCode_ERROR_CODE_INVALID_STATUS              ErrorCode = 407
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:   "ERROR_CODE_UNSPECIFIED",
		1:   "ERROR_CODE_INDEX_NOT_PRESENT_ON_FIELD",
		2:   "ERROR_CODE_API_ERROR_RETRY",
		3:   "ERROR_CODE_INVALID_TOKEN",
		4:   "ERROR_CODE_DB_QUERY_TIMEOUT",
		5:   "ERROR_CODE_DB_CAPACITY_EXCEEDED",
		6:   "ERROR_CODE_DB_SESSION_EXPIRED",
		7:   "ERROR_CODE_DB_SESSION_LIMIT_EXCEEDED",
		8:   "ERROR_CODE_DB_RATE_LIMIT_EXCEEDED",
		14:  "ERROR_CODE_GENERATE_ACCOUNT_NUMBER",
		9:   "ERROR_CODE_APIKEY_ID_ERROR",
		10:  "ERROR_CODE_APIKEY_ID_RETRY_EXCEEDED",
		12:  "ERROR_CODE_AUTH_ERROR",
		13:  "ERROR_CODE_INVALID_SHOW_TOKEN",
		16:  "ERROR_CODE_INTERNAL_ERROR",
		15:  "ERROR_CODE_PROTO_VALIDATION_FAILED",
		17:  "ERROR_CODE_UNAUTHORIZED",
		18:  "ERROR_CODE_REDIS_ERROR",
		21:  "ERROR_CODE_DB_ERROR",
		22:  "ERROR_CODE_DB_TYPE_ERROR",
		23:  "ERROR_CODE_DB_VERSION_MISMATCH",
		100: "ERROR_CODE_AUTH_FORBIDDEN",
		101: "ERROR_CODE_AUTH_EXPIRED",
		400: "ERROR_CODE_INVALID_FIELD",
		401: "ERROR_CODE_RESOURCE_ALREADY_EXIST",
		402: "ERROR_CODE_RESOURCE_NOT_FOUND",
		404: "ERROR_CODE_RESOURCE_ERROR",
		405: "ERROR_CODE_INVALID_REQUEST_BODY",
		406: "ERROR_CODE_RESOURCE_UPDATE_NOT_ALLOWED",
		407: "ERROR_CODE_INVALID_STATUS",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_UNSPECIFIED":                 0,
		"ERROR_CODE_INDEX_NOT_PRESENT_ON_FIELD":  1,
		"ERROR_CODE_API_ERROR_RETRY":             2,
		"ERROR_CODE_INVALID_TOKEN":               3,
		"ERROR_CODE_DB_QUERY_TIMEOUT":            4,
		"ERROR_CODE_DB_CAPACITY_EXCEEDED":        5,
		"ERROR_CODE_DB_SESSION_EXPIRED":          6,
		"ERROR_CODE_DB_SESSION_LIMIT_EXCEEDED":   7,
		"ERROR_CODE_DB_RATE_LIMIT_EXCEEDED":      8,
		"ERROR_CODE_GENERATE_ACCOUNT_NUMBER":     14,
		"ERROR_CODE_APIKEY_ID_ERROR":             9,
		"ERROR_CODE_APIKEY_ID_RETRY_EXCEEDED":    10,
		"ERROR_CODE_AUTH_ERROR":                  12,
		"ERROR_CODE_INVALID_SHOW_TOKEN":          13,
		"ERROR_CODE_INTERNAL_ERROR":              16,
		"ERROR_CODE_PROTO_VALIDATION_FAILED":     15,
		"ERROR_CODE_UNAUTHORIZED":                17,
		"ERROR_CODE_REDIS_ERROR":                 18,
		"ERROR_CODE_DB_ERROR":                    21,
		"ERROR_CODE_DB_TYPE_ERROR":               22,
		"ERROR_CODE_DB_VERSION_MISMATCH":         23,
		"ERROR_CODE_AUTH_FORBIDDEN":              100,
		"ERROR_CODE_AUTH_EXPIRED":                101,
		"ERROR_CODE_INVALID_FIELD":               400,
		"ERROR_CODE_RESOURCE_ALREADY_EXIST":      401,
		"ERROR_CODE_RESOURCE_NOT_FOUND":          402,
		"ERROR_CODE_RESOURCE_ERROR":              404,
		"ERROR_CODE_INVALID_REQUEST_BODY":        405,
		"ERROR_CODE_RESOURCE_UPDATE_NOT_ALLOWED": 406,
		"ERROR_CODE_INVALID_STATUS":              407,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

type GrpcError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=proto.ErrorCode" json:"code,omitempty"`
	Message    string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FieldName  string    `protobuf:"bytes,3,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	SysMessage string    `protobuf:"bytes,4,opt,name=sys_message,json=sysMessage,proto3" json:"sys_message,omitempty"`
	HttpCode   int32     `protobuf:"varint,5,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
}

func (x *GrpcError) Reset() {
	*x = GrpcError{}
	mi := &file_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrpcError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcError) ProtoMessage() {}

func (x *GrpcError) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcError.ProtoReflect.Descriptor instead.
func (*GrpcError) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcError) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *GrpcError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GrpcError) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *GrpcError) GetSysMessage() string {
	if x != nil {
		return x.SysMessage
	}
	return ""
}

func (x *GrpcError) GetHttpCode() int32 {
	if x != nil {
		return x.HttpCode
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FieldName string `protobuf:"bytes,3,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_error_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

type HTTPError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string                 `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ClientId  string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Method    string                 `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Status    int32                  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Error     *Error                 `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *HTTPError) Reset() {
	*x = HTTPError{}
	mi := &file_error_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HTTPError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPError) ProtoMessage() {}

func (x *HTTPError) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPError.ProtoReflect.Descriptor instead.
func (*HTTPError) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{2}
}

func (x *HTTPError) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *HTTPError) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *HTTPError) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HTTPError) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HTTPError) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *HTTPError) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SysMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SysMessage) Reset() {
	*x = SysMessage{}
	mi := &file_error_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SysMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysMessage) ProtoMessage() {}

func (x *SysMessage) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysMessage.ProtoReflect.Descriptor instead.
func (*SysMessage) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{3}
}

func (x *SysMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RequestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *RequestId) Reset() {
	*x = RequestId{}
	mi := &file_error_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestId) ProtoMessage() {}

func (x *RequestId) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestId.ProtoReflect.Descriptor instead.
func (*RequestId) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{4}
}

func (x *RequestId) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x09, 0x47, 0x72, 0x70, 0x63, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x54, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x09, 0x48, 0x54, 0x54, 0x50, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x26, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x2a, 0x8e, 0x08, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a,
	0x25, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x44, 0x45,
	0x58, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x4e,
	0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x43, 0x41, 0x50, 0x41, 0x43, 0x49, 0x54,
	0x59, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x05, 0x12, 0x21, 0x0a, 0x1d,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x06, 0x12,
	0x28, 0x0a, 0x24, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42,
	0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45,
	0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x07, 0x12, 0x25, 0x0a, 0x21, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x5f,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x08,
	0x12, 0x26, 0x0a, 0x22, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x47,
	0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x0e, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x4b, 0x45, 0x59, 0x5f, 0x49, 0x44,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x09, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x4b, 0x45, 0x59, 0x5f, 0x49, 0x44,
	0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x0a, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0c, 0x12, 0x21, 0x0a, 0x1d,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x0d, 0x12,
	0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x10, 0x12, 0x26,
	0x0a, 0x22, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x52, 0x4f,
	0x54, 0x4f, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x0f, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45,
	0x44, 0x10, 0x11, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x12, 0x12,
	0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x15, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x16, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x17, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x46, 0x4f,
	0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x65, 0x12, 0x1d, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x10, 0x90, 0x03, 0x12, 0x26, 0x0a, 0x21, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x91, 0x03, 0x12, 0x22, 0x0a,
	0x1d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x92,
	0x03, 0x12, 0x1e, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x94,
	0x03, 0x12, 0x24, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x42, 0x4f, 0x44, 0x59, 0x10, 0x95, 0x03, 0x12, 0x2b, 0x0a, 0x26, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45,
	0x44, 0x10, 0x96, 0x03, 0x12, 0x1e, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x10, 0x97, 0x03, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x76, 0x65, 0x6b, 0x61, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x69, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72,
	0x6f, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_error_proto_goTypes = []any{
	(ErrorCode)(0),                // 0: proto.ErrorCode
	(*GrpcError)(nil),             // 1: proto.GrpcError
	(*Error)(nil),                 // 2: proto.Error
	(*HTTPError)(nil),             // 3: proto.HTTPError
	(*SysMessage)(nil),            // 4: proto.SysMessage
	(*RequestId)(nil),             // 5: proto.RequestId
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_error_proto_depIdxs = []int32{
	0, // 0: proto.GrpcError.code:type_name -> proto.ErrorCode
	2, // 1: proto.HTTPError.error:type_name -> proto.Error
	6, // 2: proto.HTTPError.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
		MessageInfos:      file_error_proto_msgTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
