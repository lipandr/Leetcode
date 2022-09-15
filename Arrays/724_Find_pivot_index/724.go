package main

func PivotIndex(nums []int) int {
	var leftSum, sum int
	for _, v := range nums {
		sum += v
	}
	for i := 0; i < len(nums); i++ {
		if leftSum == sum-nums[i]-leftSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

// Time Complexity = O(n)
// Space complexity = O(1)
