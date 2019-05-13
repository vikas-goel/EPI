// Implement a random number generator that generates a random integer i
// between a and b, inclusive, given a random number generator that produces
// 0 or 1 with equal probability. All values in [a,b] should be equally likely.

package main

import "fmt"
import "math/rand"

type Nums struct {
	a, b int
}

func RandomNumber(a, b int) int {
	randomRange := b - a
	randomNumber := randomRange + 1

	for randomNumber > randomRange {
		randomNumber = 0
		for tempRange := randomRange; tempRange != 0; tempRange >>= 1 {
			randomNumber = (randomNumber << 1) | rand.Intn(2)
		}
	}

	return randomNumber + a
}

func main() {
	for _, n := range []Nums{{3, 9}, {2, 20}, {5, 500}} {
		fmt.Println("Random", n, "=", RandomNumber(n.a, n.b))
	}
}
