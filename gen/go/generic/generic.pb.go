// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generic/generic.proto

package genericpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// empty generic doc
type Generic struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Generic) Reset()         { *m = Generic{} }
func (m *Generic) String() string { return proto.CompactTextString(m) }
func (*Generic) ProtoMessage()    {}
func (*Generic) Descriptor() ([]byte, []int) {
	return fileDescriptor_f66020d33d313c29, []int{0}
}

func (m *Generic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Generic.Unmarshal(m, b)
}
func (m *Generic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Generic.Marshal(b, m, deterministic)
}
func (m *Generic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Generic.Merge(m, src)
}
func (m *Generic) XXX_Size() int {
	return xxx_messageInfo_Generic.Size(m)
}
func (m *Generic) XXX_DiscardUnknown() {
	xxx_messageInfo_Generic.DiscardUnknown(m)
}

var xxx_messageInfo_Generic proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Generic)(nil), "generic.Generic")
}

func init() { proto.RegisterFile("generic/generic.proto", fileDescriptor_f66020d33d313c29) }

var fileDescriptor_f66020d33d313c29 = []byte{
	// 82 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x4f, 0xcd, 0x4b,
	0x2d, 0xca, 0x4c, 0xd6, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x12, 0x27, 0x17, 0xbb, 0x3b, 0x84, 0xe9, 0xa4, 0xc1, 0xc5, 0x9d, 0x9c, 0x9f, 0xab, 0x07, 0x95,
	0x71, 0xe2, 0x81, 0x8a, 0x07, 0x80, 0x34, 0x04, 0x30, 0x46, 0x71, 0x42, 0x25, 0x0a, 0x92, 0x92,
	0xd8, 0xc0, 0x86, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x12, 0x2b, 0xb2, 0x5d, 0x00,
	0x00, 0x00,
}
