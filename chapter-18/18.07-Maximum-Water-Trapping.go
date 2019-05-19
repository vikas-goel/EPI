// Write a program which takes as input an integer array and returns the pair
// of entries that trap the maximum amount of water.

package main

import "fmt"

func WaterTrapping(verticals []int) (leftIndex, rightIndex int) {
	if len(verticals) < 2 {
		return
	}

	// Maximum water between given two verticals is the shorter vertical
	// multiplied by distance between them.
	// Start with two extremes and keep moving inward in the hunt of a
	// better number. In order to move in, leave the shorter vertical
	// behind.
	maximumWater, start, end := 0, 0, len(verticals)-1
	for start < end {
		height := verticals[start]
		if verticals[end] < height {
			height = verticals[end]
		}

		width := end - start
		area := height*width
		if area > maximumWater {
			maximumWater, leftIndex, rightIndex = area, start, end
		}

		if verticals[start] < verticals[end] {
			start++
		} else if verticals[start] > verticals[end] {
			end--
		} else {
			start++
			end--
		}
	}

	return
}

func main() {
	verticals := []int{1, 2, 1, 3, 4, 4, 5, 6, 2, 1, 3, 1, 3, 2, 1, 2, 4, 1}
	left, right := WaterTrapping(verticals)

	fmt.Printf("Maximum water trapping for %v is between [%v, %v].\n",
		verticals, left, right)
}
