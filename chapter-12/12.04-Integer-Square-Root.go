// Write a program which takes a nonnegative integer and returns the largest
// integer whose square is less than or equal to the given integer. For example,
// if the input is 16, return 4; if the input is 300, return 17, since
// 172 = 289 < 300 and 182 = 324 > 300.

package main

import "fmt"

func SquareRoot(number int) (sqroot int) {
	for start, end := 0, number; start <= end; {
		mid := start + (end - start) / 2
		square := mid * mid

		if square == number {
			sqroot = mid
			break
		} else if square > number {
			end = mid - 1
		} else { // square < number
			sqroot = mid
			start = mid + 1
		}
	}

	return
}

func main() {
	for _, num := range []int{16, 21, 25, 300} {
		fmt.Println("Square root(", num, ") =", SquareRoot(num))
	}
}
