package app

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero = errors.New("division by zero is not allowed")
	ErrOverflow       = errors.New("result is out of representable range")
)

// CalculatorApp is the interface describing business logic operations.
type CalculatorApp interface {
	Add(a, b float64) (float64, error)
	Subtract(a, b float64) (float64, error)
	Multiply(a, b float64) (float64, error)
	Divide(a, b float64) (float64, error)
	Negate(value float64) (float64, error)
	Percentage(value float64) (float64, error)
}

// calculatorApp implements CalculatorApp.
type calculatorApp struct{}

func NewCalculatorApp() CalculatorApp {
	return &calculatorApp{}
}

func (c *calculatorApp) Add(a, b float64) (float64, error) {
	result := a + b
	if math.IsInf(result, 0) {
		return 0, ErrOverflow
	}
	return result, nil
}

func (c *calculatorApp) Subtract(a, b float64) (float64, error) {
	result := a - b
	if math.IsInf(result, 0) {
		return 0, ErrOverflow
	}
	return result, nil
}

func (c *calculatorApp) Multiply(a, b float64) (float64, error) {
	result := a * b
	if math.IsInf(result, 0) {
		return 0, ErrOverflow
	}
	return result, nil
}

func (c *calculatorApp) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	result := a / b
	if math.IsInf(result, 0) {
		return 0, ErrOverflow
	}
	return result, nil
}

func (c *calculatorApp) Negate(value float64) (float64, error) {
	return -value, nil
}

func (c *calculatorApp) Percentage(value float64) (float64, error) {
	return value / 100, nil
}
