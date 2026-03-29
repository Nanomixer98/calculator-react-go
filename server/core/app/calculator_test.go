package app

import (
	"errors"
	"math"
	"testing"
)

func TestNewCalculatorApp(t *testing.T) {
	calc := NewCalculatorApp()
	if calc == nil {
		t.Error("NewCalculatorApp() should not return nil")
	}
}

func TestCalculatorApp_Add(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr error
	}{
		{
			name:    "positive numbers",
			a:       5,
			b:       3,
			want:    8,
			wantErr: nil,
		},
		{
			name:    "negative numbers",
			a:       -5,
			b:       -3,
			want:    -8,
			wantErr: nil,
		},
		{
			name:    "mixed signs",
			a:       5,
			b:       -3,
			want:    2,
			wantErr: nil,
		},
		{
			name:    "with zero",
			a:       5,
			b:       0,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "decimal numbers",
			a:       1.5,
			b:       2.5,
			want:    4,
			wantErr: nil,
		},
		{
			name:    "positive overflow",
			a:       math.MaxFloat64,
			b:       math.MaxFloat64,
			want:    0,
			wantErr: ErrOverflow,
		},
		{
			name:    "negative overflow",
			a:       -math.MaxFloat64,
			b:       -math.MaxFloat64,
			want:    0,
			wantErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Add(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorApp_Subtract(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr error
	}{
		{
			name:    "positive numbers",
			a:       5,
			b:       3,
			want:    2,
			wantErr: nil,
		},
		{
			name:    "negative numbers",
			a:       -5,
			b:       -3,
			want:    -2,
			wantErr: nil,
		},
		{
			name:    "result zero",
			a:       5,
			b:       5,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "subtract larger from smaller",
			a:       3,
			b:       5,
			want:    -2,
			wantErr: nil,
		},
		{
			name:    "decimal numbers",
			a:       5.5,
			b:       2.2,
			want:    3.3,
			wantErr: nil,
		},
		{
			name:    "positive overflow",
			a:       math.MaxFloat64,
			b:       -math.MaxFloat64,
			want:    0,
			wantErr: ErrOverflow,
		},
		{
			name:    "negative overflow",
			a:       -math.MaxFloat64,
			b:       math.MaxFloat64,
			want:    0,
			wantErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Subtract(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Subtract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorApp_Multiply(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr error
	}{
		{
			name:    "positive numbers",
			a:       5,
			b:       3,
			want:    15,
			wantErr: nil,
		},
		{
			name:    "negative numbers",
			a:       -5,
			b:       -3,
			want:    15,
			wantErr: nil,
		},
		{
			name:    "mixed signs",
			a:       5,
			b:       -3,
			want:    -15,
			wantErr: nil,
		},
		{
			name:    "by zero",
			a:       5,
			b:       0,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "decimal numbers",
			a:       2.5,
			b:       4,
			want:    10,
			wantErr: nil,
		},
		{
			name:    "positive overflow",
			a:       math.MaxFloat64,
			b:       2,
			want:    0,
			wantErr: ErrOverflow,
		},
		{
			name:    "negative overflow",
			a:       math.MaxFloat64,
			b:       -2,
			want:    0,
			wantErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Multiply(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorApp_Divide(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr error
	}{
		{
			name:    "positive numbers",
			a:       15,
			b:       3,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "negative numbers",
			a:       -15,
			b:       -3,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "mixed signs",
			a:       15,
			b:       -3,
			want:    -5,
			wantErr: nil,
		},
		{
			name:    "division by one",
			a:       5,
			b:       1,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "decimal result",
			a:       10,
			b:       4,
			want:    2.5,
			wantErr: nil,
		},
		{
			name:    "division by zero",
			a:       5,
			b:       0,
			want:    0,
			wantErr: ErrDivisionByZero,
		},
		{
			name:    "positive overflow",
			a:       math.MaxFloat64,
			b:       0.5,
			want:    0,
			wantErr: ErrOverflow,
		},
		{
			name:    "negative overflow",
			a:       -math.MaxFloat64,
			b:       0.5,
			want:    0,
			wantErr: ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Divide(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorApp_Negate(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name    string
		value   float64
		want    float64
		wantErr error
	}{
		{
			name:    "positive number",
			value:   5,
			want:    -5,
			wantErr: nil,
		},
		{
			name:    "negative number",
			value:   -5,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "zero",
			value:   0,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "decimal number",
			value:   3.14,
			want:    -3.14,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Negate(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Negate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorApp_Percentage(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name    string
		value   float64
		want    float64
		wantErr error
	}{
		{
			name:    "hundred",
			value:   100,
			want:    1,
			wantErr: nil,
		},
		{
			name:    "fifty",
			value:   50,
			want:    0.5,
			wantErr: nil,
		},
		{
			name:    "zero",
			value:   0,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "decimal number",
			value:   12.5,
			want:    0.125,
			wantErr: nil,
		},
		{
			name:    "negative number",
			value:   -50,
			want:    -0.5,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Percentage(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Percentage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Percentage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorApp_Exponentiate(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name     string
		base     float64
		exponent float64
		want     float64
		wantErr  error
	}{
		{
			name:     "power of 2",
			base:     2,
			exponent: 3,
			want:     8,
			wantErr:  nil,
		},
		{
			name:     "power of 1",
			base:     5,
			exponent: 1,
			want:     5,
			wantErr:  nil,
		},
		{
			name:     "power of 0",
			base:     5,
			exponent: 0,
			want:     1,
			wantErr:  nil,
		},
		{
			name:     "negative exponent",
			base:     2,
			exponent: -2,
			want:     0.25,
			wantErr:  nil,
		},
		{
			name:     "fractional exponent",
			base:     4,
			exponent: 0.5,
			want:     2,
			wantErr:  nil,
		},
		{
			name:     "zero base",
			base:     0,
			exponent: 5,
			want:     0,
			wantErr:  nil,
		},
		{
			name:     "overflow",
			base:     math.MaxFloat64,
			exponent: 2,
			want:     0,
			wantErr:  ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Exponentiate(tt.base, tt.exponent)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Exponentiate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Exponentiate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorApp_SquareRoot(t *testing.T) {
	calc := NewCalculatorApp()

	tests := []struct {
		name    string
		value   float64
		want    float64
		wantErr error
	}{
		{
			name:    "perfect square",
			value:   16,
			want:    4,
			wantErr: nil,
		},
		{
			name:    "non-perfect square",
			value:   2,
			want:    math.Sqrt(2),
			wantErr: nil,
		},
		{
			name:    "zero",
			value:   0,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "one",
			value:   1,
			want:    1,
			wantErr: nil,
		},
		{
			name:    "decimal number",
			value:   2.25,
			want:    1.5,
			wantErr: nil,
		},
		{
			name:    "negative number",
			value:   -4,
			want:    0,
			wantErr: ErrInvalidSquareRoot,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.SquareRoot(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("SquareRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("SquareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
