package TwoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("twoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func TestTwoSumEmpty(t *testing.T) {
	got := twoSum([]int{}, 0)
	if got != nil {
		t.Errorf("twoSum({}, 0) = %v, want nil", got)
	}
}

func TestTwoSumNoSolution(t *testing.T) {
	got := twoSum([]int{2, 3, 5, 7}, 11)
	if got != nil {
		t.Errorf("twoSum(%v, %v) = %v, want nil", []int{2, 3, 5, 7}, 11, got)
	}
}
