package Permutation_in_String

import "testing"

func TestCheckInclusion(t *testing.T) {
	testCases := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"abc", "cbaebabacd", true},
		{"abcd", "dcba", true},
		{"a", "a", true},
		{"a", "b", false},
	}

	for _, tc := range testCases {
		got := checkInclusion(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("checkInclusion(%v, %v) = %v; want %v", tc.s1, tc.s2, got, tc.want)
		}
	}
}
