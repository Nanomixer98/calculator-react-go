import { describe, it, expectTypeOf } from "vitest";
import type { Operator, BinaryOperationPayload, UnaryOperationPayload, CalculationResponse, ErrorResponse } from "./Types";

describe("Types", () => {
  it("should have correct Operator type", () => {
    const validOperators: Operator[] = ["+", "-", "×", "÷"];
    expectTypeOf(validOperators).toEqualTypeOf<Operator[]>();
  });

  it("should have correct BinaryOperationPayload structure", () => {
    const payload: BinaryOperationPayload = { a: 5, b: 3 };
    expectTypeOf(payload.a).toBeNumber();
    expectTypeOf(payload.b).toBeNumber();
  });

  it("should have correct UnaryOperationPayload structure", () => {
    const payload: UnaryOperationPayload = { value: 10 };
    expectTypeOf(payload.value).toBeNumber();
  });

  it("should have correct CalculationResponse structure", () => {
    const response: CalculationResponse = { result: 42 };
    expectTypeOf(response.result).toBeNumber();
  });

  it("should have correct ErrorResponse structure", () => {
    const response: ErrorResponse = { error: "Something went wrong" };
    expectTypeOf(response.error).toBeString();
  });
});
