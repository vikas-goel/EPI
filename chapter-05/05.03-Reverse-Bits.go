// Write a program that takes a 64-bit word and returns the 64-bit word
// consisting of the bits of the input word in reverse order.

package main

import "fmt"

func ReverseBits(number uint64) (reverse uint64) {
	for i, numBits := 0, 63; number != 0; i++ {
		// If the last bit of the number is 1, then prepare
		// corresponding reverse bit.
		if number & 1 == 1 {
			reverse |= (1 << uint(numBits-i))
		}
		number >>= 1
	}

	return
}

func main() {
	var x uint64 = 37675876856878867
	fmt.Printf("%b <->\n%b\n", x, ReverseBits(x))
}
