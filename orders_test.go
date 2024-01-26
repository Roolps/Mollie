package mollie

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateOrderWithAllRequiredData(t *testing.T) {
	envFile, _ := godotenv.Read(".env")
	mollieTest.Secret = envFile["MOLLIE_SECRET"]
	order, err := mollieTest.CreateOrder(&CreateOrderParameters{
		Amount: PaymentAmount{
			Currency: UnitedStatesDollar,
			Value:    "20.00",
		},
		OrderNumber: "value-4gb-26011446",
		Locale:      EnglishUnitedStates,
		Lines: []*OrderLine{
			{
				Name:     "Value Cloud 4GB FRA",
				Quantity: 1,
				UnitPrice: PaymentAmount{
					Currency: UnitedStatesDollar,
					Value:    "20.00",
				},
				TotalAmount: PaymentAmount{
					Currency: UnitedStatesDollar,
					Value:    "20.00",
				},
				VATRate: "0",
				VATAmount: PaymentAmount{
					Currency: UnitedStatesDollar,
					Value:    "0.00",
				},
			},
		},
		BillingAddress: &Address{
			GivenName:       "Yara",
			FamilyName:      "Durand",
			StreetAndNumber: "Example address 1",
			City:            "Example",
			PostalCode:      "0000",
			Country:         "BE",
			Email:           "someone@example.com",
		},
		Method: "creditcard",
		Payment: map[string]string{
			"cardToken": "tkn_wmHFsvtest",
		},
		RedirectURL: "https://clubnode.com/tracking/value-4gb-26011446",
	})
	if err != nil {
		t.Errorf("create payment failed: %v", err)
	}
	log.Println(order)
}
