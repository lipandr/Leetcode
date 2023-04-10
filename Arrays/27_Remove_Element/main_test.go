package RemoveElement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		inputNums []int
		inputVal  int
		want      int
		wantNums  []int
	}{
		{[]int{}, 3, 0, []int{}},
		{[]int{3}, 3, 0, []int{}},
		{[]int{1, 2, 3}, 4, 3, []int{1, 2, 3}},
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}

	for _, tc := range tests {
		got := removeElement(tc.inputNums, tc.inputVal)
		if got != tc.want {
			t.Errorf("removeElement(%v, %v) = %v; want %v", tc.inputNums, tc.inputVal, got, tc.want)
		}
		if !reflect.DeepEqual(tc.inputNums[:tc.want], tc.wantNums) {
			t.Errorf("removeElement(%v, %v) ; \n- nums: %v;     want: %v", tc.inputNums, tc.inputVal, tc.inputNums[:tc.want], tc.wantNums)
		}
	}
}
