package main

import (
	"fmt"
)

func doCalculate(num1 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = num1 + num1
	case "-":
		result = num1 - num1
	case "*":
		result = num1 * num1
	case "/":
		if num1 != 0 {
			result = num1 / num1
		} else {
			fmt.Println("Error in the code: Number is divided by zero")
		}
	default:
		fmt.Println("Invalid operator provided")
	}

	return result
}

func main() {
	var result float64
	result = doCalculate(100, "+")
	fmt.Println(result)
	result = doCalculate(50, "-")
	fmt.Println(result)
	result = doCalculate(20, "/")
	fmt.Println(result)
	result = doCalculate(10, "*")
	fmt.Println(result)
}
