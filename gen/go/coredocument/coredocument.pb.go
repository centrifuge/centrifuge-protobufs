// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coredocument/coredocument.proto

package coredocumentpb

import (
	fmt "fmt"
	proto1 "github.com/centrifuge/precise-proofs/proofs/proto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Action defines the set of actions a collaborator can/have per document.
type Action int32

const (
	Action_ACTION_INVALID Action = 0
	// read_sign represents reading as well the sign the documents. We will pick this one when requesting the signatures.
	Action_ACTION_READ_SIGN Action = 1
	// read represents just reading the doc/fields
	Action_ACTION_READ Action = 2
)

var Action_name = map[int32]string{
	0: "ACTION_INVALID",
	1: "ACTION_READ_SIGN",
	2: "ACTION_READ",
}

var Action_value = map[string]int32{
	"ACTION_INVALID":   0,
	"ACTION_READ_SIGN": 1,
	"ACTION_READ":      2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{0}
}

type FieldMatchType int32

const (
	FieldMatchType_FIELD_MATCH_TYPE_INVALID FieldMatchType = 0
	FieldMatchType_FIELD_MATCH_TYPE_PREFIX  FieldMatchType = 1
	FieldMatchType_FIELD_MATCH_TYPE_EXACT   FieldMatchType = 2
)

var FieldMatchType_name = map[int32]string{
	0: "FIELD_MATCH_TYPE_INVALID",
	1: "FIELD_MATCH_TYPE_PREFIX",
	2: "FIELD_MATCH_TYPE_EXACT",
}

var FieldMatchType_value = map[string]int32{
	"FIELD_MATCH_TYPE_INVALID": 0,
	"FIELD_MATCH_TYPE_PREFIX":  1,
	"FIELD_MATCH_TYPE_EXACT":   2,
}

func (x FieldMatchType) String() string {
	return proto.EnumName(FieldMatchType_name, int32(x))
}

func (FieldMatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{1}
}

type TransitionAction int32

const (
	TransitionAction_TRANSITION_ACTION_INVALID TransitionAction = 0
	TransitionAction_TRANSITION_ACTION_EDIT    TransitionAction = 1
)

var TransitionAction_name = map[int32]string{
	0: "TRANSITION_ACTION_INVALID",
	1: "TRANSITION_ACTION_EDIT",
}

var TransitionAction_value = map[string]int32{
	"TRANSITION_ACTION_INVALID": 0,
	"TRANSITION_ACTION_EDIT":    1,
}

func (x TransitionAction) String() string {
	return proto.EnumName(TransitionAction_name, int32(x))
}

func (TransitionAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{2}
}

type FieldType int32

const (
	FieldType_FIELD_TYPE_INVALID   FieldType = 0
	FieldType_FIELD_TYPE_INTEGER   FieldType = 1
	FieldType_FIELD_TYPE_DECIMAL   FieldType = 2
	FieldType_FIELD_TYPE_STRING    FieldType = 3
	FieldType_FIELD_TYPE_BYTES     FieldType = 4
	FieldType_FIELD_TYPE_TIMESTAMP FieldType = 5
)

var FieldType_name = map[int32]string{
	0: "FIELD_TYPE_INVALID",
	1: "FIELD_TYPE_INTEGER",
	2: "FIELD_TYPE_DECIMAL",
	3: "FIELD_TYPE_STRING",
	4: "FIELD_TYPE_BYTES",
	5: "FIELD_TYPE_TIMESTAMP",
}

var FieldType_value = map[string]int32{
	"FIELD_TYPE_INVALID":   0,
	"FIELD_TYPE_INTEGER":   1,
	"FIELD_TYPE_DECIMAL":   2,
	"FIELD_TYPE_STRING":    3,
	"FIELD_TYPE_BYTES":     4,
	"FIELD_TYPE_TIMESTAMP": 5,
}

func (x FieldType) String() string {
	return proto.EnumName(FieldType_name, int32(x))
}

func (FieldType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{3}
}

// `CoreDocument` is a document that can be sent to different nodes and anchored
// on chain. It handles all the generic features native Centrifuge Documents support:
//
// * Merkle Roots for the document data
// * Signatures on document data
// * Access Control
type CoreDocument struct {
	// # Identifiers
	// CoreDocument has two kinds of identifiers, the `document_identifier` is assigned
	// once per document and stays the same for the lifetime of the document.
	// document_identifier is the first ID ever used to anchor the document on chain and
	// is used internally to store all future versions. The `previous_version`, `current_version`, and the
	// `next_version` refer only to a particular version.
	//
	// 32 byte value
	DocumentIdentifier []byte `protobuf:"bytes,9,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// previous_version refers to the previous state of the document.
	// 32 byte value
	PreviousVersion []byte `protobuf:"bytes,16,opt,name=previous_version,json=previousVersion,proto3" json:"previous_version,omitempty"`
	// current_version is the version used to refer to the current state of the document.
	// 32 byte value
	CurrentVersion []byte `protobuf:"bytes,3,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	// current_preimage is the sha256 pre-image of the current_version. It prevents current state commitment id(anchor id) from getting exposed.
	// 32 byte value
	CurrentPreimage []byte `protobuf:"bytes,23,opt,name=current_preimage,json=currentPreimage,proto3" json:"current_preimage,omitempty"`
	// next_version is the version that is going to be used for the next version if any
	// party wants to update the state.
	NextVersion []byte `protobuf:"bytes,4,opt,name=next_version,json=nextVersion,proto3" json:"next_version,omitempty"`
	// next_preimage is the sha256 pre-image of the next_version. It prevents next state commitment id(anchor id) from getting exposed.
	NextPreimage []byte `protobuf:"bytes,22,opt,name=next_preimage,json=nextPreimage,proto3" json:"next_preimage,omitempty"`
	// Signatures of the signature_root by collaborators on the document.
	SignatureData *SignatureData `protobuf:"bytes,6,opt,name=signature_data,json=signatureData,proto3" json:"signature_data,omitempty"`
	// When a document is transmitted over the wire, the type specific fields (e.g. InvoiceData) are
	// embedded in the document using the google.protobuf.Any type.
	EmbeddedData *any.Any       `protobuf:"bytes,13,opt,name=embedded_data,json=embeddedData,proto3" json:"embedded_data,omitempty"`
	Salts        []*proto1.Salt `protobuf:"bytes,15,rep,name=salts,proto3" json:"salts,omitempty"`
	// list of roles
	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	// read_rules define who may read a document and who should sign it
	ReadRules []*ReadRule `protobuf:"bytes,19,rep,name=read_rules,json=readRules,proto3" json:"read_rules,omitempty"`
	// transition rules define how a document may be manipulated
	TransitionRules []*TransitionRule `protobuf:"bytes,24,rep,name=transition_rules,json=transitionRules,proto3" json:"transition_rules,omitempty"`
	// nft list for uniqueness check
	Nfts []*NFT `protobuf:"bytes,20,rep,name=nfts,proto3" json:"nfts,omitempty"`
	// AccessTokens which have been added to this CoreDocument
	AccessTokens []*AccessToken `protobuf:"bytes,21,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
	// author of the latest update
	Author []byte `protobuf:"bytes,25,opt,name=author,proto3" json:"author,omitempty"`
	// timestamp of the latest update
	Timestamp *timestamp.Timestamp `protobuf:"bytes,26,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// anchor repository address used to anchor this document
	AnchorRepositoryUsed []byte `protobuf:"bytes,27,opt,name=anchor_repository_used,json=anchorRepositoryUsed,proto3" json:"anchor_repository_used,omitempty"`
	// custom attributes(user defined fields) for this document
	Attributes           []*Field `protobuf:"bytes,42,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoreDocument) Reset()         { *m = CoreDocument{} }
func (m *CoreDocument) String() string { return proto.CompactTextString(m) }
func (*CoreDocument) ProtoMessage()    {}
func (*CoreDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{0}
}

func (m *CoreDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreDocument.Unmarshal(m, b)
}
func (m *CoreDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreDocument.Marshal(b, m, deterministic)
}
func (m *CoreDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreDocument.Merge(m, src)
}
func (m *CoreDocument) XXX_Size() int {
	return xxx_messageInfo_CoreDocument.Size(m)
}
func (m *CoreDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreDocument.DiscardUnknown(m)
}

var xxx_messageInfo_CoreDocument proto.InternalMessageInfo

func (m *CoreDocument) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *CoreDocument) GetPreviousVersion() []byte {
	if m != nil {
		return m.PreviousVersion
	}
	return nil
}

func (m *CoreDocument) GetCurrentVersion() []byte {
	if m != nil {
		return m.CurrentVersion
	}
	return nil
}

func (m *CoreDocument) GetCurrentPreimage() []byte {
	if m != nil {
		return m.CurrentPreimage
	}
	return nil
}

func (m *CoreDocument) GetNextVersion() []byte {
	if m != nil {
		return m.NextVersion
	}
	return nil
}

func (m *CoreDocument) GetNextPreimage() []byte {
	if m != nil {
		return m.NextPreimage
	}
	return nil
}

func (m *CoreDocument) GetSignatureData() *SignatureData {
	if m != nil {
		return m.SignatureData
	}
	return nil
}

func (m *CoreDocument) GetEmbeddedData() *any.Any {
	if m != nil {
		return m.EmbeddedData
	}
	return nil
}

func (m *CoreDocument) GetSalts() []*proto1.Salt {
	if m != nil {
		return m.Salts
	}
	return nil
}

func (m *CoreDocument) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *CoreDocument) GetReadRules() []*ReadRule {
	if m != nil {
		return m.ReadRules
	}
	return nil
}

func (m *CoreDocument) GetTransitionRules() []*TransitionRule {
	if m != nil {
		return m.TransitionRules
	}
	return nil
}

func (m *CoreDocument) GetNfts() []*NFT {
	if m != nil {
		return m.Nfts
	}
	return nil
}

func (m *CoreDocument) GetAccessTokens() []*AccessToken {
	if m != nil {
		return m.AccessTokens
	}
	return nil
}

func (m *CoreDocument) GetAuthor() []byte {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *CoreDocument) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *CoreDocument) GetAnchorRepositoryUsed() []byte {
	if m != nil {
		return m.AnchorRepositoryUsed
	}
	return nil
}

func (m *CoreDocument) GetAttributes() []*Field {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type AccessToken struct {
	// The identifier is an internal 256bit word
	Identifier []byte `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// The identity granting access to the document
	Granter []byte `protobuf:"bytes,3,opt,name=granter,proto3" json:"granter,omitempty"`
	// The identity being granted access to the document
	Grantee []byte `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Role identifier is the identifier on the read rule that this token should be mapped to
	RoleIdentifier []byte `protobuf:"bytes,5,opt,name=role_identifier,json=roleIdentifier,proto3" json:"role_identifier,omitempty"`
	// Original identifier of the document
	DocumentIdentifier []byte `protobuf:"bytes,2,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// Cryptographic signature that an access token is valid
	Signature []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	// The public key of the signed message
	Key []byte `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	// The document version refers to a version of the document this token is embedded in. Its timestamp
	// will be used to verify the validity of the signature of the access token.
	DocumentVersion      []byte   `protobuf:"bytes,8,opt,name=document_version,json=documentVersion,proto3" json:"document_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessToken) Reset()         { *m = AccessToken{} }
func (m *AccessToken) String() string { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()    {}
func (*AccessToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{1}
}

func (m *AccessToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessToken.Unmarshal(m, b)
}
func (m *AccessToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessToken.Marshal(b, m, deterministic)
}
func (m *AccessToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessToken.Merge(m, src)
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

func (m *AccessToken) GetDocumentVersion() []byte {
	if m != nil {
		return m.DocumentVersion
	}
	return nil
}

// SignatureData contains the list of signatures identified by the signature_id
type SignatureData struct {
	Signatures           []*Signature `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SignatureData) Reset()         { *m = SignatureData{} }
func (m *SignatureData) String() string { return proto.CompactTextString(m) }
func (*SignatureData) ProtoMessage()    {}
func (*SignatureData) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{2}
}

func (m *SignatureData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureData.Unmarshal(m, b)
}
func (m *SignatureData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureData.Marshal(b, m, deterministic)
}
func (m *SignatureData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureData.Merge(m, src)
}
func (m *SignatureData) XXX_Size() int {
	return xxx_messageInfo_SignatureData.Size(m)
}
func (m *SignatureData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureData.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureData proto.InternalMessageInfo

func (m *SignatureData) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// Signature contains the entity ID, public key used and signature)
type Signature struct {
	// `signature_id` is a composed key signer_id+public_key used
	SignatureId []byte `protobuf:"bytes,1,opt,name=signature_id,json=signatureId,proto3" json:"signature_id,omitempty"`
	// `signer_id` is the CentrifugeID of the identity signing the document.
	SignerId []byte `protobuf:"bytes,2,opt,name=signer_id,json=signerId,proto3" json:"signer_id,omitempty"`
	// `public_key` is the public key of the `signer` used for signing the `CoreDocument`
	PublicKey []byte `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// `signature` is the actual signature of the CoreDocument
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{3}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetSignatureId() []byte {
	if m != nil {
		return m.SignatureId
	}
	return nil
}

func (m *Signature) GetSignerId() []byte {
	if m != nil {
		return m.SignerId
	}
	return nil
}

func (m *Signature) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Roles holds a list of collaborators, NFTs, and/or access tokens.
type Role struct {
	// role key which is used to identify the group internally and map the role to rules
	RoleKey []byte `protobuf:"bytes,1,opt,name=role_key,json=roleKey,proto3" json:"role_key,omitempty"`
	// collaborators holds the list of document collaborators
	Collaborators [][]byte `protobuf:"bytes,3,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	// nfts is a list of registry address/tokenID pairs.
	// For easier verification in merkle proofs, the values are simply concatenated with the first 22 bytes being the NFT registry and the remaining 32 bytes the tokenID.
	Nfts                 [][]byte `protobuf:"bytes,4,rep,name=nfts,proto3" json:"nfts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{4}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetRoleKey() []byte {
	if m != nil {
		return m.RoleKey
	}
	return nil
}

func (m *Role) GetCollaborators() [][]byte {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *Role) GetNfts() [][]byte {
	if m != nil {
		return m.Nfts
	}
	return nil
}

type ReadRule struct {
	Roles                [][]byte `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Action               Action   `protobuf:"varint,4,opt,name=action,proto3,enum=coredocument.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRule) Reset()         { *m = ReadRule{} }
func (m *ReadRule) String() string { return proto.CompactTextString(m) }
func (*ReadRule) ProtoMessage()    {}
func (*ReadRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{5}
}

func (m *ReadRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRule.Unmarshal(m, b)
}
func (m *ReadRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRule.Marshal(b, m, deterministic)
}
func (m *ReadRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRule.Merge(m, src)
}
func (m *ReadRule) XXX_Size() int {
	return xxx_messageInfo_ReadRule.Size(m)
}
func (m *ReadRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRule.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRule proto.InternalMessageInfo

func (m *ReadRule) GetRoles() [][]byte {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ReadRule) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_ACTION_INVALID
}

type TransitionRule struct {
	// rule key, to help track of the rule
	RuleKey []byte `protobuf:"bytes,1,opt,name=rule_key,json=ruleKey,proto3" json:"rule_key,omitempty"`
	// Indicates which roles can make changes or read the fields specified:
	// this list holds role keys correlated to those in the 'roles' field of the CoreDocument
	Roles [][]byte `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	// prefix or exact
	MatchType FieldMatchType `protobuf:"varint,3,opt,name=match_type,json=matchType,proto3,enum=coredocument.FieldMatchType" json:"match_type,omitempty"`
	// compact property of the field
	Field []byte `protobuf:"bytes,4,opt,name=field,proto3" json:"field,omitempty"`
	// what kind of action this rule allows
	Action               TransitionAction `protobuf:"varint,5,opt,name=action,proto3,enum=coredocument.TransitionAction" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TransitionRule) Reset()         { *m = TransitionRule{} }
func (m *TransitionRule) String() string { return proto.CompactTextString(m) }
func (*TransitionRule) ProtoMessage()    {}
func (*TransitionRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{6}
}

func (m *TransitionRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionRule.Unmarshal(m, b)
}
func (m *TransitionRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionRule.Marshal(b, m, deterministic)
}
func (m *TransitionRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionRule.Merge(m, src)
}
func (m *TransitionRule) XXX_Size() int {
	return xxx_messageInfo_TransitionRule.Size(m)
}
func (m *TransitionRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionRule.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionRule proto.InternalMessageInfo

func (m *TransitionRule) GetRuleKey() []byte {
	if m != nil {
		return m.RuleKey
	}
	return nil
}

func (m *TransitionRule) GetRoles() [][]byte {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *TransitionRule) GetMatchType() FieldMatchType {
	if m != nil {
		return m.MatchType
	}
	return FieldMatchType_FIELD_MATCH_TYPE_INVALID
}

func (m *TransitionRule) GetField() []byte {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *TransitionRule) GetAction() TransitionAction {
	if m != nil {
		return m.Action
	}
	return TransitionAction_TRANSITION_ACTION_INVALID
}

type NFT struct {
	RegistryId           []byte   `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	TokenId              []byte   `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NFT) Reset()         { *m = NFT{} }
func (m *NFT) String() string { return proto.CompactTextString(m) }
func (*NFT) ProtoMessage()    {}
func (*NFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{7}
}

func (m *NFT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NFT.Unmarshal(m, b)
}
func (m *NFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NFT.Marshal(b, m, deterministic)
}
func (m *NFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFT.Merge(m, src)
}
func (m *NFT) XXX_Size() int {
	return xxx_messageInfo_NFT.Size(m)
}
func (m *NFT) XXX_DiscardUnknown() {
	xxx_messageInfo_NFT.DiscardUnknown(m)
}

var xxx_messageInfo_NFT proto.InternalMessageInfo

func (m *NFT) GetRegistryId() []byte {
	if m != nil {
		return m.RegistryId
	}
	return nil
}

func (m *NFT) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

// Field represents a custom attribute
type Field struct {
	Key      []byte    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	KeyLabel []byte    `protobuf:"bytes,2,opt,name=key_label,json=keyLabel,proto3" json:"key_label,omitempty"`
	Type     FieldType `protobuf:"varint,3,opt,name=type,proto3,enum=coredocument.FieldType" json:"type,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Field_StrVal
	//	*Field_ByteVal
	//	*Field_TimeVal
	Value                isField_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb191193ab8e2cd4, []int{8}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Field) GetKeyLabel() []byte {
	if m != nil {
		return m.KeyLabel
	}
	return nil
}

func (m *Field) GetType() FieldType {
	if m != nil {
		return m.Type
	}
	return FieldType_FIELD_TYPE_INVALID
}

type isField_Value interface {
	isField_Value()
}

type Field_StrVal struct {
	StrVal string `protobuf:"bytes,4,opt,name=str_val,json=strVal,proto3,oneof"`
}

type Field_ByteVal struct {
	ByteVal []byte `protobuf:"bytes,5,opt,name=byte_val,json=byteVal,proto3,oneof"`
}

type Field_TimeVal struct {
	TimeVal *timestamp.Timestamp `protobuf:"bytes,6,opt,name=time_val,json=timeVal,proto3,oneof"`
}

func (*Field_StrVal) isField_Value() {}

func (*Field_ByteVal) isField_Value() {}

func (*Field_TimeVal) isField_Value() {}

func (m *Field) GetValue() isField_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Field) GetStrVal() string {
	if x, ok := m.GetValue().(*Field_StrVal); ok {
		return x.StrVal
	}
	return ""
}

func (m *Field) GetByteVal() []byte {
	if x, ok := m.GetValue().(*Field_ByteVal); ok {
		return x.ByteVal
	}
	return nil
}

func (m *Field) GetTimeVal() *timestamp.Timestamp {
	if x, ok := m.GetValue().(*Field_TimeVal); ok {
		return x.TimeVal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Field) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Field_StrVal)(nil),
		(*Field_ByteVal)(nil),
		(*Field_TimeVal)(nil),
	}
}

func init() {
	proto.RegisterEnum("coredocument.Action", Action_name, Action_value)
	proto.RegisterEnum("coredocument.FieldMatchType", FieldMatchType_name, FieldMatchType_value)
	proto.RegisterEnum("coredocument.TransitionAction", TransitionAction_name, TransitionAction_value)
	proto.RegisterEnum("coredocument.FieldType", FieldType_name, FieldType_value)
	proto.RegisterType((*CoreDocument)(nil), "coredocument.CoreDocument")
	proto.RegisterType((*AccessToken)(nil), "coredocument.AccessToken")
	proto.RegisterType((*SignatureData)(nil), "coredocument.SignatureData")
	proto.RegisterType((*Signature)(nil), "coredocument.Signature")
	proto.RegisterType((*Role)(nil), "coredocument.Role")
	proto.RegisterType((*ReadRule)(nil), "coredocument.ReadRule")
	proto.RegisterType((*TransitionRule)(nil), "coredocument.TransitionRule")
	proto.RegisterType((*NFT)(nil), "coredocument.NFT")
	proto.RegisterType((*Field)(nil), "coredocument.Field")
}

func init() { proto.RegisterFile("coredocument/coredocument.proto", fileDescriptor_eb191193ab8e2cd4) }

var fileDescriptor_eb191193ab8e2cd4 = []byte{
	// 1270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcb, 0x72, 0x1b, 0x45,
	0x14, 0xcd, 0x58, 0x2f, 0xeb, 0xea, 0xe1, 0x49, 0x47, 0x91, 0xc7, 0x76, 0x1e, 0x46, 0x50, 0xc4,
	0x31, 0x20, 0x57, 0x85, 0x40, 0xa8, 0x82, 0x05, 0x7a, 0xd9, 0x99, 0xc2, 0x56, 0x54, 0xa3, 0xc1,
	0x95, 0xc0, 0x62, 0x68, 0x69, 0xda, 0xf2, 0x94, 0x47, 0x33, 0xaa, 0xee, 0x1e, 0x17, 0xfa, 0x03,
	0xbe, 0x80, 0x05, 0x2b, 0x3e, 0x81, 0x35, 0xbb, 0xec, 0xf9, 0x04, 0xf8, 0x8c, 0xec, 0xa9, 0xee,
	0x79, 0x68, 0x46, 0x8e, 0x61, 0x25, 0xf5, 0xb9, 0xe7, 0x9e, 0xb9, 0x7d, 0x6f, 0xdf, 0x03, 0x8f,
	0xa7, 0x3e, 0x25, 0xb6, 0x3f, 0x0d, 0xe6, 0xc4, 0xe3, 0x47, 0xe9, 0x43, 0x7b, 0x41, 0x7d, 0xee,
	0xa3, 0x6a, 0x1a, 0xdb, 0xdd, 0x99, 0xf9, 0xfe, 0xcc, 0x25, 0x47, 0x32, 0x36, 0x09, 0x2e, 0x8e,
	0xb0, 0xb7, 0x0c, 0x89, 0xbb, 0x8f, 0xd7, 0x43, 0xdc, 0x99, 0x13, 0xc6, 0xf1, 0x7c, 0x11, 0x11,
	0x9e, 0x2c, 0x28, 0x99, 0x3a, 0x8c, 0x7c, 0xb6, 0xa0, 0xbe, 0x7f, 0xc1, 0x8e, 0x56, 0x3f, 0xdc,
	0x0f, 0x0f, 0x11, 0xf1, 0xe3, 0xff, 0x22, 0x32, 0xec, 0x46, 0xa5, 0xb5, 0xfe, 0x2e, 0x41, 0xb5,
	0xe7, 0x53, 0xd2, 0x8f, 0xaa, 0x43, 0x47, 0x70, 0x2f, 0xae, 0xd4, 0x72, 0x6c, 0xe2, 0x71, 0xe7,
	0xc2, 0x21, 0x54, 0x2b, 0xef, 0x2b, 0x07, 0x55, 0x03, 0xc5, 0x21, 0x3d, 0x89, 0xa0, 0xa7, 0xa0,
	0x2e, 0x28, 0xb9, 0x76, 0xfc, 0x80, 0x59, 0xd7, 0x84, 0x32, 0xc7, 0xf7, 0x34, 0x55, 0xb2, 0xb7,
	0x62, 0xfc, 0x3c, 0x84, 0xd1, 0x13, 0xd8, 0x9a, 0x06, 0x94, 0x0a, 0xe9, 0x98, 0x99, 0x93, 0xcc,
	0x7a, 0x04, 0xc7, 0xc4, 0xa7, 0xa0, 0xc6, 0xc4, 0x05, 0x25, 0xce, 0x1c, 0xcf, 0x88, 0xb6, 0x1d,
	0x6a, 0x46, 0xf8, 0x28, 0x82, 0xd1, 0x07, 0x50, 0xf5, 0xc8, 0xcf, 0x2b, 0xc1, 0xbc, 0xa4, 0x55,
	0x04, 0x16, 0xab, 0x7d, 0x08, 0x35, 0x49, 0x49, 0xa4, 0x9a, 0x92, 0x23, 0xf3, 0x12, 0x1d, 0x1d,
	0xea, 0xcc, 0x99, 0x79, 0x98, 0x07, 0x94, 0x58, 0x36, 0xe6, 0x58, 0x2b, 0xee, 0x2b, 0x07, 0x95,
	0x67, 0x7b, 0xed, 0xcc, 0x40, 0xc7, 0x31, 0xa7, 0x8f, 0x39, 0xee, 0x16, 0x7e, 0x7f, 0xfb, 0x0e,
	0x14, 0xa3, 0xc6, 0xd2, 0x28, 0xfa, 0x16, 0x6a, 0x64, 0x3e, 0x21, 0xb6, 0x4d, 0xec, 0x50, 0xa9,
	0x26, 0x95, 0x1a, 0xed, 0x70, 0xba, 0xed, 0x78, 0xba, 0xed, 0x8e, 0xb7, 0x8c, 0x25, 0xaa, 0x71,
	0x86, 0x54, 0x68, 0x41, 0x41, 0xcc, 0x88, 0x69, 0x5b, 0xfb, 0xb9, 0x83, 0xca, 0xb3, 0x6a, 0x3b,
	0x1c, 0x5f, 0x7b, 0x8c, 0x5d, 0x6e, 0x84, 0x21, 0xf4, 0x0d, 0x14, 0xa8, 0xef, 0x12, 0xa6, 0x29,
	0x92, 0x83, 0xb2, 0x75, 0x1a, 0xbe, 0x4b, 0xba, 0xe8, 0x8f, 0xb7, 0xef, 0x60, 0xff, 0xcf, 0xb7,
	0xef, 0x60, 0x53, 0x50, 0xad, 0x2b, 0xb2, 0x34, 0xc2, 0x24, 0xf4, 0x05, 0x00, 0x25, 0xd8, 0xb6,
	0x68, 0x20, 0x24, 0xee, 0x49, 0x89, 0xe6, 0x9a, 0x04, 0xc1, 0xb6, 0x11, 0xb8, 0xc4, 0x28, 0xd3,
	0xe8, 0x1f, 0x43, 0x27, 0xa0, 0x72, 0x8a, 0x3d, 0xe6, 0x70, 0xc7, 0xf7, 0xa2, 0x64, 0x4d, 0x26,
	0x3f, 0xc8, 0x26, 0x9b, 0x09, 0x4b, 0x4a, 0x6c, 0xf1, 0xcc, 0x59, 0x54, 0x9f, 0xf7, 0x2e, 0x38,
	0xd3, 0x1a, 0x32, 0xf9, 0x6e, 0x36, 0x79, 0x78, 0x6c, 0x76, 0xef, 0x27, 0xb5, 0x57, 0x28, 0x99,
	0x39, 0x8c, 0xd3, 0xa5, 0xe5, 0xd8, 0x86, 0xcc, 0x42, 0x26, 0xd4, 0xf0, 0x74, 0x4a, 0x18, 0xb3,
	0xb8, 0x7f, 0x45, 0x3c, 0xa6, 0xdd, 0x97, 0x32, 0x3b, 0x59, 0x99, 0x8e, 0xa4, 0x98, 0x82, 0xd1,
	0x6d, 0x24, 0x72, 0xb0, 0x7a, 0xd5, 0x46, 0x15, 0xaf, 0x28, 0x0c, 0x35, 0xa1, 0x88, 0x03, 0x7e,
	0xe9, 0x53, 0x6d, 0x47, 0x3e, 0x90, 0xe8, 0x84, 0xbe, 0x82, 0x72, 0xb2, 0x87, 0xda, 0xae, 0x9c,
	0xe5, 0xee, 0x8d, 0x59, 0x9a, 0x31, 0xc3, 0x58, 0x91, 0xd1, 0x73, 0x68, 0x62, 0x6f, 0x7a, 0xe9,
	0x53, 0x8b, 0x92, 0x85, 0xcf, 0x1c, 0xee, 0xd3, 0xa5, 0x15, 0x30, 0x62, 0x6b, 0x7b, 0xf2, 0x0b,
	0x8d, 0x30, 0x6a, 0x24, 0xc1, 0xef, 0x19, 0xb1, 0x91, 0x0e, 0x80, 0x39, 0xa7, 0xce, 0x24, 0xe0,
	0x84, 0x69, 0x87, 0xf2, 0x6a, 0xf7, 0xb2, 0x57, 0x3b, 0x76, 0x88, 0x6b, 0xa7, 0x2f, 0x75, 0x45,
	0x96, 0xd6, 0x25, 0x66, 0x97, 0xc4, 0x36, 0x52, 0xc9, 0xad, 0x5f, 0x37, 0xa0, 0x92, 0x6a, 0x03,
	0x7a, 0x94, 0xbe, 0xbe, 0xa6, 0xc8, 0x22, 0x52, 0x08, 0xd2, 0xa0, 0x34, 0xa3, 0xd8, 0xe3, 0x84,
	0x46, 0x9b, 0x19, 0x1f, 0x57, 0x11, 0x12, 0xad, 0x58, 0x7c, 0x14, 0x5b, 0x2d, 0x5f, 0x57, 0x4a,
	0xb8, 0x10, 0x6e, 0xb5, 0x80, 0x53, 0x4e, 0x71, 0x8b, 0xb5, 0x6c, 0xdc, 0x6a, 0x2d, 0x0f, 0xa0,
	0x9c, 0x6c, 0x96, 0x5c, 0xc7, 0xaa, 0xb1, 0x02, 0x90, 0x0a, 0xb9, 0x2b, 0xb2, 0xd4, 0x4a, 0x12,
	0x17, 0x7f, 0x85, 0x6d, 0x24, 0x1f, 0x88, 0xfd, 0x60, 0x33, 0xb4, 0x8d, 0x18, 0x8f, 0x3c, 0xa1,
	0xf5, 0x13, 0xd4, 0x32, 0xab, 0x8c, 0x5e, 0x01, 0x24, 0xd2, 0xf1, 0x4e, 0x6d, 0xdf, 0xb2, 0xfb,
	0xdd, 0xa6, 0x68, 0xfc, 0x73, 0xd1, 0xf8, 0xea, 0xca, 0x34, 0x1c, 0xdb, 0x48, 0x49, 0xb4, 0x7e,
	0x51, 0xa0, 0x9c, 0x64, 0x08, 0x9b, 0x4a, 0x33, 0xa3, 0xd6, 0x57, 0x12, 0x4c, 0xb7, 0xd1, 0x5e,
	0x78, 0x5b, 0x42, 0x45, 0x3c, 0x6c, 0xca, 0x66, 0x08, 0xe8, 0x36, 0x7a, 0x08, 0xb0, 0x08, 0x26,
	0xae, 0x33, 0x15, 0x4b, 0x1c, 0xcd, 0xa6, 0x1c, 0x22, 0xdf, 0x91, 0x65, 0xb6, 0x53, 0xf9, 0xb5,
	0x4e, 0xb5, 0x7e, 0x84, 0xbc, 0xf0, 0x03, 0xb4, 0xb3, 0xf2, 0x81, 0xa8, 0x80, 0x92, 0x38, 0x0b,
	0x81, 0x8f, 0xa0, 0x36, 0xf5, 0x5d, 0x17, 0x4f, 0x7c, 0x8a, 0xb9, 0x4f, 0x99, 0x96, 0xdb, 0xcf,
	0x1d, 0x54, 0x8d, 0x2c, 0x88, 0x50, 0xb4, 0xb5, 0x79, 0x19, 0x94, 0xff, 0x5b, 0x43, 0xd8, 0x8c,
	0x9d, 0x02, 0x35, 0x62, 0x4f, 0xda, 0x90, 0x84, 0xc8, 0x6b, 0x3e, 0x85, 0x22, 0x9e, 0xf2, 0xd8,
	0x9c, 0xeb, 0xcf, 0x1a, 0xeb, 0x6b, 0x2a, 0x6d, 0x21, 0xe2, 0xb4, 0xfe, 0x52, 0xa0, 0x9e, 0x75,
	0x0f, 0x59, 0x77, 0xb0, 0x56, 0x77, 0x10, 0xd6, 0xfd, 0xfe, 0x2f, 0x7e, 0x0d, 0x30, 0xc7, 0x7c,
	0x7a, 0x69, 0xf1, 0xe5, 0x82, 0xc8, 0x6e, 0xd5, 0xd7, 0x0d, 0x4a, 0x6e, 0xd0, 0x99, 0x20, 0x99,
	0xcb, 0x05, 0x31, 0xca, 0xf3, 0xf8, 0xaf, 0x90, 0xbc, 0x10, 0xc1, 0xa8, 0x8f, 0xe1, 0x01, 0x7d,
	0x99, 0x5c, 0xa2, 0x20, 0xe5, 0x1e, 0xdd, 0xe6, 0x77, 0x6b, 0xd7, 0xe9, 0x40, 0x6e, 0x78, 0x6c,
	0xa2, 0xc7, 0x19, 0x1b, 0x8b, 0x37, 0x2f, 0x86, 0x74, 0x5b, 0xdc, 0x51, 0x7a, 0xd9, 0x6a, 0xf8,
	0x25, 0x79, 0xd6, 0xed, 0xd6, 0x3f, 0x0a, 0x14, 0x64, 0xb9, 0xf1, 0x93, 0x57, 0x56, 0x4f, 0x7e,
	0x0f, 0xca, 0x62, 0xf5, 0x5d, 0x3c, 0x21, 0x6e, 0xfc, 0x68, 0xae, 0xc8, 0xf2, 0x54, 0x9c, 0xd1,
	0x27, 0x90, 0x4f, 0x35, 0x60, 0xfb, 0x3d, 0x0d, 0x90, 0x77, 0x97, 0x24, 0xb4, 0x03, 0x25, 0xc6,
	0xa9, 0x75, 0x8d, 0x5d, 0x79, 0xf1, 0xf2, 0xcb, 0x3b, 0x46, 0x91, 0x71, 0x7a, 0x8e, 0x5d, 0xb4,
	0x07, 0x9b, 0x93, 0x25, 0x27, 0x32, 0x26, 0x57, 0xfb, 0xe5, 0x1d, 0xa3, 0x24, 0x10, 0x11, 0x7c,
	0x01, 0x9b, 0xc2, 0xf0, 0x64, 0xb0, 0xf8, 0x7f, 0xe6, 0x28, 0x12, 0x05, 0xfb, 0x1c, 0xbb, 0xdd,
	0x12, 0x14, 0xae, 0xb1, 0x1b, 0x90, 0xc3, 0x1e, 0x14, 0xc3, 0xa6, 0x21, 0x04, 0xf5, 0x4e, 0xcf,
	0xd4, 0x5f, 0x0d, 0x2d, 0x7d, 0x78, 0xde, 0x39, 0xd5, 0xfb, 0xea, 0x1d, 0xd4, 0x00, 0x35, 0xc2,
	0x8c, 0x41, 0xa7, 0x6f, 0x8d, 0xf5, 0x93, 0xa1, 0xaa, 0xa0, 0x2d, 0xa8, 0xa4, 0x50, 0x75, 0xe3,
	0x70, 0x06, 0xf5, 0xec, 0x48, 0xd1, 0x03, 0xd0, 0x8e, 0xf5, 0xc1, 0x69, 0xdf, 0x3a, 0xeb, 0x98,
	0xbd, 0x97, 0x96, 0xf9, 0x66, 0x34, 0x48, 0xc9, 0xee, 0xc1, 0xf6, 0x8d, 0xe8, 0xc8, 0x18, 0x1c,
	0xeb, 0xaf, 0x55, 0x05, 0xed, 0x42, 0xf3, 0x46, 0x70, 0xf0, 0xba, 0xd3, 0x33, 0xd5, 0x8d, 0xc3,
	0x33, 0x50, 0xd7, 0x87, 0x8d, 0x1e, 0xc2, 0x8e, 0x69, 0x74, 0x86, 0x63, 0x5d, 0x56, 0x74, 0xe3,
	0x0a, 0xbb, 0xd0, 0xbc, 0x19, 0x1e, 0xf4, 0x75, 0x53, 0x55, 0x0e, 0x7f, 0x53, 0xa0, 0x9c, 0x8c,
	0x02, 0x35, 0x01, 0x85, 0x1f, 0x5e, 0xab, 0x76, 0x1d, 0x37, 0x07, 0x27, 0x03, 0x43, 0x55, 0xd6,
	0xf0, 0xfe, 0xa0, 0xa7, 0x9f, 0x75, 0x4e, 0xd5, 0x0d, 0x74, 0x1f, 0xee, 0xa6, 0xf0, 0xb1, 0x69,
	0xe8, 0xc3, 0x13, 0x35, 0x27, 0x7a, 0x99, 0x82, 0xbb, 0x6f, 0xcc, 0xc1, 0x58, 0xcd, 0x23, 0x0d,
	0x1a, 0x29, 0xd4, 0xd4, 0xcf, 0x06, 0x63, 0xb3, 0x73, 0x36, 0x52, 0x0b, 0xdd, 0x17, 0xa0, 0x4e,
	0xfd, 0x79, 0xe6, 0xdd, 0x74, 0xef, 0xf6, 0x52, 0xa7, 0x91, 0x98, 0xf0, 0x48, 0xf9, 0xa1, 0x9e,
	0xa6, 0x2c, 0x26, 0x93, 0xa2, 0x1c, 0xfd, 0xe7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x87, 0x99,
	0x1b, 0x6c, 0x1b, 0x0b, 0x00, 0x00,
}
