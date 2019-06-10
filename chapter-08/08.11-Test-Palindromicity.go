// Write a program that tests whether a singly linked list is palindromic.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
var InitList = list.InitList

func IsPalindromic(list *List) (result bool) {
	// Split the list into two halves. Reverse the second half.
	// Start from both ends and compare the nodes.
	// Once done, re-reverse the second half to maintain the original list.

	if list == nil {
		return false
	} else if list.Next == nil {
		return true
	}

	mid, last := list, list.Next
	for last != nil && last.Next != nil {
		mid, last = mid.Next, last.Next.Next
	}

	// The mid pointer at this stage points to
	// the middle node if there are odd number of nodes and
	// last node of the first half if there are even number of nodes.
	// So, reverse the list starting next of the mid pointer.
	frontHead, rearHead := list, reverseList(mid.Next)
	midNext, tail := mid.Next, rearHead

	midNext.Next = nil
	result = true
	for frontHead != nil && rearHead != nil && frontHead != rearHead {
		if frontHead.Value != rearHead.Value {
			result = false
			break
		}

		frontHead, rearHead = frontHead.Next, rearHead.Next
	}

	// Reverse the second half again to get revert the changes.
	reverseList(tail)
	mid.Next = midNext

	return result
}

func reverseList(list *List) (reverse *List) {
	if list == nil || list.Next == nil {
		return list
	}

	left := list
	list = list.Next
	left.Next = nil

	for list.Next != nil {
		right := list.Next
		list.Next = left
		left = list
		list = right
	}

	list.Next = left

	return list
}

func main() {
	list := make([]*List, 0)
	list = append(list, InitList(0))
	list = append(list, InitList(0, 1))
	list = append(list, InitList(1, 1))
	list = append(list, InitList(0, 1, 0))
	list = append(list, InitList(0, 1, 1))
	list = append(list, InitList(0, 1, 2, 3, 4, 3, 2, 1, 0))
	list = append(list, InitList(0, 1, 2, 3, 4, 4, 3, 2, 1, 0))
	list = append(list, InitList(0, 1, 2, 3, 4, 5, 3, 2, 1, 0))
	list = append(list, InitList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))

	for _, l := range list {
		fmt.Print("Is Palindromic ")
		l.Print()
		result := IsPalindromic(l)
		fmt.Print("? ", result, " ")
		l.Print()
		fmt.Println()
	}
}
