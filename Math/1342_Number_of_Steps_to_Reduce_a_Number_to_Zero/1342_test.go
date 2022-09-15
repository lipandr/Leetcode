package main

import "testing"

func TestNumberOfSteps(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  int
	}{
		{"Example 1", 14, 6},
		{"Example 2", 8, 4},
		{"Example 3", 123, 12},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberOfSteps(tt.input)
			if got != tt.want {
				t.Errorf("%#v got %d want %d", tt.name, got, tt.want)
			}
		})
	}
}
