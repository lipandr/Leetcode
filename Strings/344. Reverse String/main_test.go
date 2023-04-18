package ReverseString

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("world"), []byte("dlrow")},
		{[]byte("12345"), []byte("54321")},
		{[]byte(""), []byte("")},
	}

	for _, test := range tests {
		inputCopy := make([]byte, len(test.input))
		copy(inputCopy, test.input) // create a copy of the input slice to avoid modifying it
		reverseString(inputCopy)
		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("Expected %v but got %v", test.expected, inputCopy)
		}
	}
}
