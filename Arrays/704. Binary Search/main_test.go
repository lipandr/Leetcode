package BinarySearch

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	// search for an existing value
	idx := search(nums, 9)
	if idx != 4 {
		t.Errorf("Expected index 2 but got %d", idx)
	}

	// search for a non-existing value
	idx = search(nums, 2)
	if idx != -1 {
		t.Errorf("Expected index -1 but got %d", idx)
	}
}
