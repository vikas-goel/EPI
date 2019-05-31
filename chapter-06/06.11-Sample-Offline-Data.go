// Implement an algorithm that takes as input an array of distinct elements and
// a size, and returns a subset of the given size of the array elements. All
// subsets should be equally likely. Return the result in input array itself.

package main

import (
	"fmt"
	"math/rand"
)

func RandomSubset(set []int, size int) {
	// Generate a random number and swap with the i-th index.
	for i := 0; i < size; i++ {
		random := rand.Intn(size-i) + i
		if i != random {
			set[i], set[random] = set[random], set[i]
		}
	}
}

func main() {
	for _, array := range [][]int{{3, 7, 5, 111}, {5, 6, 7, 8, 14, 15}} {
		size := len(array)/2+1
		fmt.Printf("Random subset(%v) %v = ", size, array)
		RandomSubset(array, size)
		fmt.Println(array[0:size])
	}
}
