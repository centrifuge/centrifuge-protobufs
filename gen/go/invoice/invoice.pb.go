// Code generated by protoc-gen-go. DO NOT EDIT.
// source: invoice/invoice.proto

package invoicepb

import (
	fmt "fmt"
	common "github.com/centrifuge/centrifuge-protobufs/gen/go/common"
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

// InvoiceData is the default invoice schema
type InvoiceData struct {
	// invoice number or reference number
	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	// invoice status
	Status                   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	SenderInvoiceId          string `protobuf:"bytes,3,opt,name=sender_invoice_id,json=senderInvoiceId,proto3" json:"sender_invoice_id,omitempty"`
	RecipientInvoiceId       string `protobuf:"bytes,4,opt,name=recipient_invoice_id,json=recipientInvoiceId,proto3" json:"recipient_invoice_id,omitempty"`
	SenderCompanyName        string `protobuf:"bytes,5,opt,name=sender_company_name,json=senderCompanyName,proto3" json:"sender_company_name,omitempty"`
	SenderContactPersonName  string `protobuf:"bytes,6,opt,name=sender_contact_person_name,json=senderContactPersonName,proto3" json:"sender_contact_person_name,omitempty"`
	SenderStreet1            string `protobuf:"bytes,7,opt,name=sender_street1,json=senderStreet1,proto3" json:"sender_street1,omitempty"`
	SenderStreet2            string `protobuf:"bytes,8,opt,name=sender_street2,json=senderStreet2,proto3" json:"sender_street2,omitempty"`
	SenderCity               string `protobuf:"bytes,9,opt,name=sender_city,json=senderCity,proto3" json:"sender_city,omitempty"`
	SenderZipcode            string `protobuf:"bytes,10,opt,name=sender_zipcode,json=senderZipcode,proto3" json:"sender_zipcode,omitempty"`
	SenderState              string `protobuf:"bytes,11,opt,name=sender_state,json=senderState,proto3" json:"sender_state,omitempty"`
	SenderCountry            string `protobuf:"bytes,12,opt,name=sender_country,json=senderCountry,proto3" json:"sender_country,omitempty"`
	BillToCompanyName        string `protobuf:"bytes,43,opt,name=bill_to_company_name,json=billToCompanyName,proto3" json:"bill_to_company_name,omitempty"`
	BillToContactPersonName  string `protobuf:"bytes,44,opt,name=bill_to_contact_person_name,json=billToContactPersonName,proto3" json:"bill_to_contact_person_name,omitempty"`
	BillToStreet1            string `protobuf:"bytes,15,opt,name=bill_to_street1,json=billToStreet1,proto3" json:"bill_to_street1,omitempty"`
	BillToStreet2            string `protobuf:"bytes,16,opt,name=bill_to_street2,json=billToStreet2,proto3" json:"bill_to_street2,omitempty"`
	BillToCity               string `protobuf:"bytes,17,opt,name=bill_to_city,json=billToCity,proto3" json:"bill_to_city,omitempty"`
	BillToZipcode            string `protobuf:"bytes,18,opt,name=bill_to_zipcode,json=billToZipcode,proto3" json:"bill_to_zipcode,omitempty"`
	BillToState              string `protobuf:"bytes,50,opt,name=bill_to_state,json=billToState,proto3" json:"bill_to_state,omitempty"`
	BillToCountry            string `protobuf:"bytes,20,opt,name=bill_to_country,json=billToCountry,proto3" json:"bill_to_country,omitempty"`
	BillToVatNumber          string `protobuf:"bytes,21,opt,name=bill_to_vat_number,json=billToVatNumber,proto3" json:"bill_to_vat_number,omitempty"`
	BillToLocalTaxId         string `protobuf:"bytes,60,opt,name=bill_to_local_tax_id,json=billToLocalTaxId,proto3" json:"bill_to_local_tax_id,omitempty"`
	RemitToCompanyName       string `protobuf:"bytes,23,opt,name=remit_to_company_name,json=remitToCompanyName,proto3" json:"remit_to_company_name,omitempty"`
	RemitToContactPersonName string `protobuf:"bytes,24,opt,name=remit_to_contact_person_name,json=remitToContactPersonName,proto3" json:"remit_to_contact_person_name,omitempty"`
	RemitToStreet1           string `protobuf:"bytes,25,opt,name=remit_to_street1,json=remitToStreet1,proto3" json:"remit_to_street1,omitempty"`
	RemitToStreet2           string `protobuf:"bytes,26,opt,name=remit_to_street2,json=remitToStreet2,proto3" json:"remit_to_street2,omitempty"`
	RemitToCity              string `protobuf:"bytes,27,opt,name=remit_to_city,json=remitToCity,proto3" json:"remit_to_city,omitempty"`
	RemitToZipcode           string `protobuf:"bytes,28,opt,name=remit_to_zipcode,json=remitToZipcode,proto3" json:"remit_to_zipcode,omitempty"`
	RemitToState             string `protobuf:"bytes,30,opt,name=remit_to_state,json=remitToState,proto3" json:"remit_to_state,omitempty"`
	RemitToCountry           string `protobuf:"bytes,31,opt,name=remit_to_country,json=remitToCountry,proto3" json:"remit_to_country,omitempty"`
	RemitToVatNumber         string `protobuf:"bytes,32,opt,name=remit_to_vat_number,json=remitToVatNumber,proto3" json:"remit_to_vat_number,omitempty"`
	RemitToLocalTaxId        string `protobuf:"bytes,33,opt,name=remit_to_local_tax_id,json=remitToLocalTaxId,proto3" json:"remit_to_local_tax_id,omitempty"`
	RemitToTaxCountry        string `protobuf:"bytes,34,opt,name=remit_to_tax_country,json=remitToTaxCountry,proto3" json:"remit_to_tax_country,omitempty"`
	ShipToCompanyName        string `protobuf:"bytes,35,opt,name=ship_to_company_name,json=shipToCompanyName,proto3" json:"ship_to_company_name,omitempty"`
	ShipToContactPersonName  string `protobuf:"bytes,36,opt,name=ship_to_contact_person_name,json=shipToContactPersonName,proto3" json:"ship_to_contact_person_name,omitempty"`
	ShipToStreet1            string `protobuf:"bytes,37,opt,name=ship_to_street1,json=shipToStreet1,proto3" json:"ship_to_street1,omitempty"`
	ShipToStreet2            string `protobuf:"bytes,38,opt,name=ship_to_street2,json=shipToStreet2,proto3" json:"ship_to_street2,omitempty"`
	ShipToCity               string `protobuf:"bytes,39,opt,name=ship_to_city,json=shipToCity,proto3" json:"ship_to_city,omitempty"`
	ShipToZipcode            string `protobuf:"bytes,40,opt,name=ship_to_zipcode,json=shipToZipcode,proto3" json:"ship_to_zipcode,omitempty"`
	ShipToState              string `protobuf:"bytes,41,opt,name=ship_to_state,json=shipToState,proto3" json:"ship_to_state,omitempty"`
	ShipToCountry            string `protobuf:"bytes,42,opt,name=ship_to_country,json=shipToCountry,proto3" json:"ship_to_country,omitempty"`
	Currency                 string `protobuf:"bytes,13,opt,name=currency,proto3" json:"currency,omitempty"`
	GrossAmount              []byte `protobuf:"bytes,14,opt,name=gross_amount,json=grossAmount,proto3" json:"gross_amount,omitempty"`
	NetAmount                []byte `protobuf:"bytes,45,opt,name=net_amount,json=netAmount,proto3" json:"net_amount,omitempty"`
	TaxAmount                []byte `protobuf:"bytes,46,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate                  []byte `protobuf:"bytes,47,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	//bool tax_on_line_level = 48;
	// centrifuge ID of the recipient
	Recipient []byte `protobuf:"bytes,49,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// centrifuge ID of the sender
	Sender []byte `protobuf:"bytes,19,opt,name=sender,proto3" json:"sender,omitempty"`
	// centrifuge ID of the payee
	Payee          []byte `protobuf:"bytes,51,opt,name=payee,proto3" json:"payee,omitempty"`
	Comment        string `protobuf:"bytes,52,opt,name=comment,proto3" json:"comment,omitempty"`
	ShippingTerms  string `protobuf:"bytes,53,opt,name=shipping_terms,json=shippingTerms,proto3" json:"shipping_terms,omitempty"`
	RequesterEmail string `protobuf:"bytes,54,opt,name=requester_email,json=requesterEmail,proto3" json:"requester_email,omitempty"`
	RequesterName  string `protobuf:"bytes,55,opt,name=requester_name,json=requesterName,proto3" json:"requester_name,omitempty"`
	//number of the delivery note
	DeliveryNumber string `protobuf:"bytes,56,opt,name=delivery_number,json=deliveryNumber,proto3" json:"delivery_number,omitempty"`
	//bool is_credit_note = 57;
	CreditNoteInvoiceNumber string                     `protobuf:"bytes,58,opt,name=credit_note_invoice_number,json=creditNoteInvoiceNumber,proto3" json:"credit_note_invoice_number,omitempty"`
	CreditForInvoiceDate    *timestamp.Timestamp       `protobuf:"bytes,59,opt,name=credit_for_invoice_date,json=creditForInvoiceDate,proto3" json:"credit_for_invoice_date,omitempty"`
	DateDue                 *timestamp.Timestamp       `protobuf:"bytes,22,opt,name=date_due,json=dateDue,proto3" json:"date_due,omitempty"`
	DatePaid                *timestamp.Timestamp       `protobuf:"bytes,61,opt,name=date_paid,json=datePaid,proto3" json:"date_paid,omitempty"`
	DateUpdated             *timestamp.Timestamp       `protobuf:"bytes,62,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	DateCreated             *timestamp.Timestamp       `protobuf:"bytes,63,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	Attachments             []*common.BinaryAttachment `protobuf:"bytes,64,rep,name=attachments,proto3" json:"attachments,omitempty"`
	LineItems               []*LineItem                `protobuf:"bytes,65,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
	PaymentDetails          []*common.PaymentDetails   `protobuf:"bytes,66,rep,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	TaxItems                []*TaxItem                 `protobuf:"bytes,67,rep,name=tax_items,json=taxItems,proto3" json:"tax_items,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                   `json:"-"`
	XXX_unrecognized        []byte                     `json:"-"`
	XXX_sizecache           int32                      `json:"-"`
}

func (m *InvoiceData) Reset()         { *m = InvoiceData{} }
func (m *InvoiceData) String() string { return proto.CompactTextString(m) }
func (*InvoiceData) ProtoMessage()    {}
func (*InvoiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3e2b5ce0fcadd51, []int{0}
}

func (m *InvoiceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceData.Unmarshal(m, b)
}
func (m *InvoiceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceData.Marshal(b, m, deterministic)
}
func (m *InvoiceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceData.Merge(m, src)
}
func (m *InvoiceData) XXX_Size() int {
	return xxx_messageInfo_InvoiceData.Size(m)
}
func (m *InvoiceData) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceData.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceData proto.InternalMessageInfo

func (m *InvoiceData) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *InvoiceData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *InvoiceData) GetSenderInvoiceId() string {
	if m != nil {
		return m.SenderInvoiceId
	}
	return ""
}

func (m *InvoiceData) GetRecipientInvoiceId() string {
	if m != nil {
		return m.RecipientInvoiceId
	}
	return ""
}

func (m *InvoiceData) GetSenderCompanyName() string {
	if m != nil {
		return m.SenderCompanyName
	}
	return ""
}

func (m *InvoiceData) GetSenderContactPersonName() string {
	if m != nil {
		return m.SenderContactPersonName
	}
	return ""
}

func (m *InvoiceData) GetSenderStreet1() string {
	if m != nil {
		return m.SenderStreet1
	}
	return ""
}

func (m *InvoiceData) GetSenderStreet2() string {
	if m != nil {
		return m.SenderStreet2
	}
	return ""
}

func (m *InvoiceData) GetSenderCity() string {
	if m != nil {
		return m.SenderCity
	}
	return ""
}

func (m *InvoiceData) GetSenderZipcode() string {
	if m != nil {
		return m.SenderZipcode
	}
	return ""
}

func (m *InvoiceData) GetSenderState() string {
	if m != nil {
		return m.SenderState
	}
	return ""
}

func (m *InvoiceData) GetSenderCountry() string {
	if m != nil {
		return m.SenderCountry
	}
	return ""
}

func (m *InvoiceData) GetBillToCompanyName() string {
	if m != nil {
		return m.BillToCompanyName
	}
	return ""
}

func (m *InvoiceData) GetBillToContactPersonName() string {
	if m != nil {
		return m.BillToContactPersonName
	}
	return ""
}

func (m *InvoiceData) GetBillToStreet1() string {
	if m != nil {
		return m.BillToStreet1
	}
	return ""
}

func (m *InvoiceData) GetBillToStreet2() string {
	if m != nil {
		return m.BillToStreet2
	}
	return ""
}

func (m *InvoiceData) GetBillToCity() string {
	if m != nil {
		return m.BillToCity
	}
	return ""
}

func (m *InvoiceData) GetBillToZipcode() string {
	if m != nil {
		return m.BillToZipcode
	}
	return ""
}

func (m *InvoiceData) GetBillToState() string {
	if m != nil {
		return m.BillToState
	}
	return ""
}

func (m *InvoiceData) GetBillToCountry() string {
	if m != nil {
		return m.BillToCountry
	}
	return ""
}

func (m *InvoiceData) GetBillToVatNumber() string {
	if m != nil {
		return m.BillToVatNumber
	}
	return ""
}

func (m *InvoiceData) GetBillToLocalTaxId() string {
	if m != nil {
		return m.BillToLocalTaxId
	}
	return ""
}

func (m *InvoiceData) GetRemitToCompanyName() string {
	if m != nil {
		return m.RemitToCompanyName
	}
	return ""
}

func (m *InvoiceData) GetRemitToContactPersonName() string {
	if m != nil {
		return m.RemitToContactPersonName
	}
	return ""
}

func (m *InvoiceData) GetRemitToStreet1() string {
	if m != nil {
		return m.RemitToStreet1
	}
	return ""
}

func (m *InvoiceData) GetRemitToStreet2() string {
	if m != nil {
		return m.RemitToStreet2
	}
	return ""
}

func (m *InvoiceData) GetRemitToCity() string {
	if m != nil {
		return m.RemitToCity
	}
	return ""
}

func (m *InvoiceData) GetRemitToZipcode() string {
	if m != nil {
		return m.RemitToZipcode
	}
	return ""
}

func (m *InvoiceData) GetRemitToState() string {
	if m != nil {
		return m.RemitToState
	}
	return ""
}

func (m *InvoiceData) GetRemitToCountry() string {
	if m != nil {
		return m.RemitToCountry
	}
	return ""
}

func (m *InvoiceData) GetRemitToVatNumber() string {
	if m != nil {
		return m.RemitToVatNumber
	}
	return ""
}

func (m *InvoiceData) GetRemitToLocalTaxId() string {
	if m != nil {
		return m.RemitToLocalTaxId
	}
	return ""
}

func (m *InvoiceData) GetRemitToTaxCountry() string {
	if m != nil {
		return m.RemitToTaxCountry
	}
	return ""
}

func (m *InvoiceData) GetShipToCompanyName() string {
	if m != nil {
		return m.ShipToCompanyName
	}
	return ""
}

func (m *InvoiceData) GetShipToContactPersonName() string {
	if m != nil {
		return m.ShipToContactPersonName
	}
	return ""
}

func (m *InvoiceData) GetShipToStreet1() string {
	if m != nil {
		return m.ShipToStreet1
	}
	return ""
}

func (m *InvoiceData) GetShipToStreet2() string {
	if m != nil {
		return m.ShipToStreet2
	}
	return ""
}

func (m *InvoiceData) GetShipToCity() string {
	if m != nil {
		return m.ShipToCity
	}
	return ""
}

func (m *InvoiceData) GetShipToZipcode() string {
	if m != nil {
		return m.ShipToZipcode
	}
	return ""
}

func (m *InvoiceData) GetShipToState() string {
	if m != nil {
		return m.ShipToState
	}
	return ""
}

func (m *InvoiceData) GetShipToCountry() string {
	if m != nil {
		return m.ShipToCountry
	}
	return ""
}

func (m *InvoiceData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *InvoiceData) GetGrossAmount() []byte {
	if m != nil {
		return m.GrossAmount
	}
	return nil
}

func (m *InvoiceData) GetNetAmount() []byte {
	if m != nil {
		return m.NetAmount
	}
	return nil
}

func (m *InvoiceData) GetTaxAmount() []byte {
	if m != nil {
		return m.TaxAmount
	}
	return nil
}

func (m *InvoiceData) GetTaxRate() []byte {
	if m != nil {
		return m.TaxRate
	}
	return nil
}

func (m *InvoiceData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *InvoiceData) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *InvoiceData) GetPayee() []byte {
	if m != nil {
		return m.Payee
	}
	return nil
}

func (m *InvoiceData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *InvoiceData) GetShippingTerms() string {
	if m != nil {
		return m.ShippingTerms
	}
	return ""
}

func (m *InvoiceData) GetRequesterEmail() string {
	if m != nil {
		return m.RequesterEmail
	}
	return ""
}

func (m *InvoiceData) GetRequesterName() string {
	if m != nil {
		return m.RequesterName
	}
	return ""
}

func (m *InvoiceData) GetDeliveryNumber() string {
	if m != nil {
		return m.DeliveryNumber
	}
	return ""
}

func (m *InvoiceData) GetCreditNoteInvoiceNumber() string {
	if m != nil {
		return m.CreditNoteInvoiceNumber
	}
	return ""
}

func (m *InvoiceData) GetCreditForInvoiceDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreditForInvoiceDate
	}
	return nil
}

func (m *InvoiceData) GetDateDue() *timestamp.Timestamp {
	if m != nil {
		return m.DateDue
	}
	return nil
}

func (m *InvoiceData) GetDatePaid() *timestamp.Timestamp {
	if m != nil {
		return m.DatePaid
	}
	return nil
}

func (m *InvoiceData) GetDateUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.DateUpdated
	}
	return nil
}

func (m *InvoiceData) GetDateCreated() *timestamp.Timestamp {
	if m != nil {
		return m.DateCreated
	}
	return nil
}

func (m *InvoiceData) GetAttachments() []*common.BinaryAttachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *InvoiceData) GetLineItems() []*LineItem {
	if m != nil {
		return m.LineItems
	}
	return nil
}

func (m *InvoiceData) GetPaymentDetails() []*common.PaymentDetails {
	if m != nil {
		return m.PaymentDetails
	}
	return nil
}

func (m *InvoiceData) GetTaxItems() []*TaxItem {
	if m != nil {
		return m.TaxItems
	}
	return nil
}

type TaxItem struct {
	ItemNumber           string   `protobuf:"bytes,1,opt,name=item_number,json=itemNumber,proto3" json:"item_number,omitempty"`
	InvoiceItemNumber    string   `protobuf:"bytes,2,opt,name=invoice_item_number,json=invoiceItemNumber,proto3" json:"invoice_item_number,omitempty"`
	TaxAmount            []byte   `protobuf:"bytes,3,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate              []byte   `protobuf:"bytes,4,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	TaxCode              []byte   `protobuf:"bytes,5,opt,name=tax_code,json=taxCode,proto3" json:"tax_code,omitempty"`
	TaxBaseAmount        []byte   `protobuf:"bytes,6,opt,name=tax_base_amount,json=taxBaseAmount,proto3" json:"tax_base_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaxItem) Reset()         { *m = TaxItem{} }
func (m *TaxItem) String() string { return proto.CompactTextString(m) }
func (*TaxItem) ProtoMessage()    {}
func (*TaxItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3e2b5ce0fcadd51, []int{1}
}

func (m *TaxItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaxItem.Unmarshal(m, b)
}
func (m *TaxItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaxItem.Marshal(b, m, deterministic)
}
func (m *TaxItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaxItem.Merge(m, src)
}
func (m *TaxItem) XXX_Size() int {
	return xxx_messageInfo_TaxItem.Size(m)
}
func (m *TaxItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TaxItem.DiscardUnknown(m)
}

var xxx_messageInfo_TaxItem proto.InternalMessageInfo

func (m *TaxItem) GetItemNumber() string {
	if m != nil {
		return m.ItemNumber
	}
	return ""
}

func (m *TaxItem) GetInvoiceItemNumber() string {
	if m != nil {
		return m.InvoiceItemNumber
	}
	return ""
}

func (m *TaxItem) GetTaxAmount() []byte {
	if m != nil {
		return m.TaxAmount
	}
	return nil
}

func (m *TaxItem) GetTaxRate() []byte {
	if m != nil {
		return m.TaxRate
	}
	return nil
}

func (m *TaxItem) GetTaxCode() []byte {
	if m != nil {
		return m.TaxCode
	}
	return nil
}

func (m *TaxItem) GetTaxBaseAmount() []byte {
	if m != nil {
		return m.TaxBaseAmount
	}
	return nil
}

type LineItem struct {
	ItemNumber    string `protobuf:"bytes,1,opt,name=item_number,json=itemNumber,proto3" json:"item_number,omitempty"`
	Description   string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SenderPartNo  string `protobuf:"bytes,3,opt,name=sender_part_no,json=senderPartNo,proto3" json:"sender_part_no,omitempty"`
	PricePerUnit  []byte `protobuf:"bytes,4,opt,name=price_per_unit,json=pricePerUnit,proto3" json:"price_per_unit,omitempty"`
	Quantity      []byte `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UnitOfMeasure string `protobuf:"bytes,6,opt,name=unit_of_measure,json=unitOfMeasure,proto3" json:"unit_of_measure,omitempty"`
	NetWeight     []byte `protobuf:"bytes,7,opt,name=net_weight,json=netWeight,proto3" json:"net_weight,omitempty"`
	TaxAmount     []byte `protobuf:"bytes,8,opt,name=tax_amount,json=taxAmount,proto3" json:"tax_amount,omitempty"`
	TaxRate       []byte `protobuf:"bytes,9,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	TaxCode       []byte `protobuf:"bytes,10,opt,name=tax_code,json=taxCode,proto3" json:"tax_code,omitempty"`
	//the total amount of the line item
	TotalAmount             []byte   `protobuf:"bytes,11,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	PurchaseOrderNumber     string   `protobuf:"bytes,12,opt,name=purchase_order_number,json=purchaseOrderNumber,proto3" json:"purchase_order_number,omitempty"`
	PurchaseOrderItemNumber string   `protobuf:"bytes,13,opt,name=purchase_order_item_number,json=purchaseOrderItemNumber,proto3" json:"purchase_order_item_number,omitempty"`
	DeliveryNoteNumber      string   `protobuf:"bytes,14,opt,name=delivery_note_number,json=deliveryNoteNumber,proto3" json:"delivery_note_number,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *LineItem) Reset()         { *m = LineItem{} }
func (m *LineItem) String() string { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()    {}
func (*LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3e2b5ce0fcadd51, []int{2}
}

func (m *LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineItem.Unmarshal(m, b)
}
func (m *LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineItem.Marshal(b, m, deterministic)
}
func (m *LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineItem.Merge(m, src)
}
func (m *LineItem) XXX_Size() int {
	return xxx_messageInfo_LineItem.Size(m)
}
func (m *LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_LineItem proto.InternalMessageInfo

func (m *LineItem) GetItemNumber() string {
	if m != nil {
		return m.ItemNumber
	}
	return ""
}

func (m *LineItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LineItem) GetSenderPartNo() string {
	if m != nil {
		return m.SenderPartNo
	}
	return ""
}

func (m *LineItem) GetPricePerUnit() []byte {
	if m != nil {
		return m.PricePerUnit
	}
	return nil
}

func (m *LineItem) GetQuantity() []byte {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *LineItem) GetUnitOfMeasure() string {
	if m != nil {
		return m.UnitOfMeasure
	}
	return ""
}

func (m *LineItem) GetNetWeight() []byte {
	if m != nil {
		return m.NetWeight
	}
	return nil
}

func (m *LineItem) GetTaxAmount() []byte {
	if m != nil {
		return m.TaxAmount
	}
	return nil
}

func (m *LineItem) GetTaxRate() []byte {
	if m != nil {
		return m.TaxRate
	}
	return nil
}

func (m *LineItem) GetTaxCode() []byte {
	if m != nil {
		return m.TaxCode
	}
	return nil
}

func (m *LineItem) GetTotalAmount() []byte {
	if m != nil {
		return m.TotalAmount
	}
	return nil
}

func (m *LineItem) GetPurchaseOrderNumber() string {
	if m != nil {
		return m.PurchaseOrderNumber
	}
	return ""
}

func (m *LineItem) GetPurchaseOrderItemNumber() string {
	if m != nil {
		return m.PurchaseOrderItemNumber
	}
	return ""
}

func (m *LineItem) GetDeliveryNoteNumber() string {
	if m != nil {
		return m.DeliveryNoteNumber
	}
	return ""
}

func init() {
	proto.RegisterType((*InvoiceData)(nil), "invoice.InvoiceData")
	proto.RegisterType((*TaxItem)(nil), "invoice.TaxItem")
	proto.RegisterType((*LineItem)(nil), "invoice.LineItem")
}

func init() { proto.RegisterFile("invoice/invoice.proto", fileDescriptor_b3e2b5ce0fcadd51) }

var fileDescriptor_b3e2b5ce0fcadd51 = []byte{
	// 1413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdb, 0x72, 0x13, 0x47,
	0x13, 0x2e, 0xff, 0x80, 0x6d, 0xb5, 0x64, 0xcb, 0x5e, 0xdb, 0x30, 0x08, 0xff, 0x41, 0x38, 0x1c,
	0x14, 0x08, 0x92, 0x11, 0x21, 0xa4, 0x38, 0x06, 0x9b, 0xa4, 0xca, 0x55, 0xc4, 0x28, 0x8a, 0x49,
	0xaa, 0xb8, 0xd9, 0x1a, 0xef, 0x8e, 0xed, 0xa9, 0xd2, 0xee, 0x2c, 0xb3, 0xb3, 0xc4, 0xca, 0x8b,
	0xe4, 0xd1, 0x72, 0x9b, 0x47, 0x49, 0xf5, 0x1c, 0x76, 0x67, 0x2d, 0x82, 0x73, 0x65, 0x4f, 0xf7,
	0xf7, 0x75, 0xcf, 0x7e, 0xad, 0xee, 0x69, 0xd8, 0xe0, 0xe9, 0x47, 0xc1, 0x23, 0x36, 0xb0, 0x7f,
	0xfb, 0x99, 0x14, 0x4a, 0x04, 0x0b, 0xf6, 0xd8, 0xd9, 0x8c, 0x44, 0x92, 0x88, 0x74, 0x10, 0x8b,
	0xa8, 0x48, 0x58, 0xaa, 0x42, 0x73, 0x36, 0xb0, 0xce, 0xf5, 0x63, 0x21, 0x8e, 0x27, 0x6c, 0xa0,
	0x4f, 0x87, 0xc5, 0xd1, 0x40, 0xf1, 0x84, 0xe5, 0x8a, 0x26, 0x99, 0x01, 0x6c, 0xfd, 0xbd, 0x01,
	0xcd, 0x3d, 0x13, 0xea, 0x35, 0x55, 0x34, 0xb8, 0x0c, 0xf3, 0x69, 0x91, 0x1c, 0x32, 0x49, 0xe6,
	0xba, 0x73, 0xbd, 0xc6, 0xd8, 0x9e, 0xd0, 0x9e, 0x2b, 0xaa, 0x8a, 0x9c, 0xfc, 0xcf, 0xd8, 0xcd,
	0x29, 0xb8, 0x0b, 0xab, 0x39, 0x4b, 0x63, 0x26, 0x43, 0x7b, 0xa1, 0x90, 0xc7, 0xe4, 0x82, 0x86,
	0xb4, 0x8d, 0xc3, 0x46, 0xdf, 0x8b, 0x83, 0x6d, 0x58, 0x97, 0x2c, 0xe2, 0x19, 0xc7, 0x6b, 0x7a,
	0xf0, 0x8b, 0x1a, 0x1e, 0x94, 0xbe, 0x8a, 0xd1, 0x87, 0x35, 0x1b, 0x3d, 0x12, 0x49, 0x46, 0xd3,
	0x69, 0x98, 0xd2, 0x84, 0x91, 0x4b, 0x9a, 0x60, 0x13, 0xef, 0x1a, 0xcf, 0x3e, 0x4d, 0x58, 0xf0,
	0x14, 0x3a, 0x25, 0x3e, 0x55, 0x34, 0x52, 0x61, 0xc6, 0x64, 0x2e, 0x52, 0x43, 0x9b, 0xd7, 0xb4,
	0x2b, 0x8e, 0xa6, 0x01, 0x23, 0xed, 0xd7, 0xe4, 0x5b, 0xb0, 0x6c, 0xc9, 0xb9, 0x92, 0x8c, 0xa9,
	0x07, 0x64, 0x41, 0x13, 0x96, 0x8c, 0xf5, 0x17, 0x63, 0x9c, 0x81, 0x0d, 0xc9, 0xe2, 0x2c, 0x6c,
	0x18, 0x5c, 0x87, 0xa6, 0xbb, 0x0a, 0x57, 0x53, 0xd2, 0xd0, 0x18, 0xb0, 0xb9, 0xb9, 0x9a, 0x7a,
	0x71, 0xfe, 0xe0, 0x59, 0x24, 0x62, 0x46, 0xc0, 0x8f, 0xf3, 0xde, 0x18, 0x83, 0x1b, 0xd0, 0x2a,
	0xd3, 0x51, 0xc5, 0x48, 0x53, 0x83, 0x9a, 0x2e, 0x19, 0x55, 0xfe, 0xc5, 0x23, 0x51, 0xa4, 0x4a,
	0x4e, 0x49, 0xcb, 0x8f, 0xb4, 0x6b, 0x8c, 0xc1, 0x00, 0xd6, 0x0f, 0xf9, 0x64, 0x12, 0x2a, 0x51,
	0x57, 0xf3, 0x9e, 0x51, 0x13, 0x7d, 0x07, 0xc2, 0x57, 0xf3, 0x19, 0x5c, 0xab, 0x08, 0xb3, 0x72,
	0x7e, 0x6d, 0xe4, 0x74, 0xbc, 0xb3, 0x72, 0xde, 0x86, 0xb6, 0x63, 0x3b, 0x3d, 0xdb, 0xe6, 0x5a,
	0x86, 0xe1, 0xf4, 0x9c, 0xc1, 0x0d, 0xc9, 0xca, 0x2c, 0x6e, 0x18, 0x74, 0xa1, 0x55, 0xde, 0x06,
	0x15, 0x5d, 0x35, 0x8a, 0xda, 0xf4, 0xa8, 0xa8, 0x17, 0xc9, 0x49, 0x1a, 0xf8, 0x91, 0x9c, 0xa4,
	0x5b, 0xb0, 0x54, 0x65, 0x44, 0x4d, 0x87, 0x46, 0x53, 0x97, 0x0f, 0x35, 0xf5, 0x62, 0x39, 0x51,
	0xd7, 0xfd, 0x58, 0x4e, 0xd4, 0x7b, 0x10, 0x38, 0xdc, 0x47, 0xaa, 0x42, 0xdb, 0x3b, 0x1b, 0xa6,
	0x01, 0x0c, 0xf4, 0x57, 0xaa, 0xf6, 0x4d, 0x13, 0xf5, 0xab, 0x0a, 0x4c, 0x44, 0x44, 0x27, 0xa1,
	0xa2, 0xa7, 0xd8, 0x00, 0xcf, 0x34, 0x7c, 0xc5, 0xc0, 0xdf, 0xa0, 0xe7, 0x80, 0x9e, 0xee, 0xc5,
	0xc1, 0x03, 0xd8, 0x90, 0x2c, 0xe1, 0x6a, 0xa6, 0x64, 0x57, 0x5c, 0xc7, 0x24, 0x5c, 0xd5, 0x6b,
	0xf6, 0x02, 0x36, 0x3d, 0xca, 0x6c, 0xd1, 0x88, 0x66, 0x92, 0x92, 0x79, 0xb6, 0x6a, 0x3d, 0x58,
	0x29, 0xf9, 0xae, 0x6c, 0x57, 0x35, 0x67, 0xd9, 0x72, 0x5c, 0xdd, 0x66, 0x91, 0x43, 0xd2, 0xf9,
	0x04, 0x72, 0x88, 0x7a, 0x57, 0x77, 0xc2, 0xd2, 0x5d, 0x33, 0x7a, 0xbb, 0x4b, 0x60, 0xed, 0xfc,
	0x68, 0xae, 0x78, 0x9b, 0xb5, 0x68, 0xae, 0x7a, 0x37, 0x61, 0xd9, 0xcb, 0x8b, 0xe5, 0xfb, 0x42,
	0xe3, 0x5a, 0x65, 0x56, 0xac, 0x9f, 0x1f, 0xcf, 0x15, 0xf0, 0x7a, 0x2d, 0x9e, 0xab, 0xe0, 0x7d,
	0x58, 0x2b, 0x91, 0x5e, 0x09, 0xbb, 0xa6, 0x26, 0x16, 0x5c, 0xd5, 0x70, 0xdb, 0xab, 0x49, 0xad,
	0x88, 0x37, 0x4c, 0x1b, 0x59, 0x82, 0x57, 0xc5, 0x01, 0x8e, 0x3d, 0xcb, 0x40, 0xac, 0xbb, 0xce,
	0x56, 0x8d, 0x70, 0x40, 0x4f, 0xbd, 0x46, 0xcd, 0x4f, 0x78, 0x36, 0x53, 0xf5, 0x2f, 0xed, 0xd8,
	0x3b, 0xe1, 0xd9, 0x4c, 0xa3, 0x56, 0x84, 0xd9, 0x9a, 0xdf, 0xb4, 0x73, 0xcf, 0xf2, 0x3e, 0xd1,
	0xa8, 0x8e, 0xed, 0x2a, 0x7e, 0xcb, 0xce, 0x0f, 0xcd, 0xf0, 0x1a, 0xb5, 0x8e, 0x1b, 0x92, 0xdb,
	0xb3, 0x38, 0xdd, 0xa8, 0xe5, 0x6d, 0xb0, 0xda, 0x77, 0xec, 0xe8, 0x33, 0xe9, 0x6d, 0xa3, 0x3a,
	0x84, 0xab, 0x75, 0xcf, 0x8f, 0xe4, 0x35, 0x6a, 0x95, 0x11, 0x2b, 0xfd, 0x95, 0x1d, 0x7e, 0x36,
	0x9f, 0x6d, 0xd4, 0xea, 0xdb, 0x8d, 0xb0, 0x77, 0xfd, 0x58, 0x4e, 0xd4, 0x0e, 0x2c, 0x46, 0x85,
	0x94, 0x2c, 0x8d, 0xa6, 0x64, 0x49, 0x03, 0xca, 0x33, 0xce, 0xd8, 0x63, 0x29, 0xf2, 0x3c, 0xa4,
	0x09, 0xa2, 0xc9, 0x72, 0x77, 0xae, 0xd7, 0x1a, 0x37, 0xb5, 0xed, 0x95, 0x36, 0x05, 0xff, 0x07,
	0x48, 0x99, 0x72, 0x80, 0xfb, 0x1a, 0xd0, 0x48, 0x99, 0xaa, 0xdc, 0x58, 0x5a, 0xeb, 0xee, 0x1b,
	0xb7, 0xa2, 0xa7, 0xd6, 0x7d, 0x15, 0x16, 0xd1, 0x2d, 0xf1, 0x1b, 0x06, 0xda, 0xb9, 0xa0, 0xe8,
	0xe9, 0x18, 0xef, 0xbf, 0x09, 0x8d, 0xf2, 0xe1, 0x23, 0x0f, 0x0c, 0xb1, 0x34, 0xe8, 0x67, 0x57,
	0x0f, 0x71, 0xb2, 0xa6, 0x5d, 0xf6, 0x14, 0xac, 0xc3, 0xa5, 0x8c, 0x4e, 0x19, 0x23, 0x0f, 0xb5,
	0xd9, 0x1c, 0x02, 0x02, 0x0b, 0xf8, 0xfa, 0x63, 0xa4, 0x6f, 0xf4, 0x27, 0xba, 0xa3, 0x7e, 0x22,
	0x4e, 0x78, 0x96, 0xf1, 0xf4, 0x38, 0x54, 0x4c, 0x26, 0x39, 0x79, 0x54, 0x89, 0x84, 0xd6, 0x03,
	0x34, 0x06, 0x77, 0xa0, 0x2d, 0xd9, 0x87, 0x82, 0xe5, 0x8a, 0xc9, 0x90, 0x25, 0x94, 0x4f, 0xc8,
	0xb7, 0xae, 0x69, 0xac, 0xf9, 0x07, 0xb4, 0x62, 0xbc, 0x0a, 0xa8, 0x7f, 0x64, 0x8f, 0x4d, 0xbc,
	0xd2, 0xaa, 0x7f, 0x5a, 0x77, 0xa0, 0x1d, 0xb3, 0x09, 0xff, 0xc8, 0xe4, 0xd4, 0xf5, 0xd5, 0x77,
	0x26, 0x9e, 0x33, 0xdb, 0xae, 0x7a, 0x0a, 0x9d, 0x48, 0xb2, 0x98, 0xab, 0x30, 0x15, 0x8a, 0x95,
	0xcb, 0x81, 0xe5, 0x3c, 0x31, 0x3f, 0x60, 0x83, 0xd8, 0x17, 0x8a, 0xd9, 0x0d, 0xc1, 0x92, 0x7f,
	0x06, 0xeb, 0x0a, 0x8f, 0x44, 0xb5, 0x87, 0xc4, 0x28, 0xf6, 0xd3, 0xee, 0x5c, 0xaf, 0x39, 0xec,
	0xf4, 0xcd, 0x1a, 0xd4, 0x77, 0x6b, 0x50, 0xff, 0xc0, 0xad, 0x41, 0xe3, 0x75, 0x43, 0xfd, 0x51,
	0xc8, 0x6a, 0x0d, 0x62, 0xc1, 0x23, 0x58, 0x44, 0x7e, 0x18, 0x17, 0x8c, 0x5c, 0x3e, 0x37, 0xc6,
	0x02, 0x62, 0x5f, 0x17, 0x2c, 0x78, 0x0c, 0x0d, 0x4d, 0xcb, 0x28, 0x8f, 0xc9, 0xf3, 0x73, 0x79,
	0x3a, 0xc7, 0x88, 0xf2, 0x38, 0x78, 0x0e, 0x2d, 0x4d, 0x2c, 0x32, 0xfc, 0x13, 0x93, 0x17, 0xe7,
	0x72, 0x9b, 0x08, 0x7c, 0x67, 0xe0, 0x25, 0x3d, 0x92, 0x4c, 0xd3, 0x5f, 0xfe, 0x37, 0xfa, 0xae,
	0x81, 0x07, 0x4f, 0xa0, 0x49, 0x95, 0xa2, 0xd1, 0x09, 0xfe, 0x56, 0x72, 0xf2, 0x7d, 0xf7, 0x42,
	0xaf, 0x39, 0x24, 0x7d, 0xbb, 0x49, 0xee, 0xf0, 0x94, 0xca, 0xe9, 0xab, 0x12, 0x30, 0xf6, 0xc1,
	0xc1, 0x36, 0xc0, 0x84, 0xa7, 0x2c, 0xe4, 0x8a, 0x25, 0x39, 0x79, 0xa5, 0xa9, 0xab, 0x7d, 0xb7,
	0xac, 0xbe, 0xe1, 0x29, 0xdb, 0x53, 0x2c, 0x19, 0x37, 0x26, 0xf6, 0xbf, 0x3c, 0x78, 0x09, 0xed,
	0x8c, 0x4e, 0xf5, 0xae, 0x1a, 0x33, 0x45, 0xf9, 0x24, 0x27, 0x3b, 0x9a, 0x76, 0xd9, 0x65, 0x1c,
	0x19, 0xf7, 0x6b, 0xe3, 0x1d, 0x2f, 0x67, 0xb5, 0x73, 0x70, 0x1f, 0x1a, 0x7a, 0xe6, 0xea, 0x8c,
	0xbb, 0x9a, 0xba, 0x52, 0x66, 0xc4, 0x99, 0x8b, 0x09, 0xb1, 0xe1, 0x74, 0xbe, 0xad, 0xbf, 0xe6,
	0x60, 0xc1, 0x5a, 0x71, 0x2b, 0x43, 0x5a, 0x58, 0xdb, 0x71, 0x01, 0x4d, 0xe5, 0x13, 0xbd, 0x56,
	0x6e, 0xa6, 0x1e, 0xd0, 0x2c, 0xbd, 0xab, 0xd6, 0xb5, 0x57, 0xe1, 0xeb, 0x8d, 0x7f, 0xe1, 0x73,
	0x8d, 0x7f, 0xb1, 0xde, 0xf8, 0xd6, 0xa5, 0xa7, 0xdf, 0xa5, 0xd2, 0xb5, 0x8b, 0x73, 0xef, 0x36,
	0xb4, 0xd1, 0x75, 0x48, 0x73, 0xe6, 0x22, 0xcf, 0x6b, 0xc4, 0x92, 0xa2, 0xa7, 0x3b, 0x34, 0x67,
	0x26, 0xfa, 0xd6, 0x9f, 0x17, 0x61, 0xd1, 0x29, 0x7c, 0xfe, 0xa7, 0x75, 0xa1, 0x19, 0xb3, 0x3c,
	0x92, 0x3c, 0x53, 0x5c, 0xa4, 0xf6, 0x93, 0x7c, 0x13, 0x3e, 0xad, 0x76, 0x91, 0xcc, 0xa8, 0xc4,
	0x56, 0xb4, 0x9b, 0xbc, 0xdd, 0x40, 0x47, 0x54, 0xaa, 0x7d, 0x81, 0xa8, 0x4c, 0xa2, 0x40, 0x19,
	0x93, 0x61, 0x91, 0x72, 0x65, 0xbf, 0xac, 0xa5, 0xad, 0x23, 0x26, 0xdf, 0xa5, 0x5c, 0xe1, 0xbc,
	0xfd, 0x50, 0xd0, 0x54, 0xe1, 0x0b, 0x60, 0x3e, 0xaf, 0x3c, 0xe3, 0xf7, 0x21, 0x2f, 0x14, 0x47,
	0x61, 0xc2, 0x68, 0x5e, 0x48, 0xb7, 0x9b, 0x2f, 0xa1, 0xf9, 0xed, 0xd1, 0x4f, 0xc6, 0xe8, 0x86,
	0xee, 0xef, 0x8c, 0x1f, 0x9f, 0x28, 0xbd, 0x8d, 0x9b, 0xa1, 0xfb, 0x9b, 0x36, 0x9c, 0xd1, 0x7e,
	0xf1, 0x73, 0xda, 0x37, 0xfe, 0x5d, 0x7b, 0xa8, 0x6b, 0x7f, 0x03, 0x5a, 0x4a, 0x28, 0x3a, 0x71,
	0x61, 0x9b, 0xe6, 0x2d, 0xd0, 0x36, 0x1b, 0x78, 0x08, 0x1b, 0x59, 0x21, 0xa3, 0x13, 0x2c, 0x8f,
	0x90, 0x28, 0x97, 0xd5, 0xdc, 0xac, 0xdd, 0x6b, 0xce, 0xf9, 0x16, 0x7d, 0xd5, 0x80, 0x3b, 0xc3,
	0xf1, 0x8b, 0x65, 0x1e, 0xa4, 0x2b, 0x35, 0xa2, 0xf7, 0x23, 0xdb, 0x86, 0xf5, 0x6a, 0x8c, 0xe2,
	0x7c, 0xb4, 0xb4, 0x65, 0xb3, 0x06, 0x96, 0xb3, 0x54, 0x28, 0x3b, 0x12, 0x77, 0x7a, 0xd0, 0x8c,
	0x44, 0xe2, 0x9a, 0x62, 0xa7, 0x65, 0x67, 0xdb, 0x08, 0x07, 0xc1, 0x68, 0xee, 0x7d, 0xc3, 0x3a,
	0xb2, 0xc3, 0xc3, 0x79, 0x3d, 0x1c, 0x1e, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x36, 0x5a, 0xa5,
	0x4e, 0x68, 0x0e, 0x00, 0x00,
}
