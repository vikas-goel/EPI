// Write a program that takes two cycle-free singly linked lists, and
// determines if there exists a node that is common to both lists.

package main

import (
	"fmt"
	"./list"
)

var InitList = list.InitList
var OverlappingNode = list.OverlappingNode

func main() {
	list1 := InitList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	list2 := InitList(0, 7, 8, 9)

	fmt.Printf("Overlap @ %v.\n", OverlappingNode(list1, list2))
	fmt.Printf("Overlap @ %v.\n", OverlappingNode(list1, list1.Next))
}
