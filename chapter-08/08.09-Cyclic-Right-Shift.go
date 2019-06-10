// Write a program that takes as input a singly linked list and a nonnegative
// integer k, and returns the list cyclically shifted to the right by k.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func RightShift(list *List, K int) *List {
	if K < 1 {
		return list
	}

	curHead, numNodes := list, 0

	// Find the Kth node from the head. This will be (K+1)th node in the
	// list.
	for i := 0; i < K && list != nil; i++ {
		numNodes++
		list = list.Next
	}

	if list == nil {
		// The entire list is traversed and the number of nodes is same
		// as the desired shift, then the shifting gives the original
		// list back.
		if numNodes == K {
			return curHead
		}

		// The number of desired shifts is greater than the number of
		// the nodes. Compute the effective desired shift count and
		// retraverse the list to get the Kth node from the head.
		K %= numNodes
		list = curHead
		for i := 0; i < K && list != nil; i++ {
			list = list.Next
		}
	}

	// Move the (K+1)th node till the node along with head node in
	// parallel to get the new last node in the shifted list.
	newLast := curHead
	for list.Next != nil {
		list = list.Next
		newLast = newLast.Next
	}

	// The new head of the shifted list will be next of the new last node
	// in the current list.
	newHead := newLast.Next

	// Adjust the pointers to complete the shifting.
	// The last element in the current list should point the head node of
	// the current list. The new last node should point to nil.
	list.Next, newLast.Next = curHead, nil

	return newHead
}

func main() {
	list := InitList(2, 3, 5, 3, 2, 4)
	K := 2
	list.Print()
	fmt.Printf("->Right Shift(%v)->", K)
	list = RightShift(list, K)
	list.Print()
	fmt.Println()
}
