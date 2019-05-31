// Write a program which takes an array of n integers, where A[i] denotes the
// maximum you can advance from index i, and returns whether it is possible to
// advance to the last index starting from the beginning of the array.

package main

import "fmt"

func Advance(array []int) bool {
	if len(array) == 0 || array[0] <= 0 {
		return false
	}

	furthestIndex := array[0]

	// Track how far we can go from each reachable positions.
	// Stop as soon as the end position is reachable.
	for i := 1; i <= furthestIndex && furthestIndex < len(array); i++ {
		if i + array[i] > furthestIndex {
			furthestIndex = i + array[i]
		}
	}

	return furthestIndex >= len(array)-1
}

func main() {
	for _, array := range [][]int{{3, 2, 0, 0, 2, 0, 1},
		{3, 3, 1, 0, 2, 0, 1}} {
		fmt.Println("Reachable", array, "?", Advance(array))
	}
}
