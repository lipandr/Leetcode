package SquaresOfASortedArray

import "math"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	k := len(nums) - 1
	for k >= 0 {
		if abs(nums[left]) > abs(nums[right]) {
			res[k] = nums[left] * nums[left]
			k--
			left++
		} else {
			res[k] = nums[right] * nums[right]
			k--
			right--
		}
	}
	return res
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
