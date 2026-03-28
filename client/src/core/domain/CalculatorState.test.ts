import { describe, it, expect } from "vitest";
import { initialCalculatorState } from "./CalculatorState";

describe("CalculatorState", () => {
  it("should have correct initial state", () => {
    expect(initialCalculatorState).toEqual({
      display: "0",
      previousValue: null,
      operator: null,
      waitingForOperand: false,
      error: null,
      loading: false,
    });
  });
});
