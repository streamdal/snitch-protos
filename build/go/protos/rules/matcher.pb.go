// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: rules/matcher.proto

package rules

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

type MatchType int32

const (
	MatchType_MATCH_TYPE_UNKNOWN             MatchType = 0
	MatchType_MATCH_TYPE_IS_EMPTY            MatchType = 1000
	MatchType_MATCH_TYPE_HAS_FIELD           MatchType = 1001
	MatchType_MATCH_TYPE_IS_TYPE             MatchType = 1002
	MatchType_MATCH_TYPE_STRING_CONTAINS_ANY MatchType = 1003
	MatchType_MATCH_TYPE_STRING_CONTAINS_ALL MatchType = 1004
	MatchType_MATCH_TYPE_STRING_EQUAL        MatchType = 1005
	MatchType_MATCH_TYPE_IPV4_ADDRESS        MatchType = 1006
	MatchType_MATCH_TYPE_IPV6_ADDRESS        MatchType = 1007
	MatchType_MATCH_TYPE_MAC_ADDRESS         MatchType = 1008
	MatchType_MATCH_TYPE_REGEX               MatchType = 1009
	MatchType_MATCH_TYPE_TIMESTAMP_RFC3339   MatchType = 1010
	MatchType_MATCH_TYPE_TIMESTAMP_UNIX_NANO MatchType = 1011
	MatchType_MATCH_TYPE_TIMESTAMP_UNIX      MatchType = 1012
	MatchType_MATCH_TYPE_BOOLEAN_TRUE        MatchType = 1013
	MatchType_MATCH_TYPE_BOOLEAN_FALSE       MatchType = 1014
	MatchType_MATCH_TYPE_UUID                MatchType = 1015
	// / Payloads containing values with any PII - runs all PII matchers
	MatchType_MATCH_TYPE_PII_ANY MatchType = 2000
	// Payloads containing values with a credit card number
	MatchType_MATCH_TYPE_PII_CREDIT_CARD MatchType = 2001
	// Payloads containing values with a social security number
	MatchType_MATCH_TYPE_PII_SSN MatchType = 2002
	// Payloads containing values with an email address
	MatchType_MATCH_TYPE_PII_EMAIL MatchType = 2003
	// Payloads containing values with a phone number
	MatchType_MATCH_TYPE_PII_PHONE MatchType = 2004
	// Payloads containing values with a driver's license
	MatchType_MATCH_TYPE_PII_DRIVER_LICENSE MatchType = 2005
	// Payloads containing values with a passport ID
	MatchType_MATCH_TYPE_PII_PASSPORT_ID MatchType = 2006
	// Payloads containing values with a VIN number
	MatchType_MATCH_TYPE_PII_VIN_NUMBER MatchType = 2007
	// Payloads containing values with various serial number formats
	MatchType_MATCH_TYPE_PII_SERIAL_NUMBER MatchType = 2008
	// Payloads containing fields named "login", "username", "user", "userid", "user_id", "user", "password", "pass", "passwd", "pwd"
	MatchType_MATCH_TYPE_PII_LOGIN MatchType = 2009
	// Payloads containing fields named "taxpayer_id", "tax_id", "taxpayerid", "taxid"
	MatchType_MATCH_TYPE_PII_TAXPAYER_ID MatchType = 2010
	// Payloads containing fields named "address", "street", "city", "state", "zip", "zipcode", "zip_code", "country"
	MatchType_MATCH_TYPE_PII_ADDRESS MatchType = 2011
	// Payloads containing fields named "signature", "signature_image", "signature_image_url", "signature_image_uri"
	MatchType_MATCH_TYPE_PII_SIGNATURE MatchType = 2012
	// Payloads containing values that contain GPS data or coordinates like "lat", "lon", "latitude", "longitude"
	MatchType_MATCH_TYPE_PII_GEOLOCATION MatchType = 2013
	// Payloads containing fields like "school", "university", "college", "education"
	MatchType_MATCH_TYPE_PII_EDUCATION MatchType = 2014
	// Payloads containing fields like "account", "bank", "credit", "debit", "financial", "finance"
	MatchType_MATCH_TYPE_PII_FINANCIAL MatchType = 2015
	// Payloads containing fields like "patient", "health", "healthcare", "health care", "medical"
	MatchType_MATCH_TYPE_PII_HEALTH            MatchType = 2016
	MatchType_MATCH_TYPE_NUMERIC_EQUAL_TO      MatchType = 3000
	MatchType_MATCH_TYPE_NUMERIC_GREATER_THAN  MatchType = 3001
	MatchType_MATCH_TYPE_NUMERIC_GREATER_EQUAL MatchType = 3002
	MatchType_MATCH_TYPE_NUMERIC_LESS_THAN     MatchType = 3003
	MatchType_MATCH_TYPE_NUMERIC_LESS_EQUAL    MatchType = 3004
)

// Enum value maps for MatchType.
var (
	MatchType_name = map[int32]string{
		0:    "MATCH_TYPE_UNKNOWN",
		1000: "MATCH_TYPE_IS_EMPTY",
		1001: "MATCH_TYPE_HAS_FIELD",
		1002: "MATCH_TYPE_IS_TYPE",
		1003: "MATCH_TYPE_STRING_CONTAINS_ANY",
		1004: "MATCH_TYPE_STRING_CONTAINS_ALL",
		1005: "MATCH_TYPE_STRING_EQUAL",
		1006: "MATCH_TYPE_IPV4_ADDRESS",
		1007: "MATCH_TYPE_IPV6_ADDRESS",
		1008: "MATCH_TYPE_MAC_ADDRESS",
		1009: "MATCH_TYPE_REGEX",
		1010: "MATCH_TYPE_TIMESTAMP_RFC3339",
		1011: "MATCH_TYPE_TIMESTAMP_UNIX_NANO",
		1012: "MATCH_TYPE_TIMESTAMP_UNIX",
		1013: "MATCH_TYPE_BOOLEAN_TRUE",
		1014: "MATCH_TYPE_BOOLEAN_FALSE",
		1015: "MATCH_TYPE_UUID",
		2000: "MATCH_TYPE_PII_ANY",
		2001: "MATCH_TYPE_PII_CREDIT_CARD",
		2002: "MATCH_TYPE_PII_SSN",
		2003: "MATCH_TYPE_PII_EMAIL",
		2004: "MATCH_TYPE_PII_PHONE",
		2005: "MATCH_TYPE_PII_DRIVER_LICENSE",
		2006: "MATCH_TYPE_PII_PASSPORT_ID",
		2007: "MATCH_TYPE_PII_VIN_NUMBER",
		2008: "MATCH_TYPE_PII_SERIAL_NUMBER",
		2009: "MATCH_TYPE_PII_LOGIN",
		2010: "MATCH_TYPE_PII_TAXPAYER_ID",
		2011: "MATCH_TYPE_PII_ADDRESS",
		2012: "MATCH_TYPE_PII_SIGNATURE",
		2013: "MATCH_TYPE_PII_GEOLOCATION",
		2014: "MATCH_TYPE_PII_EDUCATION",
		2015: "MATCH_TYPE_PII_FINANCIAL",
		2016: "MATCH_TYPE_PII_HEALTH",
		3000: "MATCH_TYPE_NUMERIC_EQUAL_TO",
		3001: "MATCH_TYPE_NUMERIC_GREATER_THAN",
		3002: "MATCH_TYPE_NUMERIC_GREATER_EQUAL",
		3003: "MATCH_TYPE_NUMERIC_LESS_THAN",
		3004: "MATCH_TYPE_NUMERIC_LESS_EQUAL",
	}
	MatchType_value = map[string]int32{
		"MATCH_TYPE_UNKNOWN":               0,
		"MATCH_TYPE_IS_EMPTY":              1000,
		"MATCH_TYPE_HAS_FIELD":             1001,
		"MATCH_TYPE_IS_TYPE":               1002,
		"MATCH_TYPE_STRING_CONTAINS_ANY":   1003,
		"MATCH_TYPE_STRING_CONTAINS_ALL":   1004,
		"MATCH_TYPE_STRING_EQUAL":          1005,
		"MATCH_TYPE_IPV4_ADDRESS":          1006,
		"MATCH_TYPE_IPV6_ADDRESS":          1007,
		"MATCH_TYPE_MAC_ADDRESS":           1008,
		"MATCH_TYPE_REGEX":                 1009,
		"MATCH_TYPE_TIMESTAMP_RFC3339":     1010,
		"MATCH_TYPE_TIMESTAMP_UNIX_NANO":   1011,
		"MATCH_TYPE_TIMESTAMP_UNIX":        1012,
		"MATCH_TYPE_BOOLEAN_TRUE":          1013,
		"MATCH_TYPE_BOOLEAN_FALSE":         1014,
		"MATCH_TYPE_UUID":                  1015,
		"MATCH_TYPE_PII_ANY":               2000,
		"MATCH_TYPE_PII_CREDIT_CARD":       2001,
		"MATCH_TYPE_PII_SSN":               2002,
		"MATCH_TYPE_PII_EMAIL":             2003,
		"MATCH_TYPE_PII_PHONE":             2004,
		"MATCH_TYPE_PII_DRIVER_LICENSE":    2005,
		"MATCH_TYPE_PII_PASSPORT_ID":       2006,
		"MATCH_TYPE_PII_VIN_NUMBER":        2007,
		"MATCH_TYPE_PII_SERIAL_NUMBER":     2008,
		"MATCH_TYPE_PII_LOGIN":             2009,
		"MATCH_TYPE_PII_TAXPAYER_ID":       2010,
		"MATCH_TYPE_PII_ADDRESS":           2011,
		"MATCH_TYPE_PII_SIGNATURE":         2012,
		"MATCH_TYPE_PII_GEOLOCATION":       2013,
		"MATCH_TYPE_PII_EDUCATION":         2014,
		"MATCH_TYPE_PII_FINANCIAL":         2015,
		"MATCH_TYPE_PII_HEALTH":            2016,
		"MATCH_TYPE_NUMERIC_EQUAL_TO":      3000,
		"MATCH_TYPE_NUMERIC_GREATER_THAN":  3001,
		"MATCH_TYPE_NUMERIC_GREATER_EQUAL": 3002,
		"MATCH_TYPE_NUMERIC_LESS_THAN":     3003,
		"MATCH_TYPE_NUMERIC_LESS_EQUAL":    3004,
	}
)

func (x MatchType) Enum() *MatchType {
	p := new(MatchType)
	*p = x
	return p
}

func (x MatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_rules_matcher_proto_enumTypes[0].Descriptor()
}

func (MatchType) Type() protoreflect.EnumType {
	return &file_rules_matcher_proto_enumTypes[0]
}

func (x MatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchType.Descriptor instead.
func (MatchType) EnumDescriptor() ([]byte, []int) {
	return file_rules_matcher_proto_rawDescGZIP(), []int{0}
}

type MatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []byte    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Path   string    `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Args   []string  `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"` // args determined by match_type
	Negate bool      `protobuf:"varint,4,opt,name=negate,proto3" json:"negate,omitempty"`
	Type   MatchType `protobuf:"varint,5,opt,name=type,proto3,enum=rules.MatchType" json:"type,omitempty"`
}

func (x *MatchRequest) Reset() {
	*x = MatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rules_matcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest) ProtoMessage() {}

func (x *MatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rules_matcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest.ProtoReflect.Descriptor instead.
func (*MatchRequest) Descriptor() ([]byte, []int) {
	return file_rules_matcher_proto_rawDescGZIP(), []int{0}
}

func (x *MatchRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MatchRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MatchRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *MatchRequest) GetNegate() bool {
	if x != nil {
		return x.Negate
	}
	return false
}

func (x *MatchRequest) GetType() MatchType {
	if x != nil {
		return x.Type
	}
	return MatchType_MATCH_TYPE_UNKNOWN
}

type MatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsMatch bool   `protobuf:"varint,1,opt,name=is_match,json=isMatch,proto3" json:"is_match,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MatchResponse) Reset() {
	*x = MatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rules_matcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResponse) ProtoMessage() {}

func (x *MatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rules_matcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResponse.ProtoReflect.Descriptor instead.
func (*MatchResponse) Descriptor() ([]byte, []int) {
	return file_rules_matcher_proto_rawDescGZIP(), []int{1}
}

func (x *MatchResponse) GetIsMatch() bool {
	if x != nil {
		return x.IsMatch
	}
	return false
}

func (x *MatchResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_rules_matcher_proto protoreflect.FileDescriptor

var file_rules_matcher_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x88, 0x01, 0x0a,
	0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x40, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0xbe, 0x09, 0x0a, 0x09, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x13, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x53,
	0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0xe8, 0x07, 0x12, 0x19, 0x0a, 0x14, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x10, 0xe9, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0xea, 0x07, 0x12, 0x23, 0x0a,
	0x1e, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x41, 0x4e, 0x59, 0x10,
	0xeb, 0x07, 0x12, 0x23, 0x0a, 0x1e, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53,
	0x5f, 0x41, 0x4c, 0x4c, 0x10, 0xec, 0x07, 0x12, 0x1c, 0x0a, 0x17, 0x4d, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x10, 0xed, 0x07, 0x12, 0x1c, 0x0a, 0x17, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x50, 0x56, 0x34, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53,
	0x10, 0xee, 0x07, 0x12, 0x1c, 0x0a, 0x17, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x50, 0x56, 0x36, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0xef,
	0x07, 0x12, 0x1b, 0x0a, 0x16, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4d, 0x41, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0xf0, 0x07, 0x12, 0x15,
	0x0a, 0x10, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x47,
	0x45, 0x58, 0x10, 0xf1, 0x07, 0x12, 0x21, 0x0a, 0x1c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x52, 0x46,
	0x43, 0x33, 0x33, 0x33, 0x39, 0x10, 0xf2, 0x07, 0x12, 0x23, 0x0a, 0x1e, 0x4d, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50,
	0x5f, 0x55, 0x4e, 0x49, 0x58, 0x5f, 0x4e, 0x41, 0x4e, 0x4f, 0x10, 0xf3, 0x07, 0x12, 0x1e, 0x0a,
	0x19, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x55, 0x4e, 0x49, 0x58, 0x10, 0xf4, 0x07, 0x12, 0x1c, 0x0a,
	0x17, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c,
	0x45, 0x41, 0x4e, 0x5f, 0x54, 0x52, 0x55, 0x45, 0x10, 0xf5, 0x07, 0x12, 0x1d, 0x0a, 0x18, 0x4d,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41,
	0x4e, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0xf6, 0x07, 0x12, 0x14, 0x0a, 0x0f, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x10, 0xf7, 0x07,
	0x12, 0x17, 0x0a, 0x12, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x49, 0x49, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0xd0, 0x0f, 0x12, 0x1f, 0x0a, 0x1a, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x43, 0x52, 0x45, 0x44,
	0x49, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0xd1, 0x0f, 0x12, 0x17, 0x0a, 0x12, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x53, 0x53, 0x4e,
	0x10, 0xd2, 0x0f, 0x12, 0x19, 0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0xd3, 0x0f, 0x12, 0x19,
	0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49,
	0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0xd4, 0x0f, 0x12, 0x22, 0x0a, 0x1d, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x44, 0x52, 0x49, 0x56,
	0x45, 0x52, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0xd5, 0x0f, 0x12, 0x1f, 0x0a,
	0x1a, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f,
	0x50, 0x41, 0x53, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x49, 0x44, 0x10, 0xd6, 0x0f, 0x12, 0x1e,
	0x0a, 0x19, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49,
	0x5f, 0x56, 0x49, 0x4e, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0xd7, 0x0f, 0x12, 0x21,
	0x0a, 0x1c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49,
	0x5f, 0x53, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0xd8,
	0x0f, 0x12, 0x19, 0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x49, 0x49, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0xd9, 0x0f, 0x12, 0x1f, 0x0a, 0x1a,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x54,
	0x41, 0x58, 0x50, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0xda, 0x0f, 0x12, 0x1b, 0x0a,
	0x16, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f,
	0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0xdb, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x53, 0x49, 0x47,
	0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0xdc, 0x0f, 0x12, 0x1f, 0x0a, 0x1a, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x47, 0x45, 0x4f, 0x4c,
	0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xdd, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x45, 0x44, 0x55,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xde, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x46, 0x49, 0x4e, 0x41,
	0x4e, 0x43, 0x49, 0x41, 0x4c, 0x10, 0xdf, 0x0f, 0x12, 0x1a, 0x0a, 0x15, 0x4d, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x49, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54,
	0x48, 0x10, 0xe0, 0x0f, 0x12, 0x20, 0x0a, 0x1b, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c,
	0x5f, 0x54, 0x4f, 0x10, 0xb8, 0x17, 0x12, 0x24, 0x0a, 0x1f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x47, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0xb9, 0x17, 0x12, 0x25, 0x0a, 0x20,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52,
	0x49, 0x43, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c,
	0x10, 0xba, 0x17, 0x12, 0x21, 0x0a, 0x1c, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54,
	0x48, 0x41, 0x4e, 0x10, 0xbb, 0x17, 0x12, 0x22, 0x0a, 0x1d, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x4c, 0x45, 0x53,
	0x53, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0xbc, 0x17, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64,
	0x61, 0x6c, 0x2f, 0x73, 0x6e, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rules_matcher_proto_rawDescOnce sync.Once
	file_rules_matcher_proto_rawDescData = file_rules_matcher_proto_rawDesc
)

func file_rules_matcher_proto_rawDescGZIP() []byte {
	file_rules_matcher_proto_rawDescOnce.Do(func() {
		file_rules_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_rules_matcher_proto_rawDescData)
	})
	return file_rules_matcher_proto_rawDescData
}

var file_rules_matcher_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rules_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rules_matcher_proto_goTypes = []interface{}{
	(MatchType)(0),        // 0: rules.MatchType
	(*MatchRequest)(nil),  // 1: rules.MatchRequest
	(*MatchResponse)(nil), // 2: rules.MatchResponse
}
var file_rules_matcher_proto_depIdxs = []int32{
	0, // 0: rules.MatchRequest.type:type_name -> rules.MatchType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rules_matcher_proto_init() }
func file_rules_matcher_proto_init() {
	if File_rules_matcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rules_matcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRequest); i {
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
		file_rules_matcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchResponse); i {
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
			RawDescriptor: file_rules_matcher_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rules_matcher_proto_goTypes,
		DependencyIndexes: file_rules_matcher_proto_depIdxs,
		EnumInfos:         file_rules_matcher_proto_enumTypes,
		MessageInfos:      file_rules_matcher_proto_msgTypes,
	}.Build()
	File_rules_matcher_proto = out.File
	file_rules_matcher_proto_rawDesc = nil
	file_rules_matcher_proto_goTypes = nil
	file_rules_matcher_proto_depIdxs = nil
}
