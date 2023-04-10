package ValidParentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"[", false},
		{"([{}])", true},
		{"((()))", true},
		{"(()))", false},
		{"(()))(", false},
	}

	for _, tc := range tests {
		got := isValid(tc.input)
		if got != tc.want {
			t.Errorf("isValid(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
