// Write a program that takes two strings and computes the minimum number of
// edits needed to transform the first string into the second string.

package main

import "fmt"

func Transform(from, to string) (numEdits int) {
	return NumEdits(from, to, len(from), len(to))
}

func NumEdits(str1, str2 string, end1, end2 int) (numEdits int) {
	// If length of either of the strings is 0, it takes edits equal to the
	// number of characters of the other string.
	if end1 == 0 {
		return end2
	} else if end2 == 0 {
		return end1
	}

	if str1[end1-1] == str2[end2-1] {
		// Last character of both strings is same.
		return NumEdits(str1, str2, end1-1, end2-1)
	} else {
		// Substitute last character to make it same as the other.
		edit1 := NumEdits(str1, str2, end1-1, end2-1)

		// Remove last character from str1.
		edit2 := NumEdits(str1, str2, end1-1, end2)

		// Add last character to str1.
		edit3 := NumEdits(str1, str2, end1, end2-1)

		// Find minimum of the three edits.
		numEdits = edit1
		if edit2 < numEdits {
			numEdits = edit2
		}

		if edit3 < numEdits {
			numEdits = edit3
		}

		// Increment for the edit due to operation.
		numEdits++
	}

	return
}

func main() {
	str1, str2 := "Saturday", "Sundays"
	fmt.Printf("Number of edits for (%v, %v) is %v.\n",
		str1, str2, Transform(str1, str2))

	str1, str2 = "Carthorse", "Orchestra"
	fmt.Printf("Number of edits for (%v, %v) is %v.\n",
		str1, str2, Transform(str1, str2))
}
