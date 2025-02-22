package main

import (
	"fmt"
)

func calculate(x float64, y float64, op string) (result float64, err error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("error: division by zero is not allowed")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("error: invalid operator %q", op)
	}
}

func interactiveMode() {
	var x, y float64
	var op string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Invalid input! Please enter a valid number.")
		return
	}

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&y)
	if err != nil {
		fmt.Println("Invalid input! Please enter a valid number.")
		return
	}

	result, err := calculate(x, y, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", x, op, y, result)
}

func main() {
	interactiveMode()
}
