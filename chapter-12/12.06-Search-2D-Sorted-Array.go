// Design an algorithm that takes a 2D sorted array and a number and checks
// whether that number appears in the array.

package main

import "fmt"

func Search2DSortedArray(number int, array [][]int) bool {
	if len(array) == 0 {
		return false
	}

	// Start from top-right cell and move left if the cell is bigger
	// than the number and down if the cell is smaller than the number.
	rowSize, colSize := len(array), len(array[0])
	for row, col := 0, colSize-1; row < rowSize && col >= 0; {
		if array[row][col] == number {
			return true
		} else if array[row][col] > number {
			col--
		} else { // array[row][col] < number
			row++
		}
	}

	return false
}

func main() {
	array := [][]int{{-1, 2, 4, 4, 6}, {1, 5, 5, 9, 21}, {3, 6, 6, 9, 22},
		{3, 6, 8, 10, 24}, {6, 8, 9, 12, 25}, {8, 10, 12, 13, 40}}
	fmt.Println("2D array", array)
	for i := 0; i < 15; i++ {
		fmt.Println("Exists(", i, ") ?", Search2DSortedArray(i, array))
	}
}
