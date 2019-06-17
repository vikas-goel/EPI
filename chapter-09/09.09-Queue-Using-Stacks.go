// Implement a queue given a library implementing stacks.

package main

import (
	"fmt"
	"math/rand"
	"./stack"
)

type Stack = stack.Stack
var InitStack = stack.InitStack

// Create a queue using two stacks. One stack will be used to enqueue the
// elements. The other will be used to dequeue.
type Queue struct {
	size int
	enqueue, dequeue *Stack
}

func (this *Queue) Empty() bool {
	return this.size == 0
}

func (this *Queue) Len() int {
	return this.size
}

func (this *Queue) Enqueue(element interface{}) {
	this.enqueue.Push(element)
	this.size++
}

func (this *Queue) Dequeue() interface{} {
	if this.Empty() {
		panic("Empty queue.")
	}

	// If the Dequeue stack is empty, then move all the elements of
	// Enqueue stack to it.
	if this.dequeue.Empty() {
		for !this.enqueue.Empty() {
			this.dequeue.Push(this.enqueue.Pop())
		}
	}

	this.size--
	return this.dequeue.Pop()
}

func NewQueue() *Queue {
	this := new(Queue)
	this.enqueue, this.dequeue = InitStack(), InitStack()

	return this
}

func main() {
	queue := NewQueue()

	for i :=  0; i < 100; i++ {
		operation := rand.Intn(2)
		if queue.Empty() || operation == 1 {
			element := rand.Intn(10000)
			queue.Enqueue(element)
			fmt.Println("+", element)
		} else {
			fmt.Println("\t-", queue.Dequeue().(int))
		}
	}

	for !queue.Empty() {
		fmt.Println("\t-", queue.Dequeue().(int))
	}
}
