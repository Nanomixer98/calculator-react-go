package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"calculator-server/core/app"
	"calculator-server/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "division by zero error",
			err:            app.ErrDivisionByZero,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "division by zero is not allowed",
		},
		{
			name:           "overflow error",
			err:            app.ErrOverflow,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "result is out of representable range",
		},
		{
			name:           "unknown error",
			err:            errors.New("some random error"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "unexpected error",
		},
		{
			name:           "wrapped division by zero error",
			err:            errors.New("wrapped: " + app.ErrDivisionByZero.Error()),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "unexpected error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/test", func(c *gin.Context) {
				HandleError(c, tt.err)
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			var response models.ErrorResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedError, response.Error)
		})
	}
}

func TestHandleBindError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "invalid JSON error",
			err:            errors.New("invalid JSON body"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid JSON body",
		},
		{
			name:           "missing fields error",
			err:            errors.New("both fields 'a' and 'b' are required"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  "both fields 'a' and 'b' are required",
		},
		{
			name:           "NaN values error",
			err:            errors.New("NaN values are not allowed"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  "NaN values are not allowed",
		},
		{
			name:           "infinite values error",
			err:            errors.New("infinite values are not allowed"),
			expectedStatus: http.StatusBadRequest,
			expectedError:  "infinite values are not allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/test", func(c *gin.Context) {
				HandleBindError(c, tt.err)
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			var response models.ErrorResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedError, response.Error)
		})
	}
}

func TestHandleSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		result         float64
		expectedStatus int
		expectedResult float64
	}{
		{
			name:           "positive result",
			result:         42.0,
			expectedStatus: http.StatusOK,
			expectedResult: 42.0,
		},
		{
			name:           "negative result",
			result:         -42.0,
			expectedStatus: http.StatusOK,
			expectedResult: -42.0,
		},
		{
			name:           "zero result",
			result:         0.0,
			expectedStatus: http.StatusOK,
			expectedResult: 0.0,
		},
		{
			name:           "decimal result",
			result:         3.14159,
			expectedStatus: http.StatusOK,
			expectedResult: 3.14159,
		},
		{
			name:           "large number",
			result:         1e10,
			expectedStatus: http.StatusOK,
			expectedResult: 1e10,
		},
		{
			name:           "small number",
			result:         1e-10,
			expectedStatus: http.StatusOK,
			expectedResult: 1e-10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/test", func(c *gin.Context) {
				HandleSuccess(c, tt.result)
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			var response models.CalculationResponse
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, response.Result)
		})
	}
}

func TestErrorResponse_Model(t *testing.T) {
	response := models.ErrorResponse{
		Error: "test error",
	}
	assert.Equal(t, "test error", response.Error)

	// Test JSON marshaling
	jsonData, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "test error")
}

func TestCalculationResponse_Model(t *testing.T) {
	response := models.CalculationResponse{
		Result: 42.0,
	}
	assert.Equal(t, 42.0, response.Result)

	// Test JSON marshaling
	jsonData, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "42")
}
