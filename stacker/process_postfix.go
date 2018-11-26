package stacker

import (
	"math"
)

var opPrecedence = map[string]int{
	"^": 3,
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"(": 10,
	")": 0,
}

func performCalculation(operation string, operands *floatStack) error {
	// This is where we will begin creating functions to connect the services!
	val2 := operands.Pop()
	val1 := operands.Pop()

	var calculatedVal float64
	switch operation {
	case "+":
		calculatedVal = val1 + val2
	case "-":
		calculatedVal = val1 - val2
	case "*":
		calculatedVal = val1 * val2
	case "/":
		calculatedVal = val1 / val2
	case "^":
		calculatedVal = math.Pow(val1, val2)
	}

	operands.Push(calculatedVal)
	return nil
}

func processOperator(nextOperator string, operators *stack, operands *floatStack) error {
	// stack is empty add operator
	if operators.size == 0 {
		operators.Push(nextOperator)
		return nil
	}

	nextPrecValue := opPrecedence[nextOperator]

	if nextOperator == ")" {
		for {
			if operators.Peek() == "(" {
				break
			}
			tmpOp := operators.Pop()
			if err := performCalculation(tmpOp, operands); err != nil {
				return err
			}
		}
		operators.Pop()
		return nil
	}

	for {
		if operators.size == 0 {
			break
		}

		top := operators.Peek()
		if top == "(" {
			break
		}

		topPrec := opPrecedence[top]
		if nextPrecValue > topPrec {
			break
		}

		tmpOp := operators.Pop()
		if err := performCalculation(tmpOp, operands); err != nil {
			return err
		}
	}

	operators.Push(nextOperator)

	return nil
}
