// Write a program that takes the head of a singly linked list and returns null
// if there does not exist a cycle, and the node at the start of the cycle, if
// a cycle is present. (You do not know the length of the list in advance.)

package main

import (
	"fmt"
	"./list"
)

func main() {
	list := list.InitList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	last := list.Last()
	last.Next = list.Next.Next

	fmt.Printf("List is cyclic List @ %v.\n", list.CyclicNode())
}
