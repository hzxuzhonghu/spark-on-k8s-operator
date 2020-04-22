// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/hotel_performance_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

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

// A hotel performance view.
type HotelPerformanceView struct {
	// The resource name of the hotel performance view.
	// Hotel performance view resource names have the form:
	//
	// `customers/{customer_id}/hotelPerformanceView`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelPerformanceView) Reset()         { *m = HotelPerformanceView{} }
func (m *HotelPerformanceView) String() string { return proto.CompactTextString(m) }
func (*HotelPerformanceView) ProtoMessage()    {}
func (*HotelPerformanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_performance_view_0d02987095d2b3ce, []int{0}
}
func (m *HotelPerformanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelPerformanceView.Unmarshal(m, b)
}
func (m *HotelPerformanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelPerformanceView.Marshal(b, m, deterministic)
}
func (dst *HotelPerformanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelPerformanceView.Merge(dst, src)
}
func (m *HotelPerformanceView) XXX_Size() int {
	return xxx_messageInfo_HotelPerformanceView.Size(m)
}
func (m *HotelPerformanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelPerformanceView.DiscardUnknown(m)
}

var xxx_messageInfo_HotelPerformanceView proto.InternalMessageInfo

func (m *HotelPerformanceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*HotelPerformanceView)(nil), "google.ads.googleads.v1.resources.HotelPerformanceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/hotel_performance_view.proto", fileDescriptor_hotel_performance_view_0d02987095d2b3ce)
}

var fileDescriptor_hotel_performance_view_0d02987095d2b3ce = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4a, 0xf4, 0x30,
	0x18, 0xc6, 0x69, 0x3f, 0xf8, 0xc0, 0xa2, 0x9b, 0xc1, 0x85, 0x8a, 0x0b, 0x47, 0x19, 0x70, 0x95,
	0x50, 0xdc, 0x65, 0x40, 0xc8, 0x6c, 0x46, 0x5c, 0x48, 0x99, 0x45, 0x17, 0x52, 0x28, 0xb1, 0x79,
	0x8d, 0x81, 0x36, 0x6f, 0x49, 0x6a, 0xe7, 0x06, 0x1e, 0xc4, 0xa5, 0x47, 0xf1, 0x28, 0x9e, 0x42,
	0x3a, 0x31, 0x71, 0x23, 0xba, 0x7b, 0x48, 0x7e, 0xcf, 0x1f, 0xde, 0xec, 0x5a, 0x21, 0xaa, 0x16,
	0xa8, 0x90, 0x8e, 0x7a, 0x39, 0xa9, 0x31, 0xa7, 0x16, 0x1c, 0x3e, 0xdb, 0x06, 0x1c, 0x7d, 0xc2,
	0x01, 0xda, 0xba, 0x07, 0xfb, 0x88, 0xb6, 0x13, 0xa6, 0x81, 0x7a, 0xd4, 0xb0, 0x25, 0xbd, 0xc5,
	0x01, 0x67, 0x73, 0x6f, 0x22, 0x42, 0x3a, 0x12, 0xfd, 0x64, 0xcc, 0x49, 0xf4, 0x9f, 0x9c, 0x86,
	0x8a, 0x5e, 0x53, 0x61, 0x0c, 0x0e, 0x62, 0xd0, 0x68, 0x9c, 0x0f, 0x38, 0x5f, 0x66, 0x87, 0x37,
	0x53, 0x41, 0xf1, 0x9d, 0x5f, 0x6a, 0xd8, 0xce, 0x2e, 0xb2, 0x83, 0x10, 0x51, 0x1b, 0xd1, 0xc1,
	0x51, 0x72, 0x96, 0x5c, 0xee, 0x6d, 0xf6, 0xc3, 0xe3, 0x9d, 0xe8, 0x60, 0xf5, 0x92, 0x66, 0x8b,
	0x06, 0x3b, 0xf2, 0xe7, 0x88, 0xd5, 0xf1, 0x4f, 0x25, 0xc5, 0xb4, 0xa0, 0x48, 0xee, 0x6f, 0xbf,
	0xfc, 0x0a, 0x5b, 0x61, 0x14, 0x41, 0xab, 0xa8, 0x02, 0xb3, 0xdb, 0x17, 0x8e, 0xd2, 0x6b, 0xf7,
	0xcb, 0x8d, 0x96, 0x51, 0xbd, 0xa6, 0xff, 0xd6, 0x9c, 0xbf, 0xa5, 0xf3, 0xb5, 0x8f, 0xe4, 0xd2,
	0x11, 0x2f, 0x27, 0x55, 0xe6, 0x64, 0x13, 0xc8, 0xf7, 0xc0, 0x54, 0x5c, 0xba, 0x2a, 0x32, 0x55,
	0x99, 0x57, 0x91, 0xf9, 0x48, 0x17, 0xfe, 0x83, 0x31, 0x2e, 0x1d, 0x63, 0x91, 0x62, 0xac, 0xcc,
	0x19, 0x8b, 0xdc, 0xc3, 0xff, 0xdd, 0xd8, 0xab, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x95,
	0x1c, 0x0f, 0xcf, 0x01, 0x00, 0x00,
}
