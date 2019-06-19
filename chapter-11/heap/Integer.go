package heap

import "container/heap"

type IntegerHeap struct {
	items []int
	minHeap bool
}

func (this *IntegerHeap) Len() int {
	return len(this.items)
}

func (this *IntegerHeap) Less(i, j int) bool {
	if this.minHeap {
		return this.items[i] < this.items[j]
	} else {
		return this.items[i] > this.items[j]
	}
}

func (this *IntegerHeap) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

func (this *IntegerHeap) Push(item interface{}) {
	this.items = append(this.items, item.(int))
}

func (this *IntegerHeap) Pop() (item interface{}) {
	length := len(this.items)
	item = this.items[length-1]
	this.items = this.items[:length-1]

	return
}

func (this *IntegerHeap) Peek() (item interface{}) {
	return this.items[0]
}

func newIntegerHeap(minHeap bool) *IntegerHeap {
	var this IntegerHeap
	this.items = make([]int, 0)
	this.minHeap = minHeap

	heap.Init(&this)

	return &this
}

func NewIntegerMinHeap() *IntegerHeap {
	return newIntegerHeap(true)
}

func NewIntegerMaxHeap() *IntegerHeap {
	return newIntegerHeap(false)
}
