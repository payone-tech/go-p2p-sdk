package proto

// OrderRequest protocol message
type OrderRequest struct {
	UUID       *string `json:"uuid"`
	ClientTxID *string `json:"client_tx_id"`
}

// OrderResponse protocol message
type OrderResponse struct {
	UUID         string           `json:"uuid"`
	Direction    string           `json:"direction"`
	Status       string           `json:"status"`
	CreatedTS    int64            `json:"created_ts"`
	ExpireTS     int64            `json:"expire_ts"`
	UpdatedTS    int64            `json:"updated_ts"`
	ClientTxID   string           `json:"client_tx_id"`
	ClientUserID *string          `json:"client_user_id,omitempty"`
	Sum          string           `json:"sum"`
	Pan          string           `json:"pan"`
	Name         *string          `json:"name,omitempty"`
	Bank         *BankResponse    `json:"bank,omitempty"`
	Card         *CardResponse    `json:"card,omitempty"`
	Currency     CurrencyResponse `json:"currency"`
}
