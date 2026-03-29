package controller

import (
	"calculator-server/core/app"
	"calculator-server/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleError maps domain errors to standard HTTP responses with matching status codes.
func HandleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, app.ErrDivisionByZero), errors.Is(err, app.ErrOverflow), errors.Is(err, app.ErrInvalidSquareRoot):
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "unexpected error"})
	}
}

// HandleBindError responds with an HTTP status code for generic invalid requests body or bindings.
func HandleBindError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
}

// HandleSuccess outputs a standardized success calculation response matching the application format.
func HandleSuccess(c *gin.Context, result float64) {
	c.JSON(http.StatusOK, models.CalculationResponse{Result: result})
}
