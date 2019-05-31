// Write a program that takes an integer argument and returns all the primes
// between 1 and that integer. For example, if the input is 18, you should
// return (2,3,5,7,11,13,17).

package main

import (
	"fmt"
	"math"
)

func Primes(number int) []int {
	if number < 2 {
		return nil
	}

	// Map odd numbers 3,5,... to indices 0,1,...
	numberToIndex := func(number int) int {
		return number/2-1
	}

	// Map indices 0,1,... to odd numbers 3,5,...
	indexToNumber := func(index int) int {
		return (index+1)*2+1
	}

	primes := make([]int, 1, 10)
	primes[0] = 2

	// Apart from number 2, only odd numbers can be candidate for primes.
	nonPrimes := make([]bool, numberToIndex(number+1))
	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		// If not a prime, then its multiples are multiples of primes
		// too. So, skip it.
		if nonPrimes[numberToIndex(i)] {
			continue
		}

		// Multiples of prime number are not primes.
		for j := 3; j <= number/i; j += 2 {
			nonPrimes[numberToIndex(j*i)] = true

		}
	}

	// Collect all prime numbers.
	for i := 0; i < len(nonPrimes); i++ {
		if !nonPrimes[i] {
			primes = append(primes, indexToNumber(i))
		}
	}

	return primes
}

func main() {
	number := 1000
	fmt.Printf("Primes(%v) = %v\n.", number, Primes(number))
}
