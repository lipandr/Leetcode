package SearchInsertPosition

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{}, 5, 0},
	}
	for _, c := range cases {
		got := searchInsert(c.nums, c.target)
		if got != c.want {
			t.Errorf("searchInsert(%v, %d) == %d, want %d", c.nums, c.target, got, c.want)
		}
	}
}
