// Design an algorithm that takes as input an array and a number, and determines
// if there are three entries in the array (not necessarily distinct) which add
// up to the specified number.

package main

import (
	"fmt"
	"sort"
)

func TwoSum(sortedArray []int, number int) (bool, int, int) {
	start, end := 0, len(sortedArray)-1

	// Note that an array element can be considered any number of times.
	for start <= end {
		// If extreme ends add up together.
		if sortedArray[start] + sortedArray[end] == number {
			return true, sortedArray[start], sortedArray[end]
		}

		// If the number is double of start element.
		if sortedArray[start] * 2 == number {
			return true, sortedArray[start], sortedArray[start]
		}

		// If the number is double of end element.
		if sortedArray[end] * 2 == number {
			return true, sortedArray[end], sortedArray[end]
		}

		// Move the start or end index as needed.
		if sortedArray[start] + sortedArray[end] < number {
			start++
		} else {
			end--
		}
	}

	return false, 0, 0
}

func ThreeSum(array []int, number int) (bool, int, int, int) {
	sort.Ints(array)

	for i := 0; i < len(array); i++ {
		found, num1, num2 := TwoSum(array, number-array[i])
		if found {
			return found, array[i], num1, num2
		}
	}

	return false, 0, 0, 0
}

func main() {
	array := []int{11, 2, 5, 7, 3}
	fmt.Println("Array", array)

	for _, sum := range []int {13, 21, 22, 33} {
		found, num1, num2, num3 := ThreeSum(array, sum)
		fmt.Println("Sum", sum, "= <", found, num1, num2, num3, ">")
	}
}
