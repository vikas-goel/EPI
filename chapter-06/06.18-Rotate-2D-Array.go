// Write a function that takes as input an n X n 2D array, and rotates the
// array by 90 degrees clockwise.

package main

import "fmt"

func RotateMatrix(matrix [][]int) {
	dimension := len(matrix)
	if dimension < 2 || len(matrix[0]) != dimension {
		return
	}

	for depth := 0; depth < dimension/2; depth++ {
		start, end := depth, dimension-depth-1
		for i := 0; i < end-start; i++ {
			temp := matrix[start][start+i]
			matrix[start][start+i] = matrix[end-i][start]
			matrix[end-i][start] = matrix[end][end-i]
			matrix[end][end-i] = matrix[start+i][end]
			matrix[start+i][end] = temp
		}
	}
}

func main() {
	for _, matrix := range [][][]int{
		{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}, {13,14,15,16}},
		{{1,2,3,4,5}, {6,7,8,9,10}, {11,12,13,14,15},
		{16,17,18,19,20}, {21,22,23,24,25}}} {
		fmt.Println(matrix)
		RotateMatrix(matrix)
		fmt.Println(" ->", matrix)
	}
}
