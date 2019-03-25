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
	// # Merkle roots
	// document_root the root of all roots. It's signature_root along with a list of all signatures
	DocumentRoot []byte `protobuf:"bytes,7,opt,name=document_root,json=documentRoot,proto3" json:"document_root,omitempty"`
	// signing_root is the Merkle root calculated from all fields on the document that need
	// to be signed. This is all fields except for the signatures themselves and the `document_root`.
	SigningRoot []byte `protobuf:"bytes,10,opt,name=signing_root,json=signingRoot,proto3" json:"signing_root,omitempty"`
	// previous_root is the `document_root` of the previous version of the document
	PreviousRoot []byte `protobuf:"bytes,2,opt,name=previous_root,json=previousRoot,proto3" json:"previous_root,omitempty"`
	// data_root is the Merkle root calculated for the `embedded_data` document (such as an invoice).
	DataRoot []byte `protobuf:"bytes,5,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	// signature_data of the signature_root by centrifuge identities.
	// Signatures of the signature_root by centrifuge identities.
	SignatureData *SignatureData `protobuf:"bytes,6,opt,name=signature_data,json=signatureData,proto3" json:"signature_data,omitempty"`
	// Document a serialized document
	EmbeddedData *any.Any       `protobuf:"bytes,13,opt,name=embedded_data,json=embeddedData,proto3" json:"embedded_data,omitempty"`
	Salts        []*proto1.Salt `protobuf:"bytes,15,rep,name=salts,proto3" json:"salts,omitempty"`
	// list of roles
	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	// read rules
	ReadRules []*ReadRule `protobuf:"bytes,19,rep,name=read_rules,json=readRules,proto3" json:"read_rules,omitempty"`
	// transition rules
	TransitionRules []*TransitionRule `protobuf:"bytes,24,rep,name=transition_rules,json=transitionRules,proto3" json:"transition_rules,omitempty"`
	// nft list for uniqueness check
	Nfts []*NFT `protobuf:"bytes,20,rep,name=nfts,proto3" json:"nfts,omitempty"`
	// AccessTokens which have been added to this CoreDocument
	AccessTokens []*AccessToken `protobuf:"bytes,21,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
	// author of the latest update
	Author []byte `protobuf:"bytes,25,opt,name=author,proto3" json:"author,omitempty"`
	// timestamp of the latest update
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,26,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *CoreDocument) GetDocumentRoot() []byte {
	if m != nil {
		return m.DocumentRoot
	}
	return nil
}

func (m *CoreDocument) GetSigningRoot() []byte {
	if m != nil {
		return m.SigningRoot
	}
	return nil
}

func (m *CoreDocument) GetPreviousRoot() []byte {
	if m != nil {
		return m.PreviousRoot
	}
	return nil
}

func (m *CoreDocument) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
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
	// The document version will be used to check the timestamp on the document at the time of access token issuance.
	// This time stamp will be used to validate the signature of the identity issuing the access token.
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

func init() {
	proto.RegisterEnum("coredocument.Action", Action_name, Action_value)
	proto.RegisterEnum("coredocument.FieldMatchType", FieldMatchType_name, FieldMatchType_value)
	proto.RegisterEnum("coredocument.TransitionAction", TransitionAction_name, TransitionAction_value)
	proto.RegisterType((*CoreDocument)(nil), "coredocument.CoreDocument")
	proto.RegisterType((*AccessToken)(nil), "coredocument.AccessToken")
	proto.RegisterType((*SignatureData)(nil), "coredocument.SignatureData")
	proto.RegisterType((*Signature)(nil), "coredocument.Signature")
	proto.RegisterType((*Role)(nil), "coredocument.Role")
	proto.RegisterType((*ReadRule)(nil), "coredocument.ReadRule")
	proto.RegisterType((*TransitionRule)(nil), "coredocument.TransitionRule")
	proto.RegisterType((*NFT)(nil), "coredocument.NFT")
}

func init() { proto.RegisterFile("coredocument/coredocument.proto", fileDescriptor_eb191193ab8e2cd4) }

var fileDescriptor_eb191193ab8e2cd4 = []byte{
	// 1095 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4b, 0x73, 0xe3, 0x44,
	0x10, 0xc6, 0x76, 0x9c, 0xd8, 0xed, 0x47, 0xb4, 0xb3, 0x5e, 0xaf, 0xe2, 0x7d, 0x24, 0x18, 0x8a,
	0xcd, 0xa6, 0xc0, 0xa9, 0x0a, 0xcf, 0x2a, 0x38, 0xe0, 0x57, 0x16, 0x15, 0x1b, 0x6f, 0x4a, 0x51,
	0x6d, 0x2d, 0x70, 0x10, 0xb2, 0x34, 0xf1, 0xaa, 0x62, 0x6b, 0x5c, 0xa3, 0xd1, 0x16, 0x3e, 0x72,
	0xe3, 0x17, 0x70, 0xe6, 0xc8, 0x91, 0x33, 0xb7, 0xdc, 0xf9, 0x3b, 0xb9, 0x53, 0xf3, 0x92, 0x25,
	0x87, 0x70, 0xb2, 0xfb, 0xeb, 0xaf, 0x3f, 0xf5, 0xf4, 0xcc, 0xd7, 0xb0, 0xef, 0x13, 0x8a, 0x03,
	0xe2, 0x27, 0x0b, 0x1c, 0xb1, 0xe3, 0x6c, 0xd0, 0x5b, 0x52, 0xc2, 0x08, 0xaa, 0x67, 0xb1, 0xce,
	0xde, 0x8c, 0x90, 0xd9, 0x1c, 0x1f, 0x8b, 0xdc, 0x34, 0xb9, 0x3c, 0xf6, 0xa2, 0x95, 0x24, 0x76,
	0xf6, 0x37, 0x53, 0x2c, 0x5c, 0xe0, 0x98, 0x79, 0x8b, 0xa5, 0x22, 0x3c, 0x5b, 0x52, 0xec, 0x87,
	0x31, 0xfe, 0x64, 0x49, 0x09, 0xb9, 0x8c, 0x8f, 0xd7, 0x3f, 0x8c, 0xc8, 0x40, 0x11, 0x3f, 0xfa,
	0x3f, 0x62, 0xec, 0xcd, 0x55, 0x6b, 0xdd, 0x5f, 0x2b, 0x50, 0x1f, 0x12, 0x8a, 0x47, 0xaa, 0x3b,
	0x74, 0x0c, 0xf7, 0x75, 0xa7, 0x6e, 0x18, 0xe0, 0x88, 0x85, 0x97, 0x21, 0xa6, 0x66, 0xf5, 0xa0,
	0x70, 0x58, 0xb7, 0x91, 0x4e, 0x59, 0x69, 0x06, 0x3d, 0x07, 0x63, 0x49, 0xf1, 0xbb, 0x90, 0x24,
	0xb1, 0xfb, 0x0e, 0xd3, 0x38, 0x24, 0x91, 0x69, 0x08, 0xf6, 0xae, 0xc6, 0x5f, 0x4b, 0x18, 0x3d,
	0x83, 0x5d, 0x3f, 0xa1, 0x94, 0x4b, 0x6b, 0x66, 0x49, 0x30, 0x9b, 0x0a, 0xd6, 0xc4, 0xe7, 0x60,
	0x68, 0xe2, 0x92, 0xe2, 0x70, 0xe1, 0xcd, 0xb0, 0xf9, 0x50, 0x6a, 0x2a, 0xfc, 0x5c, 0xc1, 0xe8,
	0x7d, 0xa8, 0x47, 0xf8, 0x97, 0xb5, 0xe0, 0x96, 0xa0, 0xd5, 0x38, 0xa6, 0xd5, 0x3e, 0x80, 0x86,
	0xa0, 0xa4, 0x52, 0x6d, 0xc1, 0x11, 0x75, 0xa9, 0xce, 0x11, 0x34, 0xd2, 0x73, 0x53, 0x42, 0x98,
	0xb9, 0xc3, 0x49, 0x83, 0xf2, 0x1f, 0xd7, 0x37, 0x50, 0xb0, 0xeb, 0x3a, 0x67, 0x13, 0xc2, 0xd0,
	0x21, 0xd4, 0xe3, 0x70, 0x16, 0x85, 0xd1, 0x4c, 0x52, 0x21, 0x4b, 0xad, 0xa9, 0x94, 0x60, 0x1e,
	0x41, 0x23, 0x1d, 0x8e, 0xa0, 0x16, 0x25, 0xf5, 0x4f, 0xa9, 0xaa, 0x73, 0x82, 0xdb, 0x85, 0x6a,
	0xe0, 0x31, 0x4f, 0xf2, 0xca, 0x59, 0xc9, 0x0a, 0xc7, 0x05, 0xc7, 0x82, 0x26, 0x97, 0xf7, 0x58,
	0x42, 0xb1, 0xcb, 0x51, 0x73, 0xfb, 0xa0, 0x70, 0x58, 0x3b, 0x79, 0xd4, 0xcb, 0x3d, 0xbb, 0x0b,
	0xcd, 0x19, 0x79, 0xcc, 0xd3, 0x2a, 0x8d, 0x38, 0x8b, 0xa2, 0x6f, 0xa1, 0x81, 0x17, 0x53, 0x1c,
	0x04, 0x38, 0x90, 0x4a, 0x0d, 0xa1, 0xd4, 0xea, 0xc9, 0x37, 0xd8, 0xd3, 0x6f, 0xb0, 0xd7, 0x8f,
	0x56, 0xe9, 0x18, 0x74, 0x85, 0x50, 0xe8, 0x42, 0x99, 0xbf, 0xa4, 0xd8, 0xdc, 0x3d, 0x28, 0x1d,
	0xd6, 0x4e, 0xea, 0x3d, 0xf9, 0xc8, 0x7a, 0x17, 0xde, 0x9c, 0xd9, 0x32, 0x85, 0xbe, 0x81, 0x32,
	0x25, 0x73, 0x1c, 0x9b, 0x05, 0xc1, 0x41, 0xf9, 0x3e, 0x6d, 0x32, 0xc7, 0x03, 0xf4, 0xd7, 0xf5,
	0x0d, 0x1c, 0xfc, 0x7d, 0x7d, 0x03, 0x15, 0x4e, 0x75, 0xaf, 0xf0, 0xca, 0x96, 0x45, 0xe8, 0x73,
	0x00, 0x8a, 0xbd, 0xc0, 0xa5, 0x09, 0x97, 0xb8, 0x2f, 0x24, 0xda, 0x1b, 0x12, 0xd8, 0x0b, 0xec,
	0x64, 0x8e, 0xed, 0x2a, 0x55, 0xff, 0x62, 0xf4, 0x02, 0x0c, 0x46, 0xbd, 0x28, 0x0e, 0x59, 0x48,
	0x22, 0x55, 0x6c, 0x8a, 0xe2, 0xc7, 0xf9, 0x62, 0x27, 0x65, 0x09, 0x89, 0x5d, 0x96, 0x8b, 0x79,
	0xf7, 0x5b, 0xd1, 0x25, 0x8b, 0xcd, 0x96, 0x28, 0xbe, 0x97, 0x2f, 0x9e, 0x9c, 0x3a, 0x83, 0x07,
	0x69, 0xef, 0x35, 0x8a, 0x67, 0x61, 0xcc, 0xe8, 0xca, 0x0d, 0x03, 0x5b, 0x54, 0x21, 0x07, 0x1a,
	0x9e, 0xef, 0xe3, 0x38, 0x76, 0x19, 0xb9, 0xc2, 0x51, 0x6c, 0x3e, 0x10, 0x32, 0x7b, 0x79, 0x99,
	0xbe, 0xa0, 0x38, 0x9c, 0x31, 0x68, 0xa5, 0x72, 0xb0, 0xf6, 0x9e, 0x5d, 0xf7, 0xd6, 0x94, 0x18,
	0xb5, 0x61, 0xdb, 0x4b, 0xd8, 0x5b, 0x42, 0xcd, 0x3d, 0xf1, 0x8c, 0x55, 0x84, 0xbe, 0x82, 0x6a,
	0xba, 0x2d, 0xcc, 0x8e, 0xb8, 0xcb, 0xce, 0xad, 0xbb, 0x74, 0x34, 0xc3, 0x5e, 0x93, 0xbb, 0xbf,
	0x17, 0xa1, 0x96, 0xe9, 0x02, 0x3d, 0xcd, 0x7e, 0xdd, 0x2c, 0x88, 0xaf, 0x64, 0x10, 0x64, 0xc2,
	0xce, 0x8c, 0x7a, 0x11, 0xc3, 0x54, 0xd9, 0x57, 0x87, 0xeb, 0x0c, 0x56, 0x3e, 0xd4, 0x21, 0xb7,
	0xbe, 0xb8, 0xdc, 0x8c, 0x70, 0x59, 0x5a, 0x9f, 0xc3, 0x99, 0x75, 0x72, 0xc7, 0xfe, 0x29, 0xde,
	0xb9, 0x7f, 0x1e, 0x43, 0x35, 0x7d, 0xd8, 0xc2, 0x0d, 0x75, 0x7b, 0x0d, 0x20, 0x03, 0x4a, 0x57,
	0x78, 0x25, 0xcd, 0x6c, 0xf3, 0xbf, 0x7c, 0xb7, 0xa4, 0x1f, 0xd0, 0x4b, 0xa3, 0x22, 0x77, 0x8b,
	0xc6, 0xd5, 0xe2, 0xe8, 0xfe, 0x0c, 0x8d, 0x9c, 0x93, 0xd0, 0x2b, 0x80, 0x54, 0x5a, 0x3f, 0xe9,
	0x87, 0x77, 0x58, 0x6f, 0xd0, 0xe6, 0x97, 0xf9, 0x19, 0xbf, 0xcc, 0xfa, 0xda, 0xb3, 0x61, 0x60,
	0x67, 0x24, 0xba, 0xbf, 0x15, 0xa0, 0x9a, 0x56, 0xf0, 0x5d, 0x96, 0x65, 0xaa, 0xd1, 0xd7, 0x52,
	0xcc, 0x0a, 0xd0, 0x23, 0x79, 0x5a, 0x4c, 0x79, 0x5e, 0x0e, 0xa5, 0x22, 0x01, 0x2b, 0x40, 0x4f,
	0x00, 0x96, 0xc9, 0x74, 0x1e, 0xfa, 0xdc, 0x43, 0xea, 0x6e, 0xaa, 0x12, 0xf9, 0x1e, 0xaf, 0xf2,
	0x93, 0xda, 0xda, 0x98, 0x54, 0xf7, 0x27, 0xd8, 0xe2, 0x76, 0x44, 0x7b, 0x6b, 0x1b, 0xaa, 0x06,
	0x76, 0x78, 0xcc, 0x05, 0x3e, 0x84, 0x86, 0x4f, 0xe6, 0x73, 0x6f, 0x4a, 0xa8, 0xc7, 0x08, 0x8d,
	0xcd, 0xd2, 0x41, 0xe9, 0xb0, 0x6e, 0xe7, 0x41, 0x84, 0x94, 0x69, 0xb6, 0x44, 0x52, 0xfc, 0xef,
	0x4e, 0xa0, 0xa2, 0x8d, 0x8a, 0x5a, 0x7a, 0x25, 0x14, 0x05, 0x41, 0x59, 0xfd, 0x63, 0xd8, 0xf6,
	0x7c, 0xa6, 0x37, 0x78, 0xf3, 0xa4, 0xb5, 0xe9, 0x12, 0xe1, 0x4a, 0xc5, 0xe9, 0xfe, 0x53, 0x80,
	0x66, 0xde, 0xbc, 0xa2, 0xef, 0x64, 0xa3, 0xef, 0x44, 0xf6, 0xfd, 0xdf, 0x5f, 0xfc, 0x1a, 0x60,
	0xe1, 0x31, 0xff, 0xad, 0xcb, 0x56, 0x4b, 0x2c, 0xa6, 0xd5, 0xdc, 0xdc, 0x0f, 0xa7, 0x21, 0x9e,
	0x07, 0x67, 0x9c, 0xe4, 0xac, 0x96, 0xd8, 0xae, 0x2e, 0xf4, 0x5f, 0x2e, 0x79, 0xc9, 0x93, 0x6a,
	0x8e, 0x32, 0x40, 0x5f, 0xa4, 0x87, 0x28, 0x0b, 0xb9, 0xa7, 0x77, 0xad, 0x9b, 0x8d, 0xe3, 0xf4,
	0xa1, 0x34, 0x39, 0x75, 0xd0, 0x7e, 0x6e, 0x8b, 0x68, 0xe7, 0x69, 0xc8, 0x0a, 0xf8, 0x19, 0xc5,
	0x2a, 0x59, 0x5f, 0xfe, 0x8e, 0x88, 0xad, 0xe0, 0x68, 0x08, 0xdb, 0x52, 0x14, 0x21, 0x68, 0xf6,
	0x87, 0x8e, 0xf5, 0x6a, 0xe2, 0x5a, 0x93, 0xd7, 0xfd, 0x97, 0xd6, 0xc8, 0x78, 0x0f, 0xb5, 0xc0,
	0x50, 0x98, 0x3d, 0xee, 0x8f, 0xdc, 0x0b, 0xeb, 0xc5, 0xc4, 0x28, 0xa0, 0x5d, 0xa8, 0x65, 0x50,
	0xa3, 0x78, 0x34, 0x83, 0x66, 0xfe, 0xc8, 0xe8, 0x31, 0x98, 0xa7, 0xd6, 0xf8, 0xe5, 0xc8, 0x3d,
	0xeb, 0x3b, 0xc3, 0xef, 0x5c, 0xe7, 0x87, 0xf3, 0x71, 0x46, 0xf6, 0x11, 0x3c, 0xbc, 0x95, 0x3d,
	0xb7, 0xc7, 0xa7, 0xd6, 0x1b, 0xa3, 0x80, 0x3a, 0xd0, 0xbe, 0x95, 0x1c, 0xbf, 0xe9, 0x0f, 0x1d,
	0xa3, 0x78, 0x74, 0x06, 0xc6, 0xe6, 0x30, 0xd0, 0x13, 0xd8, 0x73, 0xec, 0xfe, 0xe4, 0xc2, 0x12,
	0x1d, 0xdd, 0x3a, 0x42, 0x07, 0xda, 0xb7, 0xd3, 0xe3, 0x91, 0xe5, 0x18, 0x85, 0xc1, 0x97, 0x60,
	0xf8, 0x64, 0x91, 0x1b, 0xf6, 0xe0, 0xde, 0x30, 0x13, 0x9d, 0xf3, 0x05, 0x78, 0x5e, 0xf8, 0xb1,
	0x99, 0xa5, 0x2c, 0xa7, 0xd3, 0x6d, 0xb1, 0x19, 0x3f, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x5d, 0xab, 0x6f, 0xc3, 0x09, 0x00, 0x00,
}
