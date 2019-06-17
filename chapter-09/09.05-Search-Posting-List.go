// Write recursive and iterative routines that take a postings list, and
// compute the jump-first order. Assume each node has an integer-valued field
// that holds the order, and is initialized to -1.

package main

import (
	"fmt"
	"./stack"
)

type Stack = stack.Stack
var InitStack = stack.InitStack

type List struct {
	Order int
	Jump, Next *List
}

func SearchRecursive(posting *List) {
	if posting == nil {
		return
	}

	var search func(list *List, order int) int
	search = func(list *List, order int) int {
		if list == nil || list.Order != -1 {
			return order
		}

		list.Order = order
		order++

		order = search(list.Jump, order)
		order = search(list.Next, order)

		return order
	}

	search(posting, 0)
}

func SearchIterative(posting *List) {
	if posting == nil {
		return
	}

	order := 0
	stack := InitStack()
	stack.Push(posting)

	for !stack.Empty() {
		post := stack.Pop().(*List)
		if post == nil || post.Order != -1 {
			continue
		}

		post.Order = order
		order++

		// Process Jump first and then Next. So, push them in reverse
		// order to the stack.
		stack.Push(post.Next)
		stack.Push(post.Jump)
	}
}

func PrintList(list *List) {
	fmt.Print("[ ")
	for ; list != nil; list = list.Next {
		fmt.Print(list.Order, " ")
	}
	fmt.Println("]")
}

func main() {
	list1d := &List{Order: -1}
	list1c := &List{Order: -1, Next: list1d}
	list1b := &List{Order: -1, Next: list1c}
	list1a := &List{Order: -1, Next: list1b}
	list1a.Jump = list1c
	list1b.Jump = list1d
	list1c.Jump = list1b
	list1d.Jump = list1d

	list2d := &List{Order: -1}
	list2c := &List{Order: -1, Next: list2d}
	list2b := &List{Order: -1, Next: list2c}
	list2a := &List{Order: -1, Next: list2b}
	list2a.Jump = list2c
	list2b.Jump = list2d
	list2c.Jump = list2b
	list2d.Jump = list2d

	SearchRecursive(list1a)
	fmt.Print("Recursive: ")
	PrintList(list1a)

	SearchIterative(list2a)
	fmt.Print("Iterative: ")
	PrintList(list2a)
}
