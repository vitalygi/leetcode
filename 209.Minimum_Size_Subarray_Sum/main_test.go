package main

import (
	"math"
	"testing"
)

func minSubArrayLen(target int, nums []int) int {
	minLength := math.MaxInt32
	begin := 0
	curSum := 0
	curLength := 0
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		curLength += 1
		for curSum >= target {
			if curLength < minLength {
				minLength = curLength
			}
			curSum -= nums[begin]
			curLength -= 1
			begin += 1
		}
	}
	if minLength == math.MaxInt32 {
		return 0
	} else {
		return minLength
	}
}

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "base test",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			name:     "len is one",
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1,
		},
		{
			name:     "answer is zero",
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minSubArrayLen(test.target, test.nums)
			if result != test.expected {
				t.Errorf("Test failed. Expected %v, got %v", test.expected, result)
			}
		})
	}
}
