// Write a program that performs base conversion. The input is a string, an
// integer b2, and another integer b2. The string represents be an integer in
// base b2. The output should be the string representing the integer in base b2.
// Assume 2 <= b1, b2 <= 16. Use "A" to represent 10, "B" for 11, ..., and "F"
// for 15. (For example, if the string is "615", b1 is 7 and b2 is 13, then the
// result should be"1A7", since 6x72+1x7+5 = 1X132+10X13+7.)

package main

import "fmt"

func ConvertBase(integerString string, fromBase, toBase int) (str string) {
	if len(integerString) == 0 {
		return
	}

	startIndex := 0
	if integerString[0] == '-' {
		startIndex = 1
	}

	number := 0
	for i := startIndex; i < len(integerString); i++ {
		number *= fromBase
		if integerString[i] >= '0' && integerString[i] <= '9' {
			number += int(integerString[i] - '0')
		} else {
			number += int(integerString[i] - 'A' + 10)
		}
	}

	str = convertTo(number, toBase)
	if startIndex == 1 {
		str = "-" + str
	}

	return
}

func convertTo(decimalWholeNumber, toBase int) string {
	if decimalWholeNumber == 0 {
		return "0"
	}

	var baseConverted string = ""
	for ; decimalWholeNumber > 0; decimalWholeNumber /= toBase {
		digit := decimalWholeNumber % toBase
		if digit >= 10 {
			digit += 'A' - 10
		} else {
			digit += '0'
		}

		baseConverted = string(digit) + baseConverted
	}

	return baseConverted
}

func main() {
	number, from, to := "615", 7, 13
	fmt.Printf("Convert(\"%v\", from=%v, to=%v) = %v\n",
		number, from, to, ConvertBase(number, from, to))
}
