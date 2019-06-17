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

// GenericData for generic documents.
// Note: for now, it just contains a single field with default value always and not publicly exposed
type GenericData struct {
	Scheme               string   `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericData) Reset()         { *m = GenericData{} }
func (m *GenericData) String() string { return proto.CompactTextString(m) }
func (*GenericData) ProtoMessage()    {}
func (*GenericData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f66020d33d313c29, []int{0}
}

func (m *GenericData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericData.Unmarshal(m, b)
}
func (m *GenericData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericData.Marshal(b, m, deterministic)
}
func (m *GenericData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericData.Merge(m, src)
}
func (m *GenericData) XXX_Size() int {
	return xxx_messageInfo_GenericData.Size(m)
}
func (m *GenericData) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericData.DiscardUnknown(m)
}

var xxx_messageInfo_GenericData proto.InternalMessageInfo

func (m *GenericData) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func init() {
	proto.RegisterType((*GenericData)(nil), "generic.GenericData")
}

func init() { proto.RegisterFile("generic/generic.proto", fileDescriptor_f66020d33d313c29) }

var fileDescriptor_f66020d33d313c29 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x4f, 0xcd, 0x4b,
	0x2d, 0xca, 0x4c, 0xd6, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x2a, 0x17, 0xb7, 0x3b, 0x84, 0xe9, 0x92, 0x58, 0x92, 0x28, 0x24, 0xc6, 0xc5, 0x56, 0x9c,
	0x9c, 0x91, 0x9a, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x39, 0x69, 0x70,
	0x71, 0x27, 0xe7, 0xe7, 0xea, 0x41, 0x75, 0x39, 0xf1, 0x40, 0xf5, 0x04, 0x80, 0x0c, 0x0b, 0x60,
	0x8c, 0xe2, 0x84, 0x4a, 0x14, 0x24, 0x25, 0xb1, 0x81, 0x2d, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x80, 0xe4, 0xde, 0xdd, 0x79, 0x00, 0x00, 0x00,
}
