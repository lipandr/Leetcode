package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name  string
	input []int
	want  []int
}{
	{"Example 1", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
	{"Example 2", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
	{"Example 3", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
}

func TestRunningSum(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := RunningSum(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v got %v want %v", tt.name, tt.input, tt.want)
			}
		})
	}
}

func BenchmarkRunningSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			RunningSum(tt.input)
		}
	}
}
