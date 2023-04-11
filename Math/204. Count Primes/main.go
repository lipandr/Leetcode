package CountPrimes

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	primes := make([]bool, n)
	for i := range primes {
		primes[i] = true
	}
	primes[0], primes[1] = false, false

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}
	count := 0
	for _, prime := range primes {
		if prime {
			count++
		}
	}
	return count
}
