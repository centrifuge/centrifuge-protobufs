// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity/entity.proto

package entitypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/centrifuge/precise-proofs/proofs/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EntityRelationship allows other identities to access the Entity document.
type EntityRelationship struct {
	// owner id of the Relationship
	OwnerIdentity []byte `protobuf:"bytes,1,opt,name=owner_identity,json=ownerIdentity,proto3" json:"owner_identity,omitempty"`
	// label is a reference for the creator
	Label []byte `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// identity which should be gruanteed acess
	TargetIdentity       []byte   `protobuf:"bytes,3,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityRelationship) Reset()         { *m = EntityRelationship{} }
func (m *EntityRelationship) String() string { return proto.CompactTextString(m) }
func (*EntityRelationship) ProtoMessage()    {}
func (*EntityRelationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{0}
}
func (m *EntityRelationship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityRelationship.Unmarshal(m, b)
}
func (m *EntityRelationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityRelationship.Marshal(b, m, deterministic)
}
func (dst *EntityRelationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityRelationship.Merge(dst, src)
}
func (m *EntityRelationship) XXX_Size() int {
	return xxx_messageInfo_EntityRelationship.Size(m)
}
func (m *EntityRelationship) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityRelationship.DiscardUnknown(m)
}

var xxx_messageInfo_EntityRelationship proto.InternalMessageInfo

func (m *EntityRelationship) GetOwnerIdentity() []byte {
	if m != nil {
		return m.OwnerIdentity
	}
	return nil
}

func (m *EntityRelationship) GetLabel() []byte {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *EntityRelationship) GetTargetIdentity() []byte {
	if m != nil {
		return m.TargetIdentity
	}
	return nil
}

// EntityData is the default entity schema
type Entity struct {
	Identity  []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	LegalName string `protobuf:"bytes,2,opt,name=legal_name,json=legalName,proto3" json:"legal_name,omitempty"`
	// address
	Addresses []*Address `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// tax information
	PaymentDetails []*PaymentDetail `protobuf:"bytes,4,rep,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	// Entity contact list
	Contacts             []*Contact `protobuf:"bytes,5,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{1}
}
func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (dst *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(dst, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Entity) GetLegalName() string {
	if m != nil {
		return m.LegalName
	}
	return ""
}

func (m *Entity) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Entity) GetPaymentDetails() []*PaymentDetail {
	if m != nil {
		return m.PaymentDetails
	}
	return nil
}

func (m *Entity) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type EntityId struct {
	EntityId             string   `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	EntityName           string   `protobuf:"bytes,2,opt,name=entity_name,json=entityName,proto3" json:"entity_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityId) Reset()         { *m = EntityId{} }
func (m *EntityId) String() string { return proto.CompactTextString(m) }
func (*EntityId) ProtoMessage()    {}
func (*EntityId) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{2}
}
func (m *EntityId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityId.Unmarshal(m, b)
}
func (m *EntityId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityId.Marshal(b, m, deterministic)
}
func (dst *EntityId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityId.Merge(dst, src)
}
func (m *EntityId) XXX_Size() int {
	return xxx_messageInfo_EntityId.Size(m)
}
func (m *EntityId) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityId.DiscardUnknown(m)
}

var xxx_messageInfo_EntityId proto.InternalMessageInfo

func (m *EntityId) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *EntityId) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

type Address struct {
	IsMain               bool     `protobuf:"varint,1,opt,name=is_main,json=isMain,proto3" json:"is_main,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Zip                  string   `protobuf:"bytes,3,opt,name=zip,proto3" json:"zip,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	AddressLine1         string   `protobuf:"bytes,5,opt,name=address_line1,json=addressLine1,proto3" json:"address_line1,omitempty"`
	AddressLine2         string   `protobuf:"bytes,6,opt,name=address_line2,json=addressLine2,proto3" json:"address_line2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{3}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetIsMain() bool {
	if m != nil {
		return m.IsMain
	}
	return false
}

func (m *Address) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetAddressLine1() string {
	if m != nil {
		return m.AddressLine1
	}
	return ""
}

func (m *Address) GetAddressLine2() string {
	if m != nil {
		return m.AddressLine2
	}
	return ""
}

type BankPaymentMethod struct {
	Identifier           []byte   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Address              *Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	HolderName           string   `protobuf:"bytes,3,opt,name=holder_name,json=holderName,proto3" json:"holder_name,omitempty"`
	BankKey              string   `protobuf:"bytes,4,opt,name=bank_key,json=bankKey,proto3" json:"bank_key,omitempty"`
	BankAccountNumber    string   `protobuf:"bytes,5,opt,name=bank_account_number,json=bankAccountNumber,proto3" json:"bank_account_number,omitempty"`
	Currency             string   `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankPaymentMethod) Reset()         { *m = BankPaymentMethod{} }
func (m *BankPaymentMethod) String() string { return proto.CompactTextString(m) }
func (*BankPaymentMethod) ProtoMessage()    {}
func (*BankPaymentMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{4}
}
func (m *BankPaymentMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankPaymentMethod.Unmarshal(m, b)
}
func (m *BankPaymentMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankPaymentMethod.Marshal(b, m, deterministic)
}
func (dst *BankPaymentMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankPaymentMethod.Merge(dst, src)
}
func (m *BankPaymentMethod) XXX_Size() int {
	return xxx_messageInfo_BankPaymentMethod.Size(m)
}
func (m *BankPaymentMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_BankPaymentMethod.DiscardUnknown(m)
}

var xxx_messageInfo_BankPaymentMethod proto.InternalMessageInfo

func (m *BankPaymentMethod) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *BankPaymentMethod) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *BankPaymentMethod) GetHolderName() string {
	if m != nil {
		return m.HolderName
	}
	return ""
}

func (m *BankPaymentMethod) GetBankKey() string {
	if m != nil {
		return m.BankKey
	}
	return ""
}

func (m *BankPaymentMethod) GetBankAccountNumber() string {
	if m != nil {
		return m.BankAccountNumber
	}
	return ""
}

func (m *BankPaymentMethod) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type CryptoPaymentMethod struct {
	Identifier           []byte   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	AccountAddress       string   `protobuf:"bytes,2,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoPaymentMethod) Reset()         { *m = CryptoPaymentMethod{} }
func (m *CryptoPaymentMethod) String() string { return proto.CompactTextString(m) }
func (*CryptoPaymentMethod) ProtoMessage()    {}
func (*CryptoPaymentMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{5}
}
func (m *CryptoPaymentMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoPaymentMethod.Unmarshal(m, b)
}
func (m *CryptoPaymentMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoPaymentMethod.Marshal(b, m, deterministic)
}
func (dst *CryptoPaymentMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoPaymentMethod.Merge(dst, src)
}
func (m *CryptoPaymentMethod) XXX_Size() int {
	return xxx_messageInfo_CryptoPaymentMethod.Size(m)
}
func (m *CryptoPaymentMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoPaymentMethod.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoPaymentMethod proto.InternalMessageInfo

func (m *CryptoPaymentMethod) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *CryptoPaymentMethod) GetAccountAddress() string {
	if m != nil {
		return m.AccountAddress
	}
	return ""
}

type PayPalMethod struct {
	Identifier           []byte   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayPalMethod) Reset()         { *m = PayPalMethod{} }
func (m *PayPalMethod) String() string { return proto.CompactTextString(m) }
func (*PayPalMethod) ProtoMessage()    {}
func (*PayPalMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{6}
}
func (m *PayPalMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayPalMethod.Unmarshal(m, b)
}
func (m *PayPalMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayPalMethod.Marshal(b, m, deterministic)
}
func (dst *PayPalMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayPalMethod.Merge(dst, src)
}
func (m *PayPalMethod) XXX_Size() int {
	return xxx_messageInfo_PayPalMethod.Size(m)
}
func (m *PayPalMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_PayPalMethod.DiscardUnknown(m)
}

var xxx_messageInfo_PayPalMethod proto.InternalMessageInfo

func (m *PayPalMethod) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *PayPalMethod) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type PaymentDetail struct {
	// fields for bank accounts and ethereum wallets
	Predefined bool `protobuf:"varint,1,opt,name=predefined,proto3" json:"predefined,omitempty"`
	// Types that are valid to be assigned to PaymentMethod:
	//	*PaymentDetail_BankPaymentMethod
	//	*PaymentDetail_CryptoPaymentMethod
	//	*PaymentDetail_PayPalMethod
	PaymentMethod        isPaymentDetail_PaymentMethod `protobuf_oneof:"payment_method"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *PaymentDetail) Reset()         { *m = PaymentDetail{} }
func (m *PaymentDetail) String() string { return proto.CompactTextString(m) }
func (*PaymentDetail) ProtoMessage()    {}
func (*PaymentDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{7}
}
func (m *PaymentDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentDetail.Unmarshal(m, b)
}
func (m *PaymentDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentDetail.Marshal(b, m, deterministic)
}
func (dst *PaymentDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentDetail.Merge(dst, src)
}
func (m *PaymentDetail) XXX_Size() int {
	return xxx_messageInfo_PaymentDetail.Size(m)
}
func (m *PaymentDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentDetail.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentDetail proto.InternalMessageInfo

func (m *PaymentDetail) GetPredefined() bool {
	if m != nil {
		return m.Predefined
	}
	return false
}

type isPaymentDetail_PaymentMethod interface {
	isPaymentDetail_PaymentMethod()
}

type PaymentDetail_BankPaymentMethod struct {
	BankPaymentMethod *BankPaymentMethod `protobuf:"bytes,2,opt,name=bank_payment_method,json=bankPaymentMethod,proto3,oneof"`
}

type PaymentDetail_CryptoPaymentMethod struct {
	CryptoPaymentMethod *CryptoPaymentMethod `protobuf:"bytes,3,opt,name=crypto_payment_method,json=cryptoPaymentMethod,proto3,oneof"`
}

type PaymentDetail_PayPalMethod struct {
	PayPalMethod *PayPalMethod `protobuf:"bytes,4,opt,name=pay_pal_method,json=payPalMethod,proto3,oneof"`
}

func (*PaymentDetail_BankPaymentMethod) isPaymentDetail_PaymentMethod() {}

func (*PaymentDetail_CryptoPaymentMethod) isPaymentDetail_PaymentMethod() {}

func (*PaymentDetail_PayPalMethod) isPaymentDetail_PaymentMethod() {}

func (m *PaymentDetail) GetPaymentMethod() isPaymentDetail_PaymentMethod {
	if m != nil {
		return m.PaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetBankPaymentMethod() *BankPaymentMethod {
	if x, ok := m.GetPaymentMethod().(*PaymentDetail_BankPaymentMethod); ok {
		return x.BankPaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetCryptoPaymentMethod() *CryptoPaymentMethod {
	if x, ok := m.GetPaymentMethod().(*PaymentDetail_CryptoPaymentMethod); ok {
		return x.CryptoPaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetPayPalMethod() *PayPalMethod {
	if x, ok := m.GetPaymentMethod().(*PaymentDetail_PayPalMethod); ok {
		return x.PayPalMethod
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PaymentDetail) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PaymentDetail_OneofMarshaler, _PaymentDetail_OneofUnmarshaler, _PaymentDetail_OneofSizer, []interface{}{
		(*PaymentDetail_BankPaymentMethod)(nil),
		(*PaymentDetail_CryptoPaymentMethod)(nil),
		(*PaymentDetail_PayPalMethod)(nil),
	}
}

func _PaymentDetail_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PaymentDetail)
	// payment_method
	switch x := m.PaymentMethod.(type) {
	case *PaymentDetail_BankPaymentMethod:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BankPaymentMethod); err != nil {
			return err
		}
	case *PaymentDetail_CryptoPaymentMethod:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CryptoPaymentMethod); err != nil {
			return err
		}
	case *PaymentDetail_PayPalMethod:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PayPalMethod); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PaymentDetail.PaymentMethod has unexpected type %T", x)
	}
	return nil
}

func _PaymentDetail_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PaymentDetail)
	switch tag {
	case 2: // payment_method.bank_payment_method
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BankPaymentMethod)
		err := b.DecodeMessage(msg)
		m.PaymentMethod = &PaymentDetail_BankPaymentMethod{msg}
		return true, err
	case 3: // payment_method.crypto_payment_method
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoPaymentMethod)
		err := b.DecodeMessage(msg)
		m.PaymentMethod = &PaymentDetail_CryptoPaymentMethod{msg}
		return true, err
	case 4: // payment_method.pay_pal_method
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PayPalMethod)
		err := b.DecodeMessage(msg)
		m.PaymentMethod = &PaymentDetail_PayPalMethod{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PaymentDetail_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PaymentDetail)
	// payment_method
	switch x := m.PaymentMethod.(type) {
	case *PaymentDetail_BankPaymentMethod:
		s := proto.Size(x.BankPaymentMethod)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PaymentDetail_CryptoPaymentMethod:
		s := proto.Size(x.CryptoPaymentMethod)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PaymentDetail_PayPalMethod:
		s := proto.Size(x.PayPalMethod)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Contact struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_1ae0536e488d4d08, []int{8}
}
func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (dst *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(dst, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Contact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Contact) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*EntityRelationship)(nil), "entity.EntityRelationship")
	proto.RegisterType((*Entity)(nil), "entity.Entity")
	proto.RegisterType((*EntityId)(nil), "entity.EntityId")
	proto.RegisterType((*Address)(nil), "entity.Address")
	proto.RegisterType((*BankPaymentMethod)(nil), "entity.BankPaymentMethod")
	proto.RegisterType((*CryptoPaymentMethod)(nil), "entity.CryptoPaymentMethod")
	proto.RegisterType((*PayPalMethod)(nil), "entity.PayPalMethod")
	proto.RegisterType((*PaymentDetail)(nil), "entity.PaymentDetail")
	proto.RegisterType((*Contact)(nil), "entity.Contact")
}

func init() { proto.RegisterFile("entity/entity.proto", fileDescriptor_entity_1ae0536e488d4d08) }

var fileDescriptor_entity_1ae0536e488d4d08 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0x26, 0xe4, 0xd7, 0x43, 0x48, 0x60, 0x03, 0xaa, 0x01, 0xb5, 0x45, 0xae, 0x10, 0x54, 0x15,
	0xa0, 0xd2, 0x6b, 0x55, 0x89, 0x40, 0xa5, 0x20, 0x0a, 0x4a, 0x7d, 0xec, 0xa1, 0xd6, 0xc6, 0x1e,
	0xc8, 0x0a, 0x7b, 0xd7, 0x5a, 0x2f, 0xaa, 0xc2, 0xcb, 0xf4, 0x0d, 0xfa, 0x4a, 0xbd, 0xf5, 0x39,
	0x2a, 0xef, 0xae, 0x1d, 0x07, 0x38, 0xb4, 0xa7, 0x64, 0xbe, 0x6f, 0x77, 0xc6, 0xdf, 0x37, 0x33,
	0x0b, 0x03, 0xe4, 0x8a, 0xa9, 0xd9, 0xb1, 0xf9, 0x39, 0x4a, 0xa5, 0x50, 0x82, 0xb4, 0x4c, 0xb4,
	0xbd, 0x9f, 0x4a, 0x0c, 0x59, 0x86, 0x87, 0xa9, 0x14, 0xe2, 0x26, 0x3b, 0x9e, 0xff, 0x28, 0x61,
	0x02, 0x73, 0xc1, 0x7b, 0x00, 0xf2, 0x59, 0x5f, 0xf1, 0x31, 0xa6, 0x8a, 0x09, 0x9e, 0x4d, 0x59,
	0x4a, 0xf6, 0xa0, 0x27, 0x7e, 0x70, 0x94, 0x01, 0x8b, 0x4c, 0x42, 0xb7, 0xb6, 0x5b, 0x3b, 0xe8,
	0xfa, 0xab, 0x1a, 0xbd, 0xb0, 0x20, 0xd9, 0x80, 0x66, 0x4c, 0x27, 0x18, 0xbb, 0xcb, 0x9a, 0x35,
	0x01, 0xd9, 0x87, 0xbe, 0xa2, 0xf2, 0x16, 0xd5, 0xfc, 0x76, 0x5d, 0xf3, 0x3d, 0x03, 0x17, 0xd7,
	0xbd, 0xdf, 0x35, 0x68, 0x99, 0xe2, 0x64, 0x1b, 0x3a, 0x8f, 0x4a, 0x95, 0x31, 0x79, 0x09, 0x10,
	0xe3, 0x2d, 0x8d, 0x03, 0x4e, 0x13, 0xd4, 0xa5, 0x1c, 0xdf, 0xd1, 0xc8, 0x35, 0x4d, 0x90, 0x1c,
	0x82, 0x43, 0xa3, 0x48, 0x62, 0x96, 0x61, 0xe6, 0xd6, 0x77, 0xeb, 0x07, 0x2b, 0x27, 0xfd, 0x23,
	0x6b, 0xca, 0xa9, 0x21, 0xfc, 0xf9, 0x09, 0xf2, 0x09, 0xfa, 0x29, 0x9d, 0x25, 0xc8, 0x55, 0x10,
	0xa1, 0xa2, 0x2c, 0xce, 0xdc, 0x86, 0xbe, 0xb4, 0x59, 0x5c, 0x1a, 0x1b, 0xfa, 0x5c, 0xb3, 0x7e,
	0x2f, 0xad, 0x86, 0x19, 0x79, 0x07, 0x9d, 0x50, 0x70, 0x45, 0x43, 0x95, 0xb9, 0xcd, 0xc5, 0x6a,
	0x67, 0x06, 0xf7, 0xcb, 0x03, 0xde, 0x08, 0x3a, 0x46, 0xe0, 0x45, 0x44, 0x76, 0xc0, 0x31, 0xe7,
	0x02, 0x16, 0x69, 0x8d, 0x8e, 0xdf, 0xc1, 0x82, 0x7c, 0x0d, 0x2b, 0x96, 0xac, 0x88, 0x04, 0x03,
	0xe5, 0x2a, 0xbd, 0x5f, 0x35, 0x68, 0x5b, 0x35, 0xe4, 0x05, 0xb4, 0x59, 0x16, 0x24, 0x94, 0x71,
	0x9d, 0xa7, 0xe3, 0xb7, 0x58, 0x76, 0x45, 0x19, 0x5f, 0xec, 0x87, 0x53, 0xf4, 0x63, 0x0d, 0xea,
	0x0f, 0x2c, 0xd5, 0x3d, 0x70, 0xfc, 0xfc, 0x2f, 0x71, 0xa1, 0x1d, 0x8a, 0x7b, 0xae, 0xe4, 0xcc,
	0x6d, 0x68, 0xb4, 0x08, 0xc9, 0x1b, 0x58, 0xb5, 0x56, 0x05, 0x31, 0xe3, 0xf8, 0xde, 0x6d, 0x6a,
	0xbe, 0x6b, 0xc1, 0x2f, 0x39, 0xf6, 0xf8, 0xd0, 0x89, 0xdb, 0x7a, 0x72, 0xe8, 0xc4, 0xfb, 0x53,
	0x83, 0xf5, 0x21, 0xe5, 0x77, 0xd6, 0xcd, 0x2b, 0x54, 0x53, 0x11, 0x91, 0x57, 0x00, 0xa6, 0xaf,
	0x37, 0x0c, 0xa5, 0xed, 0x74, 0x05, 0x21, 0x6f, 0xa1, 0x6d, 0xb3, 0x68, 0x0d, 0xcf, 0xb4, 0xb2,
	0xe0, 0x73, 0xcb, 0xa6, 0x22, 0x8e, 0x50, 0x1a, 0xcb, 0x8c, 0x3c, 0x30, 0x90, 0x1e, 0x8c, 0x2d,
	0xe8, 0x4c, 0x28, 0xbf, 0x0b, 0xee, 0xb0, 0x94, 0x99, 0xc7, 0x97, 0x38, 0x23, 0x47, 0x30, 0xd0,
	0x14, 0x0d, 0xb5, 0xf0, 0x80, 0xdf, 0x27, 0x13, 0x94, 0x56, 0xec, 0x7a, 0x4e, 0x9d, 0x1a, 0xe6,
	0x5a, 0x13, 0xf9, 0x78, 0x86, 0xf7, 0x52, 0x22, 0x0f, 0x67, 0x56, 0x6c, 0x19, 0x7b, 0xdf, 0x61,
	0x70, 0x26, 0x67, 0xa9, 0x12, 0xff, 0xa7, 0x74, 0x1f, 0xfa, 0x45, 0xf5, 0xaa, 0x62, 0xc7, 0xef,
	0x59, 0xd8, 0x0a, 0xf6, 0xce, 0xa1, 0x3b, 0xa6, 0xb3, 0x31, 0x8d, 0xff, 0x31, 0xf1, 0x06, 0x34,
	0x31, 0xa1, 0xac, 0x1c, 0x02, 0x1d, 0x78, 0x3f, 0x97, 0x61, 0x75, 0x61, 0xb0, 0xf3, 0x3c, 0xa9,
	0xc4, 0x08, 0x6f, 0x18, 0xc7, 0xc8, 0x0e, 0x52, 0x05, 0x21, 0x97, 0xd6, 0xa3, 0x62, 0x5b, 0x12,
	0x5d, 0xde, 0xb6, 0x65, 0xab, 0x68, 0xcb, 0x93, 0x16, 0x8f, 0x96, 0x8c, 0x81, 0x8b, 0x6e, 0x7c,
	0x85, 0xcd, 0x50, 0x9b, 0xf4, 0x38, 0x5d, 0x5d, 0xa7, 0xdb, 0x29, 0x57, 0xe8, 0xa9, 0x93, 0xa3,
	0x25, 0x7f, 0x10, 0x3e, 0x63, 0xf0, 0x47, 0xc8, 0x57, 0x33, 0x48, 0x69, 0x5c, 0xe4, 0x6a, 0xe8,
	0x5c, 0x1b, 0x95, 0x3d, 0x2e, 0x5d, 0x1b, 0x2d, 0xf9, 0xdd, 0xb4, 0x12, 0x0f, 0xd7, 0xa0, 0xb7,
	0xf8, 0x25, 0x5e, 0x00, 0x6d, 0xbb, 0xc0, 0x84, 0x40, 0x43, 0xcf, 0x94, 0xd9, 0x52, 0xfd, 0x3f,
	0xb7, 0x55, 0x31, 0x15, 0x17, 0xbb, 0x69, 0x82, 0xb9, 0xd9, 0xf5, 0x8a, 0xd9, 0x39, 0x9a, 0x4e,
	0x05, 0x47, 0x3b, 0x76, 0x26, 0x18, 0xee, 0x01, 0x84, 0x22, 0xb1, 0x5f, 0x37, 0x5c, 0x31, 0x0f,
	0xc3, 0x38, 0x7f, 0x85, 0xc7, 0xb5, 0x6f, 0xf6, 0x29, 0x48, 0x27, 0x93, 0x96, 0x7e, 0x98, 0x3f,
	0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x76, 0xa9, 0xc9, 0xe5, 0xe0, 0x05, 0x00, 0x00,
}
