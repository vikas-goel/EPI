// Write a program which takes an array of strings and a set of strings, and
// return the indices of the starting and ending index of a shortest subarray of
// the given array that "covers" the set, i.e., contains all strings in the set.

package main

import (
	"fmt"
	"math"
)

func NearestDistance(array []string) (distance int) {
	distance = math.MaxInt32
	lastIndex := make(map[string]int)

	for i := 0; i < len(array); i++ {
		if index, ok := lastIndex[array[i]]; ok {
			if i - index < distance {
				distance = i - index
			}
		}

		lastIndex[array[i]] = i
	}

	return
}

func main() {
	strs := []string{"All", "work", "and", "no", "play", "makes", "for",
		"no", "work", "no", "fun", "and", "no", "results"}
	fmt.Printf("%v = %v\n", strs, NearestDistance(strs))
}
