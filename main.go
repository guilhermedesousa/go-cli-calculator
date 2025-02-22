package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scanner interface {
	Scan() string
}

type DefaultScanner struct {
	scanner *bufio.Scanner
}

func (d *DefaultScanner) Scan() string {
	d.scanner.Scan()
	return d.scanner.Text()
}

func parseInput(input string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(input), 64)
}

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

func interactiveMode(scanner Scanner) {
	fmt.Print("Enter first number: ")
	x, err1 := parseInput(scanner.Scan())
	if err1 != nil {
		fmt.Println("Invalid input! Please enter a valid number.")
		return
	}

	fmt.Print("Enter an operator (+, -, *, /): ")
	op := scanner.Scan()

	fmt.Print("Enter second number: ")
	y, err2 := parseInput(scanner.Scan())
	if err2 != nil {
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

func CLIMode() {
	x, err1 := strconv.ParseFloat(os.Args[1], 64)
	op := os.Args[2]
	y, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
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
	defaultScanner := &DefaultScanner{scanner: bufio.NewScanner(os.Stdin)}
	if len(os.Args) == 4 {
		CLIMode()
	} else {
		interactiveMode(defaultScanner)
	}
}
