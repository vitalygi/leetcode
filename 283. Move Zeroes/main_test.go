package main

import (
	"reflect"
	"testing"
)

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

func moveZeroes(nums []int) {
	for i, j := 0, 0; j != len(nums); j++ {
		if nums[i] == 0 {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		} else {
			i++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "basic case",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "no zeros",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			moveZeroes(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Test case '%s' failed: expected %v, got %v. \nInput: %v", tt.name, tt.expected,
					tt.input, input)
			}
		})
	}
}
