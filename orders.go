package mollie

import (
	"encoding/json"
	"time"
)

// https://docs.mollie.com/reference/v2/orders-api/create-order

type CreateOrderParameters struct {
	Amount          PaymentAmount     `json:"amount"`
	OrderNumber     string            `json:"orderNumber"`
	Lines           []*OrderLine      `json:"lines"`
	BillingAddress  *Address          `json:"billingAddress"`
	ShippingAddress *Address          `json:"shippingAddress,omitempty"`
	RedirectURL     string            `json:"redirectUrl"`
	CancelURL       string            `json:"cancelUrl,omitempty"`
	Locale          Locale            `json:"locale"`
	Method          string            `json:"method,omitempty"`
	Payment         map[string]string `json:"payment,omitempty"`
}

type OrderLine struct {
	Type           ProductType    `json:"type,omitempty"`
	Name           string         `json:"name"`
	Quantity       int            `json:"quantity"`
	UnitPrice      PaymentAmount  `json:"unitPrice"`
	DiscountAmount *PaymentAmount `json:"discountAmount,omitempty"`
	TotalAmount    PaymentAmount  `json:"totalAmount"`
	VATRate        string         `json:"vatRate"`
	VATAmount      PaymentAmount  `json:"vatAmount"`
}

type ProductType string

const (
	ProductTypePhysical    ProductType = "physical"
	ProductTypeDiscount    ProductType = "discount"
	ProductTypeDigital     ProductType = "digital"
	ProductTypeShippingFee ProductType = "shipping_fee"
	ProductTypeStoreCredit ProductType = "store_credit"
	ProductTypeGiftCard    ProductType = "gift_card"
	ProductTypeSurcharge   ProductType = "surcharge"
)

type Address struct {
	OrganizationName string `json:"organizationName,omitempty"`
	Title            string `json:"title,omitempty"`
	GivenName        string `json:"givenName"`
	FamilyName       string `json:"familyName"`
	StreetAndNumber  string `json:"streetAndNumber"`
	StreetAdditional string `json:"streetAdditional,omitempty"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	Region           string `json:"region,omitempty"`
	Country          string `json:"country"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
}

type Order struct {
	Resource        string          `json:"resource"`
	ID              string          `json:"id"`
	ProfileId       string          `json:"profileId"`
	Method          string          `json:"method"`
	Amount          PaymentAmount   `json:"amount"`
	Status          string          `json:"status"`
	IsCancelable    bool            `json:"isCancelable"`
	Metadata        map[string]any  `json:"metadata"`
	CreatedAt       time.Time       `json:"createdAt"`
	ExpiresAt       time.Time       `json:"expiresAt"`
	Mode            string          `json:"mode"`
	Locale          Locale          `json:"locale"`
	BillingAddress  Address         `json:"billingAddress"`
	ConsumerDOB     string          `json:"consumerDateOfBirth"`
	OrderNumber     string          `json:"orderNumber"`
	ShippingAddress Address         `json:"shippingAddress"`
	RedirectURL     string          `json:"redirectUrl"`
	WehookURL       string          `json:"webhookUrl"`
	Lines           []OrderLine     `json:"lines"`
	Links           map[string]Link `json:"_links"`
}

func (c *APIClient) CreateOrder(param *CreateOrderParameters) (*Order, error) {
	raw, _ := json.Marshal(param)
	raw, err := c.request("orders", "POST", raw)
	if err != nil {
		return nil, err
	}
	o := &Order{}
	json.Unmarshal(raw, o)
	return o, nil
}
