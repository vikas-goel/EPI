// You are given an array of n integers, each between 0 and n -1, inclusive.
// Exactly one element appears twice, implying that exactly one number between
// 0 and n-1 is missing from the array. Compute the duplicate & missing numbers.

package main

import "fmt"

func Find(array []int) (missing, duplicate int) {
	if len(array) == 0 {
		return
	}

	// 1. Compute XOR of array = xorA
	// 2. Compute XOR of 1..n = xorN
	// 3. Compute xorA and xorN to find XOR of missing & duplicate = xorNums
	// 4. Find the least significant bit in xorNums that is 1
	// 5. Find XOR of array elements with bit 1 on = xorKbit
	// 6. Process array to see whether xorKbit is missing or duplicate
	// 7. XOR xorKbit and xorNums to find the other desired number

	xorA, xorN, xorNums, xorKbit := 0, 0, 0, 0

	// Step 1 and 2
	for i := 0; i < len(array); i++ {
		xorA ^= array[i]
		xorN ^= i
	}

	// Step 3
	xorNums = xorA ^ xorN

	// Step 4
	kBit, copyXorNums := 1, xorNums
	for kBit = 1; copyXorNums & 1 == 0; kBit <<= 1 {
		copyXorNums >>= 1
	}

	// Step 5
	for i := 0; i < len(array); i++ {
		if array[i] & kBit != 0 {
			xorKbit ^= array[i]
		}
	}

	// Step 6
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] == xorKbit {
			count++
		}
	}

	// Step 7
	if count == 0 {
		missing = xorKbit
		duplicate = xorNums ^ xorKbit
	} else if count == 2 {
		duplicate = xorKbit
		missing = xorNums ^ xorKbit
	} else {
		panic("Invalid array")
	}

	return
}

func main() {
	array := []int{5, 3, 0, 3, 1, 2}
	missing, duplicate := Find(array)
	fmt.Printf("%v: Missing=%v, Duplicate=%v\n", array, missing, duplicate)
}
