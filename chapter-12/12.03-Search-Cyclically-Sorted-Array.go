// Design an O(log n) algorithm for finding the position of the smallest
// element in a cyclically sorted array. Assume all elements are distinct.

package main

import "fmt"

func SearchSmallest(cyclicSortedArray []int) (index int) {
	index = -1
	array := cyclicSortedArray

	// Keep shrinking the search window by moving left or right as
	// appropriate.
	for start, end := 0, len(array)-1; start <= end && index == -1; {
		mid := start + (end - start) / 2
		if array[mid] < array[mid-1] {
			index = mid
		} else if array[mid] < array[start] {
			end = mid
		} else { // array[mid] > array[end]
			start = mid
		}
	}

	return
}

func main() {
	array := []int{378, 478, 550, 631, 710, 920, 203, 220, 234, 279, 368}
	fmt.Println("Smallest entry", array, "=", SearchSmallest(array))
}
