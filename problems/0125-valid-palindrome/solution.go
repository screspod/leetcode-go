package validpalindrome

import (
	"strings"
	"unicode"
)

// Time: O(n) single pass with two pointers over the string
// Memory: O(1) only two pointer variables
// Pattern: Two pointers — shrink from both ends, skip non-alphanumeric chars, compare case-insensitively
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	s = strings.ToLower(s)

	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c))
}
