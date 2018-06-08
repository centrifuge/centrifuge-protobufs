// Code generated by protoc-gen-go. DO NOT EDIT.
// source: invoice/service.proto

package invoicepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proofs "github.com/centrifuge/precise-proofs/proofs"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type CreateInvoiceProofEnvelope struct {
	DocumentIdentifier   []byte   `protobuf:"bytes,1,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	Fields               []string `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInvoiceProofEnvelope) Reset()         { *m = CreateInvoiceProofEnvelope{} }
func (m *CreateInvoiceProofEnvelope) String() string { return proto.CompactTextString(m) }
func (*CreateInvoiceProofEnvelope) ProtoMessage()    {}
func (*CreateInvoiceProofEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_f2cb4fc01dcd6f02, []int{0}
}
func (m *CreateInvoiceProofEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInvoiceProofEnvelope.Unmarshal(m, b)
}
func (m *CreateInvoiceProofEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInvoiceProofEnvelope.Marshal(b, m, deterministic)
}
func (dst *CreateInvoiceProofEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInvoiceProofEnvelope.Merge(dst, src)
}
func (m *CreateInvoiceProofEnvelope) XXX_Size() int {
	return xxx_messageInfo_CreateInvoiceProofEnvelope.Size(m)
}
func (m *CreateInvoiceProofEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInvoiceProofEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInvoiceProofEnvelope proto.InternalMessageInfo

func (m *CreateInvoiceProofEnvelope) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *CreateInvoiceProofEnvelope) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type InvoiceProof struct {
	DocumentIdentifier   []byte          `protobuf:"bytes,1,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	FieldProofs          []*proofs.Proof `protobuf:"bytes,2,rep,name=field_proofs,json=fieldProofs" json:"field_proofs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *InvoiceProof) Reset()         { *m = InvoiceProof{} }
func (m *InvoiceProof) String() string { return proto.CompactTextString(m) }
func (*InvoiceProof) ProtoMessage()    {}
func (*InvoiceProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_f2cb4fc01dcd6f02, []int{1}
}
func (m *InvoiceProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceProof.Unmarshal(m, b)
}
func (m *InvoiceProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceProof.Marshal(b, m, deterministic)
}
func (dst *InvoiceProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceProof.Merge(dst, src)
}
func (m *InvoiceProof) XXX_Size() int {
	return xxx_messageInfo_InvoiceProof.Size(m)
}
func (m *InvoiceProof) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceProof.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceProof proto.InternalMessageInfo

func (m *InvoiceProof) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *InvoiceProof) GetFieldProofs() []*proofs.Proof {
	if m != nil {
		return m.FieldProofs
	}
	return nil
}

type AnchorInvoiceEnvelope struct {
	Document             *InvoiceDocument `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AnchorInvoiceEnvelope) Reset()         { *m = AnchorInvoiceEnvelope{} }
func (m *AnchorInvoiceEnvelope) String() string { return proto.CompactTextString(m) }
func (*AnchorInvoiceEnvelope) ProtoMessage()    {}
func (*AnchorInvoiceEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_f2cb4fc01dcd6f02, []int{2}
}
func (m *AnchorInvoiceEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchorInvoiceEnvelope.Unmarshal(m, b)
}
func (m *AnchorInvoiceEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchorInvoiceEnvelope.Marshal(b, m, deterministic)
}
func (dst *AnchorInvoiceEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorInvoiceEnvelope.Merge(dst, src)
}
func (m *AnchorInvoiceEnvelope) XXX_Size() int {
	return xxx_messageInfo_AnchorInvoiceEnvelope.Size(m)
}
func (m *AnchorInvoiceEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorInvoiceEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorInvoiceEnvelope proto.InternalMessageInfo

func (m *AnchorInvoiceEnvelope) GetDocument() *InvoiceDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type SendInvoiceEnvelope struct {
	// Centrifuge OS Entity of the recipient
	Recipients           [][]byte         `protobuf:"bytes,1,rep,name=recipients,proto3" json:"recipients,omitempty"`
	Document             *InvoiceDocument `protobuf:"bytes,10,opt,name=document" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SendInvoiceEnvelope) Reset()         { *m = SendInvoiceEnvelope{} }
func (m *SendInvoiceEnvelope) String() string { return proto.CompactTextString(m) }
func (*SendInvoiceEnvelope) ProtoMessage()    {}
func (*SendInvoiceEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_f2cb4fc01dcd6f02, []int{3}
}
func (m *SendInvoiceEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendInvoiceEnvelope.Unmarshal(m, b)
}
func (m *SendInvoiceEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendInvoiceEnvelope.Marshal(b, m, deterministic)
}
func (dst *SendInvoiceEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendInvoiceEnvelope.Merge(dst, src)
}
func (m *SendInvoiceEnvelope) XXX_Size() int {
	return xxx_messageInfo_SendInvoiceEnvelope.Size(m)
}
func (m *SendInvoiceEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_SendInvoiceEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_SendInvoiceEnvelope proto.InternalMessageInfo

func (m *SendInvoiceEnvelope) GetRecipients() [][]byte {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *SendInvoiceEnvelope) GetDocument() *InvoiceDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type GetInvoiceDocumentEnvelope struct {
	DocumentIdentifier   []byte   `protobuf:"bytes,1,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInvoiceDocumentEnvelope) Reset()         { *m = GetInvoiceDocumentEnvelope{} }
func (m *GetInvoiceDocumentEnvelope) String() string { return proto.CompactTextString(m) }
func (*GetInvoiceDocumentEnvelope) ProtoMessage()    {}
func (*GetInvoiceDocumentEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_f2cb4fc01dcd6f02, []int{4}
}
func (m *GetInvoiceDocumentEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInvoiceDocumentEnvelope.Unmarshal(m, b)
}
func (m *GetInvoiceDocumentEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInvoiceDocumentEnvelope.Marshal(b, m, deterministic)
}
func (dst *GetInvoiceDocumentEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInvoiceDocumentEnvelope.Merge(dst, src)
}
func (m *GetInvoiceDocumentEnvelope) XXX_Size() int {
	return xxx_messageInfo_GetInvoiceDocumentEnvelope.Size(m)
}
func (m *GetInvoiceDocumentEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInvoiceDocumentEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_GetInvoiceDocumentEnvelope proto.InternalMessageInfo

func (m *GetInvoiceDocumentEnvelope) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

type ReceivedInvoices struct {
	Invoices             []*InvoiceDocument `protobuf:"bytes,1,rep,name=invoices" json:"invoices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReceivedInvoices) Reset()         { *m = ReceivedInvoices{} }
func (m *ReceivedInvoices) String() string { return proto.CompactTextString(m) }
func (*ReceivedInvoices) ProtoMessage()    {}
func (*ReceivedInvoices) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_f2cb4fc01dcd6f02, []int{5}
}
func (m *ReceivedInvoices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedInvoices.Unmarshal(m, b)
}
func (m *ReceivedInvoices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedInvoices.Marshal(b, m, deterministic)
}
func (dst *ReceivedInvoices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedInvoices.Merge(dst, src)
}
func (m *ReceivedInvoices) XXX_Size() int {
	return xxx_messageInfo_ReceivedInvoices.Size(m)
}
func (m *ReceivedInvoices) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedInvoices.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedInvoices proto.InternalMessageInfo

func (m *ReceivedInvoices) GetInvoices() []*InvoiceDocument {
	if m != nil {
		return m.Invoices
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateInvoiceProofEnvelope)(nil), "invoice.CreateInvoiceProofEnvelope")
	proto.RegisterType((*InvoiceProof)(nil), "invoice.InvoiceProof")
	proto.RegisterType((*AnchorInvoiceEnvelope)(nil), "invoice.AnchorInvoiceEnvelope")
	proto.RegisterType((*SendInvoiceEnvelope)(nil), "invoice.SendInvoiceEnvelope")
	proto.RegisterType((*GetInvoiceDocumentEnvelope)(nil), "invoice.GetInvoiceDocumentEnvelope")
	proto.RegisterType((*ReceivedInvoices)(nil), "invoice.ReceivedInvoices")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InvoiceDocumentService service

type InvoiceDocumentServiceClient interface {
	CreateInvoiceProof(ctx context.Context, in *CreateInvoiceProofEnvelope, opts ...grpc.CallOption) (*InvoiceProof, error)
	AnchorInvoiceDocument(ctx context.Context, in *AnchorInvoiceEnvelope, opts ...grpc.CallOption) (*InvoiceDocument, error)
	SendInvoiceDocument(ctx context.Context, in *SendInvoiceEnvelope, opts ...grpc.CallOption) (*InvoiceDocument, error)
	GetInvoiceDocument(ctx context.Context, in *GetInvoiceDocumentEnvelope, opts ...grpc.CallOption) (*InvoiceDocument, error)
	GetReceivedInvoiceDocuments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReceivedInvoices, error)
}

type invoiceDocumentServiceClient struct {
	cc *grpc.ClientConn
}

func NewInvoiceDocumentServiceClient(cc *grpc.ClientConn) InvoiceDocumentServiceClient {
	return &invoiceDocumentServiceClient{cc}
}

func (c *invoiceDocumentServiceClient) CreateInvoiceProof(ctx context.Context, in *CreateInvoiceProofEnvelope, opts ...grpc.CallOption) (*InvoiceProof, error) {
	out := new(InvoiceProof)
	err := grpc.Invoke(ctx, "/invoice.InvoiceDocumentService/CreateInvoiceProof", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceDocumentServiceClient) AnchorInvoiceDocument(ctx context.Context, in *AnchorInvoiceEnvelope, opts ...grpc.CallOption) (*InvoiceDocument, error) {
	out := new(InvoiceDocument)
	err := grpc.Invoke(ctx, "/invoice.InvoiceDocumentService/AnchorInvoiceDocument", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceDocumentServiceClient) SendInvoiceDocument(ctx context.Context, in *SendInvoiceEnvelope, opts ...grpc.CallOption) (*InvoiceDocument, error) {
	out := new(InvoiceDocument)
	err := grpc.Invoke(ctx, "/invoice.InvoiceDocumentService/SendInvoiceDocument", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceDocumentServiceClient) GetInvoiceDocument(ctx context.Context, in *GetInvoiceDocumentEnvelope, opts ...grpc.CallOption) (*InvoiceDocument, error) {
	out := new(InvoiceDocument)
	err := grpc.Invoke(ctx, "/invoice.InvoiceDocumentService/GetInvoiceDocument", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoiceDocumentServiceClient) GetReceivedInvoiceDocuments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReceivedInvoices, error) {
	out := new(ReceivedInvoices)
	err := grpc.Invoke(ctx, "/invoice.InvoiceDocumentService/GetReceivedInvoiceDocuments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InvoiceDocumentService service

type InvoiceDocumentServiceServer interface {
	CreateInvoiceProof(context.Context, *CreateInvoiceProofEnvelope) (*InvoiceProof, error)
	AnchorInvoiceDocument(context.Context, *AnchorInvoiceEnvelope) (*InvoiceDocument, error)
	SendInvoiceDocument(context.Context, *SendInvoiceEnvelope) (*InvoiceDocument, error)
	GetInvoiceDocument(context.Context, *GetInvoiceDocumentEnvelope) (*InvoiceDocument, error)
	GetReceivedInvoiceDocuments(context.Context, *empty.Empty) (*ReceivedInvoices, error)
}

func RegisterInvoiceDocumentServiceServer(s *grpc.Server, srv InvoiceDocumentServiceServer) {
	s.RegisterService(&_InvoiceDocumentService_serviceDesc, srv)
}

func _InvoiceDocumentService_CreateInvoiceProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceProofEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceDocumentServiceServer).CreateInvoiceProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoice.InvoiceDocumentService/CreateInvoiceProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceDocumentServiceServer).CreateInvoiceProof(ctx, req.(*CreateInvoiceProofEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceDocumentService_AnchorInvoiceDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnchorInvoiceEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceDocumentServiceServer).AnchorInvoiceDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoice.InvoiceDocumentService/AnchorInvoiceDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceDocumentServiceServer).AnchorInvoiceDocument(ctx, req.(*AnchorInvoiceEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceDocumentService_SendInvoiceDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendInvoiceEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceDocumentServiceServer).SendInvoiceDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoice.InvoiceDocumentService/SendInvoiceDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceDocumentServiceServer).SendInvoiceDocument(ctx, req.(*SendInvoiceEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceDocumentService_GetInvoiceDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceDocumentEnvelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceDocumentServiceServer).GetInvoiceDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoice.InvoiceDocumentService/GetInvoiceDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceDocumentServiceServer).GetInvoiceDocument(ctx, req.(*GetInvoiceDocumentEnvelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _InvoiceDocumentService_GetReceivedInvoiceDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceDocumentServiceServer).GetReceivedInvoiceDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoice.InvoiceDocumentService/GetReceivedInvoiceDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceDocumentServiceServer).GetReceivedInvoiceDocuments(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _InvoiceDocumentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "invoice.InvoiceDocumentService",
	HandlerType: (*InvoiceDocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvoiceProof",
			Handler:    _InvoiceDocumentService_CreateInvoiceProof_Handler,
		},
		{
			MethodName: "AnchorInvoiceDocument",
			Handler:    _InvoiceDocumentService_AnchorInvoiceDocument_Handler,
		},
		{
			MethodName: "SendInvoiceDocument",
			Handler:    _InvoiceDocumentService_SendInvoiceDocument_Handler,
		},
		{
			MethodName: "GetInvoiceDocument",
			Handler:    _InvoiceDocumentService_GetInvoiceDocument_Handler,
		},
		{
			MethodName: "GetReceivedInvoiceDocuments",
			Handler:    _InvoiceDocumentService_GetReceivedInvoiceDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice/service.proto",
}

func init() { proto.RegisterFile("invoice/service.proto", fileDescriptor_service_f2cb4fc01dcd6f02) }

var fileDescriptor_service_f2cb4fc01dcd6f02 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x55, 0x36, 0x31, 0x98, 0xdb, 0x01, 0xf2, 0xb6, 0xaa, 0x64, 0xd5, 0x64, 0x0c, 0x87, 0x69,
	0xa2, 0x09, 0x14, 0x2e, 0x83, 0x53, 0x81, 0x69, 0x4c, 0x62, 0x12, 0xca, 0x0e, 0x48, 0x5c, 0xa6,
	0x34, 0xf9, 0x25, 0xb3, 0xe8, 0xec, 0x10, 0x7b, 0x45, 0x13, 0x42, 0x48, 0x9c, 0x39, 0x95, 0x23,
	0x9f, 0x80, 0xcf, 0xc3, 0x37, 0x40, 0x7c, 0x10, 0x14, 0xdb, 0x09, 0x5d, 0xfa, 0x47, 0xc0, 0xc9,
	0x89, 0x7f, 0xcf, 0xef, 0x3d, 0xfb, 0xf7, 0x6c, 0xb4, 0xc9, 0xf8, 0x48, 0xb0, 0x08, 0x7c, 0x09,
	0xf9, 0x88, 0x45, 0xe0, 0x65, 0xb9, 0x50, 0x02, 0x5f, 0xb5, 0xd3, 0xee, 0x56, 0x2a, 0x44, 0x3a,
	0x04, 0x5f, 0x4f, 0x0f, 0xce, 0x13, 0x1f, 0xce, 0x32, 0x75, 0x61, 0x50, 0x6e, 0xc7, 0x16, 0xc3,
	0x8c, 0xf9, 0x21, 0xe7, 0x42, 0x85, 0x8a, 0x09, 0x2e, 0x6d, 0xb5, 0xa2, 0xb6, 0xa3, 0x9d, 0xbe,
	0x9d, 0xe5, 0x10, 0x31, 0x09, 0xdd, 0x2c, 0x17, 0x22, 0x91, 0xfe, 0xe4, 0x60, 0x21, 0xf7, 0xf4,
	0x10, 0x75, 0x53, 0xe0, 0x5d, 0xf9, 0x3e, 0x4c, 0x53, 0xc8, 0x7d, 0x91, 0x69, 0xee, 0x69, 0x1d,
	0x0a, 0xc8, 0x7d, 0x96, 0x43, 0xa8, 0xe0, 0xd0, 0xe8, 0xbc, 0x2a, 0x98, 0xf6, 0xf9, 0x08, 0x86,
	0x22, 0x03, 0xec, 0xa3, 0xf5, 0x58, 0x44, 0xe7, 0x67, 0xc0, 0xd5, 0x09, 0x8b, 0x81, 0x2b, 0x96,
	0x30, 0xc8, 0xdb, 0x0e, 0x71, 0x76, 0x9a, 0x01, 0x2e, 0x4b, 0x87, 0x55, 0x05, 0xb7, 0xd0, 0x4a,
	0xc2, 0x60, 0x18, 0xcb, 0xf6, 0x12, 0x59, 0xde, 0x59, 0x0d, 0xec, 0x1f, 0x7d, 0x87, 0x9a, 0x93,
	0x02, 0xff, 0x4e, 0x7c, 0x1f, 0x35, 0x35, 0xd5, 0x89, 0xd9, 0xb1, 0xa6, 0x6f, 0xf4, 0xd6, 0x3c,
	0xf3, 0xeb, 0x69, 0xd6, 0xa0, 0xa1, 0x21, 0xfa, 0x5b, 0xd2, 0x23, 0xb4, 0xd9, 0xe7, 0xd1, 0xa9,
	0xc8, 0xad, 0x70, 0xb5, 0xa9, 0x47, 0xe8, 0x5a, 0x29, 0xa0, 0x05, 0x1b, 0xbd, 0xb6, 0x57, 0x9e,
	0xb2, 0xc5, 0x3e, 0xb7, 0xf5, 0xa0, 0x42, 0xd2, 0xb7, 0x68, 0xfd, 0x18, 0x78, 0x5c, 0x27, 0xdb,
	0x46, 0xa8, 0xe8, 0x48, 0xc6, 0x80, 0x2b, 0xd9, 0x76, 0xc8, 0xf2, 0x4e, 0x33, 0x98, 0x98, 0xb9,
	0x24, 0x86, 0xfe, 0x5a, 0xec, 0x08, 0xb9, 0x07, 0xa0, 0x6a, 0xf5, 0xff, 0xee, 0x0a, 0x7d, 0x81,
	0x6e, 0x06, 0x10, 0x01, 0x1b, 0x41, 0xe9, 0x5f, 0x1b, 0xb3, 0x3e, 0x8c, 0xed, 0x85, 0xc6, 0x4a,
	0x64, 0xef, 0xe7, 0x15, 0xd4, 0xaa, 0x55, 0x8f, 0x4d, 0xf6, 0xf1, 0x77, 0x07, 0xe1, 0xe9, 0x28,
	0xe1, 0x3b, 0x15, 0xeb, 0xfc, 0x9c, 0xb9, 0x9b, 0x75, 0x69, 0x5d, 0xa6, 0xaf, 0xc7, 0xfd, 0x27,
	0xee, 0x9e, 0x59, 0x27, 0x49, 0x48, 0x86, 0x4c, 0x2a, 0x22, 0x12, 0x62, 0x2f, 0x01, 0x31, 0xcd,
	0x27, 0x89, 0xc8, 0x89, 0x3a, 0x05, 0x22, 0x33, 0x88, 0x8a, 0x0d, 0xc7, 0xc4, 0xa4, 0xee, 0xf3,
	0x8f, 0x5f, 0x5f, 0x97, 0xd6, 0xe9, 0xf5, 0xf2, 0x0e, 0x99, 0x6b, 0xf2, 0xd8, 0xd9, 0xc5, 0x5f,
	0x9c, 0x5a, 0x38, 0xca, 0xcd, 0xe0, 0xed, 0xca, 0xc9, 0xcc, 0xf0, 0xb8, 0x73, 0x0f, 0x89, 0xee,
	0x8d, 0xfb, 0x1d, 0xd7, 0x35, 0xab, 0x48, 0xc8, 0x89, 0xc5, 0x91, 0xb2, 0x1f, 0xda, 0xcd, 0x06,
	0xbd, 0x51, 0xb9, 0x09, 0x35, 0xb4, 0xb0, 0xf3, 0xcd, 0xb9, 0x14, 0xae, 0xca, 0x4c, 0xa7, 0x12,
	0x9b, 0x11, 0xbd, 0x05, 0x56, 0x5e, 0x8e, 0xfb, 0x0f, 0x5c, 0xbf, 0x58, 0x33, 0xcb, 0x08, 0x51,
	0xa2, 0x76, 0x5c, 0x59, 0x98, 0xab, 0x0b, 0xed, 0x0f, 0xd3, 0x35, 0xff, 0xcf, 0xa3, 0xc6, 0xe3,
	0xc2, 0xdd, 0x27, 0x84, 0xa7, 0xc3, 0x38, 0xd1, 0xd7, 0xf9, 0x49, 0x5d, 0x60, 0x71, 0x57, 0xcb,
	0xdd, 0xc5, 0xb4, 0x92, 0x4b, 0x41, 0xf9, 0x1f, 0x66, 0xe4, 0xfa, 0x23, 0xe6, 0x68, 0xeb, 0x00,
	0x54, 0x2d, 0xc1, 0x25, 0x93, 0xc4, 0x2d, 0xcf, 0xbc, 0xa4, 0x5e, 0xf9, 0xcc, 0x7a, 0xfb, 0xc5,
	0x33, 0xeb, 0xde, 0xaa, 0xc4, 0xeb, 0xe1, 0xa7, 0x1d, 0xad, 0xde, 0xc2, 0x1b, 0x93, 0xea, 0x25,
	0xea, 0x69, 0xe3, 0xcd, 0xaa, 0x9d, 0xce, 0x06, 0x83, 0x15, 0xcd, 0xfa, 0xf0, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x6e, 0x19, 0xfc, 0xec, 0x05, 0x00, 0x00,
}
