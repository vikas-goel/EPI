package main

import "fmt"

var RomanLetter map[byte]int
var RomanString map[string]int
var letterMapInitialized, stringMapInitialized bool

// Write a program which takes as input a valid Roman number string s and
// returns the integer it corresponds to.
func RomanToDecimal(roman string) (decimal int, valid bool) {
	for i, exceptionCount := 0, 0; i < len(roman)-1; i++ {
		if ok, exception := isOkay(roman[i], roman[i+1]); !ok {
			return
		} else if exception {
			if exceptionCount == 1 {
				// Second continuous exception makes the string
				// invalid roman number.
				return
			} else {
				// Exception. Subtract the current value.
				decimal -= romanLetterToDecimal(roman[i])
				exceptionCount = 1
			}
		} else {
			// No exception. Add the number.
			decimal += romanLetterToDecimal(roman[i])
			exceptionCount = 0
		}
	}

	decimal += romanLetterToDecimal(roman[len(roman)-1])
	valid = true
	return
}

// Write a program that takes as input a positive integer n and returns a
// shortest valid simple Roman number string representing n.
func DecimalToRoman(decimal int) (roman string) {
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL",
				"X", "IX", "V", "IV", "I"}
	for decimal > 0 {
		for _, r := range romans {
			if decimal >= romanStringToDecimal(r) {
				roman += string(r)
				decimal -= romanStringToDecimal(r)
				break
			}
		}
	}

	return
}

func romanLetterToDecimal(letter byte) int {
	if !letterMapInitialized {
		RomanLetter = make(map[byte]int)
		RomanLetter['I'] = 1
		RomanLetter['V'] = 5
		RomanLetter['X'] = 10
		RomanLetter['L'] = 50
		RomanLetter['C'] = 100
		RomanLetter['D'] = 500
		RomanLetter['M'] = 1000
	}

	return RomanLetter[letter]
}

func romanStringToDecimal(roman string) int {
	if !stringMapInitialized {
		RomanString = make(map[string]int)
		RomanString["I"] = 1
		RomanString["IV"] = 4
		RomanString["V"] = 5
		RomanString["IX"] = 9
		RomanString["X"] = 10
		RomanString["XL"] = 40
		RomanString["L"] = 50
		RomanString["XC"] = 90
		RomanString["C"] = 100
		RomanString["CD"] = 400
		RomanString["D"] = 500
		RomanString["CM"] = 900
		RomanString["M"] = 1000
	}

	return RomanString[roman]
}

func isOkay(left, right byte) (ok bool, exception bool) {
	if romanLetterToDecimal(left) >= romanLetterToDecimal(right) {
		ok = true
	} else if left == 'I' && (right == 'V' || right == 'X') {
		ok, exception = true, true
	} else if left == 'X' && (right == 'L' || right == 'C') {
		ok, exception = true, true
	} else if left == 'C' && (right == 'D' || right == 'M') {
		ok, exception = true, true
	}

	return
}

func main() {
	fmt.Print("[ ")
	for _, roman := range []string{"IXC",  "CDM", "LVIIII", "LIX", "IC"} {
		decimal, ok := RomanToDecimal(roman)
		if ok {
			back := DecimalToRoman(decimal)
			fmt.Printf("{%v, %v, %v} ", roman, decimal, back)
		} else {
			fmt.Printf("{%v, %v} ", roman, "Invalid")
		}
	}
	fmt.Println("]")
}
