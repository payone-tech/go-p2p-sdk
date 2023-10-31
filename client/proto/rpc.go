package proto

import (
	"encoding/json"
)

// RpcRequest protocol message
type RpcRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params,omitempty"`
}

// RpcResponse protocol message
type RpcResponse struct {
	Result json.RawMessage `json:"result,omitempty"`
	Error  json.RawMessage `json:"error,omitempty"`
}

func (r *RpcResponse) ProtoError() *Error {
	if r.Error == nil {
		return nil
	}

	var protoError *Error
	_ = json.Unmarshal(r.Error, &protoError)
	return protoError
}
