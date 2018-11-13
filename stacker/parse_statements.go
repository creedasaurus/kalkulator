package stacker

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/scanner"
)

// ProcessStatement will take a string representing some basic algebra and calculate it
func ProcessStatement(statement string) (float64, error) {
	operatorStack := &stack{arr: make([]string, 0, 10)}
	operandStack := &floatStack{arr: make([]float64, 0, 10)}

	sReader := strings.NewReader(statement)

	sScanner := createScanner(sReader, "input")

	lastToken := ""
	numMod := ""
	for token := sScanner.Scan(); token != scanner.EOF; token = sScanner.Scan() {
		tokenString := sScanner.TokenText()

		/**
		Operator
		**/
		if _, ok := opPrecedence[tokenString]; ok {
			// if an op that's not + or -, just add it and continue
			if tokenString != "+" && tokenString != "-" {
				// operatorStack.Push(tokenString)
				if err := processOperator(tokenString, operatorStack, operandStack); err != nil {
					return 0, err
				}
				lastToken = tokenString
				continue
			}

			// if + or -, check if lastToken was also an operator
			if _, ok := opPrecedence[lastToken]; !ok || lastToken == ")" { // if not, go ahead
				if err := processOperator(tokenString, operatorStack, operandStack); err != nil {
					return 0, err
				}
				lastToken = tokenString
				continue
			}

			// if it was an operator, dont add, but set the numMod
			numMod = tokenString
			lastToken = tokenString
			continue
		}

		/**
		Operand
		**/
		val, err := strconv.ParseFloat(numMod+tokenString, 64)
		if err != nil {
			return 0, err
		}

		operandStack.Push(val)
		lastToken = tokenString
		numMod = ""
	}

	// Process remaining operations
	for {
		if operatorStack.size == 0 {
			break
		}

		remainingOp := operatorStack.Pop()
		if err := performCalculation(remainingOp, operandStack); err != nil {
			return 0, err
		}
	}

	// if there are operands remaining, something went wrong
	if operandStack.size != 1 {
		return 0, fmt.Errorf("error computing value")
	}

	calculated := operandStack.Pop()
	return calculated, nil
}

func createScanner(r io.Reader, filename string) *scanner.Scanner {
	debug := false

	var s scanner.Scanner
	s.Init(r)
	s.Filename = filename

	s.Error = func(s *scanner.Scanner, msg string) {
		if debug {
			fmt.Println(msg)
		}
	}
	s.Mode = scanner.ScanComments | scanner.ScanFloats
	return &s
}
