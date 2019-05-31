// Write a program which takes as input an array of digits encoding a decimal
// number D and updates the array to represent the number D + 1. For example,
// if the input is (1,2,9) then you should update the array to (1,3,0). Your
// algorithm should work even if it is implemented in a language that has
// finite-precision arithmetic.

package main

import "fmt"

func Increment(digits *[]int) {
	if digits == nil || len(*digits) == 0 {
		return
	}

	// Start from right end. Keep updating the digits as needed.
	carry := true
	thisDigits := *digits
	for i := len(thisDigits)-1; i >= 0 && carry; i-- {
		if thisDigits[i] == 9 {
			thisDigits[i] = 0
		} else {
			carry = false
			thisDigits[i] += 1
		}
	}

	// If carry flag is on, then add another digit in the beginning.
	if carry {
		newDigits := make([]int, len(thisDigits)+1)
		newDigits[0] = 1
		copy(newDigits[1:], thisDigits)

		*digits = newDigits
	}
}

func main() {
	for _, array := range [][]int{{9, 9}, {1, 2, 9}} {
		fmt.Print(array, "->")
		Increment(&array)
		fmt.Println(array)
	}
}
