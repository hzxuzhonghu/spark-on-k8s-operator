// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/customer_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Set of errors that are related to requests dealing with Customer.
// Next id: 26
type CustomerErrorEnum_CustomerError int32

const (
	// Enum unspecified.
	CustomerErrorEnum_UNSPECIFIED CustomerErrorEnum_CustomerError = 0
	// The received error code is not known in this version.
	CustomerErrorEnum_UNKNOWN CustomerErrorEnum_CustomerError = 1
	// Customer status is not allowed to be changed from DRAFT and CLOSED.
	// Currency code and at least one of country code and time zone needs to be
	// set when status is changed to ENABLED.
	CustomerErrorEnum_STATUS_CHANGE_DISALLOWED CustomerErrorEnum_CustomerError = 2
	// CustomerService cannot get a customer that has not been fully set up.
	CustomerErrorEnum_ACCOUNT_NOT_SET_UP CustomerErrorEnum_CustomerError = 3
)

var CustomerErrorEnum_CustomerError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STATUS_CHANGE_DISALLOWED",
	3: "ACCOUNT_NOT_SET_UP",
}
var CustomerErrorEnum_CustomerError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"STATUS_CHANGE_DISALLOWED": 2,
	"ACCOUNT_NOT_SET_UP":       3,
}

func (x CustomerErrorEnum_CustomerError) String() string {
	return proto.EnumName(CustomerErrorEnum_CustomerError_name, int32(x))
}
func (CustomerErrorEnum_CustomerError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_customer_error_19eadbf370049eef, []int{0, 0}
}

// Container for enum describing possible customer errors.
type CustomerErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerErrorEnum) Reset()         { *m = CustomerErrorEnum{} }
func (m *CustomerErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerErrorEnum) ProtoMessage()    {}
func (*CustomerErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_error_19eadbf370049eef, []int{0}
}
func (m *CustomerErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerErrorEnum.Unmarshal(m, b)
}
func (m *CustomerErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CustomerErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerErrorEnum.Merge(dst, src)
}
func (m *CustomerErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerErrorEnum.Size(m)
}
func (m *CustomerErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomerErrorEnum)(nil), "google.ads.googleads.v1.errors.CustomerErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.CustomerErrorEnum_CustomerError", CustomerErrorEnum_CustomerError_name, CustomerErrorEnum_CustomerError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/customer_error.proto", fileDescriptor_customer_error_19eadbf370049eef)
}

var fileDescriptor_customer_error_19eadbf370049eef = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x07, 0x0a, 0x19, 0x62, 0xcd, 0x41, 0x44, 0xc6, 0x0e, 0x7d, 0x80, 0x94, 0xb2,
	0x5b, 0x3c, 0x65, 0x6d, 0x9c, 0xc3, 0x91, 0x16, 0xda, 0x6e, 0x20, 0x85, 0x52, 0xd7, 0x12, 0x06,
	0x5b, 0x32, 0x92, 0x6e, 0xf8, 0x3c, 0x1e, 0x7d, 0x14, 0x1f, 0x45, 0xf0, 0x1d, 0xa4, 0xcd, 0x56,
	0xd8, 0x41, 0x4f, 0xf9, 0xf2, 0xe7, 0xf7, 0x7d, 0xf9, 0xf2, 0x07, 0x63, 0x2e, 0x25, 0xdf, 0x54,
	0x6e, 0x51, 0x6a, 0xd7, 0xc8, 0x46, 0x1d, 0x3c, 0xb7, 0x52, 0x4a, 0x2a, 0xed, 0xae, 0xf6, 0xba,
	0x96, 0xdb, 0x4a, 0xe5, 0xed, 0x1d, 0xed, 0x94, 0xac, 0x25, 0x1c, 0x19, 0x12, 0x15, 0xa5, 0x46,
	0x9d, 0x09, 0x1d, 0x3c, 0x64, 0x4c, 0x0f, 0xc3, 0x53, 0xe8, 0x6e, 0xed, 0x16, 0x42, 0xc8, 0xba,
	0xa8, 0xd7, 0x52, 0x68, 0xe3, 0x76, 0xde, 0xc1, 0xad, 0x7f, 0x4c, 0xa5, 0x0d, 0x4f, 0xc5, 0x7e,
	0xeb, 0xac, 0xc0, 0xf5, 0xd9, 0x10, 0xde, 0x80, 0x41, 0xca, 0xe2, 0x88, 0xfa, 0xb3, 0xa7, 0x19,
	0x0d, 0xec, 0x0b, 0x38, 0x00, 0x57, 0x29, 0x7b, 0x61, 0xe1, 0x92, 0xd9, 0x3d, 0x38, 0x04, 0xf7,
	0x71, 0x42, 0x92, 0x34, 0xce, 0xfd, 0x67, 0xc2, 0xa6, 0x34, 0x0f, 0x66, 0x31, 0x99, 0xcf, 0xc3,
	0x25, 0x0d, 0x6c, 0x0b, 0xde, 0x01, 0x48, 0x7c, 0x3f, 0x4c, 0x59, 0x92, 0xb3, 0x30, 0xc9, 0x63,
	0x9a, 0xe4, 0x69, 0x64, 0xf7, 0x27, 0x3f, 0x3d, 0xe0, 0xac, 0xe4, 0x16, 0xfd, 0x5f, 0x7f, 0x02,
	0xcf, 0x9a, 0x44, 0x4d, 0xe9, 0xa8, 0xf7, 0x1a, 0x1c, 0x5d, 0x5c, 0x6e, 0x0a, 0xc1, 0x91, 0x54,
	0xdc, 0xe5, 0x95, 0x68, 0xbf, 0x74, 0xda, 0xdc, 0x6e, 0xad, 0xff, 0x5a, 0xe4, 0xa3, 0x39, 0x3e,
	0xac, 0xfe, 0x94, 0x90, 0x4f, 0x6b, 0x34, 0x35, 0x61, 0xa4, 0xd4, 0xc8, 0xc8, 0x46, 0x2d, 0x3c,
	0xd4, 0x3e, 0xa9, 0xbf, 0x4e, 0x40, 0x46, 0x4a, 0x9d, 0x75, 0x40, 0xb6, 0xf0, 0x32, 0x03, 0x7c,
	0x5b, 0x8e, 0x99, 0x62, 0x4c, 0x4a, 0x8d, 0x71, 0x87, 0x60, 0xbc, 0xf0, 0x30, 0x36, 0xd0, 0xdb,
	0x65, 0xdb, 0x6e, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x03, 0x84, 0xf7, 0x6d, 0xe5, 0x01, 0x00,
	0x00,
}
