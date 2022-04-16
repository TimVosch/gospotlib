// Code generated by protoc-gen-go. DO NOT EDIT.
// source: playlist4issues.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ClientIssue_Level int32

const (
	ClientIssue_LEVEL_UNKNOWN ClientIssue_Level = 0
	ClientIssue_LEVEL_DEBUG   ClientIssue_Level = 1
	ClientIssue_LEVEL_INFO    ClientIssue_Level = 2
	ClientIssue_LEVEL_NOTICE  ClientIssue_Level = 3
	ClientIssue_LEVEL_WARNING ClientIssue_Level = 4
	ClientIssue_LEVEL_ERROR   ClientIssue_Level = 5
)

var ClientIssue_Level_name = map[int32]string{
	0: "LEVEL_UNKNOWN",
	1: "LEVEL_DEBUG",
	2: "LEVEL_INFO",
	3: "LEVEL_NOTICE",
	4: "LEVEL_WARNING",
	5: "LEVEL_ERROR",
}
var ClientIssue_Level_value = map[string]int32{
	"LEVEL_UNKNOWN": 0,
	"LEVEL_DEBUG":   1,
	"LEVEL_INFO":    2,
	"LEVEL_NOTICE":  3,
	"LEVEL_WARNING": 4,
	"LEVEL_ERROR":   5,
}

func (x ClientIssue_Level) Enum() *ClientIssue_Level {
	p := new(ClientIssue_Level)
	*p = x
	return p
}
func (x ClientIssue_Level) String() string {
	return proto.EnumName(ClientIssue_Level_name, int32(x))
}
func (x *ClientIssue_Level) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientIssue_Level_value, data, "ClientIssue_Level")
	if err != nil {
		return err
	}
	*x = ClientIssue_Level(value)
	return nil
}
func (ClientIssue_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0, 0} }

type ClientIssue_Code int32

const (
	ClientIssue_CODE_UNKNOWN             ClientIssue_Code = 0
	ClientIssue_CODE_INDEX_OUT_OF_BOUNDS ClientIssue_Code = 1
	ClientIssue_CODE_VERSION_MISMATCH    ClientIssue_Code = 2
	ClientIssue_CODE_CACHED_CHANGE       ClientIssue_Code = 3
	ClientIssue_CODE_OFFLINE_CHANGE      ClientIssue_Code = 4
	ClientIssue_CODE_CONCURRENT_CHANGE   ClientIssue_Code = 5
)

var ClientIssue_Code_name = map[int32]string{
	0: "CODE_UNKNOWN",
	1: "CODE_INDEX_OUT_OF_BOUNDS",
	2: "CODE_VERSION_MISMATCH",
	3: "CODE_CACHED_CHANGE",
	4: "CODE_OFFLINE_CHANGE",
	5: "CODE_CONCURRENT_CHANGE",
}
var ClientIssue_Code_value = map[string]int32{
	"CODE_UNKNOWN":             0,
	"CODE_INDEX_OUT_OF_BOUNDS": 1,
	"CODE_VERSION_MISMATCH":    2,
	"CODE_CACHED_CHANGE":       3,
	"CODE_OFFLINE_CHANGE":      4,
	"CODE_CONCURRENT_CHANGE":   5,
}

func (x ClientIssue_Code) Enum() *ClientIssue_Code {
	p := new(ClientIssue_Code)
	*p = x
	return p
}
func (x ClientIssue_Code) String() string {
	return proto.EnumName(ClientIssue_Code_name, int32(x))
}
func (x *ClientIssue_Code) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientIssue_Code_value, data, "ClientIssue_Code")
	if err != nil {
		return err
	}
	*x = ClientIssue_Code(value)
	return nil
}
func (ClientIssue_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0, 1} }

type ClientResolveAction_Code int32

const (
	ClientResolveAction_CODE_UNKNOWN               ClientResolveAction_Code = 0
	ClientResolveAction_CODE_NO_ACTION             ClientResolveAction_Code = 1
	ClientResolveAction_CODE_RETRY                 ClientResolveAction_Code = 2
	ClientResolveAction_CODE_RELOAD                ClientResolveAction_Code = 3
	ClientResolveAction_CODE_DISCARD_LOCAL_CHANGES ClientResolveAction_Code = 4
	ClientResolveAction_CODE_SEND_DUMP             ClientResolveAction_Code = 5
	ClientResolveAction_CODE_DISPLAY_ERROR_MESSAGE ClientResolveAction_Code = 6
)

var ClientResolveAction_Code_name = map[int32]string{
	0: "CODE_UNKNOWN",
	1: "CODE_NO_ACTION",
	2: "CODE_RETRY",
	3: "CODE_RELOAD",
	4: "CODE_DISCARD_LOCAL_CHANGES",
	5: "CODE_SEND_DUMP",
	6: "CODE_DISPLAY_ERROR_MESSAGE",
}
var ClientResolveAction_Code_value = map[string]int32{
	"CODE_UNKNOWN":               0,
	"CODE_NO_ACTION":             1,
	"CODE_RETRY":                 2,
	"CODE_RELOAD":                3,
	"CODE_DISCARD_LOCAL_CHANGES": 4,
	"CODE_SEND_DUMP":             5,
	"CODE_DISPLAY_ERROR_MESSAGE": 6,
}

func (x ClientResolveAction_Code) Enum() *ClientResolveAction_Code {
	p := new(ClientResolveAction_Code)
	*p = x
	return p
}
func (x ClientResolveAction_Code) String() string {
	return proto.EnumName(ClientResolveAction_Code_name, int32(x))
}
func (x *ClientResolveAction_Code) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientResolveAction_Code_value, data, "ClientResolveAction_Code")
	if err != nil {
		return err
	}
	*x = ClientResolveAction_Code(value)
	return nil
}
func (ClientResolveAction_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1, 0} }

type ClientResolveAction_Initiator int32

const (
	ClientResolveAction_INITIATOR_UNKNOWN ClientResolveAction_Initiator = 0
	ClientResolveAction_INITIATOR_SERVER  ClientResolveAction_Initiator = 1
	ClientResolveAction_INITIATOR_CLIENT  ClientResolveAction_Initiator = 2
)

var ClientResolveAction_Initiator_name = map[int32]string{
	0: "INITIATOR_UNKNOWN",
	1: "INITIATOR_SERVER",
	2: "INITIATOR_CLIENT",
}
var ClientResolveAction_Initiator_value = map[string]int32{
	"INITIATOR_UNKNOWN": 0,
	"INITIATOR_SERVER":  1,
	"INITIATOR_CLIENT":  2,
}

func (x ClientResolveAction_Initiator) Enum() *ClientResolveAction_Initiator {
	p := new(ClientResolveAction_Initiator)
	*p = x
	return p
}
func (x ClientResolveAction_Initiator) String() string {
	return proto.EnumName(ClientResolveAction_Initiator_name, int32(x))
}
func (x *ClientResolveAction_Initiator) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientResolveAction_Initiator_value, data, "ClientResolveAction_Initiator")
	if err != nil {
		return err
	}
	*x = ClientResolveAction_Initiator(value)
	return nil
}
func (ClientResolveAction_Initiator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor10, []int{1, 1}
}

type ClientIssue struct {
	Level            *ClientIssue_Level `protobuf:"varint,1,opt,name=level,enum=Spotify.ClientIssue_Level" json:"level,omitempty"`
	Code             *ClientIssue_Code  `protobuf:"varint,2,opt,name=code,enum=Spotify.ClientIssue_Code" json:"code,omitempty"`
	RepeatCount      *int32             `protobuf:"varint,3,opt,name=repeatCount" json:"repeatCount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ClientIssue) Reset()                    { *m = ClientIssue{} }
func (m *ClientIssue) String() string            { return proto.CompactTextString(m) }
func (*ClientIssue) ProtoMessage()               {}
func (*ClientIssue) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *ClientIssue) GetLevel() ClientIssue_Level {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return ClientIssue_LEVEL_UNKNOWN
}

func (m *ClientIssue) GetCode() ClientIssue_Code {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ClientIssue_CODE_UNKNOWN
}

func (m *ClientIssue) GetRepeatCount() int32 {
	if m != nil && m.RepeatCount != nil {
		return *m.RepeatCount
	}
	return 0
}

type ClientResolveAction struct {
	Code             *ClientResolveAction_Code      `protobuf:"varint,1,opt,name=code,enum=Spotify.ClientResolveAction_Code" json:"code,omitempty"`
	Initiator        *ClientResolveAction_Initiator `protobuf:"varint,2,opt,name=initiator,enum=Spotify.ClientResolveAction_Initiator" json:"initiator,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *ClientResolveAction) Reset()                    { *m = ClientResolveAction{} }
func (m *ClientResolveAction) String() string            { return proto.CompactTextString(m) }
func (*ClientResolveAction) ProtoMessage()               {}
func (*ClientResolveAction) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *ClientResolveAction) GetCode() ClientResolveAction_Code {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ClientResolveAction_CODE_UNKNOWN
}

func (m *ClientResolveAction) GetInitiator() ClientResolveAction_Initiator {
	if m != nil && m.Initiator != nil {
		return *m.Initiator
	}
	return ClientResolveAction_INITIATOR_UNKNOWN
}

func init() {
	proto.RegisterType((*ClientIssue)(nil), "Spotify.ClientIssue")
	proto.RegisterType((*ClientResolveAction)(nil), "Spotify.ClientResolveAction")
	proto.RegisterEnum("Spotify.ClientIssue_Level", ClientIssue_Level_name, ClientIssue_Level_value)
	proto.RegisterEnum("Spotify.ClientIssue_Code", ClientIssue_Code_name, ClientIssue_Code_value)
	proto.RegisterEnum("Spotify.ClientResolveAction_Code", ClientResolveAction_Code_name, ClientResolveAction_Code_value)
	proto.RegisterEnum("Spotify.ClientResolveAction_Initiator", ClientResolveAction_Initiator_name, ClientResolveAction_Initiator_value)
}

func init() { proto.RegisterFile("playlist4issues.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0xdb, 0x80, 0x76, 0x0a, 0xc5, 0xf3, 0xe8, 0xe8, 0x2a, 0x84, 0x4a, 0x2e, 0xd0,
	0x6e, 0xa8, 0x10, 0x82, 0x07, 0xf0, 0x6c, 0xb7, 0xb5, 0x48, 0x8f, 0x27, 0x27, 0xe9, 0xd8, 0x55,
	0x54, 0x6d, 0x41, 0x8a, 0x14, 0x35, 0xa5, 0xcd, 0x26, 0xed, 0x05, 0x78, 0x0c, 0x24, 0x24, 0x1e,
	0x14, 0xc5, 0x49, 0xd7, 0x0c, 0x01, 0x97, 0xfe, 0xce, 0xff, 0x1f, 0x1f, 0xfd, 0x3f, 0xf4, 0xd7,
	0xd9, 0xf2, 0x2e, 0x4b, 0xb7, 0xc5, 0xc7, 0x74, 0xbb, 0xbd, 0x49, 0xb6, 0xe3, 0xf5, 0x26, 0x2f,
	0x72, 0xfa, 0x24, 0x58, 0xe7, 0x45, 0xfa, 0xf5, 0xce, 0xfb, 0xd5, 0x86, 0x2e, 0xcf, 0xd2, 0x64,
	0x55, 0xa8, 0x72, 0x4e, 0xdf, 0x83, 0x9b, 0x25, 0xb7, 0x49, 0x36, 0x70, 0x46, 0xce, 0x69, 0xef,
	0xc3, 0x70, 0x5c, 0x0b, 0xc7, 0x0d, 0xd1, 0xd8, 0x2f, 0x15, 0xa6, 0x12, 0xd2, 0x77, 0xd0, 0xb9,
	0xca, 0xaf, 0x93, 0x41, 0xcb, 0x1a, 0x4e, 0xfe, 0x6a, 0xe0, 0xf9, 0x75, 0x62, 0xac, 0x8c, 0x8e,
	0xa0, 0xbb, 0x49, 0xd6, 0xc9, 0xb2, 0xe0, 0xf9, 0xcd, 0xaa, 0x18, 0xb4, 0x47, 0xce, 0xa9, 0x6b,
	0x9a, 0xc8, 0xfb, 0x06, 0xae, 0xfd, 0x80, 0x1e, 0xc2, 0x33, 0x5f, 0x2e, 0xa4, 0x1f, 0x47, 0xf8,
	0x19, 0xf5, 0x05, 0x92, 0x47, 0xf4, 0x39, 0x74, 0x2b, 0x24, 0xe4, 0x59, 0x34, 0x25, 0x0e, 0xed,
	0x01, 0x54, 0x40, 0xe1, 0x44, 0x93, 0x16, 0x25, 0xf0, 0xb4, 0x7a, 0xa3, 0x0e, 0x15, 0x97, 0xa4,
	0xbd, 0xdf, 0x72, 0xc1, 0x0c, 0x2a, 0x9c, 0x92, 0xce, 0x7e, 0x8b, 0x34, 0x46, 0x1b, 0xe2, 0x7a,
	0x3f, 0x1c, 0xe8, 0x94, 0x37, 0x96, 0x76, 0xae, 0x85, 0x6c, 0xfc, 0xf8, 0x0a, 0x06, 0x96, 0x28,
	0x14, 0xf2, 0x4b, 0xac, 0xa3, 0x30, 0xd6, 0x93, 0xf8, 0x4c, 0x47, 0x28, 0x02, 0xe2, 0xd0, 0x13,
	0xe8, 0xdb, 0xe9, 0x42, 0x9a, 0x40, 0x69, 0x8c, 0xe7, 0x2a, 0x98, 0xb3, 0x90, 0xcf, 0x48, 0x8b,
	0x1e, 0x03, 0xb5, 0x23, 0xce, 0xf8, 0x4c, 0x8a, 0x98, 0xcf, 0x18, 0x4e, 0xcb, 0x7b, 0x5e, 0xc2,
	0x91, 0xe5, 0x7a, 0x32, 0xf1, 0x15, 0xca, 0xdd, 0xa0, 0x43, 0x87, 0x70, 0x5c, 0x19, 0x34, 0xf2,
	0xc8, 0x18, 0x89, 0xe1, 0x6e, 0xe6, 0x7a, 0xdf, 0xdb, 0x70, 0x54, 0x05, 0x6a, 0x92, 0x6d, 0x9e,
	0xdd, 0x26, 0xec, 0xaa, 0x48, 0xf3, 0x15, 0xfd, 0x54, 0x87, 0x5f, 0xb5, 0xf5, 0xe6, 0x8f, 0xf0,
	0x1f, 0x68, 0x9b, 0x25, 0x08, 0x38, 0x48, 0x57, 0x69, 0x91, 0x2e, 0x8b, 0x7c, 0x53, 0x17, 0xf7,
	0xf6, 0xbf, 0x5e, 0xb5, 0x53, 0x9b, 0xbd, 0xd1, 0xfb, 0xf9, 0xef, 0xd4, 0x28, 0xf4, 0x2c, 0x41,
	0x1d, 0x33, 0x1e, 0x2a, 0x8d, 0x55, 0x55, 0x96, 0x19, 0x19, 0x9a, 0x4b, 0xd2, 0x2a, 0x5b, 0xa8,
	0xdf, 0xbe, 0x66, 0x82, 0xb4, 0xe9, 0x6b, 0x18, 0x5a, 0x20, 0x54, 0xc0, 0x99, 0x11, 0xb1, 0xaf,
	0x39, 0xf3, 0xeb, 0x0c, 0x02, 0xd2, 0xb9, 0x5f, 0x1a, 0x48, 0x14, 0xb1, 0x88, 0xe6, 0xe7, 0xc4,
	0x6d, 0x7a, 0xce, 0x7d, 0x76, 0x59, 0x35, 0x1a, 0xcf, 0x65, 0x10, 0xb0, 0xa9, 0x24, 0x8f, 0x3d,
	0x84, 0x83, 0xfb, 0xdb, 0x69, 0x1f, 0x0e, 0x15, 0xaa, 0x50, 0xb1, 0x50, 0x9b, 0xc6, 0xb1, 0x2f,
	0x80, 0xec, 0x71, 0x20, 0xcd, 0x42, 0x1a, 0xe2, 0x3c, 0xa4, 0xdc, 0x57, 0x12, 0x43, 0xd2, 0xfa,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x85, 0xa5, 0x8d, 0x1d, 0x50, 0x03, 0x00, 0x00,
}
