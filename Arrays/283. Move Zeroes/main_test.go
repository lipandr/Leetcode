package MoveZeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	t.Run("moves zeroes to end of array", func(t *testing.T) {
		input := []int{0, 1, 0, 3, 12}
		expected := []int{1, 3, 12, 0, 0}
		moveZeroes(input)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Expected %v but got %v", expected, input)
		}
	})

	t.Run("does not modify array without zeroes", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		moveZeroes(input)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Expected %v but got %v", expected, input)
		}
	})

	t.Run("handles empty array", func(t *testing.T) {
		var input []int
		var expected []int
		moveZeroes(input)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Expected %v but got %v", expected, input)
		}
	})
}
