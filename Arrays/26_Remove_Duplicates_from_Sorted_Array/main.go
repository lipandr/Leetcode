package RemoveDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var unique = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[unique] = nums[i]
			unique++
		}
	}
	return unique
}
