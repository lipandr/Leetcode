package RemoveDuplicatesFromSortedArray

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	// Test case 1: empty input slice
	var nums []int
	if removeDuplicates(nums) != 0 {
		t.Errorf("Expected result to be 0 for empty input slice, but got %d", removeDuplicates(nums))
	}

	// Test case 2: no duplicates
	nums = []int{1, 2, 3, 4}
	if removeDuplicates(nums) != len(nums) {
		t.Errorf("Expected result to be %d for slice with no duplicates, but got %d", len(nums), removeDuplicates(nums))
	}

	// Test case 3: all elements are duplicates
	nums = []int{1, 1, 1, 1, 1}
	if removeDuplicates(nums) != 1 {
		t.Errorf("Expected result to be 1 for slice with all elements duplicates, but got %d", removeDuplicates(nums))
	}

	// Test case 4: general case with duplicates
	nums = []int{1, 2, 2, 3, 4, 4, 4, 5}
	expectedUnique := 5
	if removeDuplicates(nums) != expectedUnique {
		t.Errorf("Expected result to be %d for slice with duplicates, but got %d", expectedUnique, removeDuplicates(nums))
	}
}
