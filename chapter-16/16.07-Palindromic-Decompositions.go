// Compute all palindromic decompositions of a given string. For example, if
// the string is "0204451881", then the decomposition "020", "44", "5", "1881"
// is palindromic, as is "020", "44", "5", "1", "88", "1". However, "02044,
// "5", "1881" is not a palindromic decomposition.

package main

import "fmt"

func isPalindrome(array []byte, start, end int) bool {
	if start == end {
		return true
	}

	for i, mid := start, start + (end + 1 - start) / 2; i < mid; i++ {
		if array[i] != array[end-i+start] {
			return false
		}
	}

	return true
}

func PalindromicDecompositions(str string) (result [][]string) {
	strArrary := []byte(str)
	var palindromes []string

	var decompose func(start int)
	decompose = func(start int) {
		// Base case.
		if start >= len(strArrary) {
			// Make a copy and push to the result.
			clone := make([]string, len(palindromes))
			copy(clone, palindromes)
			result = append(result, clone)

			return
		}

		// For each palindromic prefix, repeat the process for
		// remaining suffix.
		for i := start; i < len(strArrary); i++ {
			if isPalindrome(strArrary, start, i) {
				palin := string(strArrary[start:i+1])

				// Push this palindromic sub string to the list.
				palindromes = append(palindromes, palin)

				decompose(i+1)

				// Pop the last one before pushing a new item.
				palindromes = palindromes[:len(palindromes)-1]
			}
		}
	}

	decompose(0)
	return
}

func main() {
	userString := "0204451881"
	PalindromicDecompositions(userString)
	fmt.Printf("%v\n", PalindromicDecompositions(userString))
}
