// Write a program that counts how many ways you can go from the top-left to
// the bottom-right in a 2D array.

package main

import "fmt"

type Cell struct {
	row, col int
}

func Traverse(rows, columns int) int {
	if rows == 0 || columns == 0 {
		return 0
	}

	paths := make([][]int, rows)
	for i := 0; i < rows; i++ {
		paths[i] = make([]int, columns)
	}

	return NumWays(paths, rows-1, columns-1)
}

func NumWays(paths [][]int, row, col int) int {
	if row < 0 || col < 0 {
		return 0
	} else if paths[row][col] > 0 {
		return paths[row][col]
	} else if row == 0 && col == 0 {
		paths[row][col] = 1
		return paths[row][col]
	}

	waysLeft := NumWays(paths, row-1, col)
	waysUp := NumWays(paths, row, col-1)
	if waysLeft + waysUp > 0 {
		paths[row][col] = waysLeft + waysUp
	}

	return paths[row][col]
}

func main() {
	rows, columns := 5, 5
	fmt.Printf("Number of traversal ways for %vx%v is %v.\n",
		rows, columns, Traverse(rows, columns))
}
