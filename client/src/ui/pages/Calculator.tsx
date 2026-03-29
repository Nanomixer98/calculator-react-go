import { rootRestCalculatorApi } from "../../adapters/rest/RestCalculatorApi";
import { useCalculator } from "../../app/hooks/useCalculator";
import { CalculatorDisplay } from "../components/custom/CalculatorDisplay";
import { CalculatorKeypad } from "../components/custom/CalculatorKeypad";

export default function Calculator() {
  // Injecting the REST Adapter into the UI UseCase interactor hook
  const {
    state,
    inputDigit,
    inputDecimal,
    clear,
    performOperation,
    calculate,
    toggleSign,
    percentage,
    squareRoot,
  } = useCalculator(rootRestCalculatorApi);

  return (
    <main className="min-h-screen flex items-center justify-center bg-neutral-950 p-4">
      <div className="w-full max-w-xs bg-neutral-900 rounded-3xl p-4 shadow-2xl relative overflow-hidden">
        <CalculatorDisplay
          display={state.display}
          error={state.error}
          loading={state.loading}
          previousValue={state.previousValue}
          operator={state.operator}
          waitingForOperand={state.waitingForOperand}
        />
        <CalculatorKeypad
          operator={state.operator}
          loading={state.loading}
          onClear={clear}
          onToggleSign={toggleSign}
          onPercentage={percentage}
          onSquareRoot={squareRoot}
          onPerformOperation={performOperation}
          onInputDigit={inputDigit}
          onInputDecimal={inputDecimal}
          onCalculate={calculate}
        />
      </div>
    </main>
  );
}
