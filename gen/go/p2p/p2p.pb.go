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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{0}
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{0}
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{1}
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{2}
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{3}
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{4}
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{5}
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
	// Original identifier of the document
	DocumentIdentifier []byte `protobuf:"bytes,2,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// if set, access will be limited to the specific version
	ReadVersion []byte `protobuf:"bytes,3,opt,name=read_version,json=readVersion,proto3" json:"read_version,omitempty"`
	// Signature
	// Cryptographic signature that an access token is valid
	Signature []byte `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
	// The public key of the signed message
	Key                  []byte   `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessToken) Reset()         { *m = AccessToken{} }
func (m *AccessToken) String() string { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()    {}
func (*AccessToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_ee7f52c193b68750, []int{6}
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

func (m *AccessToken) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *AccessToken) GetReadVersion() []byte {
	if m != nil {
		return m.ReadVersion
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{7}
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
	return fileDescriptor_p2p_ee7f52c193b68750, []int{8}
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

func init() { proto.RegisterFile("p2p/p2p.proto", fileDescriptor_p2p_ee7f52c193b68750) }

var fileDescriptor_p2p_ee7f52c193b68750 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x6e, 0x7a, 0x5a, 0x67, 0x92, 0x9e, 0x93, 0x6e, 0xcf, 0xa1, 0x56, 0xa9, 0x20, 0x18,
	0x81, 0x0a, 0x12, 0x6d, 0x65, 0x7e, 0xee, 0xd3, 0xd4, 0x05, 0x17, 0x48, 0xc2, 0x26, 0x14, 0xc1,
	0x8d, 0xe5, 0x7a, 0x27, 0xad, 0x55, 0xba, 0xbb, 0xac, 0xb7, 0x45, 0x79, 0x06, 0x5e, 0x82, 0xf7,
	0xe1, 0x39, 0x78, 0x0f, 0xe4, 0xb5, 0x9d, 0x38, 0x91, 0x7a, 0x03, 0x77, 0xb3, 0xdf, 0x37, 0x3b,
	0xfe, 0x66, 0xe6, 0xf3, 0xc2, 0x9a, 0xf4, 0xe4, 0x9e, 0xf4, 0xe4, 0xae, 0x54, 0x42, 0x0b, 0x52,
	0x93, 0x9e, 0xdc, 0xba, 0x1b, 0x0b, 0x85, 0x4c, 0xc4, 0x57, 0x97, 0xc8, 0xf5, 0x5e, 0xf5, 0x90,
	0x67, 0xb9, 0xdf, 0x2c, 0x58, 0x79, 0x85, 0x11, 0x43, 0x45, 0x9e, 0x00, 0xe1, 0xa8, 0xbf, 0x0a,
	0x75, 0x11, 0x26, 0x0c, 0xb9, 0x4e, 0xc6, 0x09, 0x2a, 0xc7, 0x6a, 0x5b, 0x3b, 0x6b, 0x74, 0xbd,
	0x60, 0x82, 0x29, 0x41, 0xee, 0x41, 0x93, 0x0b, 0x86, 0xe1, 0x35, 0xaa, 0x34, 0x11, 0xdc, 0x59,
	0x6a, 0x5b, 0x3b, 0x75, 0xda, 0xc8, 0xb0, 0x93, 0x1c, 0x22, 0xb7, 0xa1, 0x9e, 0x22, 0x67, 0xa8,
	0xc2, 0x84, 0x39, 0xb5, 0xb6, 0xb5, 0xd3, 0xa4, 0x76, 0x0e, 0x04, 0x8c, 0x10, 0x58, 0xd6, 0x13,
	0x89, 0xce, 0xb2, 0xb9, 0x67, 0x62, 0xb7, 0x0b, 0xb6, 0xcf, 0xaf, 0xf1, 0xb3, 0x90, 0x48, 0xee,
	0xc3, 0xca, 0xb9, 0x11, 0x66, 0x24, 0x34, 0xbc, 0xc6, 0x6e, 0xd6, 0x5b, 0xae, 0x95, 0x16, 0x54,
	0x56, 0xe4, 0x54, 0xb0, 0x89, 0xf9, 0x78, 0x93, 0x9a, 0xd8, 0x3d, 0x86, 0xd6, 0x30, 0x39, 0xe3,
	0x91, 0xbe, 0x52, 0x48, 0xf1, 0xcb, 0x15, 0xa6, 0x9a, 0xbc, 0x00, 0xbb, 0x6c, 0xbc, 0x28, 0xb7,
	0xb5, 0x3b, 0x37, 0x8d, 0xae, 0x50, 0x78, 0x58, 0x1c, 0xe8, 0x34, 0xd7, 0x3d, 0x86, 0xf5, 0x4a,
	0xad, 0x54, 0x0a, 0x9e, 0x22, 0x79, 0x0e, 0xf5, 0xb4, 0x04, 0x8b, 0x6a, 0x9b, 0xf3, 0xd5, 0x66,
	0x77, 0x66, 0x99, 0x6e, 0x1f, 0xfe, 0xef, 0xf0, 0xf8, 0x5c, 0xa8, 0xe9, 0x77, 0xfe, 0x50, 0xdc,
	0x33, 0xb8, 0xb5, 0x58, 0xb0, 0x50, 0xb8, 0x05, 0x76, 0x14, 0xc7, 0x28, 0x35, 0x32, 0x53, 0xd1,
	0xa6, 0xd3, 0xb3, 0xfb, 0xd3, 0x82, 0x46, 0x27, 0x8e, 0x31, 0x4d, 0x47, 0xe2, 0x02, 0x39, 0xb9,
	0x03, 0xb0, 0xb0, 0xee, 0x26, 0xad, 0x20, 0xc4, 0x81, 0xd5, 0x33, 0x15, 0x71, 0x8d, 0xca, 0x59,
	0x35, 0x64, 0x79, 0x9c, 0x31, 0xe8, 0xd8, 0x55, 0x06, 0xc9, 0x1e, 0x6c, 0x94, 0x2a, 0xab, 0x5e,
	0xca, 0xb7, 0x44, 0x4a, 0x6a, 0xde, 0x4c, 0x0a, 0x23, 0x36, 0x35, 0x53, 0x6e, 0x96, 0x46, 0x86,
	0x95, 0x66, 0xda, 0xae, 0x4e, 0xbd, 0x6e, 0xf8, 0x19, 0x40, 0x5a, 0x50, 0xbb, 0xc0, 0x89, 0x03,
	0x06, 0xcf, 0x42, 0xf7, 0x87, 0x05, 0xe4, 0x25, 0xea, 0xc5, 0x61, 0xdf, 0x20, 0xcd, 0xba, 0x51,
	0xda, 0x3e, 0x34, 0x22, 0x33, 0xae, 0xd0, 0xd8, 0x35, 0xeb, 0xe1, 0x1f, 0xef, 0x5f, 0x63, 0xc6,
	0x62, 0x8c, 0x13, 0x89, 0x14, 0xa2, 0x69, 0x4c, 0xf6, 0xe1, 0x3f, 0x3e, 0xd6, 0xa1, 0xc2, 0xb3,
	0x24, 0xd5, 0x6a, 0x12, 0x46, 0x8c, 0x29, 0x4c, 0xd3, 0xa2, 0x29, 0xc2, 0xc7, 0x9a, 0x16, 0x54,
	0x27, 0x67, 0x48, 0x1b, 0x9a, 0xd9, 0x0d, 0x9d, 0x2d, 0x24, 0xfb, 0x57, 0x96, 0xf3, 0x2d, 0xf0,
	0xb1, 0x36, 0x3b, 0x0a, 0x98, 0xfb, 0x16, 0x36, 0xe6, 0x9a, 0x29, 0x16, 0xfd, 0x9b, 0xd6, 0x79,
	0xfc, 0xdd, 0x02, 0x98, 0xa9, 0x27, 0x9b, 0xb0, 0xd1, 0xe9, 0x76, 0xfd, 0xe1, 0x30, 0x1c, 0x7d,
	0x1c, 0xf8, 0x61, 0xd0, 0x3b, 0xe9, 0xbc, 0x09, 0x0e, 0x5b, 0x7f, 0x91, 0x87, 0xe0, 0x56, 0x09,
	0xea, 0xbf, 0x7b, 0xef, 0x0f, 0x47, 0x3e, 0x0d, 0x4f, 0x7c, 0x1a, 0x1c, 0x05, 0xdd, 0xce, 0x28,
	0xe8, 0xf7, 0x5a, 0xd6, 0x62, 0x5e, 0xef, 0x68, 0x14, 0xf6, 0x3f, 0xf4, 0x16, 0xf3, 0x96, 0xc8,
	0x23, 0x78, 0x50, 0xcd, 0x2b, 0xe3, 0xfe, 0x6b, 0xbf, 0x37, 0x9f, 0x5a, 0x3b, 0xd8, 0x86, 0xd5,
	0x58, 0x5c, 0x66, 0x73, 0x3e, 0xb0, 0x07, 0x9e, 0x1c, 0x64, 0xcf, 0xd5, 0xc0, 0xfa, 0xf4, 0xb7,
	0xf4, 0xa4, 0x3c, 0x3d, 0x5d, 0x31, 0xcf, 0xd7, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x58,
	0x7e, 0xf1, 0xcb, 0xf5, 0x04, 0x00, 0x00,
}
