// Write a program that multiplies two nonnegative integers. The only operators
// you are allowed to use are
// - assignment,
// - the bitwise operators >>, <<, |, &, ~, ^ and
// - equality checks and Boolean combinations thereof.
// You may use loops and functions that you write yourself. These constraints
// imply, for example, that you cannot use increment or decrement, or test
// if x < y.

package main

import "fmt"

type Nums struct {
	multiplier, number uint
}

func Multiply(multiplier, number uint) (result uint) {
	// Base cases.
	if multiplier == 0 || number == 0 {
		return 0
	} else if multiplier == 1 {
		return number
	} else if number == 1 {
		return multiplier
	}

	// The add factor converts the multiplier from odd to even so that
	// the number can be doubled and the mulitplier can be halved.
	var oddFactor uint
	if multiplier & 1 == 1 {
		oddFactor = number
	}

	// Double the number, reduce the multiplier to half, and add the odd
	// factor. That is, if m is
	//	even: m x n = m/2 * 2*n
	//	 odd: m x n = m/2 * 2*n + n
	result = Multiply(multiplier >> 1, number << 1)

	// Add oddFactor.
	result = Add(result, oddFactor)

	return
}

func Add(x, y uint) (result uint) {
	var cin, k uint = 0, 1

	for tx, ty := x, y; tx != 0 || ty != 0; tx, ty = tx >> 1, ty >> 1 {
		// Get kth bits of the two operands.
		xk, yk := x & k, y & k

		// Get carry-out factor.
		cout := (xk & yk) | (xk & cin) | (yk & cin)

		// Add the three numbers in the result and compute carry forward.
		result |= (xk ^ yk ^ cin)
		cin, k = cout << 1, k << 1
	}

	result |= cin
	return
}

func main() {
	for _, n := range []Nums{{8, 7}, {4, 6}, {15,9}, {0, 6}} {
		fmt.Println("Multiply", n, "=", Multiply(n.multiplier, n.number))
	}
}
