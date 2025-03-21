package _4__Search_a_2D_Matrix

import (
	"fmt"
	"slices"
	"testing"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	l, r := 0, len(matrix)-1
	for l <= r {
		m := l + (r-l)/2
		if matrix[m][0] <= target && target <= matrix[m][len(matrix[0])-1] {
			_, isFound := slices.BinarySearch(matrix[m], target)
			return isFound
		} else if matrix[m][0] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name        string
		inputMatrix [][]int
		inputTarget int
		expected    bool
	}{
		{
			name:        "basic test",
			inputMatrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			inputTarget: 3,
			expected:    true,
		},
		{
			name:        "without target",
			inputMatrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			inputTarget: 13,
			expected:    false,
		},
		{
			name:        "one element",
			inputMatrix: [][]int{{1}},
			inputTarget: 2,
			expected:    false,
		},
		{
			name:        "two elements",
			inputMatrix: [][]int{{1}, {3}},
			inputTarget: 2,
			expected:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			testResult := searchMatrix(test.inputMatrix, test.inputTarget)
			fmt.Println(test.name, testResult)
			if testResult != test.expected {
				t.Errorf("Error. Expected: %v, got %v \nInput Matrix:%v\nTarger:%v", test.expected, testResult, test.inputMatrix, test.inputTarget)
			}
		})
	}
}
