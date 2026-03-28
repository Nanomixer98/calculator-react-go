import type { Operator } from "./Types";

export interface CalculatorState {
  display: string;
  previousValue: string | null;
  operator: Operator | null;
  waitingForOperand: boolean;
  error: string | null;
  loading: boolean;
}

export const initialCalculatorState: CalculatorState = {
  display: "0",
  previousValue: null,
  operator: null,
  waitingForOperand: false,
  error: null,
  loading: false,
};
