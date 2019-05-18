// Write a program which computes all size k subsets of {1, 2,..., n}, where k
// and n are program inputs. For example, if k = 2 and n = 5, then the result
// is the following:
// ((1, 2), (1,3), (1, 4), (1,5), (2,3), (2, 4), (2,5), (3, 4), (3,5), (4, 5))

package main

import "fmt"

type Subset []int

func ComputeSubsets(n, k int) (subsets []Subset) {
	if n < k || n <= 0 {
		return
	}

	subsetK := make(Subset, k)

	var compute func(start, size int)
	compute = func(start, size int) {
		// One subset of desired size is created. Add this to the list.
		if k == size {
			newSubset := make(Subset, k)
			copy(newSubset, subsetK)
			subsets = append(subsets, newSubset)
			return
		}

		// Check whether subset possible under given conditions.
		if n - start + 1 < k - size {
			return
		}

		// Fix one index at a time and then fill in remaining.
		for i := start; i <= n; i++ {
			subsetK[size] = i
			compute(i+1, size+1)
		}
	}

	compute(1, 0)
	return
}

func main() {
	n, k := 5, 3
	fmt.Printf("Subsets(%v, %v) = %v\n", n, k, ComputeSubsets(n, k))
}
