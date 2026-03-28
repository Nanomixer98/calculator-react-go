import type { Operator } from "../../../core/domain/Types";

interface CalculatorDisplayProps {
  display: string;
  error: string | null;
  loading: boolean;
  previousValue: string | null;
  operator: Operator | null;
  waitingForOperand: boolean;
}

export function CalculatorDisplay({
  display,
  error,
  loading,
  previousValue,
  operator,
  waitingForOperand,
}: CalculatorDisplayProps) {
  const hasPrevious = previousValue !== null && operator !== null;
  const expression = hasPrevious
    ? `${previousValue} ${operator} ${!waitingForOperand ? display : ""}`
    : "";

  return (
    <>
      {/* Loading Indicator */}
      {loading && (
        <div className="absolute top-2 right-4 text-xs text-orange-500 animate-pulse">
          Calculating...
        </div>
      )}

      {/* Display */}
      <div className="h-24 flex flex-col items-end justify-end px-4 pb-2 mb-4 relative">
        {/* Error Message */}
        {error && (
          <div className="text-red-500 text-xs mt-1 absolute top-0 text-right w-full pr-4 pt-1">
            {error}
          </div>
        )}

        {/* Ongoing Expression */}
        <div className="text-neutral-400 text-sm h-5 mb-1 truncate w-full text-right tracking-wider font-light">
          {expression}
        </div>

        <span
          className={`text-5xl font-light text-white truncate transition-opacity duration-200 ${
            loading ? "opacity-50" : "opacity-100"
          }`}
        >
          {display.length > 9 ? parseFloat(display).toExponential(4) : display}
        </span>
      </div>
    </>
  );
}
