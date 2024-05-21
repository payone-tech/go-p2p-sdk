package proto

// CreateDepositRequest protocol message
type CreateDepositRequest struct {
	ClientTxID        string  `json:"client_tx_id"`
	ClientUserID      *string `json:"client_user_id"`
	Sum               string  `json:"sum"`
	Pan               *string `json:"pan"`
	Name              *string `json:"name"`
	PaymentMethodUUID *string `json:"payment_method_uuid"`
	CurrencyUUID      string  `json:"currency_uuid"`
}

// CreateDepositResponse model.
type CreateDepositResponse struct {
	UUID          string                 `json:"uuid"`
	Status        string                 `json:"status"`
	CreatedTS     int64                  `json:"created_ts"`
	ExpireTS      int64                  `json:"expire_ts"`
	UpdatedTS     int64                  `json:"updated_ts"`
	ClientTxID    string                 `json:"client_tx_id"`
	ClientUserID  *string                `json:"client_user_id,omitempty"`
	Sum           string                 `json:"sum"`
	Pan           *string                `json:"pan,omitempty"`
	Name          *string                `json:"name,omitempty"`
	PaymentMethod *PaymentMethodResponse `json:"payment_method,omitempty"`
	Currency      CurrencyResponse       `json:"currency"`
}

// CancelDepositRequest model.
type CancelDepositRequest struct {
	UUID       *string `json:"uuid"`
	ClientTxID *string `json:"client_tx_id"`
}
