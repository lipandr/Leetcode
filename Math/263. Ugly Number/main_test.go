package UglyNumber

import "testing"

func TestIsUgly(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{0, false},  // n is 0
		{-3, false}, // n is negative
		{1, true},   // n is 1
		{2, true},   // n is only divisible by 2
		{3, true},   // n is only divisible by 3
		{5, true},   // n is only divisible by 5
		{6, true},   // n is divisible by 2 and 3
		{10, true},  // n is divisible by 2 and 5
		{14, false}, // n is not divisible by 2, 3 or 5
		{15, true},  // n is divisible by 3 and 5
	}
	for _, c := range cases {
		got := isUgly(c.in)
		if got != c.out {
			t.Errorf("isUgly(%v) == %v, want %v", c.in, got, c.out)
		}
	}
}
