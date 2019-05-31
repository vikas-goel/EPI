// Design an algorithm that creates uniformly random permutations of
// {0,1,...,n-1}. You are given a random number generator that returns
// integers in the set {0,1,...,n-1} with equal probability; use as few calls
// to it as possible.

package main

import (
	"fmt"
	"math/rand"
)

func RandomSubset(size int) []int {
	if size < 1 {
		return nil
	}

	set := make([]int, size)
	for i := 0; i < size; i++ {
		set[i] = i
	}

	// Generate a random number and swap with the i-th index.
	for i := 0; i < size; i++ {
		random := rand.Intn(size-i) + i
		if i != random {
			set[i], set[random] = set[random], set[i]
		}
	}

	return set
}

func main() {
	for _, num := range []int{6, 10} {
		fmt.Printf("Random permutation (%v) = %v.\n",
			num, RandomSubset(num))
	}
}
