package RotateArray

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	left, right := 0, n-1

	// Первый переворот от начала до конца массива
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	// Второй переворот от начала до k-1 элемента
	left, right = 0, k-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	// Третий переворот от k-го элемента до конца
	left, right = k, n-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
