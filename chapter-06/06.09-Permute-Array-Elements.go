// Given an array A of L elements and a permutation P, apply P to A.

package main

import "fmt"

func PermuteArray(P []int, A []rune) bool {
	if len(P) != len(A) {
		return false
	}

	fmt.Printf("%v -> %c = ", P, A)

	// Swap the array elements as per the permutation input.
	// Update the permutation indices to track the changes.
	for i, L := 0, len(P); i < L-1; i++ {
		next := i
		for P[next] >= 0 {
			A[i], A[P[next]] = A[P[next]], A[i]
			temp := P[next]
			P[next] -= L
			next = temp
		}
	}

	fmt.Printf("%c\n", A)
	return true
}

func main() {
	PermuteArray([]int{2, 0, 1, 3}, []rune{'a', 'b', 'c', 'd'})
	PermuteArray([]int{3, 2, 1, 0}, []rune{'a', 'b', 'c', 'd'})
	PermuteArray([]int{1, 0, 3, 4, 2}, []rune{'a', 'b', 'c', 'd', 'e'})
}
