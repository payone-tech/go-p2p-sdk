package proto

import "encoding/json"

// Error protocol message
type Error struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (e *Error) ProtoValidation() []ValidateError {
	if e.Code != 401 {
		return nil
	}

	var protoValidateErrors []ValidateError
	_ = json.Unmarshal(e.Data, &protoValidateErrors)
	return protoValidateErrors
}

// ValidateError protocol message
type ValidateError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}
