import { useState } from "react";
import type { CalculatorState } from "../../core/domain/CalculatorState";
import { initialCalculatorState } from "../../core/domain/CalculatorState";
import type { Operator } from "../../core/domain/Types";
import type { CalculatorApiPort } from "../../core/ports/CalculatorApiPort";

export function useCalculator(api: CalculatorApiPort) {
  const [state, setState] = useState<CalculatorState>(initialCalculatorState);

  const update = (partial: Partial<CalculatorState>) =>
    setState((prev) => ({ ...prev, ...partial }));

  const asyncUpdate = async (
    updateFn: () => Promise<Partial<CalculatorState>>,
  ) => {
    update({ loading: true, error: null });
    try {
      const partialResult = await updateFn();
      update({ ...partialResult, loading: false });
    } catch (err) {
      update({
        error: err instanceof Error ? err.message : "Connection Error",
        loading: false,
      });
    }
  };

  const inputDigit = (digit: string) => {
    setState((prev) => {
      if (prev.waitingForOperand) {
        return { ...prev, display: digit, waitingForOperand: false };
      }
      return {
        ...prev,
        display: prev.display === "0" ? digit : prev.display + digit,
      };
    });
  };

  const inputDecimal = () => {
    setState((prev) => {
      if (prev.waitingForOperand) {
        return { ...prev, display: "0.", waitingForOperand: false };
      }
      if (!prev.display.includes(".")) {
        return { ...prev, display: prev.display + "." };
      }
      return prev;
    });
  };

  const clear = () => {
    setState(initialCalculatorState);
  };

  const performOperation = (nextOperator: Operator) =>
    asyncUpdate(async () => {
      const inputValue = parseFloat(state.display);

      if (state.previousValue === null) {
        return {
          previousValue: state.display,
          waitingForOperand: true,
          operator: nextOperator,
          error: null,
        };
      }

      // Calculate with previous operator
      const previousValue = parseFloat(state.previousValue);
      let result = 0;

      switch (state.operator) {
        case "+":
          result = await api.add(previousValue, inputValue);
          break;
        case "-":
          result = await api.subtract(previousValue, inputValue);
          break;
        case "×":
          result = await api.multiply(previousValue, inputValue);
          break;
        case "÷":
          result = await api.divide(previousValue, inputValue);
          break;
        case "xʸ":
          result = await api.exponentiate(previousValue, inputValue);
          break;
      }

      const resultStr = String(result);
      return {
        display: resultStr,
        previousValue: resultStr,
        waitingForOperand: true,
        operator: nextOperator,
        error: null,
      };
    });

  const calculate = () =>
    asyncUpdate(async () => {
      if (!state.operator || state.previousValue === null) {
        return {};
      }

      const inputValue = parseFloat(state.display);
      const previousValue = parseFloat(state.previousValue);
      let result = 0;

      switch (state.operator) {
        case "+":
          result = await api.add(previousValue, inputValue);
          break;
        case "-":
          result = await api.subtract(previousValue, inputValue);
          break;
        case "×":
          result = await api.multiply(previousValue, inputValue);
          break;
        case "÷":
          result = await api.divide(previousValue, inputValue);
          break;
        case "xʸ":
          result = await api.exponentiate(previousValue, inputValue);
          break;
      }

      return {
        display: String(result),
        previousValue: null,
        operator: null,
        waitingForOperand: true,
        error: null,
      };
    });

  const toggleSign = () =>
    asyncUpdate(async () => {
      const result = await api.negate(parseFloat(state.display));
      return { display: String(result) };
    });

  const percentage = () =>
    asyncUpdate(async () => {
      const result = await api.percentage(parseFloat(state.display));
      return { display: String(result) };
    });

  const squareRoot = () =>
    asyncUpdate(async () => {
      const result = await api.squareRoot(parseFloat(state.display));
      return { display: String(result) };
    });

  return {
    state,
    inputDigit,
    inputDecimal,
    clear,
    performOperation,
    calculate,
    toggleSign,
    percentage,
    squareRoot,
  };
}
