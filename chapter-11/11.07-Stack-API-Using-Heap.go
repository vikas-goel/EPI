// Implement a stack API using a heap.

package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	myheap "./heap"
)

type RankItem = myheap.RankItem
type RankItemHeap = myheap.RankItemHeap

// Create a stack using a heap to store the items. The latest items must be
// popped first. So, use max-heap and sequence number for ranking the items.
type Stack struct {
	NextSequence int
	Heap *RankItemHeap
}

func (this *Stack) Empty() bool {
	return this.Heap.Len() == 0
}

func (this *Stack) Peek() int {
	if this.Empty() {
		panic("Empty stack.")
	}

	return this.Heap.Peek().Value.(int)
}

func (this *Stack) Pop() int {
	if this.Empty() {
		panic("Empty stack.")
	}

	return heap.Pop(this.Heap).(*RankItem).Value.(int)
}

func (this *Stack) Push(item int) {
	heap.Push(this.Heap, &RankItem{Rank: this.NextSequence, Value: item})
	this.NextSequence++
}

func NewStack() *Stack {
	var this Stack
	this.Heap = myheap.NewRankItemHeap()

	return &this
}

func main() {
	stack := NewStack()

	fmt.Print("[ ")

	for i := 0; i < 50; i++ {
		operation := rand.Intn(2)
		if stack.Empty() || operation == 1 {
			element := rand.Intn(10000)
			stack.Push(element)
			fmt.Printf("+%v ", element)
		} else {
			fmt.Printf("-%v ", stack.Pop())
		}
	}

	for !stack.Empty() {
		fmt.Printf("-%v ", stack.Pop())
	}

	fmt.Println("]")
}
