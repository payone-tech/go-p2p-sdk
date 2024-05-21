package proto

// Status protocol message
type Status struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// Event protocol message
type Event struct {
	Comment []byte
	ID      []byte
	Event   []byte
	Data    []byte
	Retry   []byte
}

// PaymentMethodResponse protocol message
type PaymentMethodResponse struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
}

// PaymentMethodsResponse protocol message
type PaymentMethodsResponse struct {
	Count          int                     `json:"count"`
	PaymentMethods []PaymentMethodResponse `json:"payment_methods"`
}

// CurrencyResponse protocol message
type CurrencyResponse struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Code  string `json:"code"`
}

// CardResponse protocol message
type CardResponse struct {
	UUID   string  `json:"uuid"`
	Number string  `json:"number"`
	Name   *string `json:"name"`
}

// CurrenciesResponse protocol message
type CurrenciesResponse struct {
	Count      int                `json:"count"`
	Currencies []CurrencyResponse `json:"currencies"`
}
