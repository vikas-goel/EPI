// For the purpose of this problem, define a palindromic string to be a string
// which when all the non-alphanumeric are removed it reads the same front to
// back ignoring case. For example, "A man, a plan, a canal, Panama." and
// "Able was I, ere I saw Elba!" are palindromic, but "Ray a Ray" is not.
// Implement a function which takes as input a string s and returns true if s
// is a palindromic string.

package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)

	isAlphanumeric := func(index int) bool {
		if str[index] >= 'a' && str[index] <= 'z' {
			return true
		} else if str[index] >= '0' && str[index] <= '9' {
			return true
		}

		return false
	}

	start, end := 0, len(str)-1
	for start <= end {
		// Move start/end positions till a relevant character is found.
		if !isAlphanumeric(start){
			start++
		} else if !isAlphanumeric(end) {
			end--
		} else if str[start] != str[end] {
			// Mismatch.
			return false
		} else {
			// Matching characters.
			start++
			end--
		}
	}

	return true
}

func main() {
	for _, str := range []string{"A man, a plan, a canal, Panama.",
		"Able was I, ere I saw Elba!", "Ray a Ray"} {
		fmt.Printf("Palindrome(%v)? %v\n", str, IsPalindrome(str))
	}
}
