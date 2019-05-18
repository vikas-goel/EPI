// Write a program which takes an nxn 2D array and returns the spiral ordering
// of the array.

package main

import (
	"fmt"
	"math"
)

// Create a spiral-ordered array from a nxn 2D matrix.
func MatrixToSpiral(matrix [][]int) (spiral []int) {
	spiral = make([]int, 0, 2*len(matrix))

	var writeSpiral func(int)
	writeSpiral = func(depth int) {
		start, end := depth, len(matrix)-depth-1

		if start > end {
			return
		} else if start == end {
			spiral = append(spiral, matrix[start][start])
			return
		}

		for i := start; i < end; i++ {
			spiral = append(spiral, matrix[start][i])
		}

		for i := start; i < end; i++ {
			spiral = append(spiral, matrix[i][end])
		}

		for i := end; i > start; i-- {
			spiral = append(spiral, matrix[end][i])
		}

		for i := end; i > start; i-- {
			spiral = append(spiral, matrix[i][start])
		}

		writeSpiral(depth+1)
	}

	writeSpiral(0)
	return
}

// Create a nxn 2D matrix from a spiral-ordered array.
func SpiralToMatrix(spiral []int) (matrix [][]int) {
	dimension := int(math.Sqrt(float64(len(spiral))))

	matrix = make([][]int, dimension)
	for i := 0; i < dimension; i++ {
		matrix[i] = make([]int, dimension)
	}

	var writeSpiral func(int, int)
	writeSpiral = func(index, depth int) {
		start, end := depth, len(matrix)-depth-1

		if start > end {
			return
		} else if start == end {
			matrix[start][start] = spiral[index]
			return
		}

		for i := start; i < end; i++ {
			matrix[start][i] = spiral[index]
			index++
		}

		for i := start; i < end; i++ {
			matrix[i][end] = spiral[index]
			index++
		}

		for i := end; i > start; i-- {
			matrix[end][i] = spiral[index]
			index++
		}

		for i := end; i > start; i-- {
			matrix[i][start] = spiral[index]
			index++
		}

		writeSpiral(index, depth+1)
	}

	writeSpiral(0, 0)
	return
}

func main() {
	a := [][]int{{1, 2, 3, 4},{5, 6, 7, 8},{9, 10, 11, 12},{13, 14, 15, 16}}
	as := MatrixToSpiral(a)
	am := SpiralToMatrix(as)
	fmt.Println(a, "->", as)
	fmt.Println(as, "->", am)

	b := [][]int{{1, 2, 3},{4, 5, 6},{7, 8, 9}}
	bs := MatrixToSpiral(b)
	bm := SpiralToMatrix(bs)
	fmt.Println(b, "->", bs)
	fmt.Println(bs, "->", bm)
}
