package UglyNumber

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	factor := []int{2, 3, 5}

	for _, f := range factor {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}
