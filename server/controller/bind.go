package controller

import (
	"calculator-server/models"
	"errors"
	"math"

	"github.com/gin-gonic/gin"
)

// BindBinaryRequest extracts, binds and validates a binary operation request from the gin context.
func BindBinaryRequest(c *gin.Context) (*models.BinaryOperationRequest, error) {
	var req models.BinaryOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, errors.New("invalid JSON body")
	}
	if req.A == nil || req.B == nil {
		return nil, errors.New("both fields 'a' and 'b' are required")
	}
	if math.IsNaN(*req.A) || math.IsNaN(*req.B) {
		return nil, errors.New("NaN values are not allowed")
	}
	if math.IsInf(*req.A, 0) || math.IsInf(*req.B, 0) {
		return nil, errors.New("infinite values are not allowed")
	}
	return &req, nil
}

// BindUnaryRequest extracts, binds and validates a unary operation request from the gin context.
func BindUnaryRequest(c *gin.Context) (*models.UnaryOperationRequest, error) {
	var req models.UnaryOperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, errors.New("invalid JSON body")
	}
	if req.Value == nil {
		return nil, errors.New("field 'value' is required")
	}
	if math.IsNaN(*req.Value) {
		return nil, errors.New("NaN value is not allowed")
	}
	if math.IsInf(*req.Value, 0) {
		return nil, errors.New("infinite value is not allowed")
	}
	return &req, nil
}
