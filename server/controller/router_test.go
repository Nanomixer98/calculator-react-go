package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator-server/core/app"
	"calculator-server/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	assert.NotNil(t, router)
	assert.NotNil(t, router.ctrl)
	assert.Equal(t, ctrl, router.ctrl)
}

func TestRouter_RegisterRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	engine := gin.New()
	router.RegisterRoutes(engine)

	// Test that all API routes are registered
	tests := []struct {
		name           string
		method         string
		path           string
		requestBody    interface{}
		expectedStatus int
		expectedResult float64
	}{
		{
			name:           "add endpoint",
			method:         http.MethodPost,
			path:           "/api/add",
			requestBody:    map[string]interface{}{"a": 5.0, "b": 3.0},
			expectedStatus: http.StatusOK,
			expectedResult: 8.0,
		},
		{
			name:           "subtract endpoint",
			method:         http.MethodPost,
			path:           "/api/subtract",
			requestBody:    map[string]interface{}{"a": 10.0, "b": 3.0},
			expectedStatus: http.StatusOK,
			expectedResult: 7.0,
		},
		{
			name:           "multiply endpoint",
			method:         http.MethodPost,
			path:           "/api/multiply",
			requestBody:    map[string]interface{}{"a": 5.0, "b": 3.0},
			expectedStatus: http.StatusOK,
			expectedResult: 15.0,
		},
		{
			name:           "divide endpoint",
			method:         http.MethodPost,
			path:           "/api/divide",
			requestBody:    map[string]interface{}{"a": 15.0, "b": 3.0},
			expectedStatus: http.StatusOK,
			expectedResult: 5.0,
		},
		{
			name:           "negate endpoint",
			method:         http.MethodPost,
			path:           "/api/negate",
			requestBody:    map[string]interface{}{"value": 5.0},
			expectedStatus: http.StatusOK,
			expectedResult: -5.0,
		},
		{
			name:           "percentage endpoint",
			method:         http.MethodPost,
			path:           "/api/percentage",
			requestBody:    map[string]interface{}{"value": 50.0},
			expectedStatus: http.StatusOK,
			expectedResult: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(tt.method, tt.path, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			engine.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			var response models.CalculationResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, response.Result)
		})
	}
}

func TestRouter_DivideByZero(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	engine := gin.New()
	router.RegisterRoutes(engine)

	body, _ := json.Marshal(map[string]interface{}{"a": 10.0, "b": 0.0})
	req := httptest.NewRequest(http.MethodPost, "/api/divide", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response models.ErrorResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "division by zero is not allowed", response.Error)
}

func TestRouter_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	engine := gin.New()
	router.RegisterRoutes(engine)

	body := []byte(`invalid json`)
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response models.ErrorResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response.Error, "invalid JSON")
}

func TestRouter_MissingFields(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	engine := gin.New()
	router.RegisterRoutes(engine)

	body, _ := json.Marshal(map[string]interface{}{"a": 5.0})
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response models.ErrorResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response.Error, "both fields")
}

func TestRouter_MissingValue(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	engine := gin.New()
	router.RegisterRoutes(engine)

	body, _ := json.Marshal(map[string]interface{}{})
	req := httptest.NewRequest(http.MethodPost, "/api/negate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response models.ErrorResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response.Error, "field 'value' is required")
}

func TestRouter_GroupPrefix(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	engine := gin.New()
	router.RegisterRoutes(engine)

	// Test that endpoints without /api prefix don't work
	body, _ := json.Marshal(map[string]interface{}{"a": 5.0, "b": 3.0})
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestRouter_Overflow(t *testing.T) {
	gin.SetMode(gin.TestMode)

	calcApp := app.NewCalculatorApp()
	ctrl := NewCalculatorController(calcApp)
	router := NewRouter(ctrl)

	engine := gin.New()
	router.RegisterRoutes(engine)

	// Test overflow in add
	body, _ := json.Marshal(map[string]interface{}{"a": 1e308, "b": 1e308})
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response models.ErrorResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "result is out of representable range", response.Error)
}
