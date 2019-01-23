// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p/p2p.proto

package p2ppb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import coredocument "github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AccessType int32

const (
	AccessType_ACCESS_TYPE_INVALID                   AccessType = 0
	AccessType_ACCESS_TYPE_REQUESTER_VERIFICATION    AccessType = 1
	AccessType_ACCESS_TYPE_NFT_OWNER_VERIFICATION    AccessType = 2
	AccessType_ACCESS_TYPE_ACCESS_TOKEN_VERIFICATION AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "ACCESS_TYPE_INVALID",
	1: "ACCESS_TYPE_REQUESTER_VERIFICATION",
	2: "ACCESS_TYPE_NFT_OWNER_VERIFICATION",
	3: "ACCESS_TYPE_ACCESS_TOKEN_VERIFICATION",
}
var AccessType_value = map[string]int32{
	"ACCESS_TYPE_INVALID":                   0,
	"ACCESS_TYPE_REQUESTER_VERIFICATION":    1,
	"ACCESS_TYPE_NFT_OWNER_VERIFICATION":    2,
	"ACCESS_TYPE_ACCESS_TOKEN_VERIFICATION": 3,
}

func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}
func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{0}
}

type Header struct {
	NetworkIdentifier uint32 `protobuf:"varint,1,opt,name=network_identifier,json=networkIdentifier" json:"network_identifier,omitempty"`
	NodeVersion       string `protobuf:"bytes,2,opt,name=node_version,json=nodeVersion" json:"node_version,omitempty"`
	SenderId          []byte `protobuf:"bytes,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// Body message type
	Type                 string   `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{0}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetNetworkIdentifier() uint32 {
	if m != nil {
		return m.NetworkIdentifier
	}
	return 0
}

func (m *Header) GetNodeVersion() string {
	if m != nil {
		return m.NodeVersion
	}
	return ""
}

func (m *Header) GetSenderId() []byte {
	if m != nil {
		return m.SenderId
	}
	return nil
}

func (m *Header) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Envelope struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{1}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Envelope) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SignatureRequest struct {
	Document             *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SignatureRequest) Reset()         { *m = SignatureRequest{} }
func (m *SignatureRequest) String() string { return proto.CompactTextString(m) }
func (*SignatureRequest) ProtoMessage()    {}
func (*SignatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{2}
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

func (m *SignatureRequest) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type SignatureResponse struct {
	Signature            *coredocument.Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SignatureResponse) Reset()         { *m = SignatureResponse{} }
func (m *SignatureResponse) String() string { return proto.CompactTextString(m) }
func (*SignatureResponse) ProtoMessage()    {}
func (*SignatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{3}
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

func (m *SignatureResponse) GetSignature() *coredocument.Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AnchorDocumentRequest struct {
	Document             *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AnchorDocumentRequest) Reset()         { *m = AnchorDocumentRequest{} }
func (m *AnchorDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*AnchorDocumentRequest) ProtoMessage()    {}
func (*AnchorDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{4}
}
func (m *AnchorDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchorDocumentRequest.Unmarshal(m, b)
}
func (m *AnchorDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchorDocumentRequest.Marshal(b, m, deterministic)
}
func (dst *AnchorDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorDocumentRequest.Merge(dst, src)
}
func (m *AnchorDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_AnchorDocumentRequest.Size(m)
}
func (m *AnchorDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorDocumentRequest proto.InternalMessageInfo

func (m *AnchorDocumentRequest) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

type AnchorDocumentResponse struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnchorDocumentResponse) Reset()         { *m = AnchorDocumentResponse{} }
func (m *AnchorDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*AnchorDocumentResponse) ProtoMessage()    {}
func (*AnchorDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{5}
}
func (m *AnchorDocumentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnchorDocumentResponse.Unmarshal(m, b)
}
func (m *AnchorDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnchorDocumentResponse.Marshal(b, m, deterministic)
}
func (dst *AnchorDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnchorDocumentResponse.Merge(dst, src)
}
func (m *AnchorDocumentResponse) XXX_Size() int {
	return xxx_messageInfo_AnchorDocumentResponse.Size(m)
}
func (m *AnchorDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnchorDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnchorDocumentResponse proto.InternalMessageInfo

func (m *AnchorDocumentResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

type AccessToken struct {
	// The identifier is an internal 256bit word
	Identifier []byte `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// The identity granting access to the document
	Granter []byte `protobuf:"bytes,7,opt,name=granter,proto3" json:"granter,omitempty"`
	// The identity being granted access to the document
	Grantee []byte `protobuf:"bytes,8,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Role identifier is the identifier on the read rule that this token should be mapped to
	RoleIdentifier []byte `protobuf:"bytes,9,opt,name=role_identifier,json=roleIdentifier,proto3" json:"role_identifier,omitempty"`
	// Original identifier of the document
	DocumentIdentifier []byte `protobuf:"bytes,2,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// Signature
	// Cryptographic signature that an access token is valid
	Signature []byte `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`
	// The public key of the signed message
	Key                  []byte   `protobuf:"bytes,11,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessToken) Reset()         { *m = AccessToken{} }
func (m *AccessToken) String() string { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()    {}
func (*AccessToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{6}
}
func (m *AccessToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessToken.Unmarshal(m, b)
}
func (m *AccessToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessToken.Marshal(b, m, deterministic)
}
func (dst *AccessToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessToken.Merge(dst, src)
}
func (m *AccessToken) XXX_Size() int {
	return xxx_messageInfo_AccessToken.Size(m)
}
func (m *AccessToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessToken.DiscardUnknown(m)
}

var xxx_messageInfo_AccessToken proto.InternalMessageInfo

func (m *AccessToken) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *AccessToken) GetGranter() []byte {
	if m != nil {
		return m.Granter
	}
	return nil
}

func (m *AccessToken) GetGrantee() []byte {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *AccessToken) GetRoleIdentifier() []byte {
	if m != nil {
		return m.RoleIdentifier
	}
	return nil
}

func (m *AccessToken) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *AccessToken) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AccessToken) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetDocumentRequest struct {
	DocumentIdentifier []byte `protobuf:"bytes,1,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// Version of document? or should versioning be handled one layer down in AccessTypes?
	AccessType           AccessType `protobuf:"varint,2,opt,name=access_type,json=accessType,enum=p2p.AccessType" json:"access_type,omitempty"`
	NftRegistryAddress   []byte     `protobuf:"bytes,3,opt,name=nft_registry_address,json=nftRegistryAddress,proto3" json:"nft_registry_address,omitempty"`
	NftTokenId           []byte     `protobuf:"bytes,4,opt,name=nft_token_id,json=nftTokenId,proto3" json:"nft_token_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetDocumentRequest) Reset()         { *m = GetDocumentRequest{} }
func (m *GetDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentRequest) ProtoMessage()    {}
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{7}
}
func (m *GetDocumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentRequest.Unmarshal(m, b)
}
func (m *GetDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentRequest.Marshal(b, m, deterministic)
}
func (dst *GetDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentRequest.Merge(dst, src)
}
func (m *GetDocumentRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentRequest.Size(m)
}
func (m *GetDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentRequest proto.InternalMessageInfo

func (m *GetDocumentRequest) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *GetDocumentRequest) GetAccessType() AccessType {
	if m != nil {
		return m.AccessType
	}
	return AccessType_ACCESS_TYPE_INVALID
}

func (m *GetDocumentRequest) GetNftRegistryAddress() []byte {
	if m != nil {
		return m.NftRegistryAddress
	}
	return nil
}

func (m *GetDocumentRequest) GetNftTokenId() []byte {
	if m != nil {
		return m.NftTokenId
	}
	return nil
}

type GetDocumentResponse struct {
	Document             *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetDocumentResponse) Reset()         { *m = GetDocumentResponse{} }
func (m *GetDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocumentResponse) ProtoMessage()    {}
func (*GetDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_e44b22dcbdebfb1b, []int{8}
}
func (m *GetDocumentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentResponse.Unmarshal(m, b)
}
func (m *GetDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentResponse.Marshal(b, m, deterministic)
}
func (dst *GetDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentResponse.Merge(dst, src)
}
func (m *GetDocumentResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocumentResponse.Size(m)
}
func (m *GetDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentResponse proto.InternalMessageInfo

func (m *GetDocumentResponse) GetDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.Document
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "p2p.Header")
	proto.RegisterType((*Envelope)(nil), "p2p.Envelope")
	proto.RegisterType((*SignatureRequest)(nil), "p2p.SignatureRequest")
	proto.RegisterType((*SignatureResponse)(nil), "p2p.SignatureResponse")
	proto.RegisterType((*AnchorDocumentRequest)(nil), "p2p.AnchorDocumentRequest")
	proto.RegisterType((*AnchorDocumentResponse)(nil), "p2p.AnchorDocumentResponse")
	proto.RegisterType((*AccessToken)(nil), "p2p.AccessToken")
	proto.RegisterType((*GetDocumentRequest)(nil), "p2p.GetDocumentRequest")
	proto.RegisterType((*GetDocumentResponse)(nil), "p2p.GetDocumentResponse")
	proto.RegisterEnum("p2p.AccessType", AccessType_name, AccessType_value)
}

func init() { proto.RegisterFile("p2p/p2p.proto", fileDescriptor_p2p_e44b22dcbdebfb1b) }

var fileDescriptor_p2p_e44b22dcbdebfb1b = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdf, 0x52, 0xd3, 0x4e,
	0x14, 0xfe, 0x85, 0xf2, 0x83, 0xf4, 0xb4, 0x40, 0x59, 0x54, 0x32, 0xc8, 0x68, 0x8d, 0xa3, 0xa2,
	0x33, 0x02, 0x13, 0xff, 0xdc, 0x97, 0x12, 0x34, 0xa8, 0x6d, 0xdd, 0x56, 0x1c, 0xbd, 0xc9, 0x84,
	0xe4, 0x14, 0x32, 0xc0, 0xee, 0xba, 0xd9, 0xe2, 0xf4, 0x19, 0x7c, 0x09, 0xdf, 0xc7, 0x47, 0xf1,
	0x25, 0x9c, 0x6c, 0x92, 0x36, 0xed, 0x0c, 0x37, 0x7a, 0x77, 0xf6, 0xfb, 0xbe, 0xf3, 0xff, 0x24,
	0xb0, 0x22, 0x1c, 0xb1, 0x27, 0x1c, 0xb1, 0x2b, 0x24, 0x57, 0x9c, 0x54, 0x84, 0x23, 0xb6, 0xee,
	0x87, 0x5c, 0x62, 0xc4, 0xc3, 0xd1, 0x15, 0x32, 0xb5, 0x57, 0x7e, 0x64, 0x2a, 0xfb, 0x87, 0x01,
	0x4b, 0x6f, 0x31, 0x88, 0x50, 0x92, 0xe7, 0x40, 0x18, 0xaa, 0xef, 0x5c, 0x5e, 0xf8, 0x71, 0x84,
	0x4c, 0xc5, 0xc3, 0x18, 0xa5, 0x65, 0x34, 0x8d, 0x9d, 0x15, 0xba, 0x9e, 0x33, 0xde, 0x84, 0x20,
	0x0f, 0xa0, 0xce, 0x78, 0x84, 0xfe, 0x35, 0xca, 0x24, 0xe6, 0xcc, 0x5a, 0x68, 0x1a, 0x3b, 0x55,
	0x5a, 0x4b, 0xb1, 0x93, 0x0c, 0x22, 0x77, 0xa1, 0x9a, 0x20, 0x8b, 0x50, 0xfa, 0x71, 0x64, 0x55,
	0x9a, 0xc6, 0x4e, 0x9d, 0x9a, 0x19, 0xe0, 0x45, 0x84, 0xc0, 0xa2, 0x1a, 0x0b, 0xb4, 0x16, 0xb5,
	0x9f, 0xb6, 0xed, 0x36, 0x98, 0x2e, 0xbb, 0xc6, 0x4b, 0x2e, 0x90, 0x3c, 0x84, 0xa5, 0x73, 0x5d,
	0x98, 0x2e, 0xa1, 0xe6, 0xd4, 0x76, 0xd3, 0xde, 0xb2, 0x5a, 0x69, 0x4e, 0xa5, 0x41, 0x4e, 0x79,
	0x34, 0xd6, 0xc9, 0xeb, 0x54, 0xdb, 0xf6, 0x31, 0x34, 0xfa, 0xf1, 0x19, 0x0b, 0xd4, 0x48, 0x22,
	0xc5, 0x6f, 0x23, 0x4c, 0x14, 0x79, 0x0d, 0x66, 0xd1, 0x78, 0x1e, 0x6e, 0x6b, 0x77, 0x66, 0x1a,
	0x6d, 0x2e, 0xf1, 0x30, 0x7f, 0xd0, 0x89, 0xd6, 0x3e, 0x86, 0xf5, 0x52, 0xac, 0x44, 0x70, 0x96,
	0x20, 0x79, 0x05, 0xd5, 0xa4, 0x00, 0xf3, 0x68, 0x9b, 0xb3, 0xd1, 0xa6, 0x3e, 0x53, 0xa5, 0xdd,
	0x85, 0xdb, 0x2d, 0x16, 0x9e, 0x73, 0x39, 0xc9, 0xf3, 0x8f, 0xc5, 0xbd, 0x84, 0x3b, 0xf3, 0x01,
	0xf3, 0x0a, 0xb7, 0xc0, 0x0c, 0xc2, 0x10, 0x85, 0xc2, 0x48, 0x47, 0x34, 0xe9, 0xe4, 0x6d, 0xff,
	0x36, 0xa0, 0xd6, 0x0a, 0x43, 0x4c, 0x92, 0x01, 0xbf, 0x40, 0x46, 0xee, 0x01, 0xcc, 0xad, 0xbb,
	0x4e, 0x4b, 0x08, 0xb1, 0x60, 0xf9, 0x4c, 0x06, 0x4c, 0xa1, 0xb4, 0x96, 0x35, 0x59, 0x3c, 0xa7,
	0x0c, 0x5a, 0x66, 0x99, 0x41, 0xf2, 0x04, 0xd6, 0x24, 0xbf, 0xc4, 0xf2, 0x1d, 0x55, 0xb5, 0x62,
	0x35, 0x85, 0x4b, 0x47, 0xb4, 0x07, 0x1b, 0x45, 0x3b, 0x65, 0x71, 0xb6, 0x4e, 0x52, 0x50, 0x25,
	0x87, 0xed, 0xf2, 0xec, 0x41, 0xcb, 0xa6, 0x00, 0x69, 0x40, 0xe5, 0x02, 0xc7, 0x56, 0x4d, 0xe3,
	0xa9, 0x69, 0xff, 0x32, 0x80, 0xbc, 0x41, 0x35, 0x3f, 0xf2, 0x1b, 0xf2, 0x1a, 0x37, 0xe6, 0xdd,
	0x87, 0x5a, 0xa0, 0x87, 0xe6, 0xeb, 0xa3, 0x4d, 0x0b, 0x5c, 0x75, 0xd6, 0xf4, 0x49, 0xe6, 0xc3,
	0x1c, 0x0b, 0xa4, 0x10, 0x4c, 0x6c, 0xb2, 0x0f, 0xb7, 0xd8, 0x50, 0xf9, 0x12, 0xcf, 0xe2, 0x44,
	0xc9, 0xb1, 0x1f, 0x44, 0x91, 0xc4, 0x24, 0xc9, 0xbf, 0x03, 0xc2, 0x86, 0x8a, 0xe6, 0x54, 0x2b,
	0x63, 0x48, 0x13, 0xea, 0xa9, 0x87, 0x4a, 0xd7, 0x92, 0x7e, 0x31, 0x8b, 0xd9, 0x2e, 0xd8, 0x50,
	0xe9, 0x4d, 0x79, 0x91, 0xfd, 0x01, 0x36, 0x66, 0x9a, 0xc9, 0xd7, 0xfd, 0x97, 0x07, 0xf4, 0xec,
	0xa7, 0x01, 0x30, 0xad, 0x9e, 0x6c, 0xc2, 0x46, 0xab, 0xdd, 0x76, 0xfb, 0x7d, 0x7f, 0xf0, 0xa5,
	0xe7, 0xfa, 0x5e, 0xe7, 0xa4, 0xf5, 0xde, 0x3b, 0x6c, 0xfc, 0x47, 0x1e, 0x83, 0x5d, 0x26, 0xa8,
	0xfb, 0xf1, 0x93, 0xdb, 0x1f, 0xb8, 0xd4, 0x3f, 0x71, 0xa9, 0x77, 0xe4, 0xb5, 0x5b, 0x03, 0xaf,
	0xdb, 0x69, 0x18, 0xf3, 0xba, 0xce, 0xd1, 0xc0, 0xef, 0x7e, 0xee, 0xcc, 0xeb, 0x16, 0xc8, 0x53,
	0x78, 0x54, 0xd6, 0x15, 0x76, 0xf7, 0x9d, 0xdb, 0x99, 0x95, 0x56, 0x0e, 0xb6, 0x61, 0x39, 0xe4,
	0x57, 0xe9, 0x9c, 0x0f, 0xcc, 0x9e, 0x23, 0x7a, 0xe9, 0x4f, 0xab, 0x67, 0x7c, 0xfd, 0x5f, 0x38,
	0x42, 0x9c, 0x9e, 0x2e, 0xe9, 0x9f, 0xd8, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x26, 0xa1,
	0xdf, 0x34, 0xfb, 0x04, 0x00, 0x00,
}
