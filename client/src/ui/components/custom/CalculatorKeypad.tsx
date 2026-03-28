import { Button } from "@/ui/components/ui/button";
import type { Operator } from "../../../core/domain/Types";

interface CalculatorKeypadProps {
  operator: Operator | null;
  loading: boolean;
  onClear: () => void;
  onToggleSign: () => void;
  onPercentage: () => void;
  onPerformOperation: (op: Operator) => void;
  onInputDigit: (digit: string) => void;
  onInputDecimal: () => void;
  onCalculate: () => void;
}

export function CalculatorKeypad({
  operator,
  loading,
  onClear,
  onToggleSign,
  onPercentage,
  onPerformOperation,
  onInputDigit,
  onInputDecimal,
  onCalculate,
}: CalculatorKeypadProps) {
  return (
    <div className="grid grid-cols-4 gap-3">
      {/* Row 1 */}
      <Button
        variant="secondary"
        className="h-16 text-xl font-medium rounded-full bg-neutral-500 hover:bg-neutral-400 text-black"
        onClick={onClear}
        disabled={loading}
      >
        AC
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-xl font-medium rounded-full bg-neutral-500 hover:bg-neutral-400 text-black"
        onClick={onToggleSign}
        disabled={loading}
      >
        +/-
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-xl font-medium rounded-full bg-neutral-500 hover:bg-neutral-400 text-black"
        onClick={onPercentage}
        disabled={loading}
      >
        %
      </Button>
      <Button
        className={`h-16 text-2xl font-medium rounded-full ${
          operator === "÷"
            ? "bg-white text-orange-500"
            : "bg-orange-500 hover:bg-orange-400 text-white"
        }`}
        onClick={() => onPerformOperation("÷")}
        disabled={loading}
      >
        ÷
      </Button>

      {/* Row 2 */}
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("7")}
        disabled={loading}
      >
        7
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("8")}
        disabled={loading}
      >
        8
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("9")}
        disabled={loading}
      >
        9
      </Button>
      <Button
        className={`h-16 text-2xl font-medium rounded-full ${
          operator === "×"
            ? "bg-white text-orange-500"
            : "bg-orange-500 hover:bg-orange-400 text-white"
        }`}
        onClick={() => onPerformOperation("×")}
        disabled={loading}
      >
        ×
      </Button>

      {/* Row 3 */}
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("4")}
        disabled={loading}
      >
        4
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("5")}
        disabled={loading}
      >
        5
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("6")}
        disabled={loading}
      >
        6
      </Button>
      <Button
        className={`h-16 text-2xl font-medium rounded-full ${
          operator === "-"
            ? "bg-white text-orange-500"
            : "bg-orange-500 hover:bg-orange-400 text-white"
        }`}
        onClick={() => onPerformOperation("-")}
        disabled={loading}
      >
        −
      </Button>

      {/* Row 4 */}
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("1")}
        disabled={loading}
      >
        1
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("2")}
        disabled={loading}
      >
        2
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={() => onInputDigit("3")}
        disabled={loading}
      >
        3
      </Button>
      <Button
        className={`h-16 text-2xl font-medium rounded-full ${
          operator === "+"
            ? "bg-white text-orange-500"
            : "bg-orange-500 hover:bg-orange-400 text-white"
        }`}
        onClick={() => onPerformOperation("+")}
        disabled={loading}
      >
        +
      </Button>

      {/* Row 5 */}
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white col-span-2"
        onClick={() => onInputDigit("0")}
        disabled={loading}
      >
        0
      </Button>
      <Button
        variant="secondary"
        className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
        onClick={onInputDecimal}
        disabled={loading}
      >
        .
      </Button>
      <Button
        className="h-16 text-2xl font-medium rounded-full bg-orange-500 hover:bg-orange-400 text-white"
        onClick={onCalculate}
        disabled={loading}
      >
        =
      </Button>
    </div>
  );
}
