package mollie

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

var mollieTest = &APIClient{
	Version: "v2",
}

func TestExtractErrorFromValidErrorData(t *testing.T) {
	raw := []byte(`
	{
		"status":401,
		"title":"Unauthorized Request",
		"detail":"Missing authentication, or failed to authenticate",
		"_links":{
			"documentation":{
				"href":"https://docs.mollie.com/overview/authentication",
				"type":"text/html"
			}
		}
	}
	`)
	err := extractError(raw)
	if err == nil {
		t.Error("should return error as error data is valid")
	}
}

func TestExtractErrorFromInvalidErrorData(t *testing.T) {
	raw := []byte(`
	{
		"resource": "payment",
		"id": "tr_PSj7b45bkj",
		"mode": "test"
	}`)
	err := extractError(raw)
	if err != nil {
		t.Error("should return no error as there is error data is invalid")
	}
}

func TestCreatePaymentRequestWithInvalidSecret(t *testing.T) {
	mollieTest.Secret = "abcdefghijklmnopqrstuvwxyz"
	_, err := mollieTest.CreatePayment(&CreatePaymentParameters{})
	if err == nil {
		t.Error("create payment didn't return an error when secret key invalid")
	}
	// result is [error code 401] Unauthorized Request: Could not authenticate request
	log.Println(err)
}

func TestCreatePaymentRequestMissingRequiredFields(t *testing.T) {
	envFile, _ := godotenv.Read(".env")
	mollieTest.Secret = envFile["MOLLIE_SECRET"]
	_, err := mollieTest.CreatePayment(&CreatePaymentParameters{})
	if err == nil {
		t.Error("create payment didn't return an error despite missing required fields")
	}
	// result is [error code 422] Unprocessable Entity: The amount contains an invalid currency
	log.Println(err)
}

func TestCreatePaymentRequestWithAllRequiredData(t *testing.T) {
	envFile, _ := godotenv.Read(".env")
	mollieTest.Secret = envFile["MOLLIE_SECRET"]
	_, err := mollieTest.CreatePayment(&CreatePaymentParameters{
		Amount: PaymentAmount{
			Currency: UnitedStatesDollar,
			Value:    "10.00",
		},
		Description: "Payment for invoice #20240001",
		RedirectURL: "https://websiteurl.com/tracking/order_number",
	})
	if err != nil {
		t.Errorf("create payment failed: %v", err)
	}
}
