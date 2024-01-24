package mollie

import (
	"encoding/json"
	"time"
)

type CreatePayment struct {
	// Required create payment fields (amount, description, redirectURL)
	Amount      PaymentAmount `json:"amount"`
	Description string        `json:"description"`
	RedirectURL string        `json:"redirectUrl"`

	// Optional fields
	CancelURL                       string   `json:"cancelUrl,omitempty"`
	WebhookURL                      string   `json:"webhookUrl,omitempty"`
	Locale                          Locale   `json:"locale,omitempty"`
	Method                          []string `json:"method,omitempty"`
	RestrictPaymentMethodsToCountry string   `json:"restrictPaymentMethodsToCountry,omitempty"`
	MetaData                        any      `json:"metadata,omitempty"`
}

type PaymentAmount struct {
	// Currency must follow ISO 4217 schema
	Currency Currency `json:"currency"`

	// Value must be to two decimal places
	Value string `json:"value"`
}

type Payment struct {
	ID           string          `json:"id"`
	Resource     string          `json:"resource"`
	Mode         string          `json:"mode"`
	CreatedAt    time.Time       `json:"createdAt"`
	Amount       PaymentAmount   `json:"amount"`
	Description  string          `json:"description"`
	Method       json.RawMessage `json:"method"`
	MetaData     json.RawMessage `json:"metadata"`
	Status       string          `json:"status"`
	IsCancelable bool            `json:"isCancelable"`
	ExpiresAt    time.Time       `json:"expiresAt"`
	Details      string          `json:"details"`
	ProfileID    string          `json:"profileId"`
	SequenceType string          `json:"sequenceType"`
	RedirectURL  string          `json:"redirectUrl"`
	WebhookURL   string          `json:"webhookUrl"`
	Links        []Link          `json:"_links"`
}

type Currency string

// ISO currency values
const (
	UnitedArabEmiratesDirham Currency = "AED"
	AustralianDollar         Currency = "AUD"
	BulgarianLev             Currency = "BGN"
	BrazilianReal            Currency = "BRL"
	CanadianDollar           Currency = "CAD"
	SwissFranc               Currency = "CHF"
	CzechKoruna              Currency = "CZK"
	DanishKrone              Currency = "DKK"
	Euro                     Currency = "EUR"
	BritishPound             Currency = "GBP"
	HongKongDollar           Currency = "HKD"
	HungarianForint          Currency = "HUF"
	IsraeliNewShekel         Currency = "ILS"
	IcelandicKrona           Currency = "ISK"
	JapaneseYen              Currency = "JPY"
	MexicanPeso              Currency = "MXN"
	MalaysianRinggit         Currency = "MYR"
	NorwegianKrone           Currency = "NOK"
	NewZealandDollar         Currency = "NZD"
	PhilippinePiso           Currency = "PHP"
	PolishZloty              Currency = "PLN"
	RomanianLeu              Currency = "RON"
	RussianRuble             Currency = "RUB"
	SwedishKrona             Currency = "SEK"
	SingaporeDollar          Currency = "SGD"
	ThaiBaht                 Currency = "THB"
	NewTaiwanDollar          Currency = "TWD"
	UnitedStatesDollar       Currency = "USD"
	SouthAfricanRand         Currency = "ZAR"
)

type Locale string

// Language/locale values
const (
	EnglishUnitedStates  Locale = "en_US"
	EnglishUnitedKingdom Locale = "en_GB"
	Dutch                Locale = "nl_NL"
	DutchBelgium         Locale = "nl_BE"
	French               Locale = "fr_FR"
	FrenchBelgium        Locale = "fr_BE"
	German               Locale = "de_DE"
	GermanAustria        Locale = "de_AT"
	GermanSwitzerland    Locale = "de_CH"
	Spanish              Locale = "es_ES"
	Catalan              Locale = "ca_ES"
	Portuguese           Locale = "pt_PT"
	Italian              Locale = "it_IT"
	Norwegian            Locale = "nb_NO"
	Swedish              Locale = "sv_SE"
	Finnish              Locale = "fi_FI"
	Danish               Locale = "da_DK"
	Icelandic            Locale = "is_IS"
	Hungarian            Locale = "hu_HU"
	Polish               Locale = "pl_PL"
	Latvian              Locale = "lv_LV"
	Lithuanian           Locale = "lt_LT"
)

func (c *APIClient) CreatePayment(param *CreatePayment) (*Payment, error) {
	raw, _ := json.Marshal(param)
	raw, err := c.request("payments", "POST", raw)
	if err != nil {
		return nil, err
	}
	p := &Payment{}
	json.Unmarshal(raw, p)
	return p, nil
}
