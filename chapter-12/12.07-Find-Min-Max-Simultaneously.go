// Design an algorithm to find the min and max elements in an array. For
// example, if A = (3,2,5,1,2,4), you should return 1 for the min and 5 for the
// max in less than 2*(n-1) comparisons.

package main

import "fmt"

func MinMax(array []int) (min, max int) {
	if len(array) == 0 {
		return
	}

	min, max = array[0], array[0]
	for _, elem := range array[1:] {
		if elem < min {
			min = elem
		} else if elem > max {
			max = elem
		}
	}

	return
}

func main() {
	array := []int{3, 2, 5, 1, 2, 4}
	min, max := MinMax(array)
	fmt.Println(array, ":", min,  max)
}
