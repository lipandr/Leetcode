package CountPrimes

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{20, 8},
		{10, 4},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 2},
		{5, 2},
	}

	for _, tc := range tests {
		got := countPrimes(tc.input)
		if got != tc.want {
			t.Errorf("countPrimes(%d) = %d; want %d", tc.input, got, tc.want)
		}
	}
}
