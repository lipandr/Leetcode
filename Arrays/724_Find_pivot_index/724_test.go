package main

import "testing"

var testCases = []struct {
	name  string
	input []int
	want  int
}{
	{"Example 1", []int{1, 7, 3, 6, 5, 6}, 3},
	{"Example 2", []int{1, 2, 3}, -1},
	{"Example 3", []int{2, 1, -1}, 0},
}

func TestPivotIndex(t *testing.T) {

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := PivotIndex(tt.input)
			if got != tt.want {
				t.Errorf("%#v got %v want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkPivotIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			PivotIndex(tt.input)
		}
	}
}
