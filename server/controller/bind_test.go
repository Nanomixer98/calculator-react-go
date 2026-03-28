package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator-server/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBindBinaryRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		requestBody   interface{}
		expectedError string
		expectedA     float64
		expectedB     float64
		expectSuccess bool
	}{
		{
			name:          "valid request",
			requestBody:   map[string]interface{}{"a": 5.0, "b": 3.0},
			expectedError: "",
			expectedA:     5.0,
			expectedB:     3.0,
			expectSuccess: true,
		},
		{
			name:          "missing field a",
			requestBody:   map[string]interface{}{"b": 3.0},
			expectedError: "both fields 'a' and 'b' are required",
			expectSuccess: false,
		},
		{
			name:          "missing field b",
			requestBody:   map[string]interface{}{"a": 5.0},
			expectedError: "both fields 'a' and 'b' are required",
			expectSuccess: false,
		},
		{
			name:          "both fields missing",
			requestBody:   map[string]interface{}{},
			expectedError: "both fields 'a' and 'b' are required",
			expectSuccess: false,
		},
		{
			name:          "invalid JSON",
			requestBody:   "invalid json",
			expectedError: "invalid JSON body",
			expectSuccess: false,
		},
		{
			name:          "negative numbers",
			requestBody:   map[string]interface{}{"a": -5.0, "b": -3.0},
			expectedError: "",
			expectedA:     -5.0,
			expectedB:     -3.0,
			expectSuccess: true,
		},
		{
			name:          "zero values",
			requestBody:   map[string]interface{}{"a": 0.0, "b": 0.0},
			expectedError: "",
			expectedA:     0.0,
			expectedB:     0.0,
			expectSuccess: true,
		},
		{
			name:          "decimal values",
			requestBody:   map[string]interface{}{"a": 1.5, "b": 2.5},
			expectedError: "",
			expectedA:     1.5,
			expectedB:     2.5,
			expectSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.POST("/test", func(c *gin.Context) {
				req, err := BindBinaryRequest(c)
				if tt.expectSuccess {
					assert.NoError(t, err)
					assert.NotNil(t, req)
					assert.Equal(t, tt.expectedA, *req.A)
					assert.Equal(t, tt.expectedB, *req.B)
				} else {
					assert.Error(t, err)
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			})

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)
		})
	}
}

func TestBindBinaryRequest_NilPointerFields(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/test", func(c *gin.Context) {
		req, err := BindBinaryRequest(c)
		assert.Error(t, err)
		assert.Nil(t, req)
		assert.Contains(t, err.Error(), "both fields 'a' and 'b' are required")
	})

	// Create a request with explicit null values
	body := `{"a": null, "b": null}`
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
}

func TestBindUnaryRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		requestBody   interface{}
		expectedError string
		expectedValue float64
		expectSuccess bool
	}{
		{
			name:          "valid request",
			requestBody:   map[string]interface{}{"value": 5.0},
			expectedError: "",
			expectedValue: 5.0,
			expectSuccess: true,
		},
		{
			name:          "missing value field",
			requestBody:   map[string]interface{}{},
			expectedError: "field 'value' is required",
			expectSuccess: false,
		},
		{
			name:          "invalid JSON",
			requestBody:   "invalid json",
			expectedError: "invalid JSON body",
			expectSuccess: false,
		},
		{
			name:          "negative number",
			requestBody:   map[string]interface{}{"value": -5.0},
			expectedError: "",
			expectedValue: -5.0,
			expectSuccess: true,
		},
		{
			name:          "zero value",
			requestBody:   map[string]interface{}{"value": 0.0},
			expectedError: "",
			expectedValue: 0.0,
			expectSuccess: true,
		},
		{
			name:          "decimal value",
			requestBody:   map[string]interface{}{"value": 3.14159},
			expectedError: "",
			expectedValue: 3.14159,
			expectSuccess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.POST("/test", func(c *gin.Context) {
				req, err := BindUnaryRequest(c)
				if tt.expectSuccess {
					assert.NoError(t, err)
					assert.NotNil(t, req)
					assert.Equal(t, tt.expectedValue, *req.Value)
				} else {
					assert.Error(t, err)
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			})

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)
		})
	}
}

func TestBindUnaryRequest_NilPointerValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/test", func(c *gin.Context) {
		req, err := BindUnaryRequest(c)
		assert.Error(t, err)
		assert.Nil(t, req)
		assert.Contains(t, err.Error(), "field 'value' is required")
	})

	// Create a request with explicit null value
	body := `{"value": null}`
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
}

func TestBindBinaryRequest_ExtraFields(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/test", func(c *gin.Context) {
		req, err := BindBinaryRequest(c)
		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, 5.0, *req.A)
		assert.Equal(t, 3.0, *req.B)
	})

	// Extra fields should be ignored
	body := `{"a": 5.0, "b": 3.0, "extra": "ignored"}`
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
}

func TestBindUnaryRequest_ExtraFields(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/test", func(c *gin.Context) {
		req, err := BindUnaryRequest(c)
		assert.NoError(t, err)
		assert.NotNil(t, req)
		assert.Equal(t, 5.0, *req.Value)
	})

	// Extra fields should be ignored
	body := `{"value": 5.0, "extra": "ignored"}`
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
}

func TestBinaryOperationRequest_Model(t *testing.T) {
	a := 5.0
	b := 3.0
	req := models.BinaryOperationRequest{
		A: &a,
		B: &b,
	}
	assert.Equal(t, a, *req.A)
	assert.Equal(t, b, *req.B)
}

func TestUnaryOperationRequest_Model(t *testing.T) {
	value := 5.0
	req := models.UnaryOperationRequest{
		Value: &value,
	}
	assert.Equal(t, value, *req.Value)
}
