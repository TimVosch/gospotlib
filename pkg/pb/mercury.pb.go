// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: mercury.proto

package pb

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

type MercuryReply_CachePolicy int32

const (
	MercuryReply_CACHE_NO      MercuryReply_CachePolicy = 1
	MercuryReply_CACHE_PRIVATE MercuryReply_CachePolicy = 2
	MercuryReply_CACHE_PUBLIC  MercuryReply_CachePolicy = 3
)

// Enum value maps for MercuryReply_CachePolicy.
var (
	MercuryReply_CachePolicy_name = map[int32]string{
		1: "CACHE_NO",
		2: "CACHE_PRIVATE",
		3: "CACHE_PUBLIC",
	}
	MercuryReply_CachePolicy_value = map[string]int32{
		"CACHE_NO":      1,
		"CACHE_PRIVATE": 2,
		"CACHE_PUBLIC":  3,
	}
)

func (x MercuryReply_CachePolicy) Enum() *MercuryReply_CachePolicy {
	p := new(MercuryReply_CachePolicy)
	*p = x
	return p
}

func (x MercuryReply_CachePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MercuryReply_CachePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_mercury_proto_enumTypes[0].Descriptor()
}

func (MercuryReply_CachePolicy) Type() protoreflect.EnumType {
	return &file_mercury_proto_enumTypes[0]
}

func (x MercuryReply_CachePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MercuryReply_CachePolicy) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MercuryReply_CachePolicy(num)
	return nil
}

// Deprecated: Use MercuryReply_CachePolicy.Descriptor instead.
func (MercuryReply_CachePolicy) EnumDescriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{3, 0}
}

type MercuryMultiGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request []*MercuryRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (x *MercuryMultiGetRequest) Reset() {
	*x = MercuryMultiGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryMultiGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryMultiGetRequest) ProtoMessage() {}

func (x *MercuryMultiGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryMultiGetRequest.ProtoReflect.Descriptor instead.
func (*MercuryMultiGetRequest) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{0}
}

func (x *MercuryMultiGetRequest) GetRequest() []*MercuryRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type MercuryMultiGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply []*MercuryReply `protobuf:"bytes,1,rep,name=reply" json:"reply,omitempty"`
}

func (x *MercuryMultiGetReply) Reset() {
	*x = MercuryMultiGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryMultiGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryMultiGetReply) ProtoMessage() {}

func (x *MercuryMultiGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryMultiGetReply.ProtoReflect.Descriptor instead.
func (*MercuryMultiGetReply) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{1}
}

func (x *MercuryMultiGetReply) GetReply() []*MercuryReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type MercuryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri         *string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType *string `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body        []byte  `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Etag        []byte  `protobuf:"bytes,4,opt,name=etag" json:"etag,omitempty"`
}

func (x *MercuryRequest) Reset() {
	*x = MercuryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryRequest) ProtoMessage() {}

func (x *MercuryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryRequest.ProtoReflect.Descriptor instead.
func (*MercuryRequest) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{2}
}

func (x *MercuryRequest) GetUri() string {
	if x != nil && x.Uri != nil {
		return *x.Uri
	}
	return ""
}

func (x *MercuryRequest) GetContentType() string {
	if x != nil && x.ContentType != nil {
		return *x.ContentType
	}
	return ""
}

func (x *MercuryRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *MercuryRequest) GetEtag() []byte {
	if x != nil {
		return x.Etag
	}
	return nil
}

type MercuryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    *int32                    `protobuf:"zigzag32,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMessage *string                   `protobuf:"bytes,2,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
	CachePolicy   *MercuryReply_CachePolicy `protobuf:"varint,3,opt,name=cache_policy,json=cachePolicy,enum=pb.MercuryReply_CachePolicy" json:"cache_policy,omitempty"`
	Ttl           *int32                    `protobuf:"zigzag32,4,opt,name=ttl" json:"ttl,omitempty"`
	Etag          []byte                    `protobuf:"bytes,5,opt,name=etag" json:"etag,omitempty"`
	ContentType   *string                   `protobuf:"bytes,6,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body          []byte                    `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
}

func (x *MercuryReply) Reset() {
	*x = MercuryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MercuryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MercuryReply) ProtoMessage() {}

func (x *MercuryReply) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MercuryReply.ProtoReflect.Descriptor instead.
func (*MercuryReply) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{3}
}

func (x *MercuryReply) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *MercuryReply) GetStatusMessage() string {
	if x != nil && x.StatusMessage != nil {
		return *x.StatusMessage
	}
	return ""
}

func (x *MercuryReply) GetCachePolicy() MercuryReply_CachePolicy {
	if x != nil && x.CachePolicy != nil {
		return *x.CachePolicy
	}
	return MercuryReply_CACHE_NO
}

func (x *MercuryReply) GetTtl() int32 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

func (x *MercuryReply) GetEtag() []byte {
	if x != nil {
		return x.Etag
	}
	return nil
}

func (x *MercuryReply) GetContentType() string {
	if x != nil && x.ContentType != nil {
		return *x.ContentType
	}
	return ""
}

func (x *MercuryReply) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri         *string      `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType *string      `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Method      *string      `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	StatusCode  *int32       `protobuf:"zigzag32,4,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	UserFields  []*UserField `protobuf:"bytes,6,rep,name=user_fields,json=userFields" json:"user_fields,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{4}
}

func (x *Header) GetUri() string {
	if x != nil && x.Uri != nil {
		return *x.Uri
	}
	return ""
}

func (x *Header) GetContentType() string {
	if x != nil && x.ContentType != nil {
		return *x.ContentType
	}
	return ""
}

func (x *Header) GetMethod() string {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return ""
}

func (x *Header) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *Header) GetUserFields() []*UserField {
	if x != nil {
		return x.UserFields
	}
	return nil
}

type UserField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *UserField) Reset() {
	*x = UserField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mercury_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserField) ProtoMessage() {}

func (x *UserField) ProtoReflect() protoreflect.Message {
	mi := &file_mercury_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserField.ProtoReflect.Descriptor instead.
func (*UserField) Descriptor() ([]byte, []int) {
	return file_mercury_proto_rawDescGZIP(), []int{5}
}

func (x *UserField) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *UserField) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_mercury_proto protoreflect.FileDescriptor

var file_mercury_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x46, 0x0a, 0x16, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x14, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x6d, 0x0a, 0x0e, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0xb6, 0x02, 0x0a, 0x0c, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x40, 0x0a, 0x0b, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x4e, 0x4f, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49,
	0x43, 0x10, 0x03, 0x22, 0xa6, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x69, 0x6d, 0x76, 0x6f, 0x73, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x73, 0x70, 0x6f, 0x74, 0x6c,
	0x69, 0x62, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
}

var (
	file_mercury_proto_rawDescOnce sync.Once
	file_mercury_proto_rawDescData = file_mercury_proto_rawDesc
)

func file_mercury_proto_rawDescGZIP() []byte {
	file_mercury_proto_rawDescOnce.Do(func() {
		file_mercury_proto_rawDescData = protoimpl.X.CompressGZIP(file_mercury_proto_rawDescData)
	})
	return file_mercury_proto_rawDescData
}

var file_mercury_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mercury_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mercury_proto_goTypes = []interface{}{
	(MercuryReply_CachePolicy)(0),  // 0: pb.MercuryReply.CachePolicy
	(*MercuryMultiGetRequest)(nil), // 1: pb.MercuryMultiGetRequest
	(*MercuryMultiGetReply)(nil),   // 2: pb.MercuryMultiGetReply
	(*MercuryRequest)(nil),         // 3: pb.MercuryRequest
	(*MercuryReply)(nil),           // 4: pb.MercuryReply
	(*Header)(nil),                 // 5: pb.Header
	(*UserField)(nil),              // 6: pb.UserField
}
var file_mercury_proto_depIdxs = []int32{
	3, // 0: pb.MercuryMultiGetRequest.request:type_name -> pb.MercuryRequest
	4, // 1: pb.MercuryMultiGetReply.reply:type_name -> pb.MercuryReply
	0, // 2: pb.MercuryReply.cache_policy:type_name -> pb.MercuryReply.CachePolicy
	6, // 3: pb.Header.user_fields:type_name -> pb.UserField
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mercury_proto_init() }
func file_mercury_proto_init() {
	if File_mercury_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mercury_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryMultiGetRequest); i {
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
		file_mercury_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryMultiGetReply); i {
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
		file_mercury_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryRequest); i {
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
		file_mercury_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MercuryReply); i {
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
		file_mercury_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_mercury_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserField); i {
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
			RawDescriptor: file_mercury_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mercury_proto_goTypes,
		DependencyIndexes: file_mercury_proto_depIdxs,
		EnumInfos:         file_mercury_proto_enumTypes,
		MessageInfos:      file_mercury_proto_msgTypes,
	}.Build()
	File_mercury_proto = out.File
	file_mercury_proto_rawDesc = nil
	file_mercury_proto_goTypes = nil
	file_mercury_proto_depIdxs = nil
}
