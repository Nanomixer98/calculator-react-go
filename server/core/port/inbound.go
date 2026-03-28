package port

type CalculatorApp interface {
	Add(a, b float64) (float64, error)
	Subtract(a, b float64) (float64, error)
	Multiply(a, b float64) (float64, error)
	Divide(a, b float64) (float64, error)
	Negate(value float64) (float64, error)
	Percentage(value float64) (float64, error)
}
