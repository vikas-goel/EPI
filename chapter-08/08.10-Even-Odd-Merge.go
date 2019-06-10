// Consider a singly linked list whose nodesare numbered starting at 0. Define
// the evenodd merge of the list to be the list consisting of the even-numbered
// nodes followed by the odd-numbered nodes.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func EvenOddMerge(list *List) {
	if list == nil {
		return
	}

	var oddHead List

	for oddList := &oddHead; list.Next != nil; {
		// Move odd nodes to a separate odd-list.
		oddList.Next = list.Next

		// Connect the current and next even nodes.
		list.Next = list.Next.Next

		// Shift the odd list pointer to the newly connected one.
		oddList = oddList.Next
		oddList.Next = nil

		// Shift the even list pointer to the newly connected one.
		if list.Next != nil {
			list = list.Next
		}
	}

	// Append the odd-list to the end of the even-list.
	list.Next = oddHead.Next
}

func main() {
	list := InitList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	list.Print()
	fmt.Print("->Even-Odd-Merge->")
	EvenOddMerge(list)
	list.Print()
	fmt.Println()
}
