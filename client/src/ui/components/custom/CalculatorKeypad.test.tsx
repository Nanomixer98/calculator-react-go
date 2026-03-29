import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import "@testing-library/jest-dom";
import { CalculatorKeypad } from "./CalculatorKeypad";
import type { Operator } from "../../../core/domain/Types";

describe("CalculatorKeypad", () => {
  const defaultProps = {
    operator: null as Operator | null,
    loading: false,
    onClear: vi.fn(),
    onToggleSign: vi.fn(),
    onPercentage: vi.fn(),
    onSquareRoot: vi.fn(),
    onPerformOperation: vi.fn(),
    onInputDigit: vi.fn(),
    onInputDecimal: vi.fn(),
    onCalculate: vi.fn(),
  };

  it("should render all buttons", () => {
    render(<CalculatorKeypad {...defaultProps} />);

    // Check for new operation buttons
    expect(screen.getByText("xʸ")).toBeInTheDocument();
    expect(screen.getByText("√")).toBeInTheDocument();

    // Check for operation buttons
    expect(screen.getByText("AC")).toBeInTheDocument();
    expect(screen.getByText("+/-")).toBeInTheDocument();
    expect(screen.getByText("%")).toBeInTheDocument();
    expect(screen.getByText("÷")).toBeInTheDocument();

    // Check for digit buttons
    expect(screen.getByText("0")).toBeInTheDocument();
    expect(screen.getByText("1")).toBeInTheDocument();
    expect(screen.getByText("2")).toBeInTheDocument();
    expect(screen.getByText("3")).toBeInTheDocument();
    expect(screen.getByText("4")).toBeInTheDocument();
    expect(screen.getByText("5")).toBeInTheDocument();
    expect(screen.getByText("6")).toBeInTheDocument();
    expect(screen.getByText("7")).toBeInTheDocument();
    expect(screen.getByText("8")).toBeInTheDocument();
    expect(screen.getByText("9")).toBeInTheDocument();

    // Check for decimal and equals
    expect(screen.getByText(".")).toBeInTheDocument();
    expect(screen.getByText("=")).toBeInTheDocument();
  });

  it("should call onClear when AC button is clicked", () => {
    const onClear = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onClear={onClear} />);

    fireEvent.click(screen.getByText("AC"));
    expect(onClear).toHaveBeenCalled();
  });

  it("should call onToggleSign when +/- button is clicked", () => {
    const onToggleSign = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onToggleSign={onToggleSign} />);

    fireEvent.click(screen.getByText("+/-"));
    expect(onToggleSign).toHaveBeenCalled();
  });

  it("should call onPercentage when % button is clicked", () => {
    const onPercentage = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onPercentage={onPercentage} />);

    fireEvent.click(screen.getByText("%"));
    expect(onPercentage).toHaveBeenCalled();
  });

  it("should call onSquareRoot when √ button is clicked", () => {
    const onSquareRoot = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onSquareRoot={onSquareRoot} />);

    fireEvent.click(screen.getByText("√"));
    expect(onSquareRoot).toHaveBeenCalled();
  });

  it("should call onPerformOperation with correct operator", () => {
    const onPerformOperation = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onPerformOperation={onPerformOperation} />);

    fireEvent.click(screen.getByText("xʸ"));
    expect(onPerformOperation).toHaveBeenCalledWith("xʸ");

    fireEvent.click(screen.getByText("÷"));
    expect(onPerformOperation).toHaveBeenCalledWith("÷");

    fireEvent.click(screen.getByText("×"));
    expect(onPerformOperation).toHaveBeenCalledWith("×");

    fireEvent.click(screen.getByText("−"));
    expect(onPerformOperation).toHaveBeenCalledWith("-");

    fireEvent.click(screen.getByText("+"));
    expect(onPerformOperation).toHaveBeenCalledWith("+");
  });

  it("should call onInputDigit with correct digit for all number buttons", () => {
    const onInputDigit = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onInputDigit={onInputDigit} />);

    // Test all digit buttons
    const digits = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"];
    digits.forEach((digit) => {
      fireEvent.click(screen.getByText(digit));
      expect(onInputDigit).toHaveBeenCalledWith(digit);
    });
  });

  it("should call onInputDecimal when . button is clicked", () => {
    const onInputDecimal = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onInputDecimal={onInputDecimal} />);

    fireEvent.click(screen.getByText("."));
    expect(onInputDecimal).toHaveBeenCalled();
  });

  it("should call onCalculate when = button is clicked", () => {
    const onCalculate = vi.fn();
    render(<CalculatorKeypad {...defaultProps} onCalculate={onCalculate} />);

    fireEvent.click(screen.getByText("="));
    expect(onCalculate).toHaveBeenCalled();
  });

  it("should disable all buttons when loading", () => {
    render(<CalculatorKeypad {...defaultProps} loading={true} />);

    const buttons = screen.getAllByRole("button");
    buttons.forEach((button) => {
      expect(button).toBeDisabled();
    });
  });

  it("should highlight active operator", () => {
    const { rerender } = render(<CalculatorKeypad {...defaultProps} operator="+" />);

    // Check that + button has active class styling (bg-white text-orange-500)
    const plusButton = screen.getByText("+");
    expect(plusButton.className).toContain("bg-white");
    expect(plusButton.className).toContain("text-orange-500");

    // Rerender with different operator
    rerender(<CalculatorKeypad {...defaultProps} operator="÷" />);
    const divideButton = screen.getByText("÷");
    expect(divideButton.className).toContain("bg-white");
    expect(divideButton.className).toContain("text-orange-500");

    // Test multiply operator
    rerender(<CalculatorKeypad {...defaultProps} operator="×" />);
    const multiplyButton = screen.getByText("×");
    expect(multiplyButton.className).toContain("bg-white");
    expect(multiplyButton.className).toContain("text-orange-500");

    // Test minus operator
    rerender(<CalculatorKeypad {...defaultProps} operator="-" />);
    const minusButton = screen.getByText("−");
    expect(minusButton.className).toContain("bg-white");
    expect(minusButton.className).toContain("text-orange-500");

    // Test exponentiation operator
    rerender(<CalculatorKeypad {...defaultProps} operator="xʸ" />);
    const exponentiateButton = screen.getByText("xʸ");
    expect(exponentiateButton.className).toContain("bg-white");
    expect(exponentiateButton.className).toContain("text-purple-600");
  });
});
