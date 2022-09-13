package main

import "fmt"

func RunningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func main() {
	fmt.Println(RunningSum([]int{1, 2, 3, 4}))
	fmt.Println(RunningSum([]int{1, 1, 1, 1, 1}))
}
