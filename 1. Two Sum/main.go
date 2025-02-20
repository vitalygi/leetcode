package main

import (
	"slices"
)

// working with BinarySearch
func twoSum(nums []int, target int) []int {
	numsCopy := slices.Clone(nums)
	answer := make([]int, 0)
	slices.Sort(numsCopy)
	for i := 0; i != len(numsCopy)-1; i++ {
		currentItem := numsCopy[i]
		necessaryItemIndex, isFound := slices.BinarySearch(numsCopy[i:], target-currentItem)
		if isFound {
			for index, item := range nums {
				if item == numsCopy[necessaryItemIndex+i] || item == currentItem {
					answer = append(answer, index)
				}
				if len(answer) == 2 {
					return answer
				}

			}
		}
	}
	return nil
}
