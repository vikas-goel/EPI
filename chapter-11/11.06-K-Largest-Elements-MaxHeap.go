// Given a max-heap, represented as an array A, design an algorithm that
// computes the k largest elements stored in the max-heap. You cannot modify
// the heap.

package main

import (
	"container/heap"
	"fmt"
	myheap "./heap"
)

type HeapItem = myheap.RankItem
type LargestHeap = myheap.RankItemHeap
var NewLargestHeap = myheap.NewRankItemHeap

func kLargestElement(maxheap []int, k int) []int {
	if k < 1 || k > len(maxheap) {
		return nil
	}

	// Track the k-largest items through another max-heap.
	// Push children of traversing item in the largest max-heap.
	largestHeap := NewLargestHeap()

	// The first item (i.e. root) in the given max-heap is the largest
	// in the max-heap. So, push this item in the largest max-heap.
	heap.Push(largestHeap, &HeapItem{0, maxheap[0]})

	// Prepare an array to capture the k-largest items.
	largestItems := make([]int, 0, k)

	// Iterate K-times.
	for ; largestHeap.Len() > 0 && k > 0; k-- {
		// Remove the largest item from the current largest max-heap.
		item := heap.Pop(largestHeap).(*HeapItem)

		// Add the item to the array.
		largestItems = append(largestItems, item.Rank)

		// Add both children, if valid.
		left := item.Value.(int) * 2 + 1
		right := left + 1

		if left < len(maxheap) {
			heap.Push(largestHeap, &HeapItem{left, maxheap[left]})
		}

		if right < len(maxheap) {
			heap.Push(largestHeap, &HeapItem{right, maxheap[right]})
		}
	}

	return largestItems
}

func main() {
	k := 9
	maxHeap := []int{561, 314, 401, 28, 156, 359, 271, 11, 3}
	fmt.Println("Largest(", k, maxHeap, ") =", kLargestElement(maxHeap, k))
}
