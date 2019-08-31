// Write a program that takes an array A and an index i into A, and rearranges
// the elements such that all elements less than A[i] (the "pivot") appear
// first, followed by elements equal to the pivot, followed by elements greater
// than the pivot.

package main

import "fmt"

func Rearrange(pivotIndex int, array []int) {
	if pivotIndex < 0 || pivotIndex >= len(array) {
		return
	}

	pivot := array[pivotIndex]
	low, mid, high := 0, 0, len(array)-1
	for mid <= high {
		if array[mid] < pivot {
			array[low], array[mid] = array[mid], array[low]
			low++
			mid++
		} else if array[mid] == pivot {
			mid++
		} else { // if array[mid] > pivot
			array[mid], array[high] = array[high], array[mid]
			high--
		}
	}
}

func main() {
	array := []int{0, 1, 2, 0, 2, 1, 1}
	pivot := 0
	fmt.Print(array, "->rearrange(", pivot, ")->")
	Rearrange(pivot, array)
	fmt.Println(array)
}
