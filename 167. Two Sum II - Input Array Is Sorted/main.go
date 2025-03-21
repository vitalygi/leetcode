package main

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; true; {
		ijsum := numbers[i] + numbers[j]
		if ijsum == target {
			return []int{i + 1, j + 1}
		} else if ijsum > target {
			j--
			continue
		} else if ijsum < target {
			i++
			continue
		}
	}
	return []int{}
}
