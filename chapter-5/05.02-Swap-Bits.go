// Implement code that takes as input a 64-bit integer and swaps the bits at
// indices i and j.

package main

import "fmt"

func SwapBits(number int64, i, j uint) int64 {
	iBit := (number >> i) & 1
	jBit := (number >> j) & 1

	// If both bits are same, then the swap will make no difference.
	if iBit != jBit {
		// Reverse the ith and jth bits using XOR operation.
		number ^= (1 << j) | (1 << i)
	}

	return number
}

func main() {
	var x int64 = 37675876856878867
	var i, j uint = 23, 54
	fmt.Printf("%b (%v <-> %v)\n%b\n", x, i, j, SwapBits(x, i, j))
}
