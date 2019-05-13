// The parity of a binary word is 1 if the number of 1s in the word is odd;
// otherwise, it is 0. Compute parity of a very large number of 64-bit words.

package main

import "fmt"

func Parity(num int64) bool {
	num = (num >> 32) ^ (num & 0xFFFFFFFF)
	num = (num >> 16) ^ (num & 0xFFFF)
	num = (num >> 8) ^ (num & 0xFF)
	num = (num >> 4) ^ (num & 0xF)
	num = (num >> 2) ^ (num & 3)
	num = (num >> 1) ^ (num & 1)
	return num == 1
}


func main() {
	for _, x := range []int64{2, 7, 20, 35} {
		fmt.Printf("%v:%b:%v\n", x, x, Parity(x))
	}
}
