// Design an efficient algorithm that takes a sorted array of distinct integers,
// and returns an index i such that the element at index i equals i. For
// example, when the input is (-2, 0, 2, 3, 6, 7, 9) your algorithm should
// return 2 or 3.

package main

import "fmt"

func IndexEntry(array []int) (index int) {
	index = -1

	// Keep shrinking the search window by moving left or right as
	// appropriate.
	for start, end := 0, len(array)-1; start <= end && index == -1; {
		mid := start + (end - start) / 2
		if array[mid] == mid {
			index = mid
		} else if array[mid] > mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return
}

func main() {
	array := []int{-2, 0, 2, 3, 6, 7, 9}
	fmt.Println("Index and Entry", array, "=", IndexEntry(array))
}
