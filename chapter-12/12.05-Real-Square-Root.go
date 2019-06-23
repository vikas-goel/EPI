// Implement a function which takes as input a floating point value and returns
// its square root.

package main

import "fmt"

func SquareRoot(number float64) (sqroot float64) {
	tolerance := 0.00001

	start, end := 1.0, number
	if number < 1.0 {
		start, end = number, 1.0
	}

	for start < end {
		mid := start + 0.5 * (end - start)
		square := mid * mid

		diff := square - number
		if diff < -1 * tolerance {
			start = mid
		} else if square > number {
			end = mid
		} else {
			sqroot = mid
			break
		}
	}

	return
}

func main() {
	for _, num := range []float64{0.25, 8.9, 9.0} {
		fmt.Println("Square root(", num, ") =", SquareRoot(num))
	}
}
