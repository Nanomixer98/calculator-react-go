import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function Calculator() {
  const [display, setDisplay] = useState("0");
  const [previousValue, setPreviousValue] = useState<string | null>(null);
  const [operator, setOperator] = useState<string | null>(null);
  const [waitingForOperand, setWaitingForOperand] = useState(false);

  const inputDigit = (digit: string) => {
    if (waitingForOperand) {
      setDisplay(digit);
      setWaitingForOperand(false);
    } else {
      setDisplay(display === "0" ? digit : display + digit);
    }
  };

  const inputDecimal = () => {
    if (waitingForOperand) {
      setDisplay("0.");
      setWaitingForOperand(false);
      return;
    }
    if (!display.includes(".")) {
      setDisplay(display + ".");
    }
  };

  const clear = () => {
    setDisplay("0");
    setPreviousValue(null);
    setOperator(null);
    setWaitingForOperand(false);
  };

  const performOperation = (nextOperator: string) => {
    const inputValue = parseFloat(display);

    if (previousValue === null) {
      setPreviousValue(display);
    } else if (operator) {
      const currentValue = parseFloat(previousValue);
      let result: number;

      switch (operator) {
        case "+":
          result = currentValue + inputValue;
          break;
        case "-":
          result = currentValue - inputValue;
          break;
        case "×":
          result = currentValue * inputValue;
          break;
        case "÷":
          result = inputValue !== 0 ? currentValue / inputValue : 0;
          break;
        default:
          result = inputValue;
      }

      setDisplay(String(result));
      setPreviousValue(String(result));
    }

    setWaitingForOperand(true);
    setOperator(nextOperator);
  };

  const calculate = () => {
    if (!operator || previousValue === null) return;

    const inputValue = parseFloat(display);
    const currentValue = parseFloat(previousValue);
    let result: number;

    switch (operator) {
      case "+":
        result = currentValue + inputValue;
        break;
      case "-":
        result = currentValue - inputValue;
        break;
      case "×":
        result = currentValue * inputValue;
        break;
      case "÷":
        result = inputValue !== 0 ? currentValue / inputValue : 0;
        break;
      default:
        result = inputValue;
    }

    setDisplay(String(result));
    setPreviousValue(null);
    setOperator(null);
    setWaitingForOperand(true);
  };

  const toggleSign = () => {
    setDisplay(String(parseFloat(display) * -1));
  };

  const percentage = () => {
    setDisplay(String(parseFloat(display) / 100));
  };

  return (
    <main className="min-h-screen flex items-center justify-center bg-neutral-950 p-4">
      <div className="w-full max-w-xs bg-neutral-900 rounded-3xl p-4 shadow-2xl">
        {/* Display */}
        <div className="h-24 flex items-end justify-end px-4 pb-2 mb-4">
          <span className="text-5xl font-light text-white truncate">
            {display.length > 9
              ? parseFloat(display).toExponential(4)
              : display}
          </span>
        </div>

        {/* Button Grid */}
        <div className="grid grid-cols-4 gap-3">
          {/* Row 1 */}
          <Button
            variant="secondary"
            className="h-16 text-xl font-medium rounded-full bg-neutral-500 hover:bg-neutral-400 text-black"
            onClick={clear}
          >
            AC
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-xl font-medium rounded-full bg-neutral-500 hover:bg-neutral-400 text-black"
            onClick={toggleSign}
          >
            +/-
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-xl font-medium rounded-full bg-neutral-500 hover:bg-neutral-400 text-black"
            onClick={percentage}
          >
            %
          </Button>
          <Button
            className={`h-16 text-2xl font-medium rounded-full ${
              operator === "÷"
                ? "bg-white text-orange-500"
                : "bg-orange-500 hover:bg-orange-400 text-white"
            }`}
            onClick={() => performOperation("÷")}
          >
            ÷
          </Button>

          {/* Row 2 */}
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("7")}
          >
            7
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("8")}
          >
            8
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("9")}
          >
            9
          </Button>
          <Button
            className={`h-16 text-2xl font-medium rounded-full ${
              operator === "×"
                ? "bg-white text-orange-500"
                : "bg-orange-500 hover:bg-orange-400 text-white"
            }`}
            onClick={() => performOperation("×")}
          >
            ×
          </Button>

          {/* Row 3 */}
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("4")}
          >
            4
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("5")}
          >
            5
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("6")}
          >
            6
          </Button>
          <Button
            className={`h-16 text-2xl font-medium rounded-full ${
              operator === "-"
                ? "bg-white text-orange-500"
                : "bg-orange-500 hover:bg-orange-400 text-white"
            }`}
            onClick={() => performOperation("-")}
          >
            −
          </Button>

          {/* Row 4 */}
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("1")}
          >
            1
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("2")}
          >
            2
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={() => inputDigit("3")}
          >
            3
          </Button>
          <Button
            className={`h-16 text-2xl font-medium rounded-full ${
              operator === "+"
                ? "bg-white text-orange-500"
                : "bg-orange-500 hover:bg-orange-400 text-white"
            }`}
            onClick={() => performOperation("+")}
          >
            +
          </Button>

          {/* Row 5 */}
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white col-span-2"
            onClick={() => inputDigit("0")}
          >
            0
          </Button>
          <Button
            variant="secondary"
            className="h-16 text-2xl font-medium rounded-full bg-neutral-700 hover:bg-neutral-600 text-white"
            onClick={inputDecimal}
          >
            .
          </Button>
          <Button
            className="h-16 text-2xl font-medium rounded-full bg-orange-500 hover:bg-orange-400 text-white"
            onClick={calculate}
          >
            =
          </Button>
        </div>
      </div>
    </main>
  );
}
