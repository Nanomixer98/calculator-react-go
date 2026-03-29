package app

import (
	"calculator-server/core/port"
	"errors"
	"math"
)

var (
	ErrDivisionByZero    = errors.New("division by zero is not allowed")
	ErrOverflow          = errors.New("result is out of representable range")
	ErrInvalidSquareRoot = errors.New("square root of negative number is not allowed")
)

// CalculatorApp is the interface describing business logic operations.

// calculatorApp implements CalculatorApp.
type calculatorApp struct{}

func NewCalculatorApp() port.CalculatorApp {
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

func (c *calculatorApp) Exponentiate(base, exponent float64) (float64, error) {
	result := math.Pow(base, exponent)
	if math.IsInf(result, 0) {
		return 0, ErrOverflow
	}
	if math.IsNaN(result) {
		return 0, ErrOverflow
	}
	return result, nil
}

func (c *calculatorApp) SquareRoot(value float64) (float64, error) {
	if value < 0 {
		return 0, ErrInvalidSquareRoot
	}
	result := math.Sqrt(value)
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
