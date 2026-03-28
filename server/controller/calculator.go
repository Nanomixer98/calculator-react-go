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
