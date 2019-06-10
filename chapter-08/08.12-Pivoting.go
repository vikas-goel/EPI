// Implement a function which takes as input a singly linked list and an
// integer k and performs a pivot of the list with respect to k. The relative
// ordering of nodes that appear before k, and after k, must remain unchanged;
// the same must hold for nodes holding keys equal to k.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func Pivot(list *List, k int) *List {
	// Maintain 3 separate lists for lesser, equal, and greater values.
	var before, after, pivot List
	beforeTail, afterTail, pivotTail := &before, &after, &pivot

	// Keep appending to the correct list so as to maintain the order.
	for ; list != nil; list = list.Next {
		if list.Value < k {
			beforeTail.Next = list
			beforeTail = beforeTail.Next
		} else if list.Value == k {
			pivotTail.Next = list
			pivotTail = pivotTail.Next
		} else {
			afterTail.Next = list
			afterTail = afterTail.Next
		}
	}

	// Join the three lists in the correct order.
	afterTail.Next = nil
	pivotTail.Next = after.Next
	beforeTail.Next = pivot.Next

	return before.Next
}

func main() {
	pivot := 7
	list := make([]*List, 0)
	list = append(list, InitList(3, 2, 2, 11, 5, 11))
	list = append(list, InitList(7, 3, 5, 2, 1, 7, 2, 7))
	list = append(list, InitList(7, 7, 8, 7, 11, 9, 7))
	list = append(list, InitList(3, 2, 2, 11, 7, 5, 11))
	list = append(list, InitList(7, 3, 5, 2, 1, 7, 8, 7, 11, 9, 2, 7))

	for _, l := range list {
		l.Print()
		fmt.Printf("->Pivot(%v)->", pivot)
		l = Pivot(l, pivot)
		l.Print()
		fmt.Println()
	}
}
