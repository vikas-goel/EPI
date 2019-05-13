// Write a program that takes an integer and determines if that integer's
// representation as a decimal string is a palindrome.

package main

import "fmt"

func CheckPalindrome(number int) bool {
	if number < 0 {
		return false
	} else if number < 10 {
		return true
	}

	msdMask := 1
	for temp := number; temp > 9; temp /= 10 {
		msdMask *= 10
	}

	for number > 9 {
		// Compare MSB and LSB.
		if number / msdMask != number % 10 {
			return false
		}

		// Remove MSB and LSB from the number.
		number = (number % msdMask) / 10
		msdMask /= 100
	}

	return true
}

func main() {
	for _, n := range []int{-1, 7, 12, 44, 2147447412, 27447, 21474547412} {
		fmt.Println("Palindrome(", n, ")?", CheckPalindrome(n))
	}
}
