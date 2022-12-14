package main

import "testing"

var TestCases = []struct {
	testName string
	input    [][]int
	want     int
}{
	{"Example 1",
		[][]int{{1, 2, 3}, {3, 2, 1}},
		6},
	{"Example 2",
		[][]int{{1, 5}, {7, 3}, {3, 5}},
		10},
	{"Example 3",
		[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
		17},
}

func TestMaximumWealth(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.testName, func(t *testing.T) {
			got := MaximumWealth(tt.input)
			if got != tt.want {
				t.Errorf("%#v got %v want %v", tt.testName, got, tt.want)
			}
		})
	}
}

func BenchmarkMaximumWealth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range TestCases {
			MaximumWealth(tt.input)
		}
	}
}
