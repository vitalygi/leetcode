package main

import (
	"testing"
)

func isValid(s string) bool {
	stack := make([]rune, 0, 10)
	brackets := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, r := range s {
		openBracket, iF := brackets[r]
		if iF {
			if stack[len(stack)-1] == openBracket {
				stack = append(stack[:len(stack)-1])
			} else {
				return false
			}

		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			"base test",
			"()",
			true,
		},
		{
			"base test",
			"()[]{}",
			true,
		},
		{
			"false test",
			"(]",
			false,
		},
		{
			"inned brackets",
			"([])",
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isValid(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
