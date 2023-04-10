package IntegertoRoman

import "testing"

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1977, "MCMLXXVII"},
		{1994, "MCMXCIV"},
	}

	for _, tc := range testCases {
		result := intToRoman(tc.input)
		if result != tc.expected {
			t.Errorf("Expected intToRoman(%d) to be %s, but got %s", tc.input, tc.expected, result)
		}
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	testCases := []struct {
		name string
		num  int
	}{
		{name: "small", num: 58},
		{name: "medium", num: 1994},
		{name: "large", num: 3999},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					intToRoman(tc.num)
				}
			})
		})
	}
}

func BenchmarkIntToRomanHashMap(b *testing.B) {
	testCases := []struct {
		name string
		num  int
	}{
		{name: "small", num: 58},
		{name: "medium", num: 1994},
		{name: "large", num: 3999},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					intToRomanHashMap(tc.num)
				}
			})
		})
	}
}
