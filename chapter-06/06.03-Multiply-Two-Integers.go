// Write a program that takes two arrays representing integers, and returns an
// integer representing their product. For example, since
// 193707721 X -761838257287 = -147573952589676412927, if the inputs are
// (1,9,3,7,0,7,7,2,1} and (-7,6,1,8,3,8,2,5,7,2,8,7), your function should
// return (-1,4,7,5,7,3,9,5,2,5,8,9,6,7,6,4,1,2,9,2,7).

package main

import "fmt"

func Multiply(num1, num2 []int) (result []int) {
	if len(num1) == 0 || len(num2) == 0 {
		return
	}

	// Find the sign of the result.
	sign := 1
	if (num1[0] < 0 && num2[0] >= 0) || (num2[0] < 0 && num1[0] >= 0) {
		sign = -1
	}

	// The sign is stored. Convert the numbers to absolute values.
	if num1[0] < 0 {
		num1[0] *= -1
	}

	if num2[0] < 0 {
		num2[0] *= -1
	}

	// Multiply the two positive numbers.
	result = make([]int, len(num1)+len(num2))
	for i := len(num1)-1; i >=0; i-- {
		for j := len(num2)-1; j >= 0; j-- {
			result[i+j+1] += num1[i] * num2[j]
			result[i+j] += result[i+j+1] / 10
			result[i+j+1] %= 10
		}
	}

	// Remove any leading zeroes from the array.
	leadingZeroes := 0
	for i := 0; i < len(result) && result[leadingZeroes] == 0; i++ {
		leadingZeroes++
	}

	if leadingZeroes > 0 {
		result = result[leadingZeroes:]
	}

	// Multiply the result with the sign to get the final result.
	result[0] *= sign

	return
}

func main() {
	for _, n := range [][][]int{{
		{1,9,3,7,0,7,7,2,1}, {-7,6,1,8,3,8,2,5,7,2,8,7}},
		{{1,2,3}, {4,5,6,7}}} {
		fmt.Print(n[0], "x", n[1], "=")
		fmt.Println(Multiply(n[0], n[1]))
	}
}
