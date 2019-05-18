// Implement a Sudoku solver.

package main

import (
	"fmt"
	"math"
)

func SudokuSolver(matrix [][]int) {
	const EmptyCell = 0

	matrixSize := len(matrix)
	if len(matrix[0]) != matrixSize {
		fmt.Println("Invalid Sudoku")
		return
	}

	zoneSize := int(math.Sqrt(float64(matrixSize)))

	validCell := func(row, col, cell int) bool {
		// Check each cell of the corresponding row and column.
		for i := 0; i < matrixSize; i++ {
			if matrix[i][col] == cell || matrix[row][i] == cell {
				return false
			}
		}

		// Check each cell of the sub-matrix.
		iOffset, jOffset := row/zoneSize, col/zoneSize
		for i := iOffset*zoneSize; i < (iOffset+1)*zoneSize; i++ {
			for j := 0; j < zoneSize; j++ {
				if matrix[i][jOffset*zoneSize+j] == cell {
					return false
				}
			}
		}

		return true
	}

	var sudokuHelper func(int, int) bool
	sudokuHelper = func(row, col int) bool {
		if col == matrixSize {
			// All columns of this row is already taken care of.
			// So, move to the next row, 1st column.
			col, row = 0, row+1

			if row == matrixSize {
				// All empty cells are filled with valid
				// numbers.
				return true
			}
		}

		if matrix[row][col] != EmptyCell {
			return sudokuHelper(row, col+1)
		}

		for cell := 1; cell <= matrixSize; cell++ {
			if validCell(row, col, cell) {
				matrix[row][col] = cell
				if sudokuHelper(row, col+1) {
					return true
				}
				matrix[row][col] = EmptyCell
			}
		}

		return false
	}

	sudokuHelper(0, 0)
}

func main() {
	matrix := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9}}
	fmt.Println(matrix)
	SudokuSolver(matrix)
	fmt.Println(matrix)
}
