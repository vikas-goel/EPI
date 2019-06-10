// Write a program that takes as input a singly linked list of integers in
// sorted order, and removes duplicates from it. The list should be sorted.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func RemoveDuplicates(list *List) {
	// Do nothing if the node is nil pointer or tail.
	if list == nil {
		return
	}

	for list.Next != nil {
		if list.Value == list.Next.Value {
			list.Next = list.Next.Next
		} else {
			list = list.Next
		}
	}
}

func main() {
	list := InitList(2, 2, 3, 5, 5, 5, 7, 11, 11)
	list.Print()
	fmt.Print("->Remove Duplicates->")
	RemoveDuplicates(list)
	list.Print()
	fmt.Println()
}
