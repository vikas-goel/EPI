// Write a program that takes two lists, assumed to be sorted, and returns
// their merge. The only field your program can change in a node is its next
// field.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func Merge(list1, list2 *List) *List {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var head List

	current := &head
	for list1 != nil && list2 != nil {
		if list1.Value <= list2.Value {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return head.Next
}

func main() {
	list1, list2 := InitList(2, 5, 7), InitList(3, 11)

	list1.Print()
	fmt.Print("+")
	list2.Print()
	fmt.Print("=")

	list := Merge(list1, list2)
	list.Print()
}
