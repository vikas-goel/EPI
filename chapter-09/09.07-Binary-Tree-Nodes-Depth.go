// Given a binary tree, return an array consisting of the keys at the same
// level. Keys should appear in the order of the corresponding nodes' depths,
// breaking ties from left to right.

package main

import (
	"fmt"
	"container/list"
)

type Node = Tree

type Tree struct {
	Value interface{}
	Left, Right *Tree
}

func NodeDepths(node *Node) [][]interface{} {
	if node == nil {
		return nil
	}

	queue := list.New()
	queue.Init()
	queue.PushBack(node)

	depthNodes := make([][]interface{}, 0)

	// Traverse the nodes at the each depth and keep the count of nodes.
	for depth := 0; queue.Len() > 0; depth++ {
		length := queue.Len()
		nodes := make([]interface{}, length)
		depthNodes = append(depthNodes, nodes)

		// For each node at the current depth, add its value to the
		// depth-node list and add the children to the queue.
		for i := 0; i < length; i++ {
			front := queue.Front()
			queue.Remove(front)
			nodes[i] = front.Value.(*Node).Value

			next := front.Value.(*Node)

			if next.Left != nil {
				queue.PushBack(next.Left)
			}

			if next.Right != nil {
				queue.PushBack(next.Right)
			}
		}
	}

	return depthNodes
}

func Tree1() *Tree {
	A, B, C := &Tree{Value:314}, &Tree{Value:6}, &Tree{Value:271}
	D, E, F := &Tree{Value:28}, &Tree{Value:0}, &Tree{Value:561}
	G, H, I := &Tree{Value:3}, &Tree{Value:17}, &Tree{Value:6}
	J, K, L := &Tree{Value:2}, &Tree{Value:1}, &Tree{Value:401}
	M, N, O := &Tree{Value:641}, &Tree{Value:257}, &Tree{Value:271}
	P := &Tree{Value:28}

	A.Left, A.Right = B, I
	B.Left, B.Right = C, F
	C.Left, C.Right = D, E
	F.Left, F.Right = nil, G
	G.Left, G.Right = H, nil
	I.Left, I.Right = J, O
	J.Left, J.Right = nil, K
	K.Left, K.Right = L, N
	L.Left, L.Right = nil, M
	O.Left, O.Right = nil, P

	return A
}

func main() {
	fmt.Println(NodeDepths(Tree1()))
}
