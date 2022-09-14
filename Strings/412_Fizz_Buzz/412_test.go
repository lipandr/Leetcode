package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
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
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			got := FizzBuzz(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v got %v want %v", tt.testName, got, tt.want)
			}
		})
	}
}
