package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryOperationRequest(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		jsonStr  string
	}{
		{
			name:    "positive numbers",
			a:       5.0,
			b:       3.0,
			jsonStr: `{"a":5,"b":3}`,
		},
		{
			name:    "negative numbers",
			a:       -5.0,
			b:       -3.0,
			jsonStr: `{"a":-5,"b":-3}`,
		},
		{
			name:    "decimal numbers",
			a:       1.5,
			b:       2.5,
			jsonStr: `{"a":1.5,"b":2.5}`,
		},
		{
			name:    "zero values",
			a:       0.0,
			b:       0.0,
			jsonStr: `{"a":0,"b":0}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := BinaryOperationRequest{
				A: &tt.a,
				B: &tt.b,
			}

			assert.Equal(t, tt.a, *req.A)
			assert.Equal(t, tt.b, *req.B)

			// Test JSON marshaling
			jsonData, err := json.Marshal(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.jsonStr, string(jsonData))
		})
	}
}

func TestBinaryOperationRequest_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name      string
		jsonStr   string
		expectedA float64
		expectedB float64
		wantErr   bool
	}{
		{
			name:      "valid json",
			jsonStr:   `{"a":5,"b":3}`,
			expectedA: 5.0,
			expectedB: 3.0,
			wantErr:   false,
		},
		{
			name:      "valid json with decimals",
			jsonStr:   `{"a":1.5,"b":2.5}`,
			expectedA: 1.5,
			expectedB: 2.5,
			wantErr:   false,
		},
		{
			name:      "valid json with negative",
			jsonStr:   `{"a":-5,"b":-3}`,
			expectedA: -5.0,
			expectedB: -3.0,
			wantErr:   false,
		},
		{
			name:    "invalid json",
			jsonStr: `{"a":5,"b":}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req BinaryOperationRequest
			err := json.Unmarshal([]byte(tt.jsonStr), &req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, req.A)
				assert.NotNil(t, req.B)
				assert.Equal(t, tt.expectedA, *req.A)
				assert.Equal(t, tt.expectedB, *req.B)
			}
		})
	}
}

func TestBinaryOperationRequest_NilPointers(t *testing.T) {
	req := BinaryOperationRequest{
		A: nil,
		B: nil,
	}

	assert.Nil(t, req.A)
	assert.Nil(t, req.B)

	// Test JSON marshaling with nil pointers
	jsonData, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Equal(t, `{"a":null,"b":null}`, string(jsonData))
}

func TestUnaryOperationRequest(t *testing.T) {
	tests := []struct {
		name         string
		value        float64
		jsonStr      string
	}{
		{
			name:         "positive number",
			value:        5.0,
			jsonStr:      `{"value":5}`,
		},
		{
			name:         "negative number",
			value:        -5.0,
			jsonStr:      `{"value":-5}`,
		},
		{
			name:         "decimal number",
			value:        3.14159,
			jsonStr:      `{"value":3.14159}`,
		},
		{
			name:         "zero",
			value:        0.0,
			jsonStr:      `{"value":0}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := UnaryOperationRequest{
				Value: &tt.value,
			}

			assert.Equal(t, tt.value, *req.Value)

			// Test JSON marshaling
			jsonData, err := json.Marshal(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.jsonStr, string(jsonData))
		})
	}
}

func TestUnaryOperationRequest_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name          string
		jsonStr       string
		expectedValue float64
		wantErr       bool
	}{
		{
			name:          "valid json",
			jsonStr:       `{"value":5}`,
			expectedValue: 5.0,
			wantErr:       false,
		},
		{
			name:          "valid json with decimal",
			jsonStr:       `{"value":3.14}`,
			expectedValue: 3.14,
			wantErr:       false,
		},
		{
			name:          "valid json with negative",
			jsonStr:       `{"value":-10}`,
			expectedValue: -10.0,
			wantErr:       false,
		},
		{
			name:    "invalid json",
			jsonStr: `{"value":}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req UnaryOperationRequest
			err := json.Unmarshal([]byte(tt.jsonStr), &req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, req.Value)
				assert.Equal(t, tt.expectedValue, *req.Value)
			}
		})
	}
}

func TestUnaryOperationRequest_NilPointer(t *testing.T) {
	req := UnaryOperationRequest{
		Value: nil,
	}

	assert.Nil(t, req.Value)

	// Test JSON marshaling with nil pointer
	jsonData, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Equal(t, `{"value":null}`, string(jsonData))
}

func TestCalculationResponse(t *testing.T) {
	tests := []struct {
		name       string
		result     float64
		jsonStr    string
	}{
		{
			name:       "positive result",
			result:     42.0,
			jsonStr:    `{"result":42}`,
		},
		{
			name:       "negative result",
			result:     -42.0,
			jsonStr:    `{"result":-42}`,
		},
		{
			name:       "decimal result",
			result:     3.14159,
			jsonStr:    `{"result":3.14159}`,
		},
		{
			name:       "zero result",
			result:     0.0,
			jsonStr:    `{"result":0}`,
		},
		{
			name:       "large number",
			result:     1e10,
			jsonStr:    `{"result":10000000000}`,
		},
		{
			name:       "small number",
			result:     1e-10,
			jsonStr:    `{"result":1e-10}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := CalculationResponse{
				Result: tt.result,
			}

			assert.Equal(t, tt.result, resp.Result)

			// Test JSON marshaling
			jsonData, err := json.Marshal(resp)
			assert.NoError(t, err)
			assert.Equal(t, tt.jsonStr, string(jsonData))
		})
	}
}

func TestCalculationResponse_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name           string
		jsonStr        string
		expectedResult float64
		wantErr        bool
	}{
		{
			name:           "valid json",
			jsonStr:        `{"result":42}`,
			expectedResult: 42.0,
			wantErr:        false,
		},
		{
			name:           "valid json with decimal",
			jsonStr:        `{"result":3.14}`,
			expectedResult: 3.14,
			wantErr:        false,
		},
		{
			name:           "valid json with negative",
			jsonStr:        `{"result":-10}`,
			expectedResult: -10.0,
			wantErr:        false,
		},
		{
			name:    "invalid json",
			jsonStr: `{"result":}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp CalculationResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &resp)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, resp.Result)
			}
		})
	}
}

func TestErrorResponse(t *testing.T) {
	tests := []struct {
		name        string
		errorMsg    string
		jsonStr     string
	}{
		{
			name:        "simple error",
			errorMsg:    "something went wrong",
			jsonStr:     `{"error":"something went wrong"}`,
		},
		{
			name:        "division by zero error",
			errorMsg:    "division by zero is not allowed",
			jsonStr:     `{"error":"division by zero is not allowed"}`,
		},
		{
			name:        "overflow error",
			errorMsg:    "result is out of representable range",
			jsonStr:     `{"error":"result is out of representable range"}`,
		},
		{
			name:        "empty error",
			errorMsg:    "",
			jsonStr:     `{"error":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := ErrorResponse{
				Error: tt.errorMsg,
			}

			assert.Equal(t, tt.errorMsg, resp.Error)

			// Test JSON marshaling
			jsonData, err := json.Marshal(resp)
			assert.NoError(t, err)
			assert.Equal(t, tt.jsonStr, string(jsonData))
		})
	}
}

func TestErrorResponse_JSONUnmarshal(t *testing.T) {
	tests := []struct {
		name            string
		jsonStr         string
		expectedError   string
		wantErr         bool
	}{
		{
			name:            "valid json",
			jsonStr:         `{"error":"something went wrong"}`,
			expectedError:   "something went wrong",
			wantErr:         false,
		},
		{
			name:            "valid json with special chars",
			jsonStr:         `{"error":"Error: 400 - Bad Request"}`,
			expectedError:   "Error: 400 - Bad Request",
			wantErr:         false,
		},
		{
			name:            "invalid json",
			jsonStr:         `{"error":}`,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp ErrorResponse
			err := json.Unmarshal([]byte(tt.jsonStr), &resp)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedError, resp.Error)
			}
		})
	}
}

func TestAllModels_EmptyStruct(t *testing.T) {
	// Test empty structs can be created
	binReq := BinaryOperationRequest{}
	unaryReq := UnaryOperationRequest{}
	calcResp := CalculationResponse{}
	errResp := ErrorResponse{}

	assert.Nil(t, binReq.A)
	assert.Nil(t, binReq.B)
	assert.Nil(t, unaryReq.Value)
	assert.Equal(t, 0.0, calcResp.Result)
	assert.Equal(t, "", errResp.Error)
}
