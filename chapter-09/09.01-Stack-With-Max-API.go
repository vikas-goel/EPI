// Design a stack that includes a max operation, in addition to push and pop.
// The max method should return the maximum value stored in the stack.

package main

import "fmt"

type Value struct {
	data, max int
}

// Stack that keeps track of the maximum value with each data.
type StackMax struct {
	size int
	value []Value
}

// Optimized Stack that keeps track of the index of the maximum value.
type StackMaxOpt struct {
	size int
	data []int
	maxLen int
	max []int
}

func (this *StackMax) Empty() bool {
	return this.size == 0
}

func (this *StackMax) Max() (max int, ok bool) {
	if this.size == 0 {
		return
	}

	return this.value[this.size-1].max, true
}

func (this *StackMax) Peek() (data int, ok bool) {
	if this.size == 0 {
		return
	}

	return this.value[this.size-1].data, true
}

func (this *StackMax) Pop() (data int, ok bool) {
	if this.size == 0 {
		return
	}

	data, ok = this.value[this.size-1].data, true
	this.size--

	return
}

func (this *StackMax) Push(data int) {
	if this.size < len(this.value) {
		this.value[this.size].data = data
	} else {
		elem := Value{data: data}
		this.value = append(this.value, elem)
	}

	this.value[this.size].max = this.value[this.size].data
	if this.size > 0 &&
		this.value[this.size-1].max > this.value[this.size].max {
		this.value[this.size].max = this.value[this.size-1].max
	}

	this.size++
}

func (this *StackMaxOpt) Empty() bool {
	return this.size == 0
}

func (this *StackMaxOpt) Max() (max int, ok bool) {
	if this.size == 0 {
		return
	}

	return this.data[this.maxLen-1], true
}

func (this *StackMaxOpt) Peek() (data int, ok bool) {
	if this.size == 0 {
		return
	}

	return this.data[this.size-1], true
}

func (this *StackMaxOpt) Pop() (data int, ok bool) {
	if this.size == 0 {
		return
	}

	data, ok = this.data[this.size-1], true

	// If the maximum index is same as the one popped, then
	// remove the max index from the stack.
	if this.max[this.maxLen-1] == this.size-1 {
		this.maxLen--
	}

	this.size--

	return
}

func (this *StackMaxOpt) Push(data int) {
	if this.size < len(this.data) {
		this.data[this.size] = data
	} else {
		this.data = append(this.data, data)
	}

	insertMax := false
	if this.size > 0 {
		if this.data[this.size] > this.data[this.max[this.maxLen-1]] {
			insertMax = true
		}
	} else {
		insertMax = true
	}

	if insertMax {
		if this.maxLen < len(this.max) {
			this.max[this.maxLen] = this.size
			this.maxLen++
		} else {
			this.max = append(this.max, this.size)
			this.maxLen++
		}
	}

	this.size++
}

func InitStackMax(data ...int) *StackMax {
	stack := new(StackMax)
	stack.value = make([]Value, len(data))
	for _, d := range data {
		stack.Push(d)
	}

	return stack
}

func InitStackMaxOpt(data ...int) *StackMaxOpt {
	stack := new(StackMaxOpt)
	stack.data = make([]int, len(data))
	stack.max = make([]int, 1)
	for _, d := range data {
		stack.Push(d)
	}

	return stack
}

func main() {
	fmt.Println(InitStackMax(2, 2, 1, 4, 5, 5, 3))

	stack := InitStackMaxOpt(2, 2, 1, 4, 5, 5, 3)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Push(1)
	fmt.Println(stack)
	stack.Push(3)
	fmt.Println(stack)
}
