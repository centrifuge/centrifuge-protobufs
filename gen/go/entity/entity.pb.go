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

type EntityRelationship struct {
	// owner of the Relationship
	Entity []byte `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// label is a reference for the creator
	Label []byte `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// target identity
	Identity             []byte   `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityRelationship) Reset()         { *m = EntityRelationship{} }
func (m *EntityRelationship) String() string { return proto.CompactTextString(m) }
func (*EntityRelationship) ProtoMessage()    {}
func (*EntityRelationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_86c9f2485a331dd9, []int{0}
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

func (m *EntityRelationship) GetEntity() []byte {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityRelationship) GetLabel() []byte {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *EntityRelationship) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

// EntityData is the default entity schema
type Entity struct {
	Identity    []byte      `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	LegalEntity string      `protobuf:"bytes,2,opt,name=legal_entity,json=legalEntity,proto3" json:"legal_entity,omitempty"`
	EntityIds   []*EntityId `protobuf:"bytes,3,rep,name=entity_ids,json=entityIds,proto3" json:"entity_ids,omitempty"`
	// address
	Addresses []*Address `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// tax information
	PaymentDetails []*PaymentDetail `protobuf:"bytes,5,rep,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	// Entity contact list
	Contacts             []*Contact `protobuf:"bytes,6,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_86c9f2485a331dd9, []int{1}
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

func (m *Entity) GetLegalEntity() string {
	if m != nil {
		return m.LegalEntity
	}
	return ""
}

func (m *Entity) GetEntityIds() []*EntityId {
	if m != nil {
		return m.EntityIds
	}
	return nil
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
	return fileDescriptor_entity_86c9f2485a331dd9, []int{2}
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
	return fileDescriptor_entity_86c9f2485a331dd9, []int{3}
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
	return fileDescriptor_entity_86c9f2485a331dd9, []int{4}
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
	return fileDescriptor_entity_86c9f2485a331dd9, []int{5}
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
	return fileDescriptor_entity_86c9f2485a331dd9, []int{6}
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
	// Types that are valid to be assigned to PaymentMethods:
	//	*PaymentDetail_BankPaymentMethod
	//	*PaymentDetail_CryptoPaymentMethod
	//	*PaymentDetail_PayPalMethod
	PaymentMethods       isPaymentDetail_PaymentMethods `protobuf_oneof:"payment_methods"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *PaymentDetail) Reset()         { *m = PaymentDetail{} }
func (m *PaymentDetail) String() string { return proto.CompactTextString(m) }
func (*PaymentDetail) ProtoMessage()    {}
func (*PaymentDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_86c9f2485a331dd9, []int{7}
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

type isPaymentDetail_PaymentMethods interface {
	isPaymentDetail_PaymentMethods()
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

func (*PaymentDetail_BankPaymentMethod) isPaymentDetail_PaymentMethods() {}

func (*PaymentDetail_CryptoPaymentMethod) isPaymentDetail_PaymentMethods() {}

func (*PaymentDetail_PayPalMethod) isPaymentDetail_PaymentMethods() {}

func (m *PaymentDetail) GetPaymentMethods() isPaymentDetail_PaymentMethods {
	if m != nil {
		return m.PaymentMethods
	}
	return nil
}

func (m *PaymentDetail) GetBankPaymentMethod() *BankPaymentMethod {
	if x, ok := m.GetPaymentMethods().(*PaymentDetail_BankPaymentMethod); ok {
		return x.BankPaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetCryptoPaymentMethod() *CryptoPaymentMethod {
	if x, ok := m.GetPaymentMethods().(*PaymentDetail_CryptoPaymentMethod); ok {
		return x.CryptoPaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetPayPalMethod() *PayPalMethod {
	if x, ok := m.GetPaymentMethods().(*PaymentDetail_PayPalMethod); ok {
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
	// payment_methods
	switch x := m.PaymentMethods.(type) {
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
		return fmt.Errorf("PaymentDetail.PaymentMethods has unexpected type %T", x)
	}
	return nil
}

func _PaymentDetail_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PaymentDetail)
	switch tag {
	case 2: // payment_methods.bank_payment_method
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BankPaymentMethod)
		err := b.DecodeMessage(msg)
		m.PaymentMethods = &PaymentDetail_BankPaymentMethod{msg}
		return true, err
	case 3: // payment_methods.crypto_payment_method
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CryptoPaymentMethod)
		err := b.DecodeMessage(msg)
		m.PaymentMethods = &PaymentDetail_CryptoPaymentMethod{msg}
		return true, err
	case 4: // payment_methods.pay_pal_method
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PayPalMethod)
		err := b.DecodeMessage(msg)
		m.PaymentMethods = &PaymentDetail_PayPalMethod{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PaymentDetail_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PaymentDetail)
	// payment_methods
	switch x := m.PaymentMethods.(type) {
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
	return fileDescriptor_entity_86c9f2485a331dd9, []int{8}
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

func init() { proto.RegisterFile("entity/entity.proto", fileDescriptor_entity_86c9f2485a331dd9) }

var fileDescriptor_entity_86c9f2485a331dd9 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x4f, 0xdb, 0x40,
	0x10, 0x25, 0xe4, 0xd3, 0x93, 0xf0, 0xb5, 0x81, 0xd6, 0x80, 0xd4, 0x52, 0x57, 0x15, 0x54, 0x15,
	0xa0, 0xd2, 0x6b, 0x55, 0x89, 0x40, 0xa5, 0x20, 0x0a, 0x4a, 0x7d, 0xec, 0x01, 0x6b, 0x63, 0x2f,
	0xcd, 0x0a, 0x7b, 0x6d, 0x79, 0x97, 0x83, 0xfb, 0x6b, 0x7a, 0xed, 0xa5, 0xea, 0xb9, 0x37, 0xfe,
	0x4c, 0xff, 0x01, 0xf7, 0xca, 0xbb, 0x63, 0xc7, 0x21, 0x1c, 0xda, 0x53, 0xfc, 0xe6, 0xcd, 0xce,
	0xe6, 0xbd, 0x99, 0x59, 0xe8, 0x33, 0xa1, 0xb8, 0xca, 0x0e, 0xcd, 0xcf, 0x41, 0x92, 0xc6, 0x2a,
	0x26, 0x2d, 0x83, 0xb6, 0x76, 0x93, 0x94, 0xf9, 0x5c, 0xb2, 0xfd, 0x24, 0x8d, 0xe3, 0x6b, 0x79,
	0x38, 0xfd, 0x51, 0xb1, 0x01, 0xe6, 0x80, 0x73, 0x05, 0xe4, 0xa3, 0x3e, 0xe2, 0xb2, 0x90, 0x2a,
	0x1e, 0x0b, 0x39, 0xe1, 0x09, 0x79, 0x02, 0x58, 0xc8, 0xae, 0xed, 0xd4, 0xf6, 0x7a, 0x2e, 0x22,
	0xb2, 0x0e, 0xcd, 0x90, 0x8e, 0x59, 0x68, 0x2f, 0xea, 0xb0, 0x01, 0x64, 0x0b, 0x3a, 0x3c, 0xc0,
	0xfc, 0xba, 0x26, 0x4a, 0xec, 0xfc, 0x58, 0x84, 0x96, 0xb9, 0x60, 0x26, 0xad, 0x36, 0x9b, 0x46,
	0x5e, 0x40, 0x2f, 0x64, 0x5f, 0x69, 0xe8, 0x21, 0x9f, 0xd7, 0xb7, 0xdc, 0xae, 0x8e, 0xe1, 0xf1,
	0x53, 0x00, 0x43, 0x7a, 0x3c, 0x90, 0x76, 0x7d, 0xa7, 0xbe, 0xd7, 0x3d, 0x5a, 0x3d, 0x40, 0xf5,
	0x26, 0xe7, 0x2c, 0x18, 0xf4, 0x7f, 0xdd, 0xdd, 0xc3, 0xce, 0xef, 0xbb, 0x7b, 0xb0, 0xca, 0x6c,
	0x17, 0x3f, 0xcf, 0x02, 0x49, 0xf6, 0xc1, 0xa2, 0x41, 0x90, 0x32, 0x29, 0x99, 0xb4, 0x1b, 0xba,
	0xc8, 0x4a, 0x51, 0xe4, 0xd8, 0x10, 0xee, 0x34, 0x83, 0x7c, 0x80, 0x95, 0x84, 0x66, 0x11, 0x13,
	0xca, 0x0b, 0x98, 0xa2, 0x3c, 0x94, 0x76, 0x53, 0x1f, 0xda, 0x28, 0x0e, 0x8d, 0x0c, 0x7d, 0xaa,
	0x59, 0x77, 0x39, 0xa9, 0x42, 0x49, 0xde, 0x40, 0xc7, 0x8f, 0x85, 0xa2, 0xbe, 0x92, 0x76, 0x6b,
	0xf6, 0xb6, 0x13, 0x13, 0x77, 0xcb, 0x04, 0x67, 0x08, 0x9d, 0x42, 0x07, 0xd9, 0xae, 0xfc, 0x7f,
	0xed, 0x96, 0xe5, 0x76, 0x0a, 0x15, 0xe4, 0x39, 0x74, 0x91, 0x14, 0x34, 0x62, 0x68, 0x16, 0xba,
	0x73, 0x49, 0x23, 0xe6, 0xfc, 0xac, 0x41, 0x1b, 0xd5, 0x90, 0xa7, 0xd0, 0xe6, 0xd2, 0x8b, 0x28,
	0x17, 0xba, 0x4e, 0xc7, 0x6d, 0x71, 0x79, 0x41, 0xb9, 0x98, 0x6d, 0xa6, 0x55, 0x34, 0x73, 0x15,
	0xea, 0xdf, 0x78, 0xa2, 0xfb, 0x68, 0xb9, 0xf9, 0x27, 0xb1, 0xa1, 0xed, 0xc7, 0xb7, 0x42, 0xa5,
	0x99, 0xdd, 0xd0, 0xd1, 0x02, 0x92, 0x97, 0xb0, 0x84, 0x56, 0x79, 0x21, 0x17, 0xec, 0xad, 0xdd,
	0xd4, 0x7c, 0x0f, 0x83, 0x9f, 0xf2, 0xd8, 0xc3, 0xa4, 0x23, 0xbb, 0x35, 0x97, 0x74, 0xe4, 0xfc,
	0xa9, 0xc1, 0xda, 0x80, 0x8a, 0x1b, 0x74, 0xf3, 0x82, 0xa9, 0x49, 0x1c, 0x90, 0x67, 0x00, 0x66,
	0x42, 0xae, 0x39, 0x4b, 0x71, 0x66, 0x2a, 0x11, 0xf2, 0x1a, 0xda, 0x58, 0x45, 0x6b, 0x78, 0xa4,
	0x95, 0x05, 0x9f, 0x5b, 0x36, 0x89, 0xc3, 0x80, 0xa5, 0xc6, 0x32, 0x23, 0x0f, 0x4c, 0x28, 0xb7,
	0x8c, 0x6c, 0x42, 0x67, 0x4c, 0xc5, 0x8d, 0x77, 0xc3, 0x4a, 0x99, 0x39, 0x3e, 0x67, 0x19, 0x39,
	0x80, 0xbe, 0xa6, 0xa8, 0xaf, 0x85, 0x7b, 0xe2, 0x36, 0x1a, 0xb3, 0x14, 0xc5, 0xae, 0xe5, 0xd4,
	0xb1, 0x61, 0x2e, 0x35, 0x91, 0x0f, 0xba, 0x7f, 0x9b, 0xa6, 0x4c, 0xf8, 0x19, 0x8a, 0x2d, 0xb1,
	0x73, 0x05, 0xfd, 0x93, 0x34, 0x4b, 0x54, 0xfc, 0x7f, 0x4a, 0x77, 0x61, 0xa5, 0xb8, 0xbd, 0xaa,
	0xd8, 0x72, 0x97, 0x31, 0x8c, 0x82, 0x9d, 0x53, 0xe8, 0x8d, 0x68, 0x36, 0xa2, 0xe1, 0x3f, 0x16,
	0x5e, 0x87, 0x26, 0x8b, 0x28, 0x2f, 0x87, 0x40, 0x03, 0xe7, 0xfb, 0x22, 0x2c, 0xcd, 0x0c, 0x76,
	0x5e, 0x27, 0x49, 0x59, 0xc0, 0xae, 0xb9, 0x60, 0x01, 0x0e, 0x52, 0x25, 0x42, 0xce, 0xd1, 0xa3,
	0x62, 0x5b, 0x22, 0x7d, 0x3d, 0xb6, 0x65, 0xb3, 0x68, 0xcb, 0x5c, 0x8b, 0x87, 0x0b, 0xc6, 0xc0,
	0x59, 0x37, 0x3e, 0xc3, 0x86, 0xaf, 0x4d, 0x7a, 0x58, 0xae, 0xae, 0xcb, 0x6d, 0x97, 0x2b, 0x34,
	0xef, 0xe4, 0x70, 0xc1, 0xed, 0xfb, 0x8f, 0x18, 0xfc, 0x1e, 0xf2, 0xd5, 0xf4, 0x12, 0x1a, 0x16,
	0xb5, 0x1a, 0xba, 0xd6, 0x7a, 0x65, 0x8f, 0x4b, 0xd7, 0x86, 0x0b, 0x6e, 0x2f, 0xa9, 0xe0, 0xc1,
	0xda, 0xf4, 0x19, 0x30, 0xa7, 0xa5, 0xe3, 0x41, 0x1b, 0x37, 0x98, 0x10, 0x68, 0xe8, 0xa1, 0x32,
	0x6b, 0xaa, 0xbf, 0x73, 0x5f, 0x15, 0x57, 0x61, 0xb1, 0x9c, 0x06, 0x4c, 0xdd, 0xae, 0x57, 0xdc,
	0xce, 0xa3, 0xc9, 0x24, 0x16, 0x0c, 0xe7, 0xce, 0x80, 0xc1, 0x2b, 0x00, 0x3f, 0x8e, 0xf0, 0xef,
	0x0d, 0xba, 0xe6, 0x65, 0x18, 0xe5, 0x8f, 0xf6, 0xa8, 0xf6, 0x05, 0xdf, 0x82, 0x64, 0x3c, 0x6e,
	0xe9, 0x77, 0xfc, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xd2, 0x38, 0x40, 0x0f, 0x06,
	0x00, 0x00,
}
