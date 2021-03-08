package tenantapigowrapper

import (
	"time"

	"github.com/dogpakk/lib/financial"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	entityOrder = "order"
)

type APISingleEntity interface {
	getEntitySingleName() string
}

type APIListEntity interface {
	getEntitySingleName() string
	getEntityListName() string
}

type EntityCommon struct {
	ID                primitive.ObjectID `json:"_id"`
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

type InvoiceableCommon struct {
	InvoicedAt         time.Time       `json:"invoicedAt"`
	PaymentDueAt       time.Time       `json:"paymentDueAt"`
	OverdueDays        int             `json:"overdueDays"`
	OutstandingBalance financial.Cents `json:"outstandingBalance"`
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

type Order struct {
	TransactionCommon
	InvoiceableCommon

	Customer OrderCustomer `json:"customer"`
	//PaymentMethod   PaymentMethod      `json:"paymentMethod"`
	PaypalOrderID string `json:"paypalOrderId"`
	//Payments        CustomerPayments   `json:"payments"`
	//ShippingMethod  ShippingMethod     `json:"shippingMethod"`
	PromoCode string `json:"promoCode"`
	//DiscountLines   []DiscountLine     `json:"discountLines"`
	Source          string             `json:"source"`
	SourceWebsiteID primitive.ObjectID `json:"sourceWebsiteId"`
	//OrderLines      OrderLines         `json:"orderLines"`
	ShippingAddress Address            `json:"shippingAddress"`
	BillingAddress  Address            `json:"billingAddress"`
	CustomerNotes   string             `json:"customerNotes"`
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

type CartAndOrderTotals struct {
	UnitsAndWeight

	// Subtotals
	SubtotalExTax  financial.Cents `json:"subtotalExTax"`
	SubtotalTax    financial.Cents `json:"subtotalTax"`
	SubtotalIncTax financial.Cents `json:"subtotalIncTax"`

	// Shipping
	ShippingExTax  financial.Cents `json:"shippingExTax"`
	ShippingTax    financial.Cents `json:"shippingTax"`
	ShippingIncTax financial.Cents `json:"shippingIncTax"`

	// Discounts
	DiscountsExTax  financial.Cents `json:"discountsExTax"`
	DiscountsTax    financial.Cents `json:"discountsTax"`
	DiscountsIncTax financial.Cents `json:"discountsIncTax"`

	// Payment surcharge
	PaymentMethodSurchargeExTax  financial.Cents `json:"paymentMethodSurchargeExTax"`
	PaymentMethodSurchargeTax    financial.Cents `json:"paymentMethodSurchargeTax"`
	PaymentMethodSurchargeIncTax financial.Cents `json:"paymentMethodSurchargeIncTax"`

	// Totals
	TotalExTax  financial.Cents `json:"totalExTax"`
	TotalTax    financial.Cents `json:"totalTax"`
	TotalIncTax financial.Cents `json:"totalIncTax"`

	HaveCalculatedPaymentMethodSurcharge bool `json:"haveCalculatedPaymentMethodSurcharge"`
}

type UnitsAndWeight struct {
	NoOfUnits int     `json:"noOfUnits"`
	NoOfItems int     `json:"noOfItems"`
	Weight    float64 `json:"weight"`
}
