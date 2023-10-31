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

// BankResponse protocol message
type BankResponse struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
}

// BanksResponse protocol message
type BanksResponse struct {
	Count int            `json:"count"`
	Banks []BankResponse `json:"banks"`
}

// CurrencyResponse protocol message
type CurrencyResponse struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Code  string `json:"code"`
}

// CardResponse protocol message
type CardResponse struct {
	UUID   string `json:"uuid"`
	Number string `json:"number"`
}

// CurrenciesResponse protocol message
type CurrenciesResponse struct {
	Count      int                `json:"count"`
	Currencies []CurrencyResponse `json:"currencies"`
}
