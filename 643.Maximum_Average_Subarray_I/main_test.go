package main

import (
	"testing"
)

func countWindowSum(nums []int, k int) int {
	sum := 0

	return sum
}

func findMaxAverage(nums []int, k int) float64 {
	maxAverage := 0
	for i := 0; i < k; i++ {
		maxAverage += nums[i]
	}

	curSum := maxAverage
	for i := k; i < len(nums); i++ {
		curSum += nums[i]
		curSum -= nums[i-k]
		if curSum > maxAverage {
			maxAverage = curSum
		}
	}
	return float64(maxAverage) / float64(k)
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{
			name:     "base test",
			nums:     []int{0, 1, 1, 3, 3},
			k:        4,
			expected: 2,
		},
		{
			name:     "one element",
			nums:     []int{5},
			k:        1,
			expected: 5,
		},
		{
			name:     "one negative element",
			nums:     []int{-1},
			k:        1,
			expected: -1,
		},
		{
			name:     "k equals len of array",
			nums:     []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891},
			k:        93,
			expected: -594.5806451612904,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findMaxAverage(test.nums, test.k)
			if result != test.expected {
				t.Errorf("Failed. Expected %v, got %v", test.expected, result)
			}
		})
	}
}
