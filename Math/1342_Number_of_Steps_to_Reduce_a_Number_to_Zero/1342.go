package main

func NumberOfSteps(num int) int {
	var steps int

	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}
		steps++
	}
	return steps
}

// Time Complexity O(log n)
// Space Complexity O(1)
