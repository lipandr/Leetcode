package main

import "strconv"

func FizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		div3 := i%3 == 0
		div5 := i%5 == 0
		if div3 && div5 {
			result = append(result, "FizzBuzz")
		} else if div5 {
			result = append(result, "Buzz")
		} else if div3 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// Time Complexity = O(n)
// Space Complexity = O(1)
