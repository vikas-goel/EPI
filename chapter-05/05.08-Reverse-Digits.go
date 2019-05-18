// Write a program which takes an integer and returns the integer corresponding
// to the digits of the input written in reverse order. For example, the
// reverse of 42 is 24, and the reverse of -314 is -413.

package main

import "fmt"

func Reverse(number int64) (reverse int64) {
	for number != 0 {
		reverse = reverse * 10 + (number % 10)
		number /= 10
	}
	return
}

func main() {
	for _, n := range []int64{1132, -314, 42} {
		fmt.Println("Reverse(", n, ") =", Reverse(n))
	}
}
