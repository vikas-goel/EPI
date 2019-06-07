// Implement a function for reversing the words in a string s.

package main

import "fmt"

func ReverseWords(str string) string {
	// Given AB, get BA; where A and B are words.
	// (AB)' = B'A'; (B')'(A')' = BA

	chars := []byte(str)

	// Reverse the entire string. That is, (AB)' = B'A'.
	start, end := 0, len(str)-1
	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}

	// Reverse each word/token with space as the delimiter.
	// That is, (B'}'(A'}' = BA.
	start = 0
	for i := 0; i <= len(str); i++ {
		if i != len(str) && chars[i] != ' ' {
			continue
		}

		// Found a delimiter. Reverse the token.
		end = i-1
		for start < end {
			chars[start], chars[end] = chars[end], chars[start]
			start++
			end--
		}

		// Skip all contiguous spaces.
		for ;i < len(str)-1 && chars[i+1] == ' '; i++ {}

		start = i+1
	}

	return string(chars)
}

func main() {
	for _, str := range []string{"ram is costly", "Alice likes Bob"} {
		fmt.Printf("%v -> %v\n", str, ReverseWords(str))
	}
}
