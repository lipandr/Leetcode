package LongestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"a", "ab", "abc"}, "a"},
	}

	for _, tc := range testCases {
		if output := longestCommonPrefix(tc.input); output != tc.expected {
			t.Errorf("Expected %s but got %s", tc.expected, output)
		}
	}
}
