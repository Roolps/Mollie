package mollie

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestListPaymentMethods(t *testing.T) {
	envFile, _ := godotenv.Read(".env")
	mollieTest.Secret = envFile["MOLLIE_SECRET"]
	methods, err := mollieTest.ListPaymentMethods(&ListPaymentMethodsParameters{})
	if err != nil {
		t.Errorf("Error getting payment methods: %v", err)
	}
	log.Println(methods)
}

// func TestToQueryStringListPaymentMethodsParameters(t *testing.T) {

// }
