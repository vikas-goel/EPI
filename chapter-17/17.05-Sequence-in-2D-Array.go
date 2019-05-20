// Write a program that takes as arguments a 2D array and a ID array, and
// checks whether the ID array occurs in the 2D array.
// Valid moves from a given cell are immediate above, below, left, and right.
// It is acceptable to visit an entry in the grid more than once.

package main

import "fmt"

func Sequence(array [][]int, pattern []int) bool {
	if len(pattern) == 0 {
		return true
	}

	if len(array) == 0 {
		return false
	}

	numRows, numCols, numPatters := len(array), len(array[0]), len(pattern)

	memo := make(map[int]map[int]bool)

	var find func(row, col, index int) bool
	find = func(row, col, index int) bool {
		// All ekements of the pattern already found.
		if index == numPatters {
			return true
		}

		// Out of range.
		if row < 0 || col < 0 || row >= numRows || col >= numCols {
			return false
		}

		// Return the result if already checked the cell against the
		// given pattern index.
		hashIndex := (row - 1) * numRows + col
		if result, ok := memo[hashIndex][index]; ok {
			return result
		} else {
			memo[hashIndex] = make(map[int]bool)
		}

		// The pattern elements at the given index does not match
		// with the given cell of the array.
		if array[row][col] != pattern[index] {
			return false
		}

		// Check for next pattern element.
		index++

		// If there exists a patter for any of the moves from here.
		if find(row+1, col, index) || find(row-1, col, index) ||
		   find(row, col+1, index) || find(row, col-1, index) {
			memo[hashIndex][index] = true
			return true
		}

		memo[hashIndex][index] = false
		return false
	}

	// Check each cell of the array as a starting point.
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if find(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	array := [][]int{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}}

	fmt.Println("2D Array = ", array)
	for _, pattern := range[][]int{{1, 3, 4, 6}, {1, 2, 3, 4}} {
		fmt.Printf("Found pattern %v in the 2D array? %v\n",
			pattern, Sequence(array, pattern))
	}
}
