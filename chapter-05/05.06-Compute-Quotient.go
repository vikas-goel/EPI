// Given two positive integers, compute their quotient, using only the addition,
// subtraction, and shifting operators.

package main

import "fmt"

type Nums struct {
	dividend, divisor uint
}

func Quotient(dividend, divisor uint) (quotient uint) {
	// Base cases.
	if dividend == 0 || divisor == 0 || divisor > dividend {
		return 0
	} else if divisor == 1 {
		return dividend
	}

	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}

	return
}

func main() {
	for _, n := range []Nums{{8, 7}, {12, 5}, {15,3}, {0, 6}} {
		fmt.Println("Quotient", n, "=", Quotient(n.dividend, n.divisor))
	}
}
