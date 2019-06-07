// Implement methods that take a string representing an integer and return the
// corresponding integer, and vice versa. Your code should handle negative
// integers. You cannot use library functions.

package main

import "fmt"

func StringToInteger(str string) (num int) {
	sign, start := 1, 0
	if str[0] == '-' {
		sign, start = -1, 1
	}

	for i := start; i < len(str); i++ {
		num = num * 10 + int(str[i] - '0')
	}

	return num*sign
}

func IntegerToString(num int) (str string) {
	sign := 1
	if num < 0 {
		sign = -1
		num *= sign
	}

	for num > 0 {
		digit := num % 10
		num /= 10
		str = string(digit + '0') + str
	}

	if sign == -1 {
		str = "-" + str
	}

	return
}

func main() {
	for _, number := range []int{123456, -123456} {
		str := IntegerToString(number)
		num := StringToInteger(str)
		fmt.Println("int(", number, ")->str(", str, ")->int(", num, ")")
	}
}
