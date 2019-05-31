// Design a program that takes as input a size k, and reads packets,
// continuously maintaining a uniform random subset of size k of the read
// packets.

package main

import (
	"fmt"
	"math/rand"
)

func RandomSubset(set []rune, size int) []rune {
	if len(set) == 0 || size == 0 {
		return nil
	}

	subset := make([]rune, 0, size)

	// Add first size elements to the subset.
	for i := 0; i < size && i < len(set); i++ {
		subset = append(subset, set[i])
	}

	// Update the subset based on random numbers.
	fmt.Print("Random indices = [")
	for i := size; i < len(set); i++ {
		selectIndex := rand.Intn(i+1)
		fmt.Print("{", selectIndex)

		// If selected the current index, then replace it with some
		// other element in the subset.
		if selectIndex == i {
			replaceIndex := rand.Intn(size)
			fmt.Print("->", replaceIndex)
			subset[replaceIndex] = set[i]
		}

		fmt.Print("}")
	}
	fmt.Println("]")

	return subset
}

func main() {
	for _, array := range [][]rune{{'p','q','r','t','u','v'}} {
		size := 2
		fmt.Printf("Random subset(%v) %c = %c.",
			size, array, RandomSubset(array, size))
	}
}
