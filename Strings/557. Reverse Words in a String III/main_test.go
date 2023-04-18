package ReverseWordsInAStringIII

import (
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"Hello World", "olleH dlroW"},
		{"The quick brown fox jumps over the lazy dog", "ehT kciuq nworb xof spmuj revo eht yzal god"},
		{"", ""},
	}

	for _, test := range tests {
		result := reverseWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected output %v but got %v", test.input, test.expected, result)
		}
	}
}
