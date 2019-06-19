// Consider a coordinate system for the Milky Way in which Earth is at (0,0,0).
// Model stars as points, and assume distances are in light years. The Milky Way
// consists of approximately 1012 stars, and their coordinates are stored in a
// file. How would you compute the k stars which are closest to Earth?

package main

import (
	"container/heap"
	"fmt"
	myheap "./heap"
)

func ClosestStars(k int, stars []int) (closest []int) {
	if k == 0 || len(stars) == 0 {
		return
	}

	h := myheap.NewIntegerMaxHeap()
	closest = make([]int, 0, k)

	// Push first k elements.
	for i := 0; i < k; i++ {
		heap.Push(h, stars[i])
	}

	// Push an element and pop the farthest at the same time.
	// The heap size will remain k.
	for i := k; i < len(stars); i++ {
		heap.Push(h, stars[i])
		heap.Pop(h)
	}

	// Pop k heap items in increasing order that are closest to the origin.
	for h.Len() > 0 {
		closest = append(closest, heap.Pop(h).(int))
	}

	return closest
}

func main() {
	k := 4
	entries := []int{3, 101, 95, 12, 25, 30, 65, 77, 17}
	fmt.Println("Sort(", k, entries, ") =", ClosestStars(k, entries))
}
