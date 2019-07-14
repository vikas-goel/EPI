// Write a program which takes as input an array and finds the distance between
// a closest pair of equal entries. For example, if s = ("All", "work", "and",
// "no", "play", "makes", "for", "no", "work", "no", "fun", "and", "no",
// "results"), then the second and third occurrences of "no" is the closest
// pair.

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
