// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/header_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible header errors.
type HeaderErrorEnum_HeaderError int32

const (
	// Enum unspecified.
	HeaderErrorEnum_UNSPECIFIED HeaderErrorEnum_HeaderError = 0
	// The received error code is not known in this version.
	HeaderErrorEnum_UNKNOWN HeaderErrorEnum_HeaderError = 1
	// The login customer id could not be validated.
	HeaderErrorEnum_INVALID_LOGIN_CUSTOMER_ID HeaderErrorEnum_HeaderError = 3
	// One or more task headers could not be parsed.
	HeaderErrorEnum_MALFORMED_TASK_INFO HeaderErrorEnum_HeaderError = 4
)

var HeaderErrorEnum_HeaderError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "INVALID_LOGIN_CUSTOMER_ID",
	4: "MALFORMED_TASK_INFO",
}
var HeaderErrorEnum_HeaderError_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"INVALID_LOGIN_CUSTOMER_ID": 3,
	"MALFORMED_TASK_INFO":       4,
}

func (x HeaderErrorEnum_HeaderError) String() string {
	return proto.EnumName(HeaderErrorEnum_HeaderError_name, int32(x))
}
func (HeaderErrorEnum_HeaderError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_header_error_eabd2af5e7269dcb, []int{0, 0}
}

// Container for enum describing possible header errors.
type HeaderErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderErrorEnum) Reset()         { *m = HeaderErrorEnum{} }
func (m *HeaderErrorEnum) String() string { return proto.CompactTextString(m) }
func (*HeaderErrorEnum) ProtoMessage()    {}
func (*HeaderErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_header_error_eabd2af5e7269dcb, []int{0}
}
func (m *HeaderErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderErrorEnum.Unmarshal(m, b)
}
func (m *HeaderErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderErrorEnum.Marshal(b, m, deterministic)
}
func (dst *HeaderErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderErrorEnum.Merge(dst, src)
}
func (m *HeaderErrorEnum) XXX_Size() int {
	return xxx_messageInfo_HeaderErrorEnum.Size(m)
}
func (m *HeaderErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HeaderErrorEnum)(nil), "google.ads.googleads.v0.errors.HeaderErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.HeaderErrorEnum_HeaderError", HeaderErrorEnum_HeaderError_name, HeaderErrorEnum_HeaderError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/header_error.proto", fileDescriptor_header_error_eabd2af5e7269dcb)
}

var fileDescriptor_header_error_eabd2af5e7269dcb = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4f, 0x83, 0x30,
	0x18, 0xc6, 0x65, 0x33, 0x9a, 0x94, 0xc3, 0x08, 0x1e, 0x8c, 0x07, 0x77, 0xe0, 0x03, 0x14, 0x8c,
	0x47, 0x4f, 0xdd, 0x28, 0xd8, 0x0c, 0x0a, 0x19, 0x03, 0x13, 0x43, 0xd2, 0xe0, 0x20, 0xd5, 0x64,
	0xa3, 0x4b, 0xab, 0x7c, 0x20, 0x8f, 0x7e, 0x10, 0x0f, 0x7e, 0x2a, 0x03, 0x75, 0xcb, 0x2e, 0xee,
	0xd4, 0xe7, 0x6d, 0x9e, 0xdf, 0xfb, 0xe7, 0x01, 0x77, 0x5c, 0x08, 0xbe, 0x69, 0xdc, 0xaa, 0x56,
	0xae, 0x96, 0xbd, 0xea, 0x3c, 0xb7, 0x91, 0x52, 0x48, 0xe5, 0xbe, 0x36, 0x55, 0xdd, 0x48, 0x36,
	0x54, 0x70, 0x27, 0xc5, 0xbb, 0xb0, 0xa7, 0xda, 0x07, 0xab, 0x5a, 0xc1, 0x03, 0x02, 0x3b, 0x0f,
	0x6a, 0xc4, 0xe9, 0xc0, 0xe4, 0x71, 0xa0, 0x70, 0x5f, 0xe3, 0xf6, 0x63, 0xeb, 0xac, 0x81, 0x79,
	0xf4, 0x65, 0x4f, 0x80, 0x99, 0xd3, 0x2c, 0xc5, 0x73, 0x12, 0x10, 0xec, 0x5b, 0x67, 0xb6, 0x09,
	0x2e, 0x73, 0xba, 0xa0, 0xc9, 0x13, 0xb5, 0x0c, 0xfb, 0x16, 0xdc, 0x10, 0x5a, 0xa0, 0x88, 0xf8,
	0x2c, 0x4a, 0x42, 0x42, 0xd9, 0x3c, 0xcf, 0x56, 0x49, 0x8c, 0x97, 0x8c, 0xf8, 0xd6, 0xd8, 0xbe,
	0x06, 0x57, 0x31, 0x8a, 0x82, 0x64, 0x19, 0x63, 0x9f, 0xad, 0x50, 0xb6, 0x60, 0x84, 0x06, 0x89,
	0x75, 0x3e, 0xfb, 0x36, 0x80, 0xb3, 0x16, 0x5b, 0x78, 0x7a, 0xbd, 0x99, 0x75, 0xb4, 0x49, 0xda,
	0x1f, 0x94, 0x1a, 0xcf, 0xfe, 0x1f, 0xc3, 0xc5, 0xa6, 0x6a, 0x39, 0x14, 0x92, 0xbb, 0xbc, 0x69,
	0x87, 0x73, 0xf7, 0xa9, 0xec, 0xde, 0xd4, 0x7f, 0x21, 0x3d, 0xe8, 0xe7, 0x73, 0x34, 0x0e, 0x11,
	0xfa, 0x1a, 0x4d, 0x43, 0xdd, 0x0c, 0xd5, 0x0a, 0x6a, 0xd9, 0xab, 0xc2, 0x83, 0xc3, 0x48, 0xf5,
	0xb3, 0x37, 0x94, 0xa8, 0x56, 0xe5, 0xc1, 0x50, 0x16, 0x5e, 0xa9, 0x0d, 0x2f, 0x17, 0xc3, 0xe0,
	0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x27, 0xed, 0x8b, 0x9c, 0x01, 0x00, 0x00,
}
