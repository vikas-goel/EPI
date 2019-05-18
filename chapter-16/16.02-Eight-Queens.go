// Write a program which returns all distinct nonattacking placements of n
// queens on an nxn chessboard, where n is an input to the program.

package main

import "fmt"

func CandidateBoards(board []int, start, end int, total, valid *int) {
	if start == end {
		*total++

		if validBoard(board) {
			fmt.Println(board)
			*valid++
		}
	} else {
		for i := start; i <= end; i++ {
			swapPositions(board, start, i)
			CandidateBoards(board, start+1, end, total, valid)
			swapPositions(board, start, i)
		}
	}
}

func validBoard(board []int) bool {
	length := len(board)
	for i := length-8; i < length-1; i++ {
		for j := i+1; j < length; j++ {
			if ! validPositions(i, j, board[i], board[j]) {
				return false
			}
		}
	}

	return true
}

func validPositions(row1, row2, col1, col2 int) bool {
	row := int8(row1 - row2)
	col := int8(col1 - col2)

	if (row == col) || (row == -1 * col) {
		return false
	}

	return true
}

func swapPositions(board []int, row1, row2 int) {
	board[row1], board[row2] = board[row2], board[row1]
}

func main() {
	initBoard := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	total, valid := 0, 0

	low, high := len(initBoard)-8, len(initBoard)-1

	CandidateBoards(initBoard, low, high, &total, &valid)

	fmt.Println("Valid boards:", valid, "of", total)
}
