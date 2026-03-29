import type { CalculatorApiPort } from "../../core/ports/CalculatorApiPort";
import type {
  BinaryOperationPayload,
  UnaryOperationPayload,
  CalculationResponse,
  ErrorResponse,
} from "../../core/domain/Types";

const API_BASE = "/api";

export class RestCalculatorApi implements CalculatorApiPort {
  private async post<T>(endpoint: string, body: unknown): Promise<T> {
    const response = await fetch(`${API_BASE}/${endpoint}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error((data as ErrorResponse).error || "An unknown server error occurred");
    }

    return data as T;
  }

  async add(a: number, b: number): Promise<number> {
    const payload: BinaryOperationPayload = { a, b };
    const { result } = await this.post<CalculationResponse>("add", payload);
    return result;
  }

  async subtract(a: number, b: number): Promise<number> {
    const payload: BinaryOperationPayload = { a, b };
    const { result } = await this.post<CalculationResponse>("subtract", payload);
    return result;
  }

  async multiply(a: number, b: number): Promise<number> {
    const payload: BinaryOperationPayload = { a, b };
    const { result } = await this.post<CalculationResponse>("multiply", payload);
    return result;
  }

  async divide(a: number, b: number): Promise<number> {
    const payload: BinaryOperationPayload = { a, b };
    const { result } = await this.post<CalculationResponse>("divide", payload);
    return result;
  }

  async exponentiate(base: number, exponent: number): Promise<number> {
    const payload: BinaryOperationPayload = { a: base, b: exponent };
    const { result } = await this.post<CalculationResponse>("exponentiate", payload);
    return result;
  }

  async squareRoot(value: number): Promise<number> {
    const payload: UnaryOperationPayload = { value };
    const { result } = await this.post<CalculationResponse>("squareroot", payload);
    return result;
  }

  async negate(value: number): Promise<number> {
    const payload: UnaryOperationPayload = { value };
    const { result } = await this.post<CalculationResponse>("negate", payload);
    return result;
  }

  async percentage(value: number): Promise<number> {
    const payload: UnaryOperationPayload = { value };
    const { result } = await this.post<CalculationResponse>("percentage", payload);
    return result;
  }
}

// Export a default instance for easier injection
export const rootRestCalculatorApi = new RestCalculatorApi();
