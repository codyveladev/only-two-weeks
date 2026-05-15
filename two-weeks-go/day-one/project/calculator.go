/*
* Date: May 14th, 2026
* Author: Cody Vela
* Description: A CLI calculator: pass two numbers and an operator as args (go run main.go 10 + 5),
* print the result. Handle divide-by-zero as a second return value.
 */
package main

import (
	"fmt"
	"os"
	"strconv"
)

func handleNumberInput(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("input must be a number: %w", err)
	}
	return num, nil
}

/*
* Known issues:
* - Integer Division
* - Cant handle float values
 */
func calculate(a int, b int, operator string) (int, error) {

	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	case "%":
		return a % b, nil
	default:
		return 0, fmt.Errorf("unknown operator provided")
	}
}

func main() {

	args := os.Args

	if len(args) != 4 {
		fmt.Println("Usage: go run main.go <num> <operator> <num>")
		return
	}
	a, err := handleNumberInput(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	operator := args[2]

	b, err := handleNumberInput(args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result is: \t%d", result)

}
