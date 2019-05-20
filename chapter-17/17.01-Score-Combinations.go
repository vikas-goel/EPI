// In an American football game, a play can lead to 2 points (safety),3
// points (field goal), or 7 points (touchdown, assuming the extra point).
// Many different combinations of 2, 3, and 7 point plays can make up a final
// score.
// Write a program that takes a final score and scores for individual plays,
// and returns the number of combinations of plays that result in the final
// score.

package main

import "fmt"

func CountWays(finalScore int) (numWays int) {
	points := []int{2, 3, 7}
	ways := make([]int, finalScore+1)
	ways[0] = 1

	// Permutation does not matter in this case. Only combinations.
	// Add contribution of each point towards the possible scores.
	for i := 0; i < len(points); i++ {
		ways[points[i]] += 1
		for j := points[i]+1; j <= finalScore; j++ {
			ways[j] += ways[j-points[i]]
		}
	}

	fmt.Println(ways)
	return ways[finalScore]
}

func main() {
	score := 12
	fmt.Printf("Number of ways for final score %v is %v.\n ",
		score, CountWays(score))
}
