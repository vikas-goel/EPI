// Solve Problem 8.4 for the case where the lists may each or both have a cycle.
// If such a node exists, return a node that appears first when traversing the
// lists. This node may not be unique if one node ends in a cycle, the first
// cycle node encountered when traversing it may be different from the first
// cycle node encountered when traversing the second list, even though the
// cycle is the same. In such cases, you may return either of the two nodes.

package main

import (
	"fmt"
	"./list"
)

type List = list.List
type Node = list.List
var InitList = list.InitList

func OverlappingNode(list1, list2 *List) *Node {
	root1, root2 := list1.CyclicNode(), list2.CyclicNode()

	if root1 == nil && root2 == nil {
		// None of them is cyclic.
		return list.OverlappingNode(list1, list2)
	} else if (root1 == nil && root2 != nil) ||
		(root1 != nil && root2 == nil) {
		// Omly one of them is cyclic. So, no overlap.
		return nil
	}

	// Both are cyclic.
	temp := root1.Next
	for ; temp != root1 && temp != root2; temp = temp.Next {}

	if temp != root2 {
		// Disjoint cycles.
		return nil
	}

	distance1, distance2 := 0, 0
	for temp := list1; temp != root1; temp = temp.Next {
		distance1++
	}

	for temp := list2; temp != root2; temp = temp.Next {
		distance2++
	}

	if distance1 > distance2 {
		for i := 0; i < distance1-distance2; i++ {
			list1 = list1.Next
		}
	} else if distance2 > distance1 {
		for i := 0; i < distance2-distance1; i++ {
			list2 = list2.Next
		}
	}

	for list1 != list2 && list1 != root1 && list2 != root2 {
		list1, list2 = list1.Next, list2.Next
	}

	// Both lists overlap before the cycle starts.
	if list1 == list2 {
		return list1
	}

	// Overlap is after cycle starts. Return either of the cyclic nodes.
	return root1
}

func main() {
	list1 := InitList(1, 2, 3, 4, 5, 6, 7, 8, 9)
	list1.Last().Next = list1.Next.Next

	list2 := InitList(0, 7, 8, 9)
	list2.Last().Next = list1.Next

	fmt.Printf("Overlap Node @ %v.\n", OverlappingNode(list1, list2))
}
