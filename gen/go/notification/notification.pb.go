// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notification/notification.proto

package notificationpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import coredocument "github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
import types "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// NotificationMessage wraps a single CoreDocument to be notified to upstream services
type NotificationMessage struct {
	EventType            uint32                     `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	CentrifugeId         []byte                     `protobuf:"bytes,2,opt,name=centrifuge_id,json=centrifugeId,proto3" json:"centrifuge_id,omitempty"`
	Recorded             *types.Timestamp           `protobuf:"bytes,3,opt,name=recorded" json:"recorded,omitempty"`
	Document             *coredocument.CoreDocument `protobuf:"bytes,4,opt,name=document" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *NotificationMessage) Reset()         { *m = NotificationMessage{} }
func (m *NotificationMessage) String() string { return proto.CompactTextString(m) }
func (*NotificationMessage) ProtoMessage()    {}
func (*NotificationMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_bab2edfdc5106481, []int{0}
}
func (m *NotificationMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationMessage.Unmarshal(m, b)
}
func (m *NotificationMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationMessage.Marshal(b, m, deterministic)
}
func (dst *NotificationMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationMessage.Merge(dst, src)
}
func (m *NotificationMessage) XXX_Size() int {
	return xxx_messageInfo_NotificationMessage.Size(m)
}
func (m *NotificationMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationMessage proto.InternalMessageInfo

func (m *NotificationMessage) GetEventType() uint32 {
	if m != nil {
		return m.EventType
	}
	return 0
}

func (m *NotificationMessage) GetCentrifugeId() []byte {
	if m != nil {
		return m.CentrifugeId
	}
	return nil
}

func (m *NotificationMessage) GetRecorded() *types.Timestamp {
	if m != nil {
		return m.Recorded
	}
	return nil
}

func (m *NotificationMessage) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

func init() {
	proto.RegisterType((*NotificationMessage)(nil), "notification.NotificationMessage")
}

func init() {
	proto.RegisterFile("notification/notification.proto", fileDescriptor_notification_bab2edfdc5106481)
}

var fileDescriptor_notification_bab2edfdc5106481 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x14, 0x84, 0x89, 0x8a, 0x68, 0xec, 0x8a, 0xd6, 0x4b, 0x29, 0x48, 0x8b, 0x5e, 0x7a, 0x4a, 0x41,
	0x41, 0xef, 0xab, 0x17, 0x0f, 0xca, 0x52, 0xf6, 0xe4, 0x65, 0x69, 0x93, 0xd7, 0x12, 0x30, 0x79,
	0x21, 0x7d, 0x15, 0xf6, 0x27, 0xfa, 0xaf, 0x64, 0xdb, 0xed, 0x1a, 0x8f, 0x33, 0x7c, 0x93, 0xcc,
	0x3c, 0x9e, 0x59, 0x24, 0xdd, 0x6a, 0x59, 0x93, 0x46, 0x5b, 0x86, 0x42, 0x38, 0x8f, 0x84, 0x71,
	0x14, 0x7a, 0x69, 0x26, 0xd1, 0x83, 0x42, 0x39, 0x18, 0xb0, 0x54, 0x86, 0x62, 0xc2, 0xd3, 0xac,
	0x43, 0xec, 0xbe, 0xa0, 0x1c, 0x55, 0x33, 0xb4, 0x25, 0x69, 0x03, 0x3d, 0xd5, 0xc6, 0x4d, 0xc0,
	0xdd, 0x0f, 0xe3, 0x37, 0x1f, 0xc1, 0x93, 0xef, 0xd0, 0xf7, 0x75, 0x07, 0xf1, 0x2d, 0xe7, 0xf0,
	0x0d, 0x96, 0x36, 0xb4, 0x75, 0x90, 0xb0, 0x9c, 0x15, 0x8b, 0xea, 0x7c, 0x74, 0xd6, 0x5b, 0x07,
	0xf1, 0x3d, 0x5f, 0x48, 0xb0, 0xe4, 0x75, 0x3b, 0x74, 0xb0, 0xd1, 0x2a, 0x39, 0xca, 0x59, 0x11,
	0x55, 0xd1, 0x9f, 0xf9, 0xa6, 0xe2, 0x27, 0x7e, 0xe6, 0x41, 0xa2, 0x57, 0xa0, 0x92, 0xe3, 0x9c,
	0x15, 0x17, 0x0f, 0xa9, 0x98, 0xfa, 0x88, 0xb9, 0x8f, 0x58, 0xcf, 0x7d, 0xaa, 0x03, 0xbb, 0xcb,
	0xcd, 0x33, 0x92, 0x93, 0x7d, 0xee, 0xdf, 0xb6, 0x17, 0xf4, 0xf0, 0xba, 0x17, 0xd5, 0x81, 0x5d,
	0x3e, 0xf3, 0x2b, 0x89, 0x46, 0x84, 0x17, 0x5a, 0x5e, 0x87, 0xe3, 0x56, 0xbb, 0x5f, 0x57, 0xec,
	0xf3, 0x32, 0x44, 0x5c, 0xd3, 0x9c, 0x8e, 0x75, 0x1e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd2,
	0x81, 0x94, 0x61, 0x7e, 0x01, 0x00, 0x00,
}
