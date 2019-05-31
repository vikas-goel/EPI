// Check whether a 9x9 2D array representing a partially completed Sudoku is
// valid. Specifically, check that no row, column, or 3 X 3 2D subarray
// contains duplicates. A 0-value in the 2D array indicates that entry is blank;
// every other entry is in [1,9].

package main

import "fmt"

func SudokuChecker(sudoku [][]int) (valid bool, row, col int) {
	grid := 3
	size := grid * grid

	if len(sudoku) != size && len(sudoku[0]) != size {
		return
	}

	for row = 0; row < size-1; row++ {
		for col = 0; col < size-1; col++ {
			// Empty cell.
			if sudoku[row][col] == 0 {
				continue
			}

			// Check remaining cells in the row.
			for i := row+1; i < size; i++ {
				if sudoku[row][col] == sudoku[i][col] {
					return
				}
			}

			// Check remaining cells in the column.
			for j := col+1; j < size; j++ {
				if sudoku[row][col] == sudoku[row][j] {
					return
				}
			}

			// Check remaining cells in the 3x3 subarray.
			rowStart, colStart := row/grid*grid, col/grid*grid
			for i := rowStart; i < rowStart+grid; i++ {
				for j := colStart; j < colStart+grid; j++ {
					if sudoku[row][col] == sudoku[i][j] {
						if i != row && j != col {
							return
						}
					}
				}
			}
		}
	}

	return true, 0, 0
}

func main() {
	sudoku := [][]int{{5,3,0,0,7,0,0,0,0}, {6,0,0,1,9,5,0,0,0},
		{0,9,8,0,0,0,0,6,0}, {8,0,0,0,6,0,0,0,3},
		{4,0,0,8,0,3,0,0,1}, {7,0,0,0,2,0,0,0,6},
		{0,6,0,0,0,0,2,8,0}, {0,0,0,4,1,9,0,0,5},
		{0,0,0,0,8,0,0,7,9}}
	valid, row, col := SudokuChecker(sudoku)
	fmt.Printf("%v\nIs valid Sudoku = %v (%v, %v)\n",
		sudoku, valid, row+1, col+1)
}
