package builder

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBuildServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	assert.NotNil(t, server)
}

func TestBuildServer_CORS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	// Test CORS preflight request (OPTIONS)
	req := httptest.NewRequest(http.MethodOptions, "/api/add", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	// Preflight should return 204 No Content
	assert.Equal(t, http.StatusNoContent, rec.Code)

	// Check CORS headers
	headers := rec.Header()
	assert.Equal(t, "*", headers.Get("Access-Control-Allow-Origin"))
	assert.Contains(t, headers.Get("Access-Control-Allow-Methods"), "POST")
	assert.Contains(t, headers.Get("Access-Control-Allow-Headers"), "Content-Type")
}

func TestBuildServer_CORSHeadersOnPOST(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	// Test CORS headers on regular POST request
	body := `{"a": 5, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	// Check CORS headers are present
	headers := rec.Header()
	assert.Equal(t, "*", headers.Get("Access-Control-Allow-Origin"))
}

func TestCorsMiddleware_OptionsRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	// Test that OPTIONS requests are handled
	req := httptest.NewRequest(http.MethodOptions, "/api/add", nil)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestCorsMiddleware_NonOptionsRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	// Test that non-OPTIONS requests pass through
	body := `{"a": 5, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	// Should get a valid response (not 204 from OPTIONS handling)
	assert.NotEqual(t, http.StatusNoContent, rec.Code)
}

func TestBuildServer_Routes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	// Test that all expected routes are available
	routes := []struct {
		method string
		path   string
	}{
		{http.MethodPost, "/api/add"},
		{http.MethodPost, "/api/subtract"},
		{http.MethodPost, "/api/multiply"},
		{http.MethodPost, "/api/divide"},
		{http.MethodPost, "/api/negate"},
		{http.MethodPost, "/api/percentage"},
	}

	for _, route := range routes {
		var found bool
		for _, r := range server.Routes() {
			if r.Method == route.method && r.Path == route.path {
				found = true
				break
			}
		}
		assert.True(t, found, "Route %s %s should be registered", route.method, route.path)
	}
}

func TestBuildServer_EndToEnd_Add(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `{"a": 5, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "8")
}

func TestBuildServer_EndToEnd_Subtract(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `{"a": 10, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/api/subtract", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "7")
}

func TestBuildServer_EndToEnd_Multiply(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `{"a": 5, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/api/multiply", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "15")
}

func TestBuildServer_EndToEnd_Divide(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `{"a": 15, "b": 3}`
	req := httptest.NewRequest(http.MethodPost, "/api/divide", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "5")
}

func TestBuildServer_EndToEnd_Negate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `{"value": 5}`
	req := httptest.NewRequest(http.MethodPost, "/api/negate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "-5")
}

func TestBuildServer_EndToEnd_Percentage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `{"value": 50}`
	req := httptest.NewRequest(http.MethodPost, "/api/percentage", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "0.5")
}

func TestBuildServer_EndToEnd_DivideByZero(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `{"a": 10, "b": 0}`
	req := httptest.NewRequest(http.MethodPost, "/api/divide", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "division by zero")
}

func TestBuildServer_EndToEnd_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	server := BuildServer()

	body := `invalid json`
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
