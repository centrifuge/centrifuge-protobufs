// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/document_common.proto

package commonpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BinaryAttachment struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//mime type of attached file
	FileType string `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	// in byte
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	//the md5 checksum of the original file for easier verification - optional
	Checksum             []byte   `protobuf:"bytes,5,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinaryAttachment) Reset()         { *m = BinaryAttachment{} }
func (m *BinaryAttachment) String() string { return proto.CompactTextString(m) }
func (*BinaryAttachment) ProtoMessage()    {}
func (*BinaryAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6061ba33c5b40c39, []int{0}
}

func (m *BinaryAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinaryAttachment.Unmarshal(m, b)
}
func (m *BinaryAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinaryAttachment.Marshal(b, m, deterministic)
}
func (m *BinaryAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinaryAttachment.Merge(m, src)
}
func (m *BinaryAttachment) XXX_Size() int {
	return xxx_messageInfo_BinaryAttachment.Size(m)
}
func (m *BinaryAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_BinaryAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_BinaryAttachment proto.InternalMessageInfo

func (m *BinaryAttachment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BinaryAttachment) GetFileType() string {
	if m != nil {
		return m.FileType
	}
	return ""
}

func (m *BinaryAttachment) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BinaryAttachment) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BinaryAttachment) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

type PaymentDetails struct {
	//identifying this payment. could be a sequential number, could be a transaction hash of the crypto payment
	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DateExecuted *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date_executed,json=dateExecuted,proto3" json:"date_executed,omitempty"`
	//centrifuge id of payee
	Payee []byte `protobuf:"bytes,3,opt,name=payee,proto3" json:"payee,omitempty"`
	//centrifuge id of payer
	Payer    []byte `protobuf:"bytes,4,opt,name=payer,proto3" json:"payer,omitempty"`
	Amount   []byte `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency string `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	//payment reference (e.g. reference field on bank transfer)
	Reference             string `protobuf:"bytes,7,opt,name=reference,proto3" json:"reference,omitempty"`
	BankName              string `protobuf:"bytes,8,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	BankAddress           string `protobuf:"bytes,9,opt,name=bank_address,json=bankAddress,proto3" json:"bank_address,omitempty"`
	BankCountry           string `protobuf:"bytes,10,opt,name=bank_country,json=bankCountry,proto3" json:"bank_country,omitempty"`
	BankAccountNumber     string `protobuf:"bytes,11,opt,name=bank_account_number,json=bankAccountNumber,proto3" json:"bank_account_number,omitempty"`
	BankAccountCurrency   string `protobuf:"bytes,12,opt,name=bank_account_currency,json=bankAccountCurrency,proto3" json:"bank_account_currency,omitempty"`
	BankAccountHolderName string `protobuf:"bytes,13,opt,name=bank_account_holder_name,json=bankAccountHolderName,proto3" json:"bank_account_holder_name,omitempty"`
	BankKey               string `protobuf:"bytes,14,opt,name=bank_key,json=bankKey,proto3" json:"bank_key,omitempty"`
	//the ID of the chain to use in URI format. e.g. "ethereum://42/<tokenaddress>"
	CryptoChainUri string `protobuf:"bytes,15,opt,name=crypto_chain_uri,json=cryptoChainUri,proto3" json:"crypto_chain_uri,omitempty"`
	//the transaction in which the payment happened
	CryptoTransactionId string `protobuf:"bytes,16,opt,name=crypto_transaction_id,json=cryptoTransactionId,proto3" json:"crypto_transaction_id,omitempty"`
	//from address
	CryptoFrom string `protobuf:"bytes,17,opt,name=crypto_from,json=cryptoFrom,proto3" json:"crypto_from,omitempty"`
	//to address
	CryptoTo             string   `protobuf:"bytes,18,opt,name=crypto_to,json=cryptoTo,proto3" json:"crypto_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentDetails) Reset()         { *m = PaymentDetails{} }
func (m *PaymentDetails) String() string { return proto.CompactTextString(m) }
func (*PaymentDetails) ProtoMessage()    {}
func (*PaymentDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_6061ba33c5b40c39, []int{1}
}

func (m *PaymentDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentDetails.Unmarshal(m, b)
}
func (m *PaymentDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentDetails.Marshal(b, m, deterministic)
}
func (m *PaymentDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentDetails.Merge(m, src)
}
func (m *PaymentDetails) XXX_Size() int {
	return xxx_messageInfo_PaymentDetails.Size(m)
}
func (m *PaymentDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentDetails.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentDetails proto.InternalMessageInfo

func (m *PaymentDetails) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentDetails) GetDateExecuted() *timestamp.Timestamp {
	if m != nil {
		return m.DateExecuted
	}
	return nil
}

func (m *PaymentDetails) GetPayee() []byte {
	if m != nil {
		return m.Payee
	}
	return nil
}

func (m *PaymentDetails) GetPayer() []byte {
	if m != nil {
		return m.Payer
	}
	return nil
}

func (m *PaymentDetails) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *PaymentDetails) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PaymentDetails) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *PaymentDetails) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

func (m *PaymentDetails) GetBankAddress() string {
	if m != nil {
		return m.BankAddress
	}
	return ""
}

func (m *PaymentDetails) GetBankCountry() string {
	if m != nil {
		return m.BankCountry
	}
	return ""
}

func (m *PaymentDetails) GetBankAccountNumber() string {
	if m != nil {
		return m.BankAccountNumber
	}
	return ""
}

func (m *PaymentDetails) GetBankAccountCurrency() string {
	if m != nil {
		return m.BankAccountCurrency
	}
	return ""
}

func (m *PaymentDetails) GetBankAccountHolderName() string {
	if m != nil {
		return m.BankAccountHolderName
	}
	return ""
}

func (m *PaymentDetails) GetBankKey() string {
	if m != nil {
		return m.BankKey
	}
	return ""
}

func (m *PaymentDetails) GetCryptoChainUri() string {
	if m != nil {
		return m.CryptoChainUri
	}
	return ""
}

func (m *PaymentDetails) GetCryptoTransactionId() string {
	if m != nil {
		return m.CryptoTransactionId
	}
	return ""
}

func (m *PaymentDetails) GetCryptoFrom() string {
	if m != nil {
		return m.CryptoFrom
	}
	return ""
}

func (m *PaymentDetails) GetCryptoTo() string {
	if m != nil {
		return m.CryptoTo
	}
	return ""
}

func init() {
	proto.RegisterType((*BinaryAttachment)(nil), "common.BinaryAttachment")
	proto.RegisterType((*PaymentDetails)(nil), "common.PaymentDetails")
}

func init() { proto.RegisterFile("common/document_common.proto", fileDescriptor_6061ba33c5b40c39) }

var fileDescriptor_6061ba33c5b40c39 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x53, 0x4f, 0x6f, 0xda, 0x30,
	0x14, 0x57, 0x18, 0x50, 0x30, 0x94, 0x51, 0xb3, 0x4e, 0x1e, 0xab, 0x54, 0xd6, 0x13, 0x97, 0x05,
	0xa9, 0x3b, 0xec, 0x38, 0x15, 0xba, 0x69, 0xd3, 0xa4, 0x0a, 0x21, 0x76, 0xd9, 0x25, 0x32, 0xce,
	0xa3, 0x58, 0x60, 0x3b, 0x72, 0x1c, 0x69, 0xd9, 0x17, 0xd8, 0x37, 0xd8, 0xe7, 0xad, 0xfc, 0x9c,
	0x40, 0x7b, 0xf3, 0xef, 0xcf, 0x7b, 0xfa, 0xfd, 0x9c, 0x98, 0x5c, 0x09, 0xa3, 0x94, 0xd1, 0xb3,
	0xd4, 0x88, 0x42, 0x81, 0x76, 0x49, 0xc0, 0x71, 0x66, 0x8d, 0x33, 0xb4, 0x1d, 0xd0, 0xf8, 0xfa,
	0xd1, 0x98, 0xc7, 0x03, 0xcc, 0x90, 0xdd, 0x14, 0xdb, 0x99, 0x93, 0x0a, 0x72, 0xc7, 0x55, 0x16,
	0x8c, 0x37, 0xff, 0x22, 0x32, 0x9c, 0x4b, 0xcd, 0x6d, 0x79, 0xe7, 0x1c, 0x17, 0x3b, 0xbf, 0x8a,
	0x52, 0xd2, 0xd4, 0x5c, 0x01, 0x8b, 0x26, 0xd1, 0xb4, 0xbb, 0xc2, 0x33, 0x7d, 0x4f, 0xba, 0x5b,
	0x79, 0x80, 0xc4, 0x95, 0x19, 0xb0, 0x06, 0x0a, 0x1d, 0x4f, 0xac, 0xcb, 0x0c, 0xfc, 0x40, 0x2e,
	0xff, 0x02, 0x7b, 0x35, 0x89, 0xa6, 0xcd, 0x15, 0x9e, 0x3d, 0x97, 0x72, 0xc7, 0x59, 0x73, 0x12,
	0x4d, 0xfb, 0x2b, 0x3c, 0xd3, 0x31, 0xe9, 0x88, 0x1d, 0x88, 0x7d, 0x5e, 0x28, 0xd6, 0x42, 0xfe,
	0x88, 0x6f, 0xfe, 0xb7, 0xc8, 0x60, 0xc9, 0x4b, 0x1f, 0xe0, 0x1e, 0x1c, 0x97, 0x87, 0x9c, 0x0e,
	0x48, 0x43, 0xa6, 0x55, 0x8a, 0x86, 0x4c, 0xe9, 0x17, 0x72, 0x9e, 0x72, 0x07, 0x09, 0xfc, 0x01,
	0x51, 0x38, 0x48, 0x31, 0x47, 0xef, 0x76, 0x1c, 0x87, 0x96, 0x71, 0xdd, 0x32, 0x5e, 0xd7, 0x2d,
	0x57, 0x7d, 0x3f, 0xf0, 0xb5, 0xf2, 0xd3, 0x37, 0xa4, 0x95, 0xf1, 0x12, 0x42, 0xd0, 0xfe, 0x2a,
	0x80, 0x9a, 0xb5, 0x55, 0xd4, 0x00, 0xe8, 0x5b, 0xd2, 0xe6, 0xca, 0x14, 0xda, 0x55, 0x49, 0x2b,
	0x84, 0x1d, 0x0a, 0x6b, 0x41, 0x8b, 0x92, 0xb5, 0xc3, 0x3d, 0xd4, 0x98, 0x5e, 0x91, 0xae, 0x85,
	0x2d, 0x78, 0x00, 0xec, 0x0c, 0xc5, 0x13, 0xe1, 0xaf, 0x70, 0xc3, 0xf5, 0x3e, 0xc1, 0xbb, 0xed,
	0x84, 0x51, 0x4f, 0x3c, 0xf8, 0xfb, 0xfd, 0x40, 0xfa, 0x28, 0xf2, 0x34, 0xb5, 0x90, 0xe7, 0xac,
	0x8b, 0x7a, 0xcf, 0x73, 0x77, 0x81, 0x3a, 0x5a, 0x84, 0xcf, 0x61, 0x4b, 0x46, 0x4e, 0x96, 0x45,
	0xa0, 0x68, 0x4c, 0x46, 0x61, 0x8b, 0x40, 0x53, 0xa2, 0x0b, 0xb5, 0x01, 0xcb, 0x7a, 0xe8, 0xbc,
	0xc0, 0x65, 0x41, 0x79, 0x40, 0x81, 0xde, 0x92, 0xcb, 0x17, 0xfe, 0x63, 0xb3, 0x3e, 0x4e, 0x8c,
	0x9e, 0x4d, 0x2c, 0xea, 0x92, 0x9f, 0x09, 0x7b, 0x31, 0xb3, 0x33, 0x87, 0x14, 0x6c, 0x68, 0x75,
	0x8e, 0x63, 0x97, 0xcf, 0xc6, 0xbe, 0xa3, 0x8a, 0x15, 0xdf, 0x11, 0xac, 0x9b, 0xec, 0xa1, 0x64,
	0x03, 0x34, 0x9e, 0x79, 0xfc, 0x13, 0x4a, 0x3a, 0x25, 0x43, 0x61, 0xcb, 0xcc, 0x99, 0x44, 0xec,
	0xb8, 0xd4, 0x49, 0x61, 0x25, 0x7b, 0x8d, 0x96, 0x41, 0xe0, 0x17, 0x9e, 0xfe, 0x65, 0xa5, 0x4f,
	0x5c, 0x39, 0x9d, 0xe5, 0x3a, 0xe7, 0xc2, 0x49, 0xa3, 0x13, 0x99, 0xb2, 0x61, 0x48, 0x1c, 0xc4,
	0xf5, 0x49, 0xfb, 0x91, 0xd2, 0x6b, 0xd2, 0xab, 0x66, 0xb6, 0xd6, 0x28, 0x76, 0x81, 0x4e, 0x12,
	0xa8, 0x6f, 0xd6, 0x28, 0xff, 0x65, 0xea, 0xa5, 0x86, 0xd1, 0xea, 0xa3, 0x86, 0x45, 0x66, 0xfe,
	0x91, 0x10, 0x61, 0x54, 0x1c, 0x5e, 0xd4, 0x7c, 0x74, 0x5f, 0x3d, 0xb8, 0x05, 0xe2, 0xa5, 0xff,
	0xe5, 0x96, 0xd1, 0xef, 0x4e, 0x90, 0xb3, 0xcd, 0xa6, 0x8d, 0x7f, 0xe1, 0xa7, 0xa7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x71, 0x39, 0xa1, 0x79, 0xa1, 0x03, 0x00, 0x00,
}
