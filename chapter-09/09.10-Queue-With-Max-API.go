// Implement a queue with enqueue, dequeue, and max operations. The max
// operation returns the maximum element currently stored in the queue.

package main

import (
	"fmt"
	"container/list"
	"math/rand"
)

type List = list.List
type Element = list.Element
type Queue = MaxQueue

type MaxQueue struct {
	// queue contains the values.
	// max container element pointers of the values in the queue.
	queue, max *List
	Less func(interface{}, interface{}) bool
}

func (this *Queue) Empty() bool {
	return this.Len() == 0
}

func (this *Queue) Len() int {
	return this.queue.Len()
}

func (this *Queue) Enqueue(value interface{}) {
	element := this.queue.PushBack(value)

	// Remove all tail element pointers from the max list whose value is
	// smaller or equal to the specified value.
	// The front element of the max list will always be the maximum
	// element in the list.
	for this.max.Len() != 0 &&
		this.Less(this.max.Back().Value.(*Element).Value, value) {
		this.max.Remove(this.max.Back())
	}

	// Push the value-element pointer to the max list.
	this.max.PushBack(element)
}

func (this *Queue) Dequeue() interface{} {
	if this.Empty() {
		panic("Empty queue.")
	}

	element := this.queue.Front()
	value := this.queue.Remove(element)

	// Remove the element pointer from the max list, if it exists.
	if this.max.Front().Value.(*Element) == element {
		this.max.Remove(this.max.Front())
	}

	return value
}

func (this *Queue) Max() interface{} {
	if this.Empty() {
		panic("Empty queue.")
	}

	return this.max.Front().Value.(*Element).Value
}

func NewQueue(less func(interface{}, interface{}) bool) *Queue {
	this := new(Queue)
	this.Less = less
	this.queue, this.max = list.New(), list.New()
	this.queue.Init()
	this.max.Init()

	return this
}

func IntegerLess(v1, v2 interface{}) bool {
	return v1.(int) <= v2.(int)
}

func main() {
	queue := NewQueue(IntegerLess)

	for i :=  0; i < 100; i++ {
		operation := rand.Intn(2)
		if queue.Len() <= 1 || operation == 1 {
			element := rand.Intn(10000)
			queue.Enqueue(element)
			fmt.Printf("+ %v, m %v\n", element, queue.Max().(int))
		} else {
			element := queue.Dequeue().(int)
			fmt.Printf("\t- %v, m %v\n", element, queue.Max().(int))
		}
	}

	for queue.Len() > 1 {
		element := queue.Dequeue().(int)
		fmt.Printf("\t- %v, m %v\n", element, queue.Max().(int))
	}
}
