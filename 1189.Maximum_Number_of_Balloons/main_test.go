package main

import (
	"testing"
)

func maxNumberOfBalloons(text string) int {
	letters := make(map[rune]int, 26)
	for _, ch := range text {
		letters[ch] += 1
	}
	return min(letters['b'], letters['a'], letters['l']/2, letters['o']/2, letters['n'])
}

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"base test",
			"nlaebolko",
			1,
		},
		{
			"2-nd base test",
			"loonbalxballpoon",
			2,
		},
		{
			"expected 0",
			"leetcode",
			0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxNumberOfBalloons(test.input)
			if result != test.expected {
				t.Errorf("Test failed. Expected %v, got %v", test.expected, result)
			}
		})
	}
}
