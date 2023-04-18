package SquaresOfASortedArray

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "all positive numbers",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "all negative numbers",
			nums: []int{-4, -3, -2, -1},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "mixed positive and negative numbers",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
