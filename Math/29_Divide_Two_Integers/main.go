package DivideTwoIntegers

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}
	result := 0
	for dividend >= divisor {
		shift := 0
		for dividend >= (divisor << shift) {
			shift++
		}
		dividend -= divisor << (shift - 1)
		result += 1 << (shift - 1)
	}
	return sign * result
}
