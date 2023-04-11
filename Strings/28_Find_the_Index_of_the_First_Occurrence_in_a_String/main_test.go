package FindTheIndexOfTheFirstOccurrenceInAString

import "testing"

func TestStrStr(t *testing.T) {
	cases := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}

	for _, c := range cases {
		actual := strStr(c.haystack, c.needle)
		if actual != c.expected {
			t.Errorf("strStr(%q, %q) == %d, expected %d", c.haystack, c.needle, actual, c.expected)
		}
	}
}
