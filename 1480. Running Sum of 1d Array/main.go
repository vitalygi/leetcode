package main

func runningSum(nums []int) []int {
	runningSumArray := make([]int, len(nums))
	runningSumArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		runningSumArray[i] = runningSumArray[i-1] + nums[i]
	}
	return runningSumArray
}
