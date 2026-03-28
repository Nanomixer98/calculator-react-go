import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import { CalculatorDisplay } from "./CalculatorDisplay";

describe("CalculatorDisplay", () => {
  it("should render display correctly", () => {
    render(
      <CalculatorDisplay
        display="123"
        error={null}
        loading={false}
        previousValue={null}
        operator={null}
        waitingForOperand={false}
      />
    );

    expect(screen.getByText("123")).toBeInTheDocument();
  });

  it("should show error instead of display when there is an error", () => {
    render(
      <CalculatorDisplay
        display="123"
        error="Network error"
        loading={false}
        previousValue={null}
        operator={null}
        waitingForOperand={false}
      />
    );

    expect(screen.getByText("Network error")).toBeInTheDocument();
  });

  it("should display the previous value and operator", () => {
    render(
      <CalculatorDisplay
        display="456"
        error={null}
        loading={false}
        previousValue="123"
        operator="+"
        waitingForOperand={true}
      />
    );

    expect(screen.getByText("123 +")).toBeInTheDocument();
  });

  it("should display full expression when waitingForOperand is false", () => {
    render(
      <CalculatorDisplay
        display="456"
        error={null}
        loading={false}
        previousValue="123"
        operator="+"
        waitingForOperand={false}
      />
    );

    expect(screen.getByText("123 + 456")).toBeInTheDocument();
  });

  it("should apply loading opacity class when loading is true", () => {
    render(
      <CalculatorDisplay
        display="123"
        error={null}
        loading={true}
        previousValue={null}
        operator={null}
        waitingForOperand={false}
      />
    );

    const display = screen.getByTestId("calculator-display");
    expect(display.className).toContain("opacity-50");
  });

  it("should show loading indicator when loading is true", () => {
    render(
      <CalculatorDisplay
        display="123"
        error={null}
        loading={true}
        previousValue={null}
        operator={null}
        waitingForOperand={false}
      />
    );

    expect(screen.getByText("Calculating...")).toBeInTheDocument();
  });
});
