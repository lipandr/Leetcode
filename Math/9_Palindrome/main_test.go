package Palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {

	// Test case for a positive palindrome integer
	if !isPalindrome(121) {
		t.Errorf("isPalindrome(121) = %v; want true", isPalindrome(121))
	}

	// Test case for a negative integer (not a palindrome)
	if isPalindrome(-121) {
		t.Errorf("isPalindrome(-121) = %v; want false", isPalindrome(-121))
	}

	// Test case for a zero ending integer
	if isPalindrome(10) {
		t.Errorf("isPalindrome(10) = %v; want false", isPalindrome(-121))
	}

	// Test case for an odd-lengthen non-palindrome integer
	if !isPalindrome(12321) {
		t.Errorf("isPalindrome(12321) = %v; want true", isPalindrome(12321))
	}

	// Test case for an even-lengthen non-palindrome integer
	if isPalindrome(1233214) {
		t.Errorf("isPalindrome(1233214) = %v; want false", isPalindrome(1233214))
	}
}
