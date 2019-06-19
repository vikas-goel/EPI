// Design an algorithm for computing the running median of a sequence.

package main

import (
	"container/heap"
	"fmt"
	myheap "./heap"
)

func Median(stream []int) (medians []float64) {
	if len(stream) == 0 {
		return
	}

	// First half of the stream is maintained in max-heap to keep the
	// biggest of first half on top.
	maxH := myheap.NewIntegerMaxHeap()

	// Second half of the stream is maintained in min-heap to keep the
	// smallest of second half on top.
	minH := myheap.NewIntegerMinHeap()

	// Keep the length of the two heaps balanced. When the two heaps are of
	// the same length, the median will be average of the top elements of
	// the two heaps.
	// Every alternate inclusion of the stream data will make one of the
	// heaps length one more than the other. In this event, the median will
	// be the top element of the heap that has one more element.

	medians = make([]float64, 0, len(stream))

	for i := 0; i < len(stream); i++ {
		// Push in the right heap starting with max-heap,
		if maxH.Len() == 0 || maxH.Peek().(int) >= stream[i] {
			heap.Push(maxH, stream[i])
		} else {
			heap.Push(minH, stream[i])
		}

		// Balance the two heaps.
		if minH.Len() > maxH.Len() + 1 {
			heap.Push(maxH, heap.Pop(minH))
		} else if maxH.Len() > minH.Len() + 1 {
			heap.Push(minH, heap.Pop(maxH))
		}

		var median float64
		if maxH.Len() == minH.Len() {
			median = float64(minH.Peek().(int)+maxH.Peek().(int))/2
		} else if maxH.Len() > minH.Len() {
			median = float64(maxH.Peek().(int))
		} else {
			median = float64(minH.Peek().(int))
		}

		medians = append(medians, median)
	}

	return
}

func main() {
	entries := []int{1, 0, 3, 5, 2, 0, 1}
	fmt.Println("Median", entries, "=", Median(entries))
}
