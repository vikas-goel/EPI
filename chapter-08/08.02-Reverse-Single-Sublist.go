// Write a program which takes a singly linked list L and two integers s and f
// as arguments, and reverses the order of the nodes from the sth node to fth
// node, inclusive. The numbering begins at 1, i.e., the head node is the first
// node. Do not allocate additional nodes.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func Reverse(list *List, start, finish int) {
	if list == nil || start >= finish || start < 1 {
		return
	}

	for i := start; i > 2 && list.Next != nil; i-- {
		list = list.Next
	}

	// Variable list now points to the element previous of the start one.
	if list.Next == nil {
		return
	}

	// Reverse start to finish pointers.
	curr, next := list.Next, list.Next.Next
	for finish -= start; finish > 0 && next != nil; finish-- {
		temp := next.Next
		next.Next = curr
		curr = next
		next = temp
	}

	// Adjust next of start pointer.
	list.Next.Next = next


	// Adjust previous of start pointer.
	list.Next = curr

	return
}

func main() {
	list := InitList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	s, f := 4, 7

	list.Print()
	fmt.Printf("->Reverse[%v, %v]->", s, f)

	Reverse(list, s, f)
	list.Print()
}
