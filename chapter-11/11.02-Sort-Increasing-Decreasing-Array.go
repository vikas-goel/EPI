// An array is said to be k-increasing-decreasing if elements repeatedly
// increase up to a certain index after which they decrease, then again
// increase, a total of k times.
// Design an efficient algorithm for sorting a k-increasing-decreasing array.

package main

import (
	"fmt"
	"./heap"
)

func reverse(array []int) {
	length := len(array)
	for i := 0; i < length/2; i++ {
		array[i], array[length-1-i] = array[length-1-i], array[i]
	}
}

func SortSlopes(array []int) []int {
	if len(array) == 0 {
		return nil
	}

	arrayList := make([][]int, 0)

	const UP, DOWN = 0, 1

	// Create sub-arrays for each up and down slopes.
	for i, start, slope := 1, 0, UP; i < len(array); i++ {
		if (slope == UP && array[i-1] > array[i]) ||
			(slope == DOWN && array[i-1] < array[i]) {
			subArray := array[start:i]

			// Change the decreasing order to increasing order
			// before adding to the array list.
			if slope == DOWN {
				reverse(subArray)
			}

			arrayList = append(arrayList, subArray)
			slope ^= UP
			start = i
		}
	}

	if len(array) > 1 {
		arrayList = append(arrayList, array[len(array)-1:])
	}

	return heap.MergeSortedArrays(arrayList)
}

func main() {
	entries := []int{57, 131, 493, 294, 221, 339, 418, 452, 442, 190}
	fmt.Println("Sort", entries, "=", SortSlopes(entries))
}
