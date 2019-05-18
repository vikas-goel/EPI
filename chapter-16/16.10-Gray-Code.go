// An n-bit Gray code is a permutation of {0,1, 2,...,2" - 1) such that the
// binary representations of successive integers in the sequence differ in
// only one place. (This is with wraparound, i.e., the last and first elements
// must also differ in only one place.) For example, both
// <(000)2, (100)2,(101)2, (111)2,(110)2, (010)2,(011)2, (001)2> =
// (0, 4, 5, 7,6, 2,3, 1) and (0, 1,3, 2,6, 7, 5, 4) are Gray codes for n = 3.
//
// Write a program which takes n as input and returns an n-bit Gray code.

package main

import "fmt"

func GrayCode(num uint) (codes []int) {
	if num <= 0 {
		return
	} else if num == 1 {
		codes = append(codes, 0, 1)
		return
	}

	// Compute gray code for the previous number and then prefix 1 to each
	// code in the reverse order.
	codes = GrayCode(num - 1)
	for i := len(codes)-1; i >= 0 ; i-- {
		newCode := codes[i] | (1 << (num - 1))
		codes = append(codes, newCode)
	}

	return
}

func main() {
	fmt.Printf("%03b\n", GrayCode(3))
}
