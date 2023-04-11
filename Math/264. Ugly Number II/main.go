package UglyNumberII

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	var i2, i3, i5 int
	var next2, next3, next5 int

	for i := 1; i < n; i++ {
		next2 = ugly[i2] * 2
		next3 = ugly[i3] * 3
		next5 = ugly[i5] * 5

		ugly[i] = min(next2, min(next3, next5))

		if ugly[i] == next2 {
			i2++
		}
		if ugly[i] == next3 {
			i3++
		}
		if ugly[i] == next5 {
			i5++
		}
	}
	return ugly[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
