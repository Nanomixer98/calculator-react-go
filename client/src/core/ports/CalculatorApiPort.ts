export interface CalculatorApiPort {
  add(a: number, b: number): Promise<number>;
  subtract(a: number, b: number): Promise<number>;
  multiply(a: number, b: number): Promise<number>;
  divide(a: number, b: number): Promise<number>;
  exponentiate(base: number, exponent: number): Promise<number>;
  squareRoot(value: number): Promise<number>;
  negate(value: number): Promise<number>;
  percentage(value: number): Promise<number>;
}
