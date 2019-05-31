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

	left, right := pivotIndex, pivotIndex
	for i := 0; i < len(array); i++ {
		if i >= left && i <= right {
			continue
		}

		if array[i] > array[right] && i < right {
			// Found a bigger element on the left side.
			// Swap it with the right pivot position and shift the
			// pivot range to one left.
			array[i], array[right] = array[right], array[i]
			left--
			right--
		} else if array[i] < array[left] && i > left {
			// Found a smaller element on the right side.
			// Swap it with the left pivot position and shift the
			// pivot range to one right.
			array[i], array[left] = array[left], array[i]
			left++
			right++
		} else if array[i] == array[left] {
			if i < left {
				// Found an element same as the pivot on the
				// left side. Move it before the left index
				// and increase the pivot index range.
				left--
				array[i], array[left] = array[left], array[i]

				// The i-th index has a new element that is
				// not processed yet. So, adjust the index.
				i--
			} else {
				// Found an element same as the pivot on the
				// right side. Move it after the right index
				// and increase the pivot index range.
				right++
				array[i], array[right] = array[right], array[i]
			}
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
