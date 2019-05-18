// Write a program which takes as input a nonnegative integer n and returns the
// first n rows of Pascal's triangle.

package main

import "fmt"

func PascalTriangle(row, initial int) (triangle [][]int) {
	triangle = make([][]int, row)
	for i := 0; i < row; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangle[i][j] = initial
			} else {
				triangle[i][j] = triangle[i-1][j-1]+triangle[i-1][j]
			}
		}
	}

	return triangle
}

func nthPascalRow(n int) (row []int) {
	row = make([]int, n)

	row[0] = 1
	for i := 1; i < n; i++ {
		row[i] = 1
		for j := i-1; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return
}

func main() {
	fmt.Println(PascalTriangle(5, 1))
	fmt.Println("Row:5", nthPascalRow(5))
}
