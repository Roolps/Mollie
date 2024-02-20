package mollie

import (
	"encoding/json"
	"fmt"
)

// Finish later
type ListPaymentMethodsParameters struct {
	SequenceType string `json:"sequenceType"`
	Locale       Locale `json:"locale"`
}

type PaymentMethod struct {
	ID            string                 `json:"id"`
	Resource      string                 `json:"resource"`
	Description   string                 `json:"description"`
	MinimumAmount PaymentAmount          `json:"minimumAmount"`
	MaximumAmount PaymentAmount          `json:"maximumAmount"`
	Image         PaymentMethodImage     `json:"image"`
	Status        string                 `json:"status"`
	Pricing       []PaymentMethodPricing `json:"pricing"`
	Links         map[string]Link        `json:"_links"`
}

type PaymentMethodImage struct {
	Size1X string `json:"size1x"`
	Size2X string `json:"size2x"`
	SVG    string `json:"svg"`
}

type PaymentMethodPricing struct {
	Description string        `json:"description"`
	Fixed       PaymentAmount `json:"fixed"`
	Variable    string        `json:"variable"`
}

// PARAMETERS DON'T CURRENTLY WORK
func (c *APIClient) ListPaymentMethods(param *ListPaymentMethodsParameters) ([]*PaymentMethod, error) {
	raw, err := c.request("methods", "GET", nil)
	if err != nil {
		return nil, err
	}
	result := ListResponse{}
	if err = json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body: %v", err)
	}
	methods := []*PaymentMethod{}
	if err = json.Unmarshal(result.Embedded["methods"], &methods); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json body: %v", err)
	}
	return methods, nil
}
