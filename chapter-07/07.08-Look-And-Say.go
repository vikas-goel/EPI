// The look-and-say sequence starts with 1. Subsequent numbers are derived by
// describing the previous number in terms of consecutive digits.
// The first eight numbers in the look-and-say sequence are
// <1, 11, 21, 1211, 111221, 312211, 13112221, 1113213211>.
// Write a program that takes as input an integer n and returns the nth integer
// in the look-and-say sequence. Return the result as a string.

package main

import "fmt"

func LookSay(n int) (nth string) {
	if n < 1 {
		return
	}

	for nth = "1"; n > 1; n-- {
		count, next := 1, ""
		for i := 1; i <= len(nth); i++ {
			// If last number or different from previous one,
			// then append the count and number to the string.
			if i == len(nth) || nth[i-1] != nth[i] {
				next += string('0'+count) + string(nth[i-1])
				count = 1
			} else {
				// Increase the count of contiguous number.
				count++
			}
		}
		nth = next
	}

	return
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v ", LookSay(i))
	}
	fmt.Println()
}
