// Define a path in the triangle to be a sequence of entries in the triangle
// in which adjacent entries in the sequence correspond to entries that are
// adjacent in the triangle. The path must start at the top, descend the
// triangle continuously, and end with an entry on the bottom row. The weight
// of a path is the sum of the entries.
// Write a program that takes as input a triangle of numbers and returns the
// weight of a minimum weight path.

package main

import "fmt"

func MinimumWeight(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	memo := make(map[int]int)

	var compute func(row, col int) int
	compute = func(row, col int) int {
		if row >= len(triangle) {
			return 0
		}

		hashIndex := row*(row+1)/2+col
		if w, ok := memo[hashIndex]; ok {
			return w
		}

		if row == len(triangle)-1 {
			memo[hashIndex] = triangle[row][col]
		} else {
			w1, w2 := compute(row+1, col), compute(row+1, col+1)
			if w1 < w2 {
				memo[hashIndex] = w1
			} else {
				memo[hashIndex] = w2
			}
			memo[hashIndex] += triangle[row][col]
		}

		return memo[hashIndex]
	}

	return compute(0, 0)
}

func main() {
	triangle := [][]int{{2}, {4,4}, {8,5,6}, {4,2,6,2}, {1,5,2,3,4}}
	fmt.Printf("Minimum weight path of %v is %v.\n",
		triangle, MinimumWeight(triangle))
}
