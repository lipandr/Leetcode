package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	var (
		maxLen      int
		left, right int
		charSet     = make(map[byte]bool)
	)
	for right < n {
		// Если текущий символ уже есть в множестве, удаляем крайний левый символ и сдвигаем указатель left
		if charSet[s[right]] {
			delete(charSet, s[left])
			left++
		} else {
			// Иначе добавляем текущий символ в множество и сдвигаем указатель right
			charSet[s[right]] = true
			right++
			// Обновляем максимальную длину, если текущая подстрока без повторов длиннее
			if curLen := right - left; curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return maxLen
}
