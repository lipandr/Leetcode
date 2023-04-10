package RemoveDuplicatesFromSortedArrayII

func removeDuplicates(nums []int) int {
	i := 0
	for _, j := range nums {
		if i == 0 || i == 1 || nums[i-2] != j {
			nums[i] = j
			i += 1
		}
	}
	return i
}
