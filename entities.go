package tenantapigowrapper

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	entityOrder = "order"
)

type APISingleEntity interface {
	getEntitySingleName() string
	getID() primitive.ObjectID
}

type APIListEntity interface {
	getEntitySingleName() string
	getEntityListName() string
}

type EntityCommon struct {
	ID                primitive.ObjectID `json:"_id"`
	BrandID           primitive.ObjectID `json:"brandId"`
	CreatedAt         time.Time          `json:"createdAt"`
	ModifiedAt        time.Time          `json:"modifiedAt"`
	StatusDescription []string           `json:"statusDescription"`
	Memo              string             `json:"memo"`
	ColourCode        string             `json:"colourCode"`
}

type TransactionCommon struct {
	EntityCommon

	Date         time.Time          `json:"date"`
	Ref          uint               `json:"ref"`
	Currency     string             `json:"currency"`
	Cancelled    bool               `json:"cancelled"`
	JournalEntry primitive.ObjectID `json:"journalEntry"`
}

type NonTransactionCommon struct {
	EntityCommon
	Inactive bool `json:"inactive"`
}

type InvoiceableCommon struct {
	InvoicedAt         time.Time `json:"invoicedAt"`
	PaymentDueAt       time.Time `json:"paymentDueAt"`
	OverdueDays        int       `json:"overdueDays"`
	OutstandingBalance int64     `json:"outstandingBalance"`
}

func (orders Orders) getEntitySingleName() string {
	return entityOrder
}

func (orders Orders) getEntityListName() string {
	return entityOrder
}

func (order Order) getEntitySingleName() string {
	return entityOrder
}

func (order Order) getID() primitive.ObjectID {
	return order.ID
}

type Order struct {
	TransactionCommon
	InvoiceableCommon

	Customer OrderCustomer `json:"customer"`
	//PaymentMethod   PaymentMethod      `json:"paymentMethod"`
	PaypalOrderID string `json:"paypalOrderId"`
	//Payments        CustomerPayments   `json:"payments"`
	ShippingMethod ShippingMethod `json:"shippingMethod"`
	PromoCode      string         `json:"promoCode"`
	//DiscountLines   []DiscountLine     `json:"discountLines"`
	Source          string             `json:"source"`
	SourceWebsiteID primitive.ObjectID `json:"sourceWebsiteId"`
	OrderLines      []OrderLine        `json:"orderLines"`
	ShippingAddress Address            `json:"shippingAddress"`
	BillingAddress  Address            `json:"billingAddress"`
	CustomerNotes   string             `json:"customerNotes"`
	DispatchNotes   string             `json:"dispatchNotes"`
	Totals          CartAndOrderTotals `json:"orderTotals"`
	//Fulfillments    StockMovements     `json:"fulfillments"`
	//Status          OrderStatus        `json:"status"`
	Stage          string `json:"stage"`
	ExternalRefIn  string `json:"externalRefIn"`
	ExternalRefOut string `json:"externalRefOut"`
	BrandedRef     string `json:"brandedRef" bson:"-"`
}

type Orders []Order

type OrderCustomer struct {
	BusinessDetails
	CustomerID   primitive.ObjectID `bson:"customerId" json:"customerId"`
	Email        string             `bson:"email" json:"email"`
	BillingEmail string             `bson:"billingEmail" json:"billingEmail"`
	Name         string             `bson:"name" json:"name"`
	PaymentTerms int                `bson:"paymentTerms" json:"paymentTerms"`
	Phone        string             `bson:"phone" json:"phone"`
}

type BusinessDetails struct {
	LegalName          string `json:"legalName"`
	LegalAddress       string `json:"legalAddress"`
	RegistrationNumber string `json:"registrationNumber"`
	TaxNumber          string `json:"taxNumber"`
}

type Address struct {
	Label               string `json:"label"`
	RecipientName       string `json:"recipientName"`
	Company             string `json:"companyName"`
	Line1               string `json:"line1"`
	Line2               string `json:"line2"`
	City                string `json:"city"`
	CountyProvinceState string `json:"countyProvinceState"`
	Postcode            string `json:"postcode"`
	Country             string `json:"country"`
	Phone               string `json:"phone"`
	DefaultBilling      bool   `json:"defaultBilling"`
	DefaultShipping     bool   `json:"defaultShipping"`
	CourierInstructions string `json:"courierInstructions"`
}

type ShippingMethod struct {
	TransactionCommon
	Name                string `json:"name"`
	ExternalServiceCode string `json:"externalServiceCode"`
}

type CartAndOrderTotals struct {
	UnitsAndWeight

	// Subtotals
	SubtotalExTax  int64 `json:"subtotalExTax"`
	SubtotalTax    int64 `json:"subtotalTax"`
	SubtotalIncTax int64 `json:"subtotalIncTax"`

	// Shipping
	ShippingExTax  int64 `json:"shippingExTax"`
	ShippingTax    int64 `json:"shippingTax"`
	ShippingIncTax int64 `json:"shippingIncTax"`

	// Discounts
	DiscountsExTax  int64 `json:"discountsExTax"`
	DiscountsTax    int64 `json:"discountsTax"`
	DiscountsIncTax int64 `json:"discountsIncTax"`

	// Payment surcharge
	PaymentMethodSurchargeExTax  int64 `json:"paymentMethodSurchargeExTax"`
	PaymentMethodSurchargeTax    int64 `json:"paymentMethodSurchargeTax"`
	PaymentMethodSurchargeIncTax int64 `json:"paymentMethodSurchargeIncTax"`

	// Totals
	TotalExTax  int64 `json:"totalExTax"`
	TotalTax    int64 `json:"totalTax"`
	TotalIncTax int64 `json:"totalIncTax"`

	HaveCalculatedPaymentMethodSurcharge bool `json:"haveCalculatedPaymentMethodSurcharge"`
}

type UnitsAndWeight struct {
	NoOfUnits int     `json:"noOfUnits"`
	NoOfItems int     `json:"noOfItems"`
	Weight    float64 `json:"weight"`
}

type OrderLine struct {
	ProductLogistics

	ProductID     primitive.ObjectID `json:"productId"`
	Sku           string             `json:"sku"`
	Name          string             `json:"name"`
	Memo          string             `json:"memo"`
	Price         int64              `json:"price"`
	WeightKg      float64            `json:"weightKg"`
	TaxCodeID     primitive.ObjectID `json:"taxCodeId"`
	TaxCode       string             `json:"taxCode"`
	TaxPercentage float64            `json:"taxPercentage"`
	TaxAmount     int64              `json:"taxAmount"`
	IncTaxPrice   int64              `json:"incTaxPrice"`
	Quantity      int                `json:"quantity"`

	LineExTaxAmount  int64   `json:"lineExTaxAmount"`
	LineTaxAmount    int64   `json:"lineTaxAmount"`
	LineIncTaxAmount int64   `json:"lineIncTaxAmount"`
	LineWeightKg     float64 `json:"lineWeightKg"`

	// for lot commitment
	LotID              primitive.ObjectID `json:"lotId"`
	QtyMoved           int                `json:"qtyDispatched"`
	QtyLeftToMove      int                `json:"qtyLeftToDispatch"`
	StockTreatment     `bson:"inline"`
	AlreadyTransformed bool               `json:"alreadyTransformed"`
	TransformedFrom    primitive.ObjectID `json:"transformedFrom"`
}

type ProductLogistics struct {
	UnitsPerCase    int `json:"unitsPerCase"`
	CasesPerLayer   int `json:"casesPerLayer"`
	LayersPerPallet int `json:"layersPerPallet"`
}

type StockTreatment struct {
	NotStocked    bool `json:"notStocked"`
	NotDispatched bool `json:"notDispatched"`
	IsComposite   bool `json:"isComposite"`
}
