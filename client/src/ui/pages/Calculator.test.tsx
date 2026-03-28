import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach, afterEach } from "vitest";
import Calculator from "./Calculator";

// Mocking useCalculator to avoid network calls and purely test UI interactions
const mockUseCalculator = vi.fn();
vi.mock("../../app/hooks/useCalculator", () => ({
  useCalculator: () => mockUseCalculator()
}));

// Mock the global adapter instance to avoid creating new ones with fetch logic
vi.mock("../../adapters/rest/RestCalculatorApi", () => ({
  rootRestCalculatorApi: {}
}));

describe("Calculator Component", () => {
  beforeEach(() => {
    mockUseCalculator.mockReturnValue({
      state: {
        display: "0",
        previousValue: null,
        operator: null,
        waitingForOperand: false,
        error: null,
        loading: false,
      },
      inputDigit: vi.fn(),
      inputDecimal: vi.fn(),
      clear: vi.fn(),
      performOperation: vi.fn(),
      calculate: vi.fn(),
      toggleSign: vi.fn(),
      percentage: vi.fn(),
    });
  });

  afterEach(() => {
    vi.clearAllMocks();
  });

  it("should render initially with 0 displayed", () => {
    render(<Calculator />);
    expect(screen.getByTestId("calculator-display")).toHaveTextContent("0");
  });

  it("should display error state if present", () => {
    mockUseCalculator.mockReturnValue({
      state: {
        display: "Error",
        previousValue: null,
        operator: null,
        waitingForOperand: false,
        error: "Division by zero",
        loading: false,
      },
      inputDigit: vi.fn(),
      inputDecimal: vi.fn(),
      clear: vi.fn(),
      performOperation: vi.fn(),
      calculate: vi.fn(),
      toggleSign: vi.fn(),
      percentage: vi.fn(),
    });

    render(<Calculator />);
    expect(screen.getByText("Error")).toBeInTheDocument();
  });

  it("should interact with keypad and call actions", () => {
    const inputDigitMock = vi.fn();
    const calculateMock = vi.fn();

    mockUseCalculator.mockReturnValue({
      state: {
        display: "5",
        previousValue: null,
        operator: null,
        waitingForOperand: false,
        error: null,
        loading: false,
      },
      inputDigit: inputDigitMock,
      inputDecimal: vi.fn(),
      clear: vi.fn(),
      performOperation: vi.fn(),
      calculate: calculateMock,
      toggleSign: vi.fn(),
      percentage: vi.fn(),
    });

    render(<Calculator />);

    // Click on number 7
    const btn7 = screen.getByText("7");
    fireEvent.click(btn7);
    expect(inputDigitMock).toHaveBeenCalledWith("7");

    // Click on equals
    const btnEquals = document.querySelector('button[data-testid="calculate"]') 
      || screen.getByText("="); // fallback if no testid
      
    if (btnEquals) {
      fireEvent.click(btnEquals);
      expect(calculateMock).toHaveBeenCalled();
    }
  });

});
