package main

import (
	"strings"
	"testing"
)

type MockScanner struct {
	lines []string
	index int
}

func (m *MockScanner) Scan() string {
	if m.index >= len(m.lines) {
		return ""
	}
	line := m.lines[m.index]
	m.index++
	return line
}

type MockWriter struct {
	Messages []string
}

func (m *MockWriter) Write(message string) {
	m.Messages = append(m.Messages, message)
}

func TestAddition(t *testing.T) {
	t.Run("calculate addition in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"10", "+", "5"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 10.00 + 5.00 = 15.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("calculate addition in cli mode", func(t *testing.T) {
		args := []string{"", "10", "+", "5"}
		writer := &MockWriter{}

		CLIMode(args, writer)

		expected := "Result: 10.00 + 5.00 = 15.00"
		assertContains(t, writer.Messages, expected)
	})
}

func TestSubtraction(t *testing.T) {
	t.Run("calculate subtraction in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"20", "-", "8"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 20.00 - 8.00 = 12.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("calculate subtraction in cli mode", func(t *testing.T) {
		args := []string{"", "20", "-", "8"}
		writer := &MockWriter{}

		CLIMode(args, writer)

		expected := "Result: 20.00 - 8.00 = 12.00"
		assertContains(t, writer.Messages, expected)
	})
}

func TestMultiplication(t *testing.T) {
	t.Run("calculate multiplication in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"3", "*", "7"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 3.00 * 7.00 = 21.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("calculate multiplication in cli mode", func(t *testing.T) {
		args := []string{"", "3", "*", "7"}
		writer := &MockWriter{}

		CLIMode(args, writer)

		expected := "Result: 3.00 * 7.00 = 21.00"
		assertContains(t, writer.Messages, expected)
	})
}

func TestDivision(t *testing.T) {
	t.Run("calculate division in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"15", "/", "3"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 15.00 / 3.00 = 5.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("calculate division in cli mode", func(t *testing.T) {
		args := []string{"", "15", "/", "3"}
		writer := &MockWriter{}

		CLIMode(args, writer)

		expected := "Result: 15.00 / 3.00 = 5.00"
		assertContains(t, writer.Messages, expected)
	})
}

func TestDivisionByZero(t *testing.T) {
	t.Run("division by zero in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"10", "/", "0"}}
		writer := &MockWriter{}

		err := interactiveMode(scanner, writer)

		expected := ErrorDivisionByZero
		assertError(t, err, expected)
	})

	t.Run("division by zero in cli mode", func(t *testing.T) {
		args := []string{"", "10", "/", "0"}
		writer := &MockWriter{}

		err := CLIMode(args, writer)

		expected := ErrorDivisionByZero
		assertError(t, err, expected)
	})
}

func TestInvalidOperator(t *testing.T) {
	t.Run("invalid operator in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"5", "x", "2"}}
		writer := &MockWriter{}

		err := interactiveMode(scanner, writer)

		expected := InvalidOperator
		assertError(t, err, expected)
	})

	t.Run("invalid operator in cli mode", func(t *testing.T) {
		args := []string{"", "5", "x", "2"}
		writer := &MockWriter{}

		err := CLIMode(args, writer)

		expected := InvalidOperator
		assertError(t, err, expected)
	})
}

func TestInvalidNumber(t *testing.T) {
	t.Run("invalid number in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"ten", "+", "5"}}
		writer := &MockWriter{}

		err := interactiveMode(scanner, writer)

		expected := InvalidNumber
		assertError(t, err, expected)
	})

	t.Run("invalid number in cli mode", func(t *testing.T) {
		args := []string{"", "ten", "+", "5"}
		writer := &MockWriter{}

		err := CLIMode(args, writer)

		expected := InvalidNumber
		assertError(t, err, expected)
	})
}

func TestWhitespaceInput(t *testing.T) {
	t.Run("whitespace input in interactive mode", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"  30  ", "-", "  10  "}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 30.00 - 10.00 = 20.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("whitespace input in cli mode", func(t *testing.T) {
		args := []string{"", "  30  ", "-", "  10  "}
		writer := &MockWriter{}

		CLIMode(args, writer)

		expected := "Result: 30.00 - 10.00 = 20.00"
		assertContains(t, writer.Messages, expected)
	})
}

func assertContains(t *testing.T, messages []string, expected string) {
	t.Helper()

	found := false
	for _, msg := range messages {
		if strings.Contains(msg, expected) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected output to contain: %q, but got %q", expected, messages)
	}
}

func assertError(t *testing.T, err, expected error) {
	t.Helper()

	if err != expected {
		t.Errorf("got error %q want %q", err, expected)
	}
}
