package RomanToInteger

import "testing"

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"single char", "I", 1},
		{"different single char", "V", 5},
		{"multiple same char", "III", 3},
		{"multiple different char", "XIV", 14},
		{"subtraction case", "IV", 4},
		{"complex case", "MCMLXXVII", 1977},
		{"complex case", "MCMXCIV", 1994},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := romanToInt(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, actual)
			}
		})
	}
}
