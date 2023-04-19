package TwoSum_II_InputArrayIsSorted

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, tt := range tests {
		if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("twoSum(%v, %v) = %v, want %v", tt.numbers, tt.target, got, tt.want)
		}
	}
}
