package controller

import (
	"bytes"
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

// MockApp is a simple mock implementation of the CalculatorApp interface
type MockApp struct {
	addResult        float64
	addErr           error
	subtractResult   float64
	subtractErr      error
	multiplyResult   float64
	multiplyErr      error
	divideResult     float64
	divideErr        error
	negateResult     float64
	negateErr        error
	percentageResult float64
	percentageErr    error
}

func (m *MockApp) Add(a, b float64) (float64, error) {
	return m.addResult, m.addErr
}

func (m *MockApp) Subtract(a, b float64) (float64, error) {
	return m.subtractResult, m.subtractErr
}

func (m *MockApp) Multiply(a, b float64) (float64, error) {
	return m.multiplyResult, m.multiplyErr
}

func (m *MockApp) Divide(a, b float64) (float64, error) {
	return m.divideResult, m.divideErr
}

func (m *MockApp) Negate(value float64) (float64, error) {
	return m.negateResult, m.negateErr
}

func (m *MockApp) Percentage(value float64) (float64, error) {
	return m.percentageResult, m.percentageErr
}

func setupControllerWithMock(mockApp *MockApp) (*CalculatorController, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	ctrl := NewCalculatorController(mockApp)
	router := gin.New()
	return ctrl, router
}

func TestNewCalculatorController(t *testing.T) {
	mockApp := &MockApp{}
	ctrl := NewCalculatorController(mockApp)
	assert.NotNil(t, ctrl)
	assert.Equal(t, mockApp, ctrl.app)
}

func TestCalculatorController_Add_Success(t *testing.T) {
	mockApp := &MockApp{addResult: 8.0, addErr: nil}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/add", ctrl.Add)

	body := `{"a": 5, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response models.CalculationResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 8.0, response.Result)
}

func TestCalculatorController_Add_InvalidJSON(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/add", ctrl.Add)

	body := `invalid json`
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Add_DivisionByZeroError(t *testing.T) {
	mockApp := &MockApp{addResult: 0, addErr: app.ErrDivisionByZero}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/add", ctrl.Add)

	body := `{"a": 5, "b": 0}`
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response models.ErrorResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "division by zero is not allowed", response.Error)
}

func TestCalculatorController_Add_OverflowError(t *testing.T) {
	mockApp := &MockApp{addResult: 0, addErr: app.ErrOverflow}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/add", ctrl.Add)

	body := `{"a": 1e308, "b": 1e308}`
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response models.ErrorResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "result is out of representable range", response.Error)
}

func TestCalculatorController_Add_MissingFields(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/add", ctrl.Add)

	body := `{"a": 5}`
	req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Subtract_Success(t *testing.T) {
	mockApp := &MockApp{subtractResult: 7.0, subtractErr: nil}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/subtract", ctrl.Subtract)

	body := `{"a": 10, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/subtract", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response models.CalculationResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 7.0, response.Result)
}

func TestCalculatorController_Subtract_InvalidJSON(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/subtract", ctrl.Subtract)

	body := `invalid json`
	req := httptest.NewRequest(http.MethodPost, "/subtract", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Subtract_OverflowError(t *testing.T) {
	mockApp := &MockApp{subtractResult: 0, subtractErr: app.ErrOverflow}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/subtract", ctrl.Subtract)

	body := `{"a": 1e308, "b": -1e308}`
	req := httptest.NewRequest(http.MethodPost, "/subtract", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Multiply_Success(t *testing.T) {
	mockApp := &MockApp{multiplyResult: 15.0, multiplyErr: nil}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/multiply", ctrl.Multiply)

	body := `{"a": 5, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/multiply", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response models.CalculationResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 15.0, response.Result)
}

func TestCalculatorController_Multiply_InvalidJSON(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/multiply", ctrl.Multiply)

	body := `invalid json`
	req := httptest.NewRequest(http.MethodPost, "/multiply", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Multiply_OverflowError(t *testing.T) {
	mockApp := &MockApp{multiplyResult: 0, multiplyErr: app.ErrOverflow}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/multiply", ctrl.Multiply)

	body := `{"a": 1e308, "b": 2}`
	req := httptest.NewRequest(http.MethodPost, "/multiply", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Divide_Success(t *testing.T) {
	mockApp := &MockApp{divideResult: 5.0, divideErr: nil}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/divide", ctrl.Divide)

	body := `{"a": 15, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/divide", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response models.CalculationResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 5.0, response.Result)
}

func TestCalculatorController_Divide_InvalidJSON(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/divide", ctrl.Divide)

	body := `invalid json`
	req := httptest.NewRequest(http.MethodPost, "/divide", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Divide_DivisionByZero(t *testing.T) {
	mockApp := &MockApp{divideResult: 0, divideErr: app.ErrDivisionByZero}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/divide", ctrl.Divide)

	body := `{"a": 10, "b": 0}`
	req := httptest.NewRequest(http.MethodPost, "/divide", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Divide_OverflowError(t *testing.T) {
	mockApp := &MockApp{divideResult: 0, divideErr: app.ErrOverflow}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/divide", ctrl.Divide)

	body := `{"a": 1e308, "b": 0.5}`
	req := httptest.NewRequest(http.MethodPost, "/divide", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Negate_Success(t *testing.T) {
	mockApp := &MockApp{negateResult: -5.0, negateErr: nil}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/negate", ctrl.Negate)

	body := `{"value": 5}`
	req := httptest.NewRequest(http.MethodPost, "/negate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response models.CalculationResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, -5.0, response.Result)
}

func TestCalculatorController_Negate_InvalidJSON(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/negate", ctrl.Negate)

	body := `invalid json`
	req := httptest.NewRequest(http.MethodPost, "/negate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Negate_UnexpectedError(t *testing.T) {
	mockApp := &MockApp{negateResult: 0, negateErr: errors.New("unexpected")}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/negate", ctrl.Negate)

	body := `{"value": 5}`
	req := httptest.NewRequest(http.MethodPost, "/negate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCalculatorController_Negate_MissingValue(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/negate", ctrl.Negate)

	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/negate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Percentage_Success(t *testing.T) {
	mockApp := &MockApp{percentageResult: 0.5, percentageErr: nil}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/percentage", ctrl.Percentage)

	body := `{"value": 50}`
	req := httptest.NewRequest(http.MethodPost, "/percentage", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response models.CalculationResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 0.5, response.Result)
}

func TestCalculatorController_Percentage_InvalidJSON(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/percentage", ctrl.Percentage)

	body := `invalid json`
	req := httptest.NewRequest(http.MethodPost, "/percentage", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCalculatorController_Percentage_UnexpectedError(t *testing.T) {
	mockApp := &MockApp{percentageResult: 0, percentageErr: errors.New("unexpected")}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/percentage", ctrl.Percentage)

	body := `{"value": 50}`
	req := httptest.NewRequest(http.MethodPost, "/percentage", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCalculatorController_Percentage_MissingValue(t *testing.T) {
	mockApp := &MockApp{}
	ctrl, router := setupControllerWithMock(mockApp)
	router.POST("/percentage", ctrl.Percentage)

	body := `{}`
	req := httptest.NewRequest(http.MethodPost, "/percentage", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
