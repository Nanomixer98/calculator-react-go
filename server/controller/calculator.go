package controller

import (
	"calculator-server/app"

	"github.com/gin-gonic/gin"
)

type CalculatorController struct {
	app app.CalculatorApp
}

func NewCalculatorController(a app.CalculatorApp) *CalculatorController {
	return &CalculatorController{app: a}
}

// Endpoints

// Add godoc
// @Summary      Addition
// @Description  Adds operand A to operand B.
// @Tags         Binary Operations
// @Accept       json
// @Produce      json
// @Param        request  body      models.BinaryOperationRequest  true  "Operands"
// @Success      200      {object}  models.CalculationResponse     "Calculation result"
// @Failure      400      {object}  models.ErrorResponse           "Error like Invalid JSON, NaN/Inf, or Overflow"
// @Router       /add [post]
func (ctrl *CalculatorController) Add(c *gin.Context) {
	req, err := BindBinaryRequest(c)
	if err != nil {
		HandleBindError(c, err)
		return
	}

	result, err := ctrl.app.Add(*req.A, *req.B)
	if err != nil {
		HandleError(c, err)
		return
	}

	HandleSuccess(c, result)
}

// Subtract godoc
// @Summary      Subtraction
// @Description  Subtracts operand B from operand A.
// @Tags         Binary Operations
// @Accept       json
// @Produce      json
// @Param        request  body      models.BinaryOperationRequest  true  "Operands"
// @Success      200      {object}  models.CalculationResponse     "Calculation result"
// @Failure      400      {object}  models.ErrorResponse
// @Router       /subtract [post]
func (ctrl *CalculatorController) Subtract(c *gin.Context) {
	req, err := BindBinaryRequest(c)
	if err != nil {
		HandleBindError(c, err)
		return
	}

	result, err := ctrl.app.Subtract(*req.A, *req.B)
	if err != nil {
		HandleError(c, err)
		return
	}

	HandleSuccess(c, result)
}

// Multiply godoc
// @Summary      Multiplication
// @Description  Multiplies operand A by operand B.
// @Tags         Binary Operations
// @Accept       json
// @Produce      json
// @Param        request  body      models.BinaryOperationRequest  true  "Operands"
// @Success      200      {object}  models.CalculationResponse     "Calculation result"
// @Failure      400      {object}  models.ErrorResponse
// @Router       /multiply [post]
func (ctrl *CalculatorController) Multiply(c *gin.Context) {
	req, err := BindBinaryRequest(c)
	if err != nil {
		HandleBindError(c, err)
		return
	}

	result, err := ctrl.app.Multiply(*req.A, *req.B)
	if err != nil {
		HandleError(c, err)
		return
	}

	HandleSuccess(c, result)
}

// Divide godoc
// @Summary      Division
// @Description  Divides operand A by operand B. Returns a 400 Error in case of division by zero.
// @Tags         Binary Operations
// @Accept       json
// @Produce      json
// @Param        request  body      models.BinaryOperationRequest  true  "Operands"
// @Success      200      {object}  models.CalculationResponse     "Calculation result"
// @Failure      400      {object}  models.ErrorResponse
// @Router       /divide [post]
func (ctrl *CalculatorController) Divide(c *gin.Context) {
	req, err := BindBinaryRequest(c)
	if err != nil {
		HandleBindError(c, err)
		return
	}

	result, err := ctrl.app.Divide(*req.A, *req.B)
	if err != nil {
		HandleError(c, err)
		return
	}

	HandleSuccess(c, result)
}

// Negate godoc
// @Summary      Negation (Sign Change)
// @Description  Returns the negative value of the input.
// @Tags         Unary Operations
// @Accept       json
// @Produce      json
// @Param        request  body      models.UnaryOperationRequest  true  "Single operand"
// @Success      200      {object}  models.CalculationResponse    "Calculation result"
// @Failure      400      {object}  models.ErrorResponse
// @Router       /negate [post]
func (ctrl *CalculatorController) Negate(c *gin.Context) {
	req, err := BindUnaryRequest(c)
	if err != nil {
		HandleBindError(c, err)
		return
	}

	result, err := ctrl.app.Negate(*req.Value)
	if err != nil {
		HandleError(c, err)
		return
	}

	HandleSuccess(c, result)
}

// Percentage godoc
// @Summary      Percentage
// @Description  Returns the percentage of the given value (value / 100).
// @Tags         Unary Operations
// @Accept       json
// @Produce      json
// @Param        request  body      models.UnaryOperationRequest  true  "Single operand"
// @Success      200      {object}  models.CalculationResponse    "Calculation result"
// @Failure      400      {object}  models.ErrorResponse
// @Router       /percentage [post]
func (ctrl *CalculatorController) Percentage(c *gin.Context) {
	req, err := BindUnaryRequest(c)
	if err != nil {
		HandleBindError(c, err)
		return
	}

	result, err := ctrl.app.Percentage(*req.Value)
	if err != nil {
		HandleError(c, err)
		return
	}

	HandleSuccess(c, result)
}
