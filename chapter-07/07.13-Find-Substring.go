// Given two strings s (the "search string") and t (the "text"), find the first
// occurrence of s in t.

package main

import (
	"fmt"
	"strings"
)

func FindSubstring(str, substr string) int {
	// Invalid substring.
	if len(substr) > len(str) {
		return -1
	}

	const BASE = 26
	strHash, substrHash, strPower := 0, 0, 1
	for i := 0; i < len(substr); i++ {
		if i > 0 {
			strPower *= BASE
		}
		strHash = strHash * BASE + int(str[i])
		substrHash = substrHash * BASE + int(substr[i])
	}

	for i := len(substr); i < len(str); i++ {
		// Checks the two substrings are actually equal or not, to
		// protect against hash collision.
		if strHash == substrHash &&
			strings.Compare(substr, str[i-len(substr):i]) == 0 {
			return i - len(substr)
		}

		// Uses rolling hash to compute the new hash code.
		strHash -= int(str[i-len(substr)]) * strPower
		strHash = strHash * BASE + int(str[i])
	}

	if strHash == substrHash &&
		strings.Compare(str[len(str)-len(substr):], substr) == 0 {
		return len(str) - len(substr)
	}

	return -1
}

func main() {
	s1, s2 := "GACGCCA", "CGC"
	fmt.Printf("(%v, %v) = %v\n", s1, s2, FindSubstring(s1, s2))
}
