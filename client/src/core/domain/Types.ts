export type Operator = "+" | "-" | "×" | "÷";

// Standard payload expected by the API endpoints
export interface BinaryOperationPayload {
  a: number;
  b: number;
}

export interface UnaryOperationPayload {
  value: number;
}

export interface CalculationResponse {
  result: number;
}

export interface ErrorResponse {
  error: string;
}
