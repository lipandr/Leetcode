package main

func MaximumWealth(accounts [][]int) int {
	var max int
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// Time Complexity = O(n x m)
// Space complexity = O(1)
