package RotateArray

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			nums:     []int{0},
			k:        0,
			expected: []int{0},
		},
		{
			nums:     []int{1, 2, 3},
			k:        4,
			expected: []int{3, 1, 2},
		},
	}

	for i, test := range tests {
		rotate(test.nums, test.k)
		if !reflect.DeepEqual(test.nums, test.expected) {
			t.Errorf("Test %d: expected %v but got %v", i+1, test.expected, test.nums)
		}
	}
}
