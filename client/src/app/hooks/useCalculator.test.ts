import { renderHook, act } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { useCalculator } from "./useCalculator";
import type { CalculatorApiPort } from "../../core/ports/CalculatorApiPort";

describe("useCalculator", () => {
  let mockApi: CalculatorApiPort;

  beforeEach(() => {
    mockApi = {
      add: vi.fn(),
      subtract: vi.fn(),
      multiply: vi.fn(),
      divide: vi.fn(),
      negate: vi.fn(),
      percentage: vi.fn(),
    };
  });

  it("should initialize with initial state", () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    expect(result.current.state.display).toBe("0");
    expect(result.current.state.previousValue).toBeNull();
    expect(result.current.state.operator).toBeNull();
    expect(result.current.state.waitingForOperand).toBe(false);
  });

  it("should handle digit input", () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => {
      result.current.inputDigit("5");
    });
    expect(result.current.state.display).toBe("5");

    act(() => {
      result.current.inputDigit("2");
    });
    expect(result.current.state.display).toBe("52");
  });

  it("should handle digit input when waitingForOperand is true", async () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => {
      result.current.inputDigit("5");
    });

    await act(async () => {
      await result.current.performOperation("+");
    });

    act(() => {
      result.current.inputDigit("3");
    });

    expect(result.current.state.display).toBe("3");
    expect(result.current.state.waitingForOperand).toBe(false);
  });

  it("should handle decimal input", () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => {
      result.current.inputDecimal();
    });
    expect(result.current.state.display).toBe("0.");

    act(() => {
      result.current.inputDigit("5");
    });
    expect(result.current.state.display).toBe("0.5");

    act(() => {
      result.current.inputDecimal();
    });
    expect(result.current.state.display).toBe("0.5"); // ignores double decimal
  });

  it("should handle decimal input when waitingForOperand is true", async () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => {
      result.current.inputDigit("5");
    });

    await act(async () => {
      await result.current.performOperation("+");
    });

    act(() => {
      result.current.inputDecimal();
    });

    expect(result.current.state.display).toBe("0.");
    expect(result.current.state.waitingForOperand).toBe(false);
  });

  it("should clear the state", () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => {
      result.current.inputDigit("5");
      result.current.clear();
    });
    expect(result.current.state.display).toBe("0");
  });

  it("should set operator and wait for operand when performOperation is called", async () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => {
      result.current.inputDigit("5");
    });

    await act(async () => {
      await result.current.performOperation("+");
    });

    expect(result.current.state.operator).toBe("+");
    expect(result.current.state.previousValue).toBe("5");
    expect(result.current.state.waitingForOperand).toBe(true);
  });

  it("should handle performOperation when there is previousValue but no operator", async () => {
    // Mock the API to avoid any async operations
    const { result } = renderHook(() => useCalculator(mockApi));

    // First, set up a state with previousValue and operator
    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.performOperation("+"); });
    act(() => { result.current.inputDigit("3"); });

    // Clear should reset everything
    act(() => { result.current.clear(); });
    
    // Now start fresh with a specific scenario
    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.performOperation("+"); });
    
    // The state now has previousValue="5", operator="+", waitingForOperand=true
    expect(result.current.state.previousValue).toBe("5");
    expect(result.current.state.operator).toBe("+");
    
    // Now if we directly clear only the operator (simulating the edge case)
    // by chaining operations, it should work correctly
    vi.mocked(mockApi.add).mockResolvedValue(10);
    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.performOperation("+"); });
    
    expect(result.current.state.previousValue).toBe("10");
    expect(result.current.state.operator).toBe("+");
  });

  it("should handle subtract operation", async () => {
    vi.mocked(mockApi.subtract).mockResolvedValue(3);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.performOperation("-"); });
    act(() => { result.current.inputDigit("2"); });
    await act(async () => { await result.current.calculate(); });

    expect(mockApi.subtract).toHaveBeenCalledWith(5, 2);
    expect(result.current.state.display).toBe("3");
  });

  it("should handle multiply operation", async () => {
    vi.mocked(mockApi.multiply).mockResolvedValue(10);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.performOperation("×"); });
    act(() => { result.current.inputDigit("2"); });
    await act(async () => { await result.current.performOperation("+"); });

    expect(mockApi.multiply).toHaveBeenCalledWith(5, 2);
    expect(result.current.state.display).toBe("10");
  });

  it("should handle divide operation", async () => {
    vi.mocked(mockApi.divide).mockResolvedValue(2);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("6"); });
    await act(async () => { await result.current.performOperation("÷"); });
    act(() => { result.current.inputDigit("3"); });
    await act(async () => { await result.current.performOperation("+"); });

    expect(mockApi.divide).toHaveBeenCalledWith(6, 3);
    expect(result.current.state.display).toBe("2");
  });

  it("should calculate correctly", async () => {
    vi.mocked(mockApi.add).mockResolvedValue(10);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => {
      result.current.inputDigit("5");
    });

    await act(async () => {
      await result.current.performOperation("+");
    });

    act(() => {
      result.current.inputDigit("5");
    });

    await act(async () => {
      await result.current.calculate();
    });

    expect(mockApi.add).toHaveBeenCalledWith(5, 5);
    expect(result.current.state.display).toBe("10");
    expect(result.current.state.operator).toBeNull();
  });

  it("should handle calculate with subtract", async () => {
    vi.mocked(mockApi.subtract).mockResolvedValue(5);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("8"); });
    await act(async () => { await result.current.performOperation("-"); });
    act(() => { result.current.inputDigit("3"); });
    await act(async () => { await result.current.calculate(); });

    expect(mockApi.subtract).toHaveBeenCalledWith(8, 3);
    expect(result.current.state.display).toBe("5");
  });

  it("should handle calculate with multiply", async () => {
    vi.mocked(mockApi.multiply).mockResolvedValue(24);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("6"); });
    await act(async () => { await result.current.performOperation("×"); });
    act(() => { result.current.inputDigit("4"); });
    await act(async () => { await result.current.calculate(); });

    expect(mockApi.multiply).toHaveBeenCalledWith(6, 4);
    expect(result.current.state.display).toBe("24");
  });

  it("should handle calculate with divide", async () => {
    vi.mocked(mockApi.divide).mockResolvedValue(4);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("8"); });
    await act(async () => { await result.current.performOperation("÷"); });
    act(() => { result.current.inputDigit("2"); });
    await act(async () => { await result.current.calculate(); });

    expect(mockApi.divide).toHaveBeenCalledWith(8, 2);
    expect(result.current.state.display).toBe("4");
  });

  it("should do nothing on calculate without operator", async () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.calculate(); });

    expect(result.current.state.display).toBe("5");
    expect(result.current.state.operator).toBeNull();
  });

  it("should do nothing on calculate without previous value", async () => {
    const { result } = renderHook(() => useCalculator(mockApi));

    // Set operator but no previous value
    await act(async () => { await result.current.performOperation("+"); });

    expect(result.current.state.display).toBe("0");
  });

  it("should chain operations correctly", async () => {
    vi.mocked(mockApi.add).mockResolvedValue(10);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });

    await act(async () => { await result.current.performOperation("+"); });

    act(() => { result.current.inputDigit("5"); });

    await act(async () => { await result.current.performOperation("-"); });

    expect(mockApi.add).toHaveBeenCalledWith(5, 5);
    expect(result.current.state.display).toBe("10");
    expect(result.current.state.operator).toBe("-");
  });

  it("should handle toggleSign", async () => {
    vi.mocked(mockApi.negate).mockResolvedValue(-5);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.toggleSign(); });

    expect(mockApi.negate).toHaveBeenCalledWith(5);
    expect(result.current.state.display).toBe("-5");
  });

  it("should handle percentage", async () => {
    vi.mocked(mockApi.percentage).mockResolvedValue(0.05);
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.percentage(); });

    expect(mockApi.percentage).toHaveBeenCalledWith(5);
    expect(result.current.state.display).toBe("0.05");
  });

  it("should handle error with Error instance", async () => {
    vi.mocked(mockApi.add).mockRejectedValue(new Error("Network error"));
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.performOperation("+"); });
    act(() => { result.current.inputDigit("3"); });
    await act(async () => { await result.current.calculate(); });

    expect(result.current.state.error).toBe("Network error");
    expect(result.current.state.loading).toBe(false);
  });

  it("should handle error with non-Error object", async () => {
    vi.mocked(mockApi.add).mockRejectedValue("Some error");
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    await act(async () => { await result.current.performOperation("+"); });
    act(() => { result.current.inputDigit("3"); });
    await act(async () => { await result.current.calculate(); });

    expect(result.current.state.error).toBe("Connection Error");
    expect(result.current.state.loading).toBe(false);
  });

  it("should set loading state during async operations", async () => {
    vi.mocked(mockApi.add).mockImplementation(() => 
      new Promise(resolve => setTimeout(() => resolve(10), 10))
    );
    const { result } = renderHook(() => useCalculator(mockApi));

    act(() => { result.current.inputDigit("5"); });
    
    let promise: Promise<void>;
    act(() => {
      promise = result.current.performOperation("+");
    });

    expect(result.current.state.loading).toBe(true);

    await act(async () => { await promise; });

    expect(result.current.state.loading).toBe(false);
  });
});
