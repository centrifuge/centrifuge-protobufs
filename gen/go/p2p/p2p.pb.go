// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p/p2p.proto

package p2ppb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import coredocument "github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// P2PMessage wraps a single CoreDocument to be transferred to another node
type P2PMessage struct {
	NetworkIdentifier  uint32 `protobuf:"varint,1,opt,name=network_identifier,json=networkIdentifier,proto3" json:"network_identifier,omitempty"`
	CentNodeVersion    string `protobuf:"bytes,2,opt,name=cent_node_version,json=centNodeVersion,proto3" json:"cent_node_version,omitempty"`
	SenderCentrifugeId []byte `protobuf:"bytes,3,opt,name=sender_centrifuge_id,json=senderCentrifugeId,proto3" json:"sender_centrifuge_id,omitempty"`
	// Open questions in the P2PMessage
	// - should we include the document schema so the client can refuse it right away?
	// - how do you determine if the node is the right recipient for the current
	//   transaction (e.g. two different nodes with different keys are used for
	//   one centrifuge ID handling different data).
	Document             *coredocument.CoreDocument `protobuf:"bytes,5,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *P2PMessage) Reset()         { *m = P2PMessage{} }
func (m *P2PMessage) String() string { return proto.CompactTextString(m) }
func (*P2PMessage) ProtoMessage()    {}
func (*P2PMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_42fc7e0ec0f6bc9d, []int{0}
}
func (m *P2PMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PMessage.Unmarshal(m, b)
}
func (m *P2PMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PMessage.Marshal(b, m, deterministic)
}
func (dst *P2PMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PMessage.Merge(dst, src)
}
func (m *P2PMessage) XXX_Size() int {
	return xxx_messageInfo_P2PMessage.Size(m)
}
func (m *P2PMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PMessage.DiscardUnknown(m)
}

var xxx_messageInfo_P2PMessage proto.InternalMessageInfo

func (m *P2PMessage) GetNetworkIdentifier() uint32 {
	if m != nil {
		return m.NetworkIdentifier
	}
	return 0
}

func (m *P2PMessage) GetCentNodeVersion() string {
	if m != nil {
		return m.CentNodeVersion
	}
	return ""
}

func (m *P2PMessage) GetSenderCentrifugeId() []byte {
	if m != nil {
		return m.SenderCentrifugeId
	}
	return nil
}

func (m *P2PMessage) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type P2PReply struct {
	CentNodeVersion      string                     `protobuf:"bytes,1,opt,name=cent_node_version,json=centNodeVersion,proto3" json:"cent_node_version,omitempty"`
	Document             *coredocument.CoreDocument `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *P2PReply) Reset()         { *m = P2PReply{} }
func (m *P2PReply) String() string { return proto.CompactTextString(m) }
func (*P2PReply) ProtoMessage()    {}
func (*P2PReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_42fc7e0ec0f6bc9d, []int{1}
}
func (m *P2PReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_P2PReply.Unmarshal(m, b)
}
func (m *P2PReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_P2PReply.Marshal(b, m, deterministic)
}
func (dst *P2PReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_P2PReply.Merge(dst, src)
}
func (m *P2PReply) XXX_Size() int {
	return xxx_messageInfo_P2PReply.Size(m)
}
func (m *P2PReply) XXX_DiscardUnknown() {
	xxx_messageInfo_P2PReply.DiscardUnknown(m)
}

var xxx_messageInfo_P2PReply proto.InternalMessageInfo

func (m *P2PReply) GetCentNodeVersion() string {
	if m != nil {
		return m.CentNodeVersion
	}
	return ""
}

func (m *P2PReply) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type CentrifugeHeader struct {
	NetworkIdentifier    uint32   `protobuf:"varint,1,opt,name=network_identifier,json=networkIdentifier,proto3" json:"network_identifier,omitempty"`
	CentNodeVersion      string   `protobuf:"bytes,2,opt,name=cent_node_version,json=centNodeVersion,proto3" json:"cent_node_version,omitempty"`
	SenderCentrifugeId   []byte   `protobuf:"bytes,3,opt,name=sender_centrifuge_id,json=senderCentrifugeId,proto3" json:"sender_centrifuge_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CentrifugeHeader) Reset()         { *m = CentrifugeHeader{} }
func (m *CentrifugeHeader) String() string { return proto.CompactTextString(m) }
func (*CentrifugeHeader) ProtoMessage()    {}
func (*CentrifugeHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_42fc7e0ec0f6bc9d, []int{2}
}
func (m *CentrifugeHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CentrifugeHeader.Unmarshal(m, b)
}
func (m *CentrifugeHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CentrifugeHeader.Marshal(b, m, deterministic)
}
func (dst *CentrifugeHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CentrifugeHeader.Merge(dst, src)
}
func (m *CentrifugeHeader) XXX_Size() int {
	return xxx_messageInfo_CentrifugeHeader.Size(m)
}
func (m *CentrifugeHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_CentrifugeHeader.DiscardUnknown(m)
}

var xxx_messageInfo_CentrifugeHeader proto.InternalMessageInfo

func (m *CentrifugeHeader) GetNetworkIdentifier() uint32 {
	if m != nil {
		return m.NetworkIdentifier
	}
	return 0
}

func (m *CentrifugeHeader) GetCentNodeVersion() string {
	if m != nil {
		return m.CentNodeVersion
	}
	return ""
}

func (m *CentrifugeHeader) GetSenderCentrifugeId() []byte {
	if m != nil {
		return m.SenderCentrifugeId
	}
	return nil
}

type SignatureRequest struct {
	Header               *CentrifugeHeader          `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Document             *coredocument.CoreDocument `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SignatureRequest) Reset()         { *m = SignatureRequest{} }
func (m *SignatureRequest) String() string { return proto.CompactTextString(m) }
func (*SignatureRequest) ProtoMessage()    {}
func (*SignatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_42fc7e0ec0f6bc9d, []int{3}
}
func (m *SignatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureRequest.Unmarshal(m, b)
}
func (m *SignatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureRequest.Marshal(b, m, deterministic)
}
func (dst *SignatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureRequest.Merge(dst, src)
}
func (m *SignatureRequest) XXX_Size() int {
	return xxx_messageInfo_SignatureRequest.Size(m)
}
func (m *SignatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureRequest proto.InternalMessageInfo

func (m *SignatureRequest) GetHeader() *CentrifugeHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignatureRequest) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type SignatureResponse struct {
	CentNodeVersion      string                  `protobuf:"bytes,1,opt,name=cent_node_version,json=centNodeVersion,proto3" json:"cent_node_version,omitempty"`
	Signature            *coredocument.Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SignatureResponse) Reset()         { *m = SignatureResponse{} }
func (m *SignatureResponse) String() string { return proto.CompactTextString(m) }
func (*SignatureResponse) ProtoMessage()    {}
func (*SignatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_42fc7e0ec0f6bc9d, []int{4}
}
func (m *SignatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureResponse.Unmarshal(m, b)
}
func (m *SignatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureResponse.Marshal(b, m, deterministic)
}
func (dst *SignatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureResponse.Merge(dst, src)
}
func (m *SignatureResponse) XXX_Size() int {
	return xxx_messageInfo_SignatureResponse.Size(m)
}
func (m *SignatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureResponse proto.InternalMessageInfo

func (m *SignatureResponse) GetCentNodeVersion() string {
	if m != nil {
		return m.CentNodeVersion
	}
	return ""
}

func (m *SignatureResponse) GetSignature() *coredocument.Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AnchDocumentRequest struct {
	Header               *CentrifugeHeader          `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Document             *coredocument.CoreDocument `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AnchDocumentRequest) Reset()         { *m = AnchDocumentRequest{} }
func (m *AnchDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*AnchDocumentRequest) ProtoMessage()    {}
func (*AnchDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_42fc7e0ec0f6bc9d, []int{5}
}
func (m *AnchDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchDocumentRequest.Unmarshal(m, b)
}
func (m *AnchDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchDocumentRequest.Marshal(b, m, deterministic)
}
func (dst *AnchDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchDocumentRequest.Merge(dst, src)
}
func (m *AnchDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_AnchDocumentRequest.Size(m)
}
func (m *AnchDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnchDocumentRequest proto.InternalMessageInfo

func (m *AnchDocumentRequest) GetHeader() *CentrifugeHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *AnchDocumentRequest) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type AnchDocumentResponse struct {
	CentNodeVersion      string   `protobuf:"bytes,1,opt,name=cent_node_version,json=centNodeVersion,proto3" json:"cent_node_version,omitempty"`
	Accepted             bool     `protobuf:"varint,2,opt,name=accepted,proto3" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnchDocumentResponse) Reset()         { *m = AnchDocumentResponse{} }
func (m *AnchDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*AnchDocumentResponse) ProtoMessage()    {}
func (*AnchDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_42fc7e0ec0f6bc9d, []int{6}
}
func (m *AnchDocumentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchDocumentResponse.Unmarshal(m, b)
}
func (m *AnchDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchDocumentResponse.Marshal(b, m, deterministic)
}
func (dst *AnchDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchDocumentResponse.Merge(dst, src)
}
func (m *AnchDocumentResponse) XXX_Size() int {
	return xxx_messageInfo_AnchDocumentResponse.Size(m)
}
func (m *AnchDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnchDocumentResponse proto.InternalMessageInfo

func (m *AnchDocumentResponse) GetCentNodeVersion() string {
	if m != nil {
		return m.CentNodeVersion
	}
	return ""
}

func (m *AnchDocumentResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func init() {
	proto.RegisterType((*P2PMessage)(nil), "p2p.P2PMessage")
	proto.RegisterType((*P2PReply)(nil), "p2p.P2PReply")
	proto.RegisterType((*CentrifugeHeader)(nil), "p2p.CentrifugeHeader")
	proto.RegisterType((*SignatureRequest)(nil), "p2p.SignatureRequest")
	proto.RegisterType((*SignatureResponse)(nil), "p2p.SignatureResponse")
	proto.RegisterType((*AnchDocumentRequest)(nil), "p2p.AnchDocumentRequest")
	proto.RegisterType((*AnchDocumentResponse)(nil), "p2p.AnchDocumentResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// P2PServiceClient is the client API for P2PService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type P2PServiceClient interface {
	// Post transmits a new version of the document to the recipient
	Post(ctx context.Context, in *P2PMessage, opts ...grpc.CallOption) (*P2PReply, error)
	// after the pre-commit the sender asks all participants to sign the new document version
	RequestDocumentSignature(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*SignatureResponse, error)
	// after all signatures are collected the sender sends the document including the signatures
	SendAnchoredDocument(ctx context.Context, in *AnchDocumentRequest, opts ...grpc.CallOption) (*AnchDocumentResponse, error)
}

type p2PServiceClient struct {
	cc *grpc.ClientConn
}

func NewP2PServiceClient(cc *grpc.ClientConn) P2PServiceClient {
	return &p2PServiceClient{cc}
}

func (c *p2PServiceClient) Post(ctx context.Context, in *P2PMessage, opts ...grpc.CallOption) (*P2PReply, error) {
	out := new(P2PReply)
	err := c.cc.Invoke(ctx, "/p2p.P2PService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PServiceClient) RequestDocumentSignature(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*SignatureResponse, error) {
	out := new(SignatureResponse)
	err := c.cc.Invoke(ctx, "/p2p.P2PService/RequestDocumentSignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PServiceClient) SendAnchoredDocument(ctx context.Context, in *AnchDocumentRequest, opts ...grpc.CallOption) (*AnchDocumentResponse, error) {
	out := new(AnchDocumentResponse)
	err := c.cc.Invoke(ctx, "/p2p.P2PService/SendAnchoredDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P2PServiceServer is the server API for P2PService service.
type P2PServiceServer interface {
	// Post transmits a new version of the document to the recipient
	Post(context.Context, *P2PMessage) (*P2PReply, error)
	// after the pre-commit the sender asks all participants to sign the new document version
	RequestDocumentSignature(context.Context, *SignatureRequest) (*SignatureResponse, error)
	// after all signatures are collected the sender sends the document including the signatures
	SendAnchoredDocument(context.Context, *AnchDocumentRequest) (*AnchDocumentResponse, error)
}

func RegisterP2PServiceServer(s *grpc.Server, srv P2PServiceServer) {
	s.RegisterService(&_P2PService_serviceDesc, srv)
}

func _P2PService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(P2PMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.P2PService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).Post(ctx, req.(*P2PMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PService_RequestDocumentSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).RequestDocumentSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.P2PService/RequestDocumentSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).RequestDocumentSignature(ctx, req.(*SignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PService_SendAnchoredDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnchDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PServiceServer).SendAnchoredDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.P2PService/SendAnchoredDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PServiceServer).SendAnchoredDocument(ctx, req.(*AnchDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _P2PService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.P2PService",
	HandlerType: (*P2PServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _P2PService_Post_Handler,
		},
		{
			MethodName: "RequestDocumentSignature",
			Handler:    _P2PService_RequestDocumentSignature_Handler,
		},
		{
			MethodName: "SendAnchoredDocument",
			Handler:    _P2PService_SendAnchoredDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "p2p/p2p.proto",
}

func init() { proto.RegisterFile("p2p/p2p.proto", fileDescriptor_p2p_42fc7e0ec0f6bc9d) }

var fileDescriptor_p2p_42fc7e0ec0f6bc9d = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x96, 0x1b, 0x5a, 0xb6, 0x53, 0xa2, 0x36, 0x26, 0xc0, 0xb2, 0x42, 0x22, 0xca, 0x01, 0x45,
	0x48, 0x4d, 0x91, 0x11, 0xdc, 0x69, 0x39, 0x10, 0x21, 0xd0, 0xca, 0x91, 0x38, 0x70, 0x60, 0x95,
	0xda, 0xd3, 0xd4, 0x82, 0xda, 0xc6, 0x76, 0x82, 0x22, 0xf1, 0x2e, 0xbc, 0x10, 0x17, 0xde, 0x08,
	0xed, 0x6f, 0x7e, 0x94, 0x4b, 0x38, 0xa0, 0x1e, 0x67, 0xbe, 0xd1, 0x7c, 0x3f, 0xbb, 0x63, 0x68,
	0x5b, 0x66, 0xcf, 0x2c, 0xb3, 0x43, 0xeb, 0x4c, 0x30, 0xb4, 0x65, 0x99, 0x4d, 0x9e, 0x0a, 0xe3,
	0x50, 0x1a, 0x31, 0xbb, 0x41, 0x1d, 0xce, 0x56, 0x8b, 0x72, 0xaa, 0xff, 0x87, 0x00, 0xa4, 0x2c,
	0xfd, 0x80, 0xde, 0x4f, 0xa6, 0x48, 0x4f, 0x81, 0x6a, 0x0c, 0x3f, 0x8c, 0xfb, 0x9a, 0x29, 0x89,
	0x3a, 0xa8, 0x2b, 0x85, 0x2e, 0x26, 0x3d, 0x32, 0x68, 0xf3, 0x4e, 0x85, 0x8c, 0x1a, 0x80, 0x3e,
	0x87, 0x8e, 0x40, 0x1d, 0x32, 0x6d, 0x24, 0x66, 0x73, 0x74, 0x5e, 0x19, 0x1d, 0xef, 0xf5, 0xc8,
	0xe0, 0x90, 0x1f, 0xe7, 0xc0, 0x47, 0x23, 0xf1, 0x53, 0xd9, 0xa6, 0x2f, 0xa0, 0xeb, 0x51, 0x4b,
	0x74, 0x59, 0x8e, 0x38, 0x75, 0x35, 0x9b, 0x62, 0xa6, 0x64, 0xdc, 0xea, 0x91, 0xc1, 0x3d, 0x4e,
	0x4b, 0xec, 0xa2, 0x81, 0x46, 0x92, 0xbe, 0x86, 0xa8, 0x56, 0x1b, 0xef, 0xf7, 0xc8, 0xe0, 0x88,
	0x25, 0xc3, 0x35, 0x0b, 0x17, 0xc6, 0xe1, 0xdb, 0xaa, 0xe0, 0xcd, 0x6c, 0x5f, 0x43, 0x94, 0xb2,
	0x94, 0xa3, 0xfd, 0xb6, 0xd8, 0xae, 0x90, 0x6c, 0x57, 0xb8, 0xca, 0xd7, 0xda, 0x81, 0xef, 0x17,
	0x81, 0x93, 0xa5, 0xf0, 0x77, 0x38, 0x91, 0xe8, 0x6e, 0x55, 0x92, 0xfd, 0x05, 0x9c, 0x8c, 0xd5,
	0x54, 0x4f, 0xc2, 0xcc, 0x21, 0xc7, 0xef, 0x33, 0xf4, 0x81, 0x9e, 0xc2, 0xc1, 0x75, 0x21, 0xb5,
	0x10, 0x75, 0xc4, 0x1e, 0x0c, 0xf3, 0x7f, 0x67, 0xd3, 0x07, 0xaf, 0x86, 0xd6, 0xc2, 0xd9, 0xdb,
	0x21, 0x9c, 0x39, 0x74, 0x56, 0xa8, 0xbd, 0x35, 0xda, 0xe3, 0x4e, 0x5f, 0xe5, 0x15, 0x1c, 0xfa,
	0x7a, 0x41, 0xc5, 0xfc, 0x68, 0x9d, 0x79, 0xb9, 0x7f, 0x39, 0xd9, 0xff, 0x09, 0xf7, 0xdf, 0x68,
	0x71, 0xdd, 0x28, 0xfa, 0xbf, 0xae, 0xbf, 0x40, 0x77, 0x9d, 0xfd, 0x1f, 0x8c, 0x27, 0x10, 0x4d,
	0x84, 0x40, 0x1b, 0x50, 0x16, 0xdc, 0x11, 0x6f, 0x6a, 0xf6, 0xbb, 0x3c, 0xdb, 0x31, 0xba, 0xb9,
	0x12, 0x48, 0x9f, 0xc1, 0x9d, 0xd4, 0xf8, 0x40, 0x8f, 0x0b, 0x37, 0xcb, 0x7b, 0x4e, 0xda, 0x75,
	0xa3, 0xbc, 0x86, 0x11, 0xc4, 0x55, 0x10, 0xb5, 0xb2, 0x26, 0x3b, 0x5a, 0x26, 0xb1, 0xf9, 0x9b,
	0x24, 0x0f, 0x37, 0xdb, 0x95, 0x93, 0xf7, 0xd0, 0x1d, 0xa3, 0x96, 0xb9, 0xcb, 0x3c, 0x8f, 0x7a,
	0x1f, 0x8d, 0x8b, 0xf9, 0x2d, 0xd1, 0x27, 0x8f, 0xb7, 0x20, 0xe5, 0xb2, 0xf3, 0x27, 0x70, 0x57,
	0x98, 0x9b, 0x1c, 0x3f, 0x8f, 0x52, 0x66, 0xd3, 0xfc, 0x69, 0x4a, 0xc9, 0xe7, 0x7d, 0xcb, 0xac,
	0xbd, 0xbc, 0x3c, 0x28, 0x9e, 0xaa, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x52, 0x34, 0xe6,
	0xcd, 0xe1, 0x04, 0x00, 0x00,
}
