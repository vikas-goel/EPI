// Write a program which takes as input a very long sequence of numbers and
// prints the numbers in sorted order. Each number is at most k away from its
// correctly sorted position. (Such an array is sometimes referred to as being
// For example, no number in the sequence (3,-1,2,6,4,5,8} is more than 2 away
// from its final sorted position.

package main

import (
	"container/heap"
	"fmt"
	myheap "./heap"
)

func SortAlmostSorted(kAway int, array []int) []int {
	if kAway < 1 || len(array) == 0 {
		return array
	}

	h := myheap.NewIntegerMinHeap()
	sorted := make([]int, 0, len(array))

	// Push first k elements.
	for i := 0; i < kAway; i++ {
		heap.Push(h, array[i])
	}

	// Push an element and pop the smallest at the same time.
	// The heap size will remain k.
	for i := kAway; i < len(array); i++ {
		heap.Push(h, array[i])
		sorted = append(sorted, heap.Pop(h).(int))
	}

	// Pop remaining items in increasing order.
	for h.Len() > 0 {
		sorted = append(sorted, heap.Pop(h).(int))
	}

	return sorted
}

func main() {
	k := 2
	entries := []int{3, -1, 2, 6, 4, 5, 8}
	fmt.Println("Sort(", k, entries, ") =", SortAlmostSorted(k, entries))
}
