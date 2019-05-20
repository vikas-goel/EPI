// Write a program that takes as input an array of numbers and returns the
// length of a longest nondecreasing subsequence in the array.

package main

import "fmt"

func LongestSubsequence(array []int) (longest int) {
	length := make([]int, len(array))
	length[0], longest = 1, 1

	for i := 1; i < len(array); i++ {
		length[i] = 0
		for j := 0; j < i-1; j++ {
			if array[j] <= array[i] && length[j] > length[i] {
				length[i] = length[j]
			}
		}
		length[i]++

		if length[i] > longest {
			longest = length[i]
		}
	}

	return
}

func main() {
	array := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9}
	fmt.Printf("Longest non-decreasing subsequence for %v is %v.\n",
		array, LongestSubsequence(array))
}
