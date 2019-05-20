// Design an efficient algorithm for computing C(n,k) which has the property
// that it never overflows if the final result fits in the integer word size.

package main

import "fmt"

func Coefficient(n, k int) int {
	if k > m {
		return 0
	} else if k == 0 || n == k {
		return 1
	}

	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, k+1)
	}

	var compute func(int, int) int
	compute = func(thisN, thisK int) int {
		if thisK == 0 || thisK == thisN {
			C[thisN][thisK] = 1
		} else if C[thisN][thisK] == 0 {
			C[thisN][thisK] += compute(thisN-1, thisK)
			C[thisN][thisK] += compute(thisN-1, thisK-1)
		}
		return C[thisN][thisK]
	}

	return compute(n, k)
}

func main() {
	n, k := 5, 2
	fmt.Printf("Coefficient(%v, %v) = %v.\n", n, k, Coefficient(n, k))
}
