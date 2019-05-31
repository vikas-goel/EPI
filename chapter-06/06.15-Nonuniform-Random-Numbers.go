// You are given n numbers as well as probabilities p(0), p(1),...,p(n-1)
// which sum up to 1. Given a random number generator that produces values in
// [0,1] uniformly, how would you generate one of the n numbers according to
// the specified probabilities? For example, if the numbers are 3,,5,7,11, and
// the probabilities are 9/18,6/18,2/18,1/18, then in 1000000 calls to your
// program, 2 should appear roughly 500000 times, 5 should appear roughly 333333
// times, 7 should appear roughly 111111 times, and 11 should appear roughly
// 55555 times.

package main

import (
	"fmt"
	"math/rand"
)

type Number struct {
	id, frequency int
	expected, actual int
}

func RandomNumber(numbers []Number, calls int) {
	if calls < 1 {
		return
	}

	// Find total of all frequencies to get probability.
	sumFrequencies := 0
	for i := 0; i < len(numbers); i++ {
		sumFrequencies += numbers[i].frequency
	}

	if sumFrequencies == 0 {
		return
	}

	// Initialize what expected count for individual ids is for given calls.
	for i := 0; i < len(numbers); i++ {
		// Be wary of integer overflow possibility due to the
		// multiplication.
		numbers[i].expected = numbers[i].frequency*calls/sumFrequencies
	}

	// Generate random numbers between [1,sumFrequencies] and check which
	// id it belongs to.
	for i := 0; i < calls; i++ {
		rand := rand.Intn(sumFrequencies) + 1
		index := 0

		// Keep iterating till the random number meets the probability.
		for ; rand > numbers[index].frequency; index++ {
			rand -= numbers[index].frequency
		}

		// Update the actual count of the id.
		numbers[index].actual++
	}

	return
}

func main() {
	calls := 1000000
	numbers := []Number{{3,9,0,0},{5,6,0,0},{7,2,0,0},{9,1,0,0}}
	RandomNumber(numbers, calls)
	fmt.Printf("%+v for %v calls.\n", numbers, calls)
}
