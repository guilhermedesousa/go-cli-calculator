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

type Writer interface {
	Write(message string)
}

type DefaultScanner struct {
	scanner *bufio.Scanner
}

func (d *DefaultScanner) Scan() string {
	d.scanner.Scan()
	return d.scanner.Text()
}

type DefaultWriter struct{}

func (d *DefaultWriter) Write(message string) {
	fmt.Print(message)
}

func parseInput(input string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(input), 64)
}

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrorDivisionByZero = DictionaryErr("Division by zero is not allowed")
	InvalidOperator     = DictionaryErr("Invalid input! Please enter a valid operator")
	InvalidNumber       = DictionaryErr("Invalid input! Please enter a valid number")
)

func calculate(x float64, y float64, op string) (float64, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, ErrorDivisionByZero
		}
		return x / y, nil
	default:
		return 0, InvalidOperator
	}
}

func interactiveMode(scanner Scanner, writer Writer) error {
	writer.Write("Enter first number: ")
	x, err1 := parseInput(scanner.Scan())

	writer.Write("Enter an operator (+, -, *, /): ")
	op := scanner.Scan()

	writer.Write("Enter second number: ")
	y, err2 := parseInput(scanner.Scan())

	if err1 != nil || err2 != nil {
		writer.Write(InvalidNumber.Error() + "\n")
		return InvalidNumber
	}

	result, err := calculate(x, y, op)
	if err != nil {
		writer.Write(err.Error() + "\n")
		return err
	}

	writer.Write(fmt.Sprintf("Result: %.2f %s %.2f = %.2f\n", x, op, y, result))
	return nil
}

func CLIMode(args []string, writer Writer) error {
	x, err1 := strconv.ParseFloat(strings.TrimSpace(args[1]), 64)
	op := args[2]
	y, err2 := strconv.ParseFloat(strings.TrimSpace(args[3]), 64)

	if err1 != nil || err2 != nil {
		return InvalidNumber
	}

	result, err := calculate(x, y, op)
	if err != nil {
		writer.Write(err.Error() + "\n")
		return err
	}

	writer.Write(fmt.Sprintf("Result: %.2f %s %.2f = %.2f\n", x, op, y, result))
	return nil
}

func main() {
	defaultScanner := &DefaultScanner{scanner: bufio.NewScanner(os.Stdin)}
	defaultWriter := &DefaultWriter{}

	if len(os.Args) == 4 {
		CLIMode(os.Args, defaultWriter)
	} else {
		interactiveMode(defaultScanner, defaultWriter)
	}
}
