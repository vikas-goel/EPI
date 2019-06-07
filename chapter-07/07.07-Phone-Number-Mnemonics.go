// Write a program which takes as input a phone number, specified as a string
// of digits, and returns all possible character sequences that correspond to
// the phone number. The cell phone keypad is specified by a mapping that takes
// a digit and returns the corresponding set of characters. The character
// sequences do not have to be legal words or phrases.

package main

import "fmt"

var KeyChars map[byte][]byte
var KeyPadInit bool

func Mnemonics(phoneNumber string) []string {
	if len(phoneNumber) == 0 {
		return nil
	}

	allMnemonics := make([]string, 0, 10)
	mnemonics := make([]byte, len(phoneNumber))

	var compute func(int)
	compute = func(index int) {
		if index == len(mnemonics) {
			allMnemonics = append(allMnemonics, string(mnemonics))
			return
		}

		for _, char := range keypad(phoneNumber[index]) {
			mnemonics[index] = char
			compute(index+1)
		}
	}

	compute(0)

	return allMnemonics
}

func keypad(digit byte) []byte {
	if !KeyPadInit {
		KeyChars = make(map[byte][]byte)
		KeyChars['0'] = []byte{'0'}
		KeyChars['1'] = []byte{'1'}
		KeyChars['2'] = []byte{'A', 'B', 'C'}
		KeyChars['3'] = []byte{'D', 'E', 'F'}
		KeyChars['4'] = []byte{'G', 'H', 'I'}
		KeyChars['5'] = []byte{'J', 'K', 'L'}
		KeyChars['6'] = []byte{'M', 'N', 'O'}
		KeyChars['7'] = []byte{'P', 'Q', 'R', 'S'}
		KeyChars['8'] = []byte{'T', 'U', 'V'}
		KeyChars['9'] = []byte{'W', 'X', 'Y', 'Z'}
		KeyPadInit = true
	}

	return KeyChars[digit]
}

func main() {
	for _, str := range []string{"2276696", "21203"} {
		fmt.Printf("%v -> %v\n", str, Mnemonics(str))
	}
}
