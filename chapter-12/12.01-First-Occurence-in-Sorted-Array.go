// Write a method that takes a sorted array and a key and returns the index of
// the first occurrence of that key in the array.

package main

import "fmt"

func FirstOccurence(k int, array []int) (index int) {
	index = -1

	// Similar to binary search except that the search continues till
	// all options exhausted by discarding the irrelevant subsets.
	for start, end := 0, len(array)-1; start <= end; {
		mid := start + (end - start) / 2
		if array[mid] == k {
			index, end = mid, mid - 1
		} else if array[mid] > k {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return
}

func main() {
	k := 108
	array := []int{-14, -10, 2, 108, 108, 243, 285, 285, 401}
	fmt.Println("First(", k, ") in", array, "=", FirstOccurence(k, array))
}
