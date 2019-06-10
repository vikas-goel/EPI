// Write a program which deletes a node in a singly linked list. The input node
// is guaranteed not to be the tail node.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
type Node = list.List
var InitList = list.InitList

func DeleteNode(node *Node) {
	// Do nothing if the node is nil pointer or tail.
	if node == nil || node.Next == nil {
		return
	}

	// Copy the successor node's value and next pointer to the one
	// supposed to be deleted and discard the successor node.
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}

func main() {
	list := InitList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	node := list.Next.Next.Next
	list.Print()
	fmt.Printf("->Delete(%v)->", node.Value)
	DeleteNode(node)
	list.Print()
	fmt.Println()
}
