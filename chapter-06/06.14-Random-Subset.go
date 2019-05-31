// Write a program that takes as input a positive integer n and a size k < n,
// and returns a size-k subset of {0,1, 2,...,n-1}. The subset should be
// represented as an array. All subsets should be equally likely and, in
// addition, all permutations of elements of the array should be equally likely.
// You may assume you have a function which takes as input a nonnegative
// integer t and returns an integer in the set {0, 1, . . . , t - 1) with
// uniform probability.

package main

import (
	"fmt"
	"math/rand"
)

func RandomSubset(n, k int) []int {
	if n < 2 || k >= n {
		return nil
	}

	cache := make(map[int]int)

	// Generate a random number and swap with the i-th index.
	for i := 0; i < k; i++ {
		random := rand.Intn(n-i) + i

		rValue, iValue := -1, -1
		if value, ok := cache[random]; ok {
			rValue = value
		}

		if value, ok := cache[i]; ok {
			iValue = value
		}

		if rValue == -1 && iValue == -1 {
			cache[random] = i
			cache[i] = random
		} else if rValue == -1 {
			cache[random] = iValue
			cache[i] = random
		} else if iValue == -1 {
			cache[random] = i
			cache[i] = rValue
		} else {
			cache[random] = iValue
			cache[i] = rValue
		}
	}

	set := make([]int, k)
	for i := 0; i < k; i++ {
		set[i] = cache[i]
	}

	return set
}

func main() {
	for _, num := range [][]int{{7,2},{10, 6},{100,10}} {
		fmt.Printf("Random permutation (%v) = %v.\n",
			num, RandomSubset(num[0], num[1]))
	}
}
