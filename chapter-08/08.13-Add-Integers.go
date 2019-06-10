// Write a program which takes two singly linked lists of digits, and returns
// the list corresponding to the sum of the integers they represent. The least
// significant digit comes first.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
type Node = list.List
var InitList = list.InitList

func Add(list1, list2 *List) *List {
	carry := 0
	var sum List
	tail := &sum

	// Add each respective node and track carry forward.
	for list1 != nil && list2 != nil {
		addValue := list1.Value + list2.Value + carry
		if addValue > 9 {
			addValue -= 10
			carry = 1
		} else {
			carry = 0
		}

		tail.Next = &Node{Value: addValue}

		list1, list2, tail = list1.Next, list2.Next, tail.Next
	}

	remaining := list1
	if list2 != nil {
		remaining = list2
	}

	// If one of the list has remaining digits, then add carry in it.
	for remaining != nil {
		addValue := remaining.Value + carry
		if addValue > 9 {
			addValue -= 10
			carry = 1
		} else {
			carry = 0
		}

		tail.Next = &Node{Value: addValue}
		tail = tail.Next
	}

	// Check whether carry forward is left.
	if carry == 1 {
		tail.Next = &Node{Value: 1}
	}

	return sum.Next
}

func main() {
	list1, list2 := InitList(3, 1, 4), InitList(7, 0, 9)
	sum := Add(list1, list2)
	list1.Print()
	fmt.Print("+")
	list2.Print()
	fmt.Print("=")
	sum.Print()
	fmt.Println()
}
