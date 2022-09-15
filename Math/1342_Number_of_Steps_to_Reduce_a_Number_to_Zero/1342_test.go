package main

import "testing"

var testCases = []struct {
	name  string
	input int
	want  int
}{
	{"Example 1", 14, 6},
	{"Example 2", 8, 4},
	{"Example 3", 123, 12},
}

func TestNumberOfSteps(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberOfSteps(tt.input)
			if got != tt.want {
				t.Errorf("%#v got %d want %d", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkNumberOfSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			NumberOfSteps(tt.input)
		}
	}
}
