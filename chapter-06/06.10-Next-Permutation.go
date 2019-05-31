// Write a program that takes as input a permutation, and returns the next
// permutation under dictionary ordering. If the permutation is the last
// permutation, return the empty array. For example, if the input is (1,0,3,2)
// your function should return (1,2,0,3). If the input is (3,2,1,0), return ().

package main

import "fmt"

func NextPermutation(array []int) []int {
	if len(array) < 2 {
		return nil
	}

	// Find the first element from the right that is smaller than its next
	// element (i.e., first element on the down slope).
	swapIndex := -1
	for i := len(array)-2; i >=0; i-- {
		if array[i] < array[i+1] {
			swapIndex = i
			break
		}
	}

	// If no such element exists, then next permutation not possible as
	// the array is in non-increasing order.
	if swapIndex == -1 {
		return nil
	}

	// Create a new array for the next permutation.
	next := make([]int, len(array))
	copy(next, array)

	// Find the smallest number on the right side of down slope number that
	// is bigger than it. Once found, swap the two numbers.
	// The right subarray remains in non-increasing order.
	for i := len(next)-1; i > swapIndex; i-- {
		if next[i] > next[swapIndex] {
			next[i], next[swapIndex] = next[swapIndex], next[i]
			break
		}
	}

	// Reverse the right subarray to make it non-decreasing to get the
	// next permutation.
	start, end := swapIndex+1, len(next)-1
	for start < end {
		next[start], next[end] = next[end], next[start]
		start++
		end--
	}

	return next
}

func main() {
	for _, array := range [][]int{{1,0,3,2}, {3,2,1,0}, {6,4,7,5,3}} {
		fmt.Printf("Next permutatation %v = %v\n",
			array, NextPermutation(array))
	}
}
