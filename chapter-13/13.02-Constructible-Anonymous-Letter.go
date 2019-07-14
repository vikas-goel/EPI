// Write a program which takes text for an anonymous letter and text for a
// magazine and determines if it is possible to write the anonymous letter
// using the magazine. The anonymous letter can be written using the magazine
// if for each character in the anonymous letter, the number of times it
// appears in the anonymous letter is no more than the number of times it
// appears in the magazine.

package main

import "fmt"

func Constuctible(letter, magazine string) bool {
	texts := make(map[byte]int)
	for i := 0; i < len(letter); i++ {
		// Ignore white spaces.
		if letter[i] == ' ' || letter[i] == '\t' || letter[i] ==  '\n' {
			continue
		}

		// Count frequency of each letter.
		texts[letter[i]]++
	}

	// Check whether the letters are present in the magazine at least
	// the corresponding frequencies times.
	for i := 0; i < len(magazine); i++ {
		if count, ok := texts[magazine[i]]; ok && count > 0 {
			texts[magazine[i]]--

			if texts[magazine[i]] == 0 {
				delete(texts, magazine[i])
			}
		}
	}

	if len(texts) > 0 {
		return false
	}

	return true
}

func main() {
	letter := "You are welcome"
	magazine := "Yellow color is my favorite along with blue"
	fmt.Printf("%v, %v = %v\n",
		letter, magazine, Constuctible(letter, magazine))
}
