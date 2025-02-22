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

func TestInteractiveMode(t *testing.T) {
	t.Run("calculate addition", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"10", "+", "5"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 10.00 + 5.00 = 15.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("calculate subtraction", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"20", "-", "8"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 20.00 - 8.00 = 12.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("calculate multiplication", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"3", "*", "7"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 3.00 * 7.00 = 21.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("calculate division", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"15", "/", "3"}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

		expected := "Result: 15.00 / 3.00 = 5.00"
		assertContains(t, writer.Messages, expected)
	})

	t.Run("division by zero", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"10", "/", "0"}}
		writer := &MockWriter{}

		err := interactiveMode(scanner, writer)

		expected := ErrorDivisionByZero
		assertError(t, err, expected)
	})

	t.Run("invalid operator", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"5", "x", "2"}}
		writer := &MockWriter{}

		err := interactiveMode(scanner, writer)

		expected := InvalidOperator
		assertError(t, err, expected)
	})

	t.Run("invalid number", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"ten", "+", "5"}}
		writer := &MockWriter{}

		err := interactiveMode(scanner, writer)

		expected := InvalidNumber
		assertError(t, err, expected)
	})

	t.Run("whitespace input", func(t *testing.T) {
		scanner := &MockScanner{lines: []string{"  30  ", "-", "  10  "}}
		writer := &MockWriter{}

		interactiveMode(scanner, writer)

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
