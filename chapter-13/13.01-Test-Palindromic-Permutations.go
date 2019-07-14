// Write a program to test whether the letters forming a string can be permuted
// to form a palindrome. For example, "edified" can be permuted to form
// "deified".

package main

import "fmt"

func TestPalindrome(str string) bool {
	chars := make(map[byte]bool)
	for i := 0; i < len(str); i++ {
		if present := chars[str[i]]; present {
			chars[str[i]] = false
		} else {
			chars[str[i]] = true
		}
	}

	countOdd := 0
	for _, value := range chars {
		if value {
			countOdd++
		}

	}

	if countOdd > 1 {
		return false
	} else {
		return true
	}
}

func main() {
	array := []string{"edified", "level", "rotator", "foobaraboof", "true"}
	for _, str := range array {
		fmt.Printf("Palindrome(%v)? %v\n", str, TestPalindrome(str))
	}
}
