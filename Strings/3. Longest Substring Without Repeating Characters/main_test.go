package LongestSubstringWithoutRepeatingCharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	// Test case 1: When the input is an empty string, expect the output to be 0
	result := lengthOfLongestSubstring("")
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: When the input is a string with no repeated characters, expect the output to be the length of the string
	result = lengthOfLongestSubstring("abcdefg")
	expected = 7
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 3: When the input is a string with repeated characters, expect the output to be the length of the longest substring without repeating characters
	result = lengthOfLongestSubstring("aabcabcbb")
	expected = 3
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 4: When the input is a string with a single character, expect the output to be 1
	result = lengthOfLongestSubstring("a")
	expected = 1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 5: When the input is a string with a two no repeated characters, expect the output to be 2
	result = lengthOfLongestSubstring("au")
	expected = 2
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
