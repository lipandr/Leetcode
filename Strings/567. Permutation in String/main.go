package Permutation_in_String

func checkInclusion(s1 string, s2 string) bool {
	var left, right, count int
	freq := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		freq[s1[i]]++
	}
	for right < len(s2) {
		if freq[s2[right]] > 0 {
			count++
		}
		freq[s2[right]]--

		if count == len(s1) {
			return true
		}
		if right-left+1 == len(s1) {
			if freq[s2[left]] >= 0 {
				count--
			}
			freq[s2[left]]++
			left++
		}
		right++
	}
	return false
}
