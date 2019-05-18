// Write a program that takes a double x and an integer y and returns x-power-y.
// You can ignore overflow and underflow.

package main

import "fmt"

type Nums struct {
	number float64
	power int
}

func Exponent(number float64, power int) (result float64) {
	result = 1.0

	// Base cases.
	if power == 0 {
		return 1.0
	} else if power == 1 {
		return number
	}

	if power < 0 {
		power = -power
		number = 1.0/number
	}

	for power != 0 {
		// For odd bit, multiply the number.
		if power & 1 == 1 {
			result *= number
		}

		// Multiple the number to self and reduce the power to half.
		power >>= 1
		number *= number
	}

	return
}

func main() {
	for _, n := range []Nums{{2.0, 3}, {3.0, -2}, {-4.0, 2}} {
		fmt.Println("Exponential", n, "=", Exponent(n.number, n.power))
	}
}
