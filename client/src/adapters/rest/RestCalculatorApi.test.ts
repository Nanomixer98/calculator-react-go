import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import { RestCalculatorApi } from "./RestCalculatorApi";

describe("RestCalculatorApi", () => {
  let api: RestCalculatorApi;
  const mockFetch = vi.fn();

  beforeEach(() => {
    api = new RestCalculatorApi();
    global.fetch = mockFetch;
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  const setupFetchMock = (data: any, ok: boolean = true) => {
    mockFetch.mockResolvedValue({
      ok,
      json: () => Promise.resolve(data),
    });
  };

  it("should handle error response", async () => {
    setupFetchMock({ error: "Division by zero" }, false);

    await expect(api.divide(10, 0)).rejects.toThrow("Division by zero");
  });

  it("should handle unknown server error", async () => {
    setupFetchMock({}, false);

    await expect(api.add(1, 2)).rejects.toThrow("An unknown server error occurred");
  });

  it("should add two numbers", async () => {
    setupFetchMock({ result: 15 });
    const result = await api.add(10, 5);
    expect(result).toBe(15);
    expect(mockFetch).toHaveBeenCalledWith("/api/add", expect.objectContaining({
      body: JSON.stringify({ a: 10, b: 5 })
    }));
  });

  it("should subtract two numbers", async () => {
    setupFetchMock({ result: 5 });
    const result = await api.subtract(10, 5);
    expect(result).toBe(5);
    expect(mockFetch).toHaveBeenCalledWith("/api/subtract", expect.objectContaining({
      body: JSON.stringify({ a: 10, b: 5 })
    }));
  });

  it("should multiply two numbers", async () => {
    setupFetchMock({ result: 50 });
    const result = await api.multiply(10, 5);
    expect(result).toBe(50);
    expect(mockFetch).toHaveBeenCalledWith("/api/multiply", expect.objectContaining({
      body: JSON.stringify({ a: 10, b: 5 })
    }));
  });

  it("should divide two numbers", async () => {
    setupFetchMock({ result: 2 });
    const result = await api.divide(10, 5);
    expect(result).toBe(2);
    expect(mockFetch).toHaveBeenCalledWith("/api/divide", expect.objectContaining({
      body: JSON.stringify({ a: 10, b: 5 })
    }));
  });

  it("should negate a number", async () => {
    setupFetchMock({ result: -10 });
    const result = await api.negate(10);
    expect(result).toBe(-10);
    expect(mockFetch).toHaveBeenCalledWith("/api/negate", expect.objectContaining({
      body: JSON.stringify({ value: 10 })
    }));
  });

  it("should calculate percentage of a number", async () => {
    setupFetchMock({ result: 0.1 });
    const result = await api.percentage(10);
    expect(result).toBe(0.1);
    expect(mockFetch).toHaveBeenCalledWith("/api/percentage", expect.objectContaining({
      body: JSON.stringify({ value: 10 })
    }));
  });
});
