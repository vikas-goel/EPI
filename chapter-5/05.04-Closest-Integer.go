// Write a program which takes as input a nonnegative integer x and returns a
// number y which is not equal to x, but has the same weight as x and their
// difference, |y- x|, is as small as possible. You can assume x is not 0, or
// all 1s.

package main

import "fmt"

// The closet number can be either side, i.e. smaller or larger, of the given
// number.
func GetNext(number int64) (next int64) {
	next = number
	currentBit := number & 1

	for i := uint(0); number != 0; i++ {
		number >>= 1
		if number & 1 != currentBit {
			// Swap the two adjacent bits.
			next ^= (1 << i) | (1 << (i+1))
			break
		}
	}

	return
}

func main() {
	var num int64 = 13948
	next := GetNext(num)
	fmt.Printf("Next(%d,%b) = (%d,%b)\n", num, num, next, next)
}
