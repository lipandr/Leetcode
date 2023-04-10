package IntegertoRoman

import "strings"

func intToRoman(num int) string {
	var result strings.Builder
	integers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(integers); i++ {
		for num >= integers[i] {
			result.WriteString(romans[i])
			num -= integers[i]
		}
	}
	return result.String()
}

func intToRomanHashMap(num int) string {
	var result strings.Builder
	romanMap := map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}

	for integer, roman := range romanMap {
		for num >= integer {
			num -= integer
			result.WriteString(roman)
		}
	}
	return result.String()
}
