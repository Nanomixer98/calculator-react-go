package models

type BinaryOperationRequest struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type UnaryOperationRequest struct {
	Value *float64 `json:"value"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
