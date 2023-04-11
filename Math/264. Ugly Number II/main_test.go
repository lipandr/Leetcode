package UglyNumberII

import "testing"

func TestNthUglyNumber(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{10, 12},
		{11, 15},
	}

	for _, tc := range testCases {
		got := nthUglyNumber(tc.n)
		if got != tc.expected {
			t.Errorf("Expected nthUglyNumber(%d) to be %d, but got %d", tc.n, tc.expected, got)
		}
	}
}
