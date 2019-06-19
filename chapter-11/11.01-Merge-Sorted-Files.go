// Write a program that takes as input a set of sorted sequences and computes
// the union of these sequences as a sorted sequence. For example, if the input
// is (3,5,7), (0,6), and (0,6,28), then the output is (0,0,3,5,6,6,7,28).

package main

import (
	"fmt"
	"./heap"
)

func main() {
	entries := [][]int{{3,5,7}, {0,6}, {0,6,28}}
	fmt.Println("Sort", entries, "=", heap.MergeSortedArrays(entries))
}
