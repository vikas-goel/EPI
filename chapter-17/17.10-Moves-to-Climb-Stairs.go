// You are climbing stairs. You can advance 1 to k steps at a time. Your
// destination is exactly n steps up. Write a program which takes as inputs n &
// k and returns the number of ways in which you can get to your destination.

package main

import "fmt"

func CountWays(n, k int) (numWays int) {
	ways := make([]int, n+1)

	for level := 1; level <= n; level++ {
		// The additional way is due to direct jump to the level
		// if it is less than or equal to k steps.
		if level <= k {
			ways[level] = 1
		}

		// For 1..k steps from the level.
		for step := 1; step < level && step <= k; step++ {
			ways[level] += ways[level-step]
		}
	}

	return ways[n]
}

func main() {
	n, k := 5, 2
	fmt.Printf("Number of ways to get destination for (%v, %v) is %v.\n",
		n, k, CountWays(n, k))
}
