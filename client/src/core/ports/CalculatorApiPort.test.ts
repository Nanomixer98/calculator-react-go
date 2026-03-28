import { describe, it, expectTypeOf, vi } from "vitest";
import type { CalculatorApiPort } from "./CalculatorApiPort";

describe("CalculatorApiPort interface", () => {
  it("should enforce correct method signatures", async () => {
    // Create a mock implementation that satisfies the interface
    const mockApi: CalculatorApiPort = {
      add: vi.fn().mockResolvedValue(0 as number),
      subtract: vi.fn().mockResolvedValue(0 as number),
      multiply: vi.fn().mockResolvedValue(0 as number),
      divide: vi.fn().mockResolvedValue(0 as number),
      negate: vi.fn().mockResolvedValue(0 as number),
      percentage: vi.fn().mockResolvedValue(0 as number),
    };

    // Verify all methods exist and return promises
    expectTypeOf(mockApi.add).toBeFunction();
    expectTypeOf(mockApi.subtract).toBeFunction();
    expectTypeOf(mockApi.multiply).toBeFunction();
    expectTypeOf(mockApi.divide).toBeFunction();
    expectTypeOf(mockApi.negate).toBeFunction();
    expectTypeOf(mockApi.percentage).toBeFunction();

    // Verify return types are Promise<number>
    expectTypeOf(mockApi.add(1, 2)).toEqualTypeOf<Promise<number>>();
    expectTypeOf(mockApi.subtract(1, 2)).toEqualTypeOf<Promise<number>>();
    expectTypeOf(mockApi.multiply(1, 2)).toEqualTypeOf<Promise<number>>();
    expectTypeOf(mockApi.divide(1, 2)).toEqualTypeOf<Promise<number>>();
    expectTypeOf(mockApi.negate(1)).toEqualTypeOf<Promise<number>>();
    expectTypeOf(mockApi.percentage(1)).toEqualTypeOf<Promise<number>>();
  });

  it("should work with actual implementations", async () => {
    // Verify the interface can be implemented
    const implementation: CalculatorApiPort = {
      add: async (a: number, b: number) => a + b,
      subtract: async (a: number, b: number) => a - b,
      multiply: async (a: number, b: number) => a * b,
      divide: async (a: number, b: number) => a / b,
      negate: async (value: number) => -value,
      percentage: async (value: number) => value / 100,
    };

    const sum = await implementation.add(2, 3);
    expectTypeOf(sum).toBeNumber();
  });
});
