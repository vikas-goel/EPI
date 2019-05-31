// Write a program which takes as input a sorted array and updates it so that
// all duplicates have been removed and the remaining elements have been
// shifted left to fill the emptied indices. Return the number of valid
// elements.

package main

import "fmt"

func DeleteDuplicates(sortedArray []int) int {
	if len(sortedArray) < 2 {
		return len(sortedArray)
	}

	// Keep track of last index where the next unique number can be
	// placed.
	lastIndex := 1
	for i := 1; i < len(sortedArray); i++ {
		// If the current number is bigger than the last one, then
		// it is the new unique number found. So, place it in the
		// right index of the unique elements.
		if sortedArray[i] > sortedArray[i-1] {
			sortedArray[lastIndex] = sortedArray[i]
			lastIndex++
		}
	}

	// Fill the remaining positions with 0.
	for i := lastIndex; i < len(sortedArray); i++ {
		sortedArray[i] = 0
	}

	// The last index also indicates the total number of unique elements.
	return lastIndex
}

func main() {
	for _, array := range [][]int{{2, 3, 5, 5, 7, 11, 11, 11, 13},
		{1, 2, 3, 4, 5, 6, 7, 8, 10, 11}} {
		fmt.Print(array)
		unique := DeleteDuplicates(array)
		fmt.Println("; Unique numbers =", unique, array)
	}
}
