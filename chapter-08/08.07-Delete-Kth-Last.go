// Given a singly linked list and an integer k, write a program to remove the
// Kth last element from the list. Your algorithm cannot use more than a few
// words of storage, regardless of the length of the list. In particular, you
// cannot assume that it is possible to record the length of the list.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func DeleteKthLast(list *List, K int) {
	// Do nothing if the node is nil pointer or tail.
	if list == nil {
		return
	}

	// Find a node Kth from the head.
	last := list
	for i := 1; i <= K && last != nil; i++ {
		last = last.Next
	}

	// If there does not exist such a node, then it's an invalid request.
	if last == nil {
		return
	}

	// Move the head and Kth-from-head pointers together till the end.
	for last.Next != nil {
		list = list.Next
		last = last.Next
	}

	// Variable list now points to (K+1)th last node.

	// Adjust the (K+1)th's next pointer to discard the Kth last node.
	list.Next = list.Next.Next
}

func main() {
	list := InitList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	K := 4
	list.Print()
	fmt.Printf("->Delete(%v)->", K)
	DeleteKthLast(list, K)
	list.Print()
	fmt.Println()
}
