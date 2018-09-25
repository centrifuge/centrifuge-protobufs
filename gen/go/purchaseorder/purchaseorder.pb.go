// Code generated by protoc-gen-go. DO NOT EDIT.
// source: purchaseorder/purchaseorder.proto

package purchaseorderpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import coredocument "github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PurchaseOrderData is the default schema for a purchase order
type PurchaseOrderData struct {
	// purchase order number or reference number
	PoNumber string `protobuf:"bytes,1,opt,name=po_number,json=poNumber,proto3" json:"po_number,omitempty"`
	// name of the ordering company
	OrderName string `protobuf:"bytes,2,opt,name=order_name,json=orderName,proto3" json:"order_name,omitempty"`
	// street and address details of the ordering company
	OrderStreet  string `protobuf:"bytes,3,opt,name=order_street,json=orderStreet,proto3" json:"order_street,omitempty"`
	OrderCity    string `protobuf:"bytes,4,opt,name=order_city,json=orderCity,proto3" json:"order_city,omitempty"`
	OrderZipcode string `protobuf:"bytes,5,opt,name=order_zipcode,json=orderZipcode,proto3" json:"order_zipcode,omitempty"`
	// country ISO code of the ordering company of this purchase order
	OrderCountry string `protobuf:"bytes,6,opt,name=order_country,json=orderCountry,proto3" json:"order_country,omitempty"`
	// name of the recipient company
	RecipientName    string `protobuf:"bytes,7,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	RecipientStreet  string `protobuf:"bytes,8,opt,name=recipient_street,json=recipientStreet,proto3" json:"recipient_street,omitempty"`
	RecipientCity    string `protobuf:"bytes,9,opt,name=recipient_city,json=recipientCity,proto3" json:"recipient_city,omitempty"`
	RecipientZipcode string `protobuf:"bytes,10,opt,name=recipient_zipcode,json=recipientZipcode,proto3" json:"recipient_zipcode,omitempty"`
	// country ISO code of the receipient of this purchase order
	RecipientCountry string `protobuf:"bytes,11,opt,name=recipient_country,json=recipientCountry,proto3" json:"recipient_country,omitempty"`
	// ISO currency code
	Currency string `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	// ordering gross amount including tax
	OrderAmount int64 `protobuf:"varint,13,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	// invoice amount excluding tax
	NetAmount int64  `protobuf:"varint,14,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount int64  `protobuf:"varint,15,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate   int64  `protobuf:"varint,16,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	Recipient []byte `protobuf:"bytes,17,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Order     []byte `protobuf:"bytes,18,opt,name=order,proto3" json:"order,omitempty"`
	// contact or requester or purchaser at the ordering company
	OrderContact string `protobuf:"bytes,19,opt,name=order_contact,json=orderContact,proto3" json:"order_contact,omitempty"`
	Comment      string `protobuf:"bytes,20,opt,name=comment,proto3" json:"comment,omitempty"`
	// requested delivery date
	DeliveryDate *timestamp.Timestamp `protobuf:"bytes,21,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	// purchase order date
	DateCreated          *timestamp.Timestamp `protobuf:"bytes,22,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	ExtraData            []byte               `protobuf:"bytes,23,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PurchaseOrderData) Reset()         { *m = PurchaseOrderData{} }
func (m *PurchaseOrderData) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderData) ProtoMessage()    {}
func (*PurchaseOrderData) Descriptor() ([]byte, []int) {
	return fileDescriptor_purchaseorder_d3f89a918fc49fe1, []int{0}
}
func (m *PurchaseOrderData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderData.Unmarshal(m, b)
}
func (m *PurchaseOrderData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderData.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderData.Merge(dst, src)
}
func (m *PurchaseOrderData) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderData.Size(m)
}
func (m *PurchaseOrderData) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderData.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderData proto.InternalMessageInfo

func (m *PurchaseOrderData) GetPoNumber() string {
	if m != nil {
		return m.PoNumber
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderStreet() string {
	if m != nil {
		return m.OrderStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCity() string {
	if m != nil {
		return m.OrderCity
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderZipcode() string {
	if m != nil {
		return m.OrderZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderCountry() string {
	if m != nil {
		return m.OrderCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientName() string {
	if m != nil {
		return m.RecipientName
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientStreet() string {
	if m != nil {
		return m.RecipientStreet
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCity() string {
	if m != nil {
		return m.RecipientCity
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientZipcode() string {
	if m != nil {
		return m.RecipientZipcode
	}
	return ""
}

func (m *PurchaseOrderData) GetRecipientCountry() string {
	if m != nil {
		return m.RecipientCountry
	}
	return ""
}

func (m *PurchaseOrderData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PurchaseOrderData) GetOrderAmount() int64 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetNetAmount() int64 {
	if m != nil {
		return m.NetAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetTaxAmount() int64 {
	if m != nil {
		return m.TaxAmount
	}
	return 0
}

func (m *PurchaseOrderData) GetTaxRate() int64 {
	if m != nil {
		return m.TaxRate
	}
	return 0
}

func (m *PurchaseOrderData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *PurchaseOrderData) GetOrder() []byte {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *PurchaseOrderData) GetOrderContact() string {
	if m != nil {
		return m.OrderContact
	}
	return ""
}

func (m *PurchaseOrderData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *PurchaseOrderData) GetDeliveryDate() *timestamp.Timestamp {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

func (m *PurchaseOrderData) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *PurchaseOrderData) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

// PurhcaseOrderDataSalts keeps track of the salts used for each PurchaseOrderData field.
type PurchaseOrderDataSalts struct {
	PoNumber             []byte   `protobuf:"bytes,1,opt,name=po_number,json=poNumber,proto3" json:"po_number,omitempty"`
	OrderName            []byte   `protobuf:"bytes,2,opt,name=order_name,json=orderName,proto3" json:"order_name,omitempty"`
	OrderStreet          []byte   `protobuf:"bytes,3,opt,name=order_street,json=orderStreet,proto3" json:"order_street,omitempty"`
	OrderCity            []byte   `protobuf:"bytes,4,opt,name=order_city,json=orderCity,proto3" json:"order_city,omitempty"`
	OrderZipcode         []byte   `protobuf:"bytes,5,opt,name=order_zipcode,json=orderZipcode,proto3" json:"order_zipcode,omitempty"`
	OrderCountry         []byte   `protobuf:"bytes,6,opt,name=order_country,json=orderCountry,proto3" json:"order_country,omitempty"`
	RecipientName        []byte   `protobuf:"bytes,7,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	RecipientStreet      []byte   `protobuf:"bytes,8,opt,name=recipient_street,json=recipientStreet,proto3" json:"recipient_street,omitempty"`
	RecipientCity        []byte   `protobuf:"bytes,9,opt,name=recipient_city,json=recipientCity,proto3" json:"recipient_city,omitempty"`
	RecipientZipcode     []byte   `protobuf:"bytes,10,opt,name=recipient_zipcode,json=recipientZipcode,proto3" json:"recipient_zipcode,omitempty"`
	RecipientCountry     []byte   `protobuf:"bytes,11,opt,name=recipient_country,json=recipientCountry,proto3" json:"recipient_country,omitempty"`
	Currency             []byte   `protobuf:"bytes,12,opt,name=currency,proto3" json:"currency,omitempty"`
	OrderAmount          []byte   `protobuf:"bytes,13,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	NetAmount            []byte   `protobuf:"bytes,14,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount            []byte   `protobuf:"bytes,15,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate              []byte   `protobuf:"bytes,16,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	Recipient            []byte   `protobuf:"bytes,17,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Order                []byte   `protobuf:"bytes,18,opt,name=order,proto3" json:"order,omitempty"`
	OrderContact         []byte   `protobuf:"bytes,19,opt,name=order_contact,json=orderContact,proto3" json:"order_contact,omitempty"`
	Comment              []byte   `protobuf:"bytes,20,opt,name=comment,proto3" json:"comment,omitempty"`
	DeliveryDate         []byte   `protobuf:"bytes,21,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	DateCreated          []byte   `protobuf:"bytes,22,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	ExtraData            []byte   `protobuf:"bytes,23,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurchaseOrderDataSalts) Reset()         { *m = PurchaseOrderDataSalts{} }
func (m *PurchaseOrderDataSalts) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderDataSalts) ProtoMessage()    {}
func (*PurchaseOrderDataSalts) Descriptor() ([]byte, []int) {
	return fileDescriptor_purchaseorder_d3f89a918fc49fe1, []int{1}
}
func (m *PurchaseOrderDataSalts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderDataSalts.Unmarshal(m, b)
}
func (m *PurchaseOrderDataSalts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderDataSalts.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderDataSalts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderDataSalts.Merge(dst, src)
}
func (m *PurchaseOrderDataSalts) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderDataSalts.Size(m)
}
func (m *PurchaseOrderDataSalts) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderDataSalts.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderDataSalts proto.InternalMessageInfo

func (m *PurchaseOrderDataSalts) GetPoNumber() []byte {
	if m != nil {
		return m.PoNumber
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrderName() []byte {
	if m != nil {
		return m.OrderName
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrderStreet() []byte {
	if m != nil {
		return m.OrderStreet
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrderCity() []byte {
	if m != nil {
		return m.OrderCity
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrderZipcode() []byte {
	if m != nil {
		return m.OrderZipcode
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrderCountry() []byte {
	if m != nil {
		return m.OrderCountry
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetRecipientName() []byte {
	if m != nil {
		return m.RecipientName
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetRecipientStreet() []byte {
	if m != nil {
		return m.RecipientStreet
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetRecipientCity() []byte {
	if m != nil {
		return m.RecipientCity
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetRecipientZipcode() []byte {
	if m != nil {
		return m.RecipientZipcode
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetRecipientCountry() []byte {
	if m != nil {
		return m.RecipientCountry
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetCurrency() []byte {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrderAmount() []byte {
	if m != nil {
		return m.OrderAmount
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetNetAmount() []byte {
	if m != nil {
		return m.NetAmount
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetTaxAmount() []byte {
	if m != nil {
		return m.TaxAmount
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetTaxRate() []byte {
	if m != nil {
		return m.TaxRate
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrder() []byte {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetOrderContact() []byte {
	if m != nil {
		return m.OrderContact
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetComment() []byte {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetDeliveryDate() []byte {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetDateCreated() []byte {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

// PurchaseOrderDocument combines the salts, data & coredocument for a purchase order.
type PurchaseOrderDocument struct {
	CoreDocument         *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=core_document,json=coreDocument,proto3" json:"core_document,omitempty"`
	Data                 *PurchaseOrderData         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Salts                *PurchaseOrderDataSalts    `protobuf:"bytes,3,opt,name=salts,proto3" json:"salts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PurchaseOrderDocument) Reset()         { *m = PurchaseOrderDocument{} }
func (m *PurchaseOrderDocument) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderDocument) ProtoMessage()    {}
func (*PurchaseOrderDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_purchaseorder_d3f89a918fc49fe1, []int{2}
}
func (m *PurchaseOrderDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderDocument.Unmarshal(m, b)
}
func (m *PurchaseOrderDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderDocument.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderDocument.Merge(dst, src)
}
func (m *PurchaseOrderDocument) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderDocument.Size(m)
}
func (m *PurchaseOrderDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderDocument.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderDocument proto.InternalMessageInfo

func (m *PurchaseOrderDocument) GetCoreDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.CoreDocument
	}
	return nil
}

func (m *PurchaseOrderDocument) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PurchaseOrderDocument) GetSalts() *PurchaseOrderDataSalts {
	if m != nil {
		return m.Salts
	}
	return nil
}

func init() {
	proto.RegisterType((*PurchaseOrderData)(nil), "purchaseorder.PurchaseOrderData")
	proto.RegisterType((*PurchaseOrderDataSalts)(nil), "purchaseorder.PurchaseOrderDataSalts")
	proto.RegisterType((*PurchaseOrderDocument)(nil), "purchaseorder.PurchaseOrderDocument")
}

func init() {
	proto.RegisterFile("purchaseorder/purchaseorder.proto", fileDescriptor_purchaseorder_d3f89a918fc49fe1)
}

var fileDescriptor_purchaseorder_d3f89a918fc49fe1 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0xfd, 0x93, 0x26, 0xce, 0xa4, 0x6d, 0x4c, 0x5b, 0x4c, 0x00, 0x35, 0x6d, 0x55,
	0x29, 0x08, 0x29, 0x91, 0x0a, 0x27, 0x10, 0x42, 0x34, 0x3d, 0x97, 0x68, 0xcb, 0xa9, 0x97, 0xc8,
	0x71, 0x4c, 0x59, 0xa9, 0xbb, 0x5e, 0x39, 0x0e, 0x4a, 0x78, 0x1f, 0x9e, 0x86, 0x17, 0xe0, 0x71,
	0x90, 0xc7, 0x71, 0x37, 0xdb, 0x56, 0xd9, 0xed, 0x81, 0xe3, 0x7c, 0xf3, 0xcd, 0x78, 0xec, 0xf5,
	0x6f, 0x4d, 0x8e, 0xd2, 0xa9, 0x16, 0x3f, 0xf8, 0x44, 0x2a, 0x3d, 0x96, 0xba, 0x97, 0x8b, 0xba,
	0xa9, 0x56, 0x46, 0xd1, 0x46, 0x4e, 0x6c, 0x1d, 0x0a, 0xa5, 0xe5, 0x58, 0x89, 0x69, 0x2c, 0x13,
	0xd3, 0x5b, 0x0e, 0x9c, 0xbf, 0x75, 0x78, 0xa3, 0xd4, 0xcd, 0xad, 0xec, 0x61, 0x34, 0x9a, 0x7e,
	0xef, 0x99, 0x28, 0x96, 0x13, 0xc3, 0xe3, 0xd4, 0x19, 0x8e, 0xff, 0x56, 0x48, 0x73, 0xb0, 0xe8,
	0xf9, 0xd5, 0xf6, 0xbc, 0xe0, 0x86, 0xd3, 0x97, 0xa4, 0x96, 0xaa, 0x61, 0x32, 0x8d, 0x47, 0x52,
	0xb3, 0xa0, 0x1d, 0x74, 0x6a, 0x61, 0x35, 0x55, 0x97, 0x18, 0xd3, 0xd7, 0x84, 0xe0, 0xea, 0xc3,
	0x84, 0xc7, 0x92, 0xad, 0x61, 0xb6, 0x86, 0xca, 0x25, 0x8f, 0x25, 0x3d, 0x22, 0xe0, 0xd2, 0x13,
	0xa3, 0xa5, 0x34, 0x6c, 0x1d, 0x0d, 0x75, 0xd4, 0xae, 0x50, 0xca, 0x3a, 0x88, 0xc8, 0xcc, 0xd9,
	0xc6, 0x52, 0x87, 0x7e, 0x64, 0xe6, 0xf4, 0x84, 0x34, 0x5c, 0xfa, 0x57, 0x94, 0x0a, 0x35, 0x96,
	0x6c, 0x13, 0x1d, 0xae, 0xed, 0xb5, 0xd3, 0x32, 0x93, 0x50, 0xd3, 0xc4, 0xe8, 0x39, 0xab, 0x2c,
	0x99, 0xfa, 0x4e, 0xa3, 0xa7, 0x64, 0x5b, 0x4b, 0x11, 0xa5, 0x91, 0x4c, 0x8c, 0x1b, 0x77, 0x0b,
	0x5d, 0x8d, 0x3b, 0x15, 0x47, 0x7e, 0x43, 0x76, 0x33, 0xdb, 0x62, 0xec, 0x2a, 0x1a, 0x77, 0xee,
	0xf4, 0xc5, 0xe8, 0xb9, 0x8e, 0x38, 0x7e, 0xed, 0x5e, 0x47, 0xdc, 0xc2, 0x5b, 0xd2, 0xcc, 0x6c,
	0x7e, 0x1b, 0x04, 0x9d, 0xd9, 0x52, 0x7e, 0x2b, 0x39, 0xb3, 0xdf, 0x4e, 0xfd, 0x9e, 0xd9, 0x6f,
	0xa9, 0x45, 0xaa, 0x62, 0xaa, 0xb5, 0x4c, 0xc4, 0x9c, 0x81, 0xfb, 0x32, 0x3e, 0xce, 0x8e, 0x9e,
	0xc7, 0xd6, 0xcd, 0x1a, 0xed, 0xa0, 0xb3, 0xbe, 0x38, 0xfa, 0x2f, 0x28, 0xd9, 0xa3, 0x4f, 0xa4,
	0xf1, 0x86, 0x6d, 0x34, 0xd4, 0x12, 0x69, 0xb2, 0xb4, 0xe1, 0x33, 0x9f, 0xde, 0x71, 0x69, 0xc3,
	0x67, 0x8b, 0xf4, 0x0b, 0x52, 0xb5, 0x69, 0xcd, 0x8d, 0x64, 0xbb, 0x98, 0xdc, 0x32, 0x7c, 0x16,
	0x72, 0x23, 0xe9, 0x2b, 0x52, 0xbb, 0x9b, 0x95, 0x35, 0xdb, 0x41, 0x07, 0xc2, 0x4c, 0xa0, 0x7b,
	0x64, 0x13, 0xa7, 0x60, 0x14, 0x33, 0x2e, 0x58, 0xfe, 0x86, 0x89, 0xe1, 0xc2, 0xb0, 0x67, 0xb9,
	0x6f, 0x88, 0x1a, 0x65, 0x64, 0x4b, 0xa8, 0xd8, 0xde, 0x69, 0xb6, 0x87, 0x69, 0x1f, 0xd2, 0xcf,
	0xa4, 0x31, 0x96, 0xb7, 0xd1, 0x4f, 0xa9, 0xe7, 0xc3, 0xb1, 0x1d, 0x69, 0xbf, 0x1d, 0x74, 0xea,
	0x67, 0xad, 0xae, 0xbb, 0xf4, 0x5d, 0x7f, 0xe9, 0xbb, 0xdf, 0xfc, 0xa5, 0x0f, 0xc1, 0x17, 0x5c,
	0xd8, 0x99, 0x3f, 0x11, 0xb0, 0x75, 0x43, 0xa1, 0x25, 0x37, 0x72, 0xcc, 0x0e, 0x0a, 0xeb, 0xeb,
	0xd6, 0xdf, 0x77, 0x76, 0x7b, 0x58, 0x72, 0x66, 0x34, 0xb7, 0x8b, 0x73, 0xf6, 0xdc, 0xed, 0x19,
	0x15, 0x0b, 0xd1, 0xf1, 0xef, 0x0a, 0x39, 0x78, 0x80, 0xd6, 0x15, 0xbf, 0x35, 0x93, 0x87, 0x7c,
	0xc1, 0x4a, 0xbe, 0xa0, 0x88, 0x2f, 0x28, 0xe2, 0x0b, 0x0a, 0xf9, 0x82, 0x32, 0x7c, 0x41, 0x29,
	0xbe, 0xa0, 0x2c, 0x5f, 0x50, 0x96, 0x2f, 0x28, 0xcd, 0x17, 0x3c, 0x85, 0x2f, 0x28, 0xc1, 0x17,
	0x14, 0xf0, 0x05, 0x45, 0x7c, 0xc1, 0x6a, 0xbe, 0x60, 0x15, 0x5f, 0xf0, 0x1f, 0xf8, 0x82, 0xd5,
	0x7c, 0x41, 0xc6, 0xd7, 0xc9, 0x63, 0x7c, 0xc1, 0x3d, 0x86, 0x8e, 0x1e, 0x61, 0x08, 0x9e, 0xc4,
	0xc9, 0x9f, 0x80, 0xec, 0xe7, 0x39, 0x59, 0xbc, 0x61, 0x16, 0x70, 0xfb, 0xa6, 0x0d, 0xfd, 0xa3,
	0x86, 0xa8, 0x58, 0x40, 0x73, 0x2f, 0x5d, 0x5f, 0x69, 0xe9, 0x4b, 0x42, 0x10, 0x4b, 0x11, 0x7d,
	0x4f, 0x36, 0x70, 0xcd, 0x35, 0xac, 0x6b, 0x77, 0xf3, 0x4f, 0xea, 0x03, 0x38, 0x43, 0x74, 0xd3,
	0x8f, 0x64, 0x73, 0x62, 0x31, 0x45, 0xb4, 0xea, 0x67, 0xa7, 0x45, 0x65, 0xc8, 0x74, 0xe8, 0x6a,
	0xce, 0x3f, 0x90, 0xa6, 0x50, 0x71, 0xbe, 0xe4, 0x9c, 0x0e, 0x96, 0xc3, 0x81, 0xfd, 0xaf, 0x0c,
	0x82, 0xeb, 0x9d, 0x9c, 0x29, 0x1d, 0x8d, 0x2a, 0xf8, 0xc7, 0x79, 0xf7, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x34, 0x5b, 0x91, 0x88, 0x09, 0x08, 0x00, 0x00,
}
