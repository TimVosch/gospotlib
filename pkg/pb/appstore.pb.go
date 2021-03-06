// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: appstore.proto

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

type RequestHeader_Platform int32

const (
	RequestHeader_WIN32_X86      RequestHeader_Platform = 0
	RequestHeader_OSX_X86        RequestHeader_Platform = 1
	RequestHeader_LINUX_X86      RequestHeader_Platform = 2
	RequestHeader_IPHONE_ARM     RequestHeader_Platform = 3
	RequestHeader_SYMBIANS60_ARM RequestHeader_Platform = 4
	RequestHeader_OSX_POWERPC    RequestHeader_Platform = 5
	RequestHeader_ANDROID_ARM    RequestHeader_Platform = 6
	RequestHeader_WINCE_ARM      RequestHeader_Platform = 7
	RequestHeader_LINUX_X86_64   RequestHeader_Platform = 8
	RequestHeader_OSX_X86_64     RequestHeader_Platform = 9
	RequestHeader_PALM_ARM       RequestHeader_Platform = 10
	RequestHeader_LINUX_SH       RequestHeader_Platform = 11
	RequestHeader_FREEBSD_X86    RequestHeader_Platform = 12
	RequestHeader_FREEBSD_X86_64 RequestHeader_Platform = 13
	RequestHeader_BLACKBERRY_ARM RequestHeader_Platform = 14
	RequestHeader_SONOS_UNKNOWN  RequestHeader_Platform = 15
	RequestHeader_LINUX_MIPS     RequestHeader_Platform = 16
	RequestHeader_LINUX_ARM      RequestHeader_Platform = 17
	RequestHeader_LOGITECH_ARM   RequestHeader_Platform = 18
	RequestHeader_LINUX_BLACKFIN RequestHeader_Platform = 19
	RequestHeader_ONKYO_ARM      RequestHeader_Platform = 21
	RequestHeader_QNXNTO_ARM     RequestHeader_Platform = 22
	RequestHeader_BADPLATFORM    RequestHeader_Platform = 255
)

// Enum value maps for RequestHeader_Platform.
var (
	RequestHeader_Platform_name = map[int32]string{
		0:   "WIN32_X86",
		1:   "OSX_X86",
		2:   "LINUX_X86",
		3:   "IPHONE_ARM",
		4:   "SYMBIANS60_ARM",
		5:   "OSX_POWERPC",
		6:   "ANDROID_ARM",
		7:   "WINCE_ARM",
		8:   "LINUX_X86_64",
		9:   "OSX_X86_64",
		10:  "PALM_ARM",
		11:  "LINUX_SH",
		12:  "FREEBSD_X86",
		13:  "FREEBSD_X86_64",
		14:  "BLACKBERRY_ARM",
		15:  "SONOS_UNKNOWN",
		16:  "LINUX_MIPS",
		17:  "LINUX_ARM",
		18:  "LOGITECH_ARM",
		19:  "LINUX_BLACKFIN",
		21:  "ONKYO_ARM",
		22:  "QNXNTO_ARM",
		255: "BADPLATFORM",
	}
	RequestHeader_Platform_value = map[string]int32{
		"WIN32_X86":      0,
		"OSX_X86":        1,
		"LINUX_X86":      2,
		"IPHONE_ARM":     3,
		"SYMBIANS60_ARM": 4,
		"OSX_POWERPC":    5,
		"ANDROID_ARM":    6,
		"WINCE_ARM":      7,
		"LINUX_X86_64":   8,
		"OSX_X86_64":     9,
		"PALM_ARM":       10,
		"LINUX_SH":       11,
		"FREEBSD_X86":    12,
		"FREEBSD_X86_64": 13,
		"BLACKBERRY_ARM": 14,
		"SONOS_UNKNOWN":  15,
		"LINUX_MIPS":     16,
		"LINUX_ARM":      17,
		"LOGITECH_ARM":   18,
		"LINUX_BLACKFIN": 19,
		"ONKYO_ARM":      21,
		"QNXNTO_ARM":     22,
		"BADPLATFORM":    255,
	}
)

func (x RequestHeader_Platform) Enum() *RequestHeader_Platform {
	p := new(RequestHeader_Platform)
	*p = x
	return p
}

func (x RequestHeader_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestHeader_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_appstore_proto_enumTypes[0].Descriptor()
}

func (RequestHeader_Platform) Type() protoreflect.EnumType {
	return &file_appstore_proto_enumTypes[0]
}

func (x RequestHeader_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestHeader_Platform) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestHeader_Platform(num)
	return nil
}

// Deprecated: Use RequestHeader_Platform.Descriptor instead.
func (RequestHeader_Platform) EnumDescriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{3, 0}
}

type RequestHeader_DeviceClass int32

const (
	RequestHeader_DESKTOP RequestHeader_DeviceClass = 1
	RequestHeader_TABLET  RequestHeader_DeviceClass = 2
	RequestHeader_MOBILE  RequestHeader_DeviceClass = 3
	RequestHeader_WEB     RequestHeader_DeviceClass = 4
	RequestHeader_TV      RequestHeader_DeviceClass = 5
)

// Enum value maps for RequestHeader_DeviceClass.
var (
	RequestHeader_DeviceClass_name = map[int32]string{
		1: "DESKTOP",
		2: "TABLET",
		3: "MOBILE",
		4: "WEB",
		5: "TV",
	}
	RequestHeader_DeviceClass_value = map[string]int32{
		"DESKTOP": 1,
		"TABLET":  2,
		"MOBILE":  3,
		"WEB":     4,
		"TV":      5,
	}
)

func (x RequestHeader_DeviceClass) Enum() *RequestHeader_DeviceClass {
	p := new(RequestHeader_DeviceClass)
	*p = x
	return p
}

func (x RequestHeader_DeviceClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestHeader_DeviceClass) Descriptor() protoreflect.EnumDescriptor {
	return file_appstore_proto_enumTypes[1].Descriptor()
}

func (RequestHeader_DeviceClass) Type() protoreflect.EnumType {
	return &file_appstore_proto_enumTypes[1]
}

func (x RequestHeader_DeviceClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RequestHeader_DeviceClass) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RequestHeader_DeviceClass(num)
	return nil
}

// Deprecated: Use RequestHeader_DeviceClass.Descriptor instead.
func (RequestHeader_DeviceClass) EnumDescriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{3, 1}
}

type AppItem_Requirement int32

const (
	AppItem_REQUIRED_INSTALL AppItem_Requirement = 1
	AppItem_LAZYLOAD         AppItem_Requirement = 2
	AppItem_OPTIONAL_INSTALL AppItem_Requirement = 3
)

// Enum value maps for AppItem_Requirement.
var (
	AppItem_Requirement_name = map[int32]string{
		1: "REQUIRED_INSTALL",
		2: "LAZYLOAD",
		3: "OPTIONAL_INSTALL",
	}
	AppItem_Requirement_value = map[string]int32{
		"REQUIRED_INSTALL": 1,
		"LAZYLOAD":         2,
		"OPTIONAL_INSTALL": 3,
	}
)

func (x AppItem_Requirement) Enum() *AppItem_Requirement {
	p := new(AppItem_Requirement)
	*p = x
	return p
}

func (x AppItem_Requirement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppItem_Requirement) Descriptor() protoreflect.EnumDescriptor {
	return file_appstore_proto_enumTypes[2].Descriptor()
}

func (AppItem_Requirement) Type() protoreflect.EnumType {
	return &file_appstore_proto_enumTypes[2]
}

func (x AppItem_Requirement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AppItem_Requirement) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AppItem_Requirement(num)
	return nil
}

// Deprecated: Use AppItem_Requirement.Descriptor instead.
func (AppItem_Requirement) EnumDescriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{4, 0}
}

type AppItem_Type int32

const (
	AppItem_APPLICATION AppItem_Type = 0
	AppItem_FRAMEWORK   AppItem_Type = 1
	AppItem_BRIDGE      AppItem_Type = 2
)

// Enum value maps for AppItem_Type.
var (
	AppItem_Type_name = map[int32]string{
		0: "APPLICATION",
		1: "FRAMEWORK",
		2: "BRIDGE",
	}
	AppItem_Type_value = map[string]int32{
		"APPLICATION": 0,
		"FRAMEWORK":   1,
		"BRIDGE":      2,
	}
)

func (x AppItem_Type) Enum() *AppItem_Type {
	p := new(AppItem_Type)
	*p = x
	return p
}

func (x AppItem_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppItem_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_appstore_proto_enumTypes[3].Descriptor()
}

func (AppItem_Type) Type() protoreflect.EnumType {
	return &file_appstore_proto_enumTypes[3]
}

func (x AppItem_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AppItem_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AppItem_Type(num)
	return nil
}

// Deprecated: Use AppItem_Type.Descriptor instead.
func (AppItem_Type) EnumDescriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{4, 1}
}

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier *string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	VersionInt *int32  `protobuf:"varint,2,opt,name=version_int,json=versionInt" json:"version_int,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{0}
}

func (x *AppInfo) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *AppInfo) GetVersionInt() int32 {
	if x != nil && x.VersionInt != nil {
		return *x.VersionInt
	}
	return 0
}

type AppInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AppInfo `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (x *AppInfoList) Reset() {
	*x = AppInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfoList) ProtoMessage() {}

func (x *AppInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfoList.ProtoReflect.Descriptor instead.
func (*AppInfoList) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{1}
}

func (x *AppInfoList) GetItems() []*AppInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type SemanticVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major *int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor *int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch *int32 `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
}

func (x *SemanticVersion) Reset() {
	*x = SemanticVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemanticVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemanticVersion) ProtoMessage() {}

func (x *SemanticVersion) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemanticVersion.ProtoReflect.Descriptor instead.
func (*SemanticVersion) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{2}
}

func (x *SemanticVersion) GetMajor() int32 {
	if x != nil && x.Major != nil {
		return *x.Major
	}
	return 0
}

func (x *SemanticVersion) GetMinor() int32 {
	if x != nil && x.Minor != nil {
		return *x.Minor
	}
	return 0
}

func (x *SemanticVersion) GetPatch() int32 {
	if x != nil && x.Patch != nil {
		return *x.Patch
	}
	return 0
}

type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market           *string                    `protobuf:"bytes,1,opt,name=market" json:"market,omitempty"`
	Platform         *RequestHeader_Platform    `protobuf:"varint,2,opt,name=platform,enum=pb.RequestHeader_Platform" json:"platform,omitempty"`
	AppInfos         *AppInfoList               `protobuf:"bytes,6,opt,name=app_infos,json=appInfos" json:"app_infos,omitempty"`
	BridgeIdentifier *string                    `protobuf:"bytes,7,opt,name=bridge_identifier,json=bridgeIdentifier" json:"bridge_identifier,omitempty"`
	BridgeVersion    *SemanticVersion           `protobuf:"bytes,8,opt,name=bridge_version,json=bridgeVersion" json:"bridge_version,omitempty"`
	DeviceClass      *RequestHeader_DeviceClass `protobuf:"varint,9,opt,name=device_class,json=deviceClass,enum=pb.RequestHeader_DeviceClass" json:"device_class,omitempty"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{3}
}

func (x *RequestHeader) GetMarket() string {
	if x != nil && x.Market != nil {
		return *x.Market
	}
	return ""
}

func (x *RequestHeader) GetPlatform() RequestHeader_Platform {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return RequestHeader_WIN32_X86
}

func (x *RequestHeader) GetAppInfos() *AppInfoList {
	if x != nil {
		return x.AppInfos
	}
	return nil
}

func (x *RequestHeader) GetBridgeIdentifier() string {
	if x != nil && x.BridgeIdentifier != nil {
		return *x.BridgeIdentifier
	}
	return ""
}

func (x *RequestHeader) GetBridgeVersion() *SemanticVersion {
	if x != nil {
		return x.BridgeVersion
	}
	return nil
}

func (x *RequestHeader) GetDeviceClass() RequestHeader_DeviceClass {
	if x != nil && x.DeviceClass != nil {
		return *x.DeviceClass
	}
	return RequestHeader_DESKTOP
}

type AppItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier    *string              `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Requirement   *AppItem_Requirement `protobuf:"varint,2,opt,name=requirement,enum=pb.AppItem_Requirement" json:"requirement,omitempty"`
	Manifest      *string              `protobuf:"bytes,4,opt,name=manifest" json:"manifest,omitempty"`
	Checksum      *string              `protobuf:"bytes,5,opt,name=checksum" json:"checksum,omitempty"`
	BundleUri     *string              `protobuf:"bytes,6,opt,name=bundle_uri,json=bundleUri" json:"bundle_uri,omitempty"`
	SmallIconUri  *string              `protobuf:"bytes,7,opt,name=small_icon_uri,json=smallIconUri" json:"small_icon_uri,omitempty"`
	LargeIconUri  *string              `protobuf:"bytes,8,opt,name=large_icon_uri,json=largeIconUri" json:"large_icon_uri,omitempty"`
	MediumIconUri *string              `protobuf:"bytes,9,opt,name=medium_icon_uri,json=mediumIconUri" json:"medium_icon_uri,omitempty"`
	BundleType    *AppItem_Type        `protobuf:"varint,10,opt,name=bundle_type,json=bundleType,enum=pb.AppItem_Type" json:"bundle_type,omitempty"`
	Version       *SemanticVersion     `protobuf:"bytes,11,opt,name=version" json:"version,omitempty"`
	TtlInSeconds  *uint32              `protobuf:"varint,12,opt,name=ttl_in_seconds,json=ttlInSeconds" json:"ttl_in_seconds,omitempty"`
	Categories    *IdentifierList      `protobuf:"bytes,13,opt,name=categories" json:"categories,omitempty"`
}

func (x *AppItem) Reset() {
	*x = AppItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppItem) ProtoMessage() {}

func (x *AppItem) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppItem.ProtoReflect.Descriptor instead.
func (*AppItem) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{4}
}

func (x *AppItem) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *AppItem) GetRequirement() AppItem_Requirement {
	if x != nil && x.Requirement != nil {
		return *x.Requirement
	}
	return AppItem_REQUIRED_INSTALL
}

func (x *AppItem) GetManifest() string {
	if x != nil && x.Manifest != nil {
		return *x.Manifest
	}
	return ""
}

func (x *AppItem) GetChecksum() string {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return ""
}

func (x *AppItem) GetBundleUri() string {
	if x != nil && x.BundleUri != nil {
		return *x.BundleUri
	}
	return ""
}

func (x *AppItem) GetSmallIconUri() string {
	if x != nil && x.SmallIconUri != nil {
		return *x.SmallIconUri
	}
	return ""
}

func (x *AppItem) GetLargeIconUri() string {
	if x != nil && x.LargeIconUri != nil {
		return *x.LargeIconUri
	}
	return ""
}

func (x *AppItem) GetMediumIconUri() string {
	if x != nil && x.MediumIconUri != nil {
		return *x.MediumIconUri
	}
	return ""
}

func (x *AppItem) GetBundleType() AppItem_Type {
	if x != nil && x.BundleType != nil {
		return *x.BundleType
	}
	return AppItem_APPLICATION
}

func (x *AppItem) GetVersion() *SemanticVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *AppItem) GetTtlInSeconds() uint32 {
	if x != nil && x.TtlInSeconds != nil {
		return *x.TtlInSeconds
	}
	return 0
}

func (x *AppItem) GetCategories() *IdentifierList {
	if x != nil {
		return x.Categories
	}
	return nil
}

type AppList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AppItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (x *AppList) Reset() {
	*x = AppList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppList) ProtoMessage() {}

func (x *AppList) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppList.ProtoReflect.Descriptor instead.
func (*AppList) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{5}
}

func (x *AppList) GetItems() []*AppItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type IdentifierList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifiers []string `protobuf:"bytes,1,rep,name=identifiers" json:"identifiers,omitempty"`
}

func (x *IdentifierList) Reset() {
	*x = IdentifierList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentifierList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentifierList) ProtoMessage() {}

func (x *IdentifierList) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentifierList.ProtoReflect.Descriptor instead.
func (*IdentifierList) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{6}
}

func (x *IdentifierList) GetIdentifiers() []string {
	if x != nil {
		return x.Identifiers
	}
	return nil
}

type BannerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json *string `protobuf:"bytes,1,opt,name=json" json:"json,omitempty"`
}

func (x *BannerConfig) Reset() {
	*x = BannerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerConfig) ProtoMessage() {}

func (x *BannerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_appstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerConfig.ProtoReflect.Descriptor instead.
func (*BannerConfig) Descriptor() ([]byte, []int) {
	return file_appstore_proto_rawDescGZIP(), []int{7}
}

func (x *BannerConfig) GetJson() string {
	if x != nil && x.Json != nil {
		return *x.Json
	}
	return ""
}

var File_appstore_proto protoreflect.FileDescriptor

var file_appstore_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x70, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x4a, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74,
	0x22, 0x30, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x53, 0x0a, 0x0f, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x22, 0x8a, 0x06, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2c, 0x0a, 0x09, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x61,
	0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x40, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x22, 0x8a, 0x03, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x0d, 0x0a, 0x09, 0x57, 0x49, 0x4e, 0x33, 0x32, 0x5f, 0x58, 0x38, 0x36, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x4f, 0x53, 0x58, 0x5f, 0x58, 0x38, 0x36, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4c,
	0x49, 0x4e, 0x55, 0x58, 0x5f, 0x58, 0x38, 0x36, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x50,
	0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x59,
	0x4d, 0x42, 0x49, 0x41, 0x4e, 0x53, 0x36, 0x30, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x04, 0x12, 0x0f,
	0x0a, 0x0b, 0x4f, 0x53, 0x58, 0x5f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x50, 0x43, 0x10, 0x05, 0x12,
	0x0f, 0x0a, 0x0b, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x06,
	0x12, 0x0d, 0x0a, 0x09, 0x57, 0x49, 0x4e, 0x43, 0x45, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x07, 0x12,
	0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x5f, 0x58, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x10,
	0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x53, 0x58, 0x5f, 0x58, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x10,
	0x09, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x4c, 0x4d, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x0a, 0x12,
	0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x5f, 0x53, 0x48, 0x10, 0x0b, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x52, 0x45, 0x45, 0x42, 0x53, 0x44, 0x5f, 0x58, 0x38, 0x36, 0x10, 0x0c, 0x12, 0x12,
	0x0a, 0x0e, 0x46, 0x52, 0x45, 0x45, 0x42, 0x53, 0x44, 0x5f, 0x58, 0x38, 0x36, 0x5f, 0x36, 0x34,
	0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x42, 0x45, 0x52, 0x52, 0x59,
	0x5f, 0x41, 0x52, 0x4d, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x4f, 0x4e, 0x4f, 0x53, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x0f, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x4e,
	0x55, 0x58, 0x5f, 0x4d, 0x49, 0x50, 0x53, 0x10, 0x10, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49, 0x4e,
	0x55, 0x58, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4f, 0x47, 0x49,
	0x54, 0x45, 0x43, 0x48, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x49,
	0x4e, 0x55, 0x58, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x46, 0x49, 0x4e, 0x10, 0x13, 0x12, 0x0d,
	0x0a, 0x09, 0x4f, 0x4e, 0x4b, 0x59, 0x4f, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x15, 0x12, 0x0e, 0x0a,
	0x0a, 0x51, 0x4e, 0x58, 0x4e, 0x54, 0x4f, 0x5f, 0x41, 0x52, 0x4d, 0x10, 0x16, 0x12, 0x10, 0x0a,
	0x0b, 0x42, 0x41, 0x44, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0xff, 0x01, 0x22,
	0x43, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x54,
	0x41, 0x42, 0x4c, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c,
	0x45, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02,
	0x54, 0x56, 0x10, 0x05, 0x22, 0xe8, 0x04, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x39, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x74,
	0x65, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x55,
	0x72, 0x69, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x63, 0x6f, 0x6e,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x49, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x72, 0x67,
	0x65, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x26,
	0x0a, 0x0f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x49,
	0x63, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x31, 0x0a, 0x0b, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x70, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x62,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x74, 0x6c, 0x5f,
	0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x74, 0x74, 0x6c, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x47, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x49, 0x4e,
	0x53, 0x54, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x41, 0x5a, 0x59, 0x4c,
	0x4f, 0x41, 0x44, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41,
	0x4c, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x10, 0x03, 0x22, 0x32, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x52, 0x49, 0x44, 0x47, 0x45, 0x10, 0x02, 0x22,
	0x2c, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x70, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x32, 0x0a,
	0x0e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x22, 0x22, 0x0a, 0x0c, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6d, 0x76, 0x6f, 0x73, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x73,
	0x70, 0x6f, 0x74, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
}

var (
	file_appstore_proto_rawDescOnce sync.Once
	file_appstore_proto_rawDescData = file_appstore_proto_rawDesc
)

func file_appstore_proto_rawDescGZIP() []byte {
	file_appstore_proto_rawDescOnce.Do(func() {
		file_appstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_appstore_proto_rawDescData)
	})
	return file_appstore_proto_rawDescData
}

var file_appstore_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_appstore_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_appstore_proto_goTypes = []interface{}{
	(RequestHeader_Platform)(0),    // 0: pb.RequestHeader.Platform
	(RequestHeader_DeviceClass)(0), // 1: pb.RequestHeader.DeviceClass
	(AppItem_Requirement)(0),       // 2: pb.AppItem.Requirement
	(AppItem_Type)(0),              // 3: pb.AppItem.Type
	(*AppInfo)(nil),                // 4: pb.AppInfo
	(*AppInfoList)(nil),            // 5: pb.AppInfoList
	(*SemanticVersion)(nil),        // 6: pb.SemanticVersion
	(*RequestHeader)(nil),          // 7: pb.RequestHeader
	(*AppItem)(nil),                // 8: pb.AppItem
	(*AppList)(nil),                // 9: pb.AppList
	(*IdentifierList)(nil),         // 10: pb.IdentifierList
	(*BannerConfig)(nil),           // 11: pb.BannerConfig
}
var file_appstore_proto_depIdxs = []int32{
	4,  // 0: pb.AppInfoList.items:type_name -> pb.AppInfo
	0,  // 1: pb.RequestHeader.platform:type_name -> pb.RequestHeader.Platform
	5,  // 2: pb.RequestHeader.app_infos:type_name -> pb.AppInfoList
	6,  // 3: pb.RequestHeader.bridge_version:type_name -> pb.SemanticVersion
	1,  // 4: pb.RequestHeader.device_class:type_name -> pb.RequestHeader.DeviceClass
	2,  // 5: pb.AppItem.requirement:type_name -> pb.AppItem.Requirement
	3,  // 6: pb.AppItem.bundle_type:type_name -> pb.AppItem.Type
	6,  // 7: pb.AppItem.version:type_name -> pb.SemanticVersion
	10, // 8: pb.AppItem.categories:type_name -> pb.IdentifierList
	8,  // 9: pb.AppList.items:type_name -> pb.AppItem
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_appstore_proto_init() }
func file_appstore_proto_init() {
	if File_appstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
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
		file_appstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfoList); i {
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
		file_appstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemanticVersion); i {
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
		file_appstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
		file_appstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppItem); i {
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
		file_appstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppList); i {
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
		file_appstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentifierList); i {
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
		file_appstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerConfig); i {
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
			RawDescriptor: file_appstore_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_appstore_proto_goTypes,
		DependencyIndexes: file_appstore_proto_depIdxs,
		EnumInfos:         file_appstore_proto_enumTypes,
		MessageInfos:      file_appstore_proto_msgTypes,
	}.Build()
	File_appstore_proto = out.File
	file_appstore_proto_rawDesc = nil
	file_appstore_proto_goTypes = nil
	file_appstore_proto_depIdxs = nil
}
