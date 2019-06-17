// Given a binary tree, return an array consisting of the keys at the same
// level. Keys should appear in the order of the corresponding nodes' depths,
// breaking ties from left to right.

package main

import (
	"fmt"
	"math/rand"
)

type Ring struct {
	head, tail, size int
	queue []interface{}
}

func (this *Ring) Enqueue(element interface{}) {
	// Grow the queue for the new element if it is full.
	if this.size == cap(this.queue) {
		if this.head != 0 {
			// Rotate the queue to move head to 0th place.
			temp := this.queue[:this.head]
			this.queue = append(this.queue[this.head:], temp...)
			this.head, this.tail = 0, this.size
		}

		// Append the new element at the end.
		this.queue = append(this.queue, element)
	} else if this.tail == len(this.queue) && this.tail < cap(this.queue) {
		// The length of the queue is exhausted but it still has
		// capacity to auto-grow the length.
		// So, append the new element at the end.
		this.queue = append(this.queue, element)
	} else {
		this.tail %= len(this.queue)
		this.queue[this.tail] = element
	}

	this.tail++
	this.size++
}

func (this *Ring) Dequeue() (element interface{}) {
	if this.size == 0 {
		panic("Empty circular queue.")
	}

	element = this.queue[this.head]

	this.head = (this.head + 1) % len(this.queue)
	this.size--

	return
}

func (this *Ring) Len() int {
	return this.size
}

func NewRing(capacity int) *Ring {
	this := new(Ring)
	this.queue = make([]interface{}, capacity, 2*capacity)
	this.head, this.tail = 0, 0

	return this
}

func main() {
	ring := NewRing(5)

	for i :=  0; i < 100; i++ {
		operation := rand.Intn(2)
		if ring.Len() == 0 || operation == 1 {
			element := rand.Intn(10000)
			ring.Enqueue(element)
			fmt.Println("Enqueue:", element, "=", ring)
		} else {
			element := ring.Dequeue().(int)
			fmt.Println("Dequeue:", element, "=", ring)
		}
	}
}
