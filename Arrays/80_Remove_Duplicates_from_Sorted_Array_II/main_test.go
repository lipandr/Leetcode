package RemoveDuplicatesFromSortedArrayII

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		output int
		result []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1}, 2, []int{1, 1}},
		{[]int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3}},
		{[]int{1, 2, 3, 3, 3, 3, 4, 4}, 6, []int{1, 2, 3, 3, 4, 4}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3}},
		{[]int{}, 0, []int{}},
	}

	for _, test := range tests {
		output := removeDuplicates(test.input)
		if output != test.output {
			t.Errorf("Test Failed: input = %v, expected = %d, output = %d", test.input, test.output, output)
		}
		if !reflect.DeepEqual(test.input[:test.output], test.result) {
			t.Errorf("Test Failed: input = %v, expected = %v, output = %v", test.input, test.result, test.input[:test.output])
		}
	}
}
