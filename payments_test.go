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
	err := extractError(401, raw)
	if err == nil {
		t.Error("should return error as error data is valid")
	}
}

func TestExtractErrorFromUnknownErrorCode(t *testing.T) {
	raw := []byte(`
	{
		"status":620,
		"title":"Error title here",
		"detail":"Lorem ipsum dolor sit amet",
		"_links":{
			"documentation":{
				"href":"https://docs.mollie.com/overview/authentication",
				"type":"text/html"
			}
		}
	}
	`)
	err := extractError(620, raw)
	if err == nil {
		t.Error("should return error as error data is valid")
	}
	log.Println(err)
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
