package TwoSum

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)
	for i, num := range nums {
		if j, ok := diffs[num]; ok {
			return []int{j, i}
		}
		diffs[target-num] = i
	}
	return []int{}
}

// Time Complexity O(n)
// Space Complexity O(n)
