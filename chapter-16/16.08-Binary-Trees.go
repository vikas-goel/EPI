// Write a program which returns all distinct binary trees with a specified
// number of nodes. For example, if the number of nodes is specified to be
// three, return all 5 trees.

package main

import "fmt"

type Node struct {
	data int
	left, right *Node
}

func makeTree(data int, left, right *Node) *Node {
	root := new(Node)
	root.data = data

	if left != nil {
		root.left = new(Node)
		root.left.data = left.data
		root.left.left = left.left
		root.left.right = left.right
	}

	if right != nil {
		root.right = new(Node)
		root.right.data = right.data
		root.right.left = right.left
		root.right.right = right.right
	}

	return root
}

func printTree(root *Node, nodeType rune) {
	if root == nil {
		return
	}
	printTree(root.left, 'L')
	fmt.Printf("%c(%v) ", nodeType, root.data)
	printTree(root.right, 'R')
}

func GenerateBinaryTrees(totalNodes int) []*Node {
	var generate func(numNodes int) (result []*Node)
	generate = func(numNodes int) (result []*Node) {
		if numNodes == 0 {
			result = append(result, nil)
			return
		}

		// Get all possible left and right sub-trees and combine them.
		for i := 0; i < numNodes; i++ {
			leftTrees := generate(i)
			rightTrees := generate(numNodes-1-i)

			for l := 0; l < len(leftTrees); l++ {
				for r := 0; r < len(rightTrees); r++ {
					result = append(result, makeTree(numNodes, leftTrees[l], rightTrees[r]))
				}
			}
		}

		return
	}

	return generate(totalNodes)
}

func main() {
	allTrees := GenerateBinaryTrees(5)
	for i := 0; i < len(allTrees); i++ {
		fmt.Printf("{ ")
		printTree(allTrees[i], 'P')
		fmt.Println("}")
	}
	fmt.Printf("Total number of trees = %v\n", len(allTrees))
}
