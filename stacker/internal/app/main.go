package main

import (
	"fmt"
	"github.com/creedasaurus/kalkulator/stacker"
)

func main() {
	val, err := stacker.ProcessStatement("1+3 + 4")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%.4f\n", val)
	val, err = stacker.ProcessStatement("1++3 + -4 / 8.4")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%.4f\n", val)

	val, err = stacker.ProcessStatement("55 / 5 ^ 3 + (6 / 3)")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%.4f\n", val)

	val, err = stacker.ProcessStatement("5 * (4 /   (1 / 2)) + 99 / 1")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%.4f\n", val)
}
