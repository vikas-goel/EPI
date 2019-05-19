// Let A be an array representing the heights of adjacent buildings of unit
// width. Design an algorithm to compute the area of the largest rectangle
// contained in this skyline.

package main

import "fmt"

const Verbose = false

func main() {
	skyline := Skyline{1, 4, 2, 5, 6, 3, 2, 6, 6, 5, 2, 1, 3}
	fmt.Println("Largest area of", skyline, "=", maxArea(skyline))
}

func maxArea(bars Skyline) (maxArea int) {
	curBar := 0
	myStack := NewStack(bars.Len())

	evalMaxBarArea := func(index int) {
		thisArea := 0

		if Verbose {
			fmt.Printf("Index = %2v, %2v", index, *myStack)
		}

		lastBar := bars.Height(myStack.Pop())
		if myStack.Empty() {
			thisArea = lastBar * index
		} else {
			thisArea = lastBar * (index - myStack.Peek() - 1)
		}

		if thisArea > maxArea {
			maxArea = thisArea
		}

		if Verbose {
			fmt.Printf(", Area = %2v\n", maxArea)
		}
	}

	for curBar < bars.Len() {
		if myStack.Empty() ||
			bars.Height(myStack.Peek()) < bars.Height(curBar) {
			myStack.Push(curBar)
			curBar += 1
		} else {
			evalMaxBarArea(curBar)
		}
	}

	for !myStack.Empty() {
		evalMaxBarArea(curBar)
	}

	return
}

type Skyline []int

func (this *Skyline) Len() int {
	return len(*this)
}

func (this *Skyline) Height(index int) int {
	return (*this)[index]
}

type Stack struct {
	size int
	elems []int
}

func NewStack(capacity int) *Stack {
	stack := new(Stack)
	stack.elems = make([]int, capacity)
	return stack
}

func (this *Stack) Push(value int) {
	this.elems[this.size] = value
	this.size++
}

func (this *Stack) Pop() int {
	if this.size > 0  {
		this.size--
		value := this.elems[this.size]
		this.elems[this.size] = 0
		return value
	}

	return 0
}

func (this *Stack) Peek() int {
	if this.size > 0  {
		return this.elems[this.size-1]
	}

	return 0
}

func (this *Stack) Len() int {
	return this.size
}

func (this *Stack) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}
