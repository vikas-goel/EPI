// The diameter of a tree is defined to be the length of a longest path in the
// tree. Design an efficient algorithm to compute the diameter of a tree.

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	edges []*Edge
}

type Edge struct {
	length float64
	root *TreeNode
}

func ComputeDiameter(t *TreeNode) float64 {
	if t == nil {
		return 0.0
	}

	_, diameter := ComputeDiameterAndHeight(t)
	return diameter
}

func ComputeDiameterAndHeight(t *TreeNode) (float64, float64) {
	maxHeights := make([]float64, 2)
	maxDiameter := math.SmallestNonzeroFloat64

	for _, e := range t.edges {
		height, diameter := ComputeDiameterAndHeight(e.root)
		if height + e.length > maxHeights[0] {
			maxHeights[1] = maxHeights[0]
			maxHeights[0] = height + e.length
		} else if height + e.length > maxHeights[1] {
			maxHeights[1] = height + e.length
		}

		if diameter > maxDiameter {
			maxDiameter = diameter
		}
	}

	if maxHeights[0] + maxHeights[1] > maxDiameter {
		maxDiameter = maxHeights[0] + maxHeights[1]
	}

	return maxHeights[0], maxDiameter
}

func main() {
	Leaf := &TreeNode{}

	EE := &Edge{6.0, Leaf}

	TD := &TreeNode{[]*Edge{EE}}
	ED := &Edge{4.0, TD}

	ECr := &Edge{3.0, Leaf}

	TC := &TreeNode{[]*Edge{ED, ECr}}
	EC := &Edge{7.0, TC}

	EA := &Edge{14.0, Leaf}

	EBrl := &Edge{2.0, Leaf}
	EBrrl := &Edge{6.0, Leaf}
	EBrrrl := &Edge{4.0, Leaf}
	EBrrrrl := &Edge{1.0, Leaf}
	EBrrrrm := &Edge{2.0, Leaf}
	EBrrrrr := &Edge{3.0, Leaf}

	TBrrrr := &TreeNode{[]*Edge{EBrrrrl, EBrrrrm, EBrrrrr}}
	EBrrrr := &Edge{2.0, TBrrrr}

	TBrrr := &TreeNode{[]*Edge{EBrrrl, EBrrrr}}
	EBrrr := &Edge{4.0, TBrrr}

	TBrr := &TreeNode{[]*Edge{EBrrl, EBrrr}}
	EBrr := &Edge{1.0, TBrr}

	TBr := &TreeNode{[]*Edge{EBrl, EBrr}}
	EBr := &Edge{3.0, TBr}

	TB := &TreeNode{[]*Edge{EC, EA, EBr}}

	fmt.Println("Diameter of the tree =", ComputeDiameter(TB))
}
