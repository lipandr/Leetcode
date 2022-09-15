package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	testName string
	input    int
	want     []string
}{
	{
		"Example 1",
		3,
		[]string{"1", "2", "Fizz"}},
	{
		"Example 2",
		5,
		[]string{"1", "2", "Fizz", "4", "Buzz"}},
	{
		"Example 3",
		15,
		[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
}

func TestFizzBuzz(t *testing.T) {

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			got := FizzBuzz(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approach 1 - %#v got %v want %v", tt.testName, got, tt.want)
			}
			got = FizzBuzzApproach2(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approach 2 - %#v got %v want %v", tt.testName, got, tt.want)
			}
			got = FizzBuzzApproach3(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Approach 3 - %#v got %v want %v", tt.testName, got, tt.want)
			}
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			FizzBuzzApproach3(tc.input)
		}
	}
}
