// Write a program that determines where to add periods to a decimal string so
// that the resulting string is a valid IP address. There may be more than one
// valid IP address corresponding to a string, in which case you should print
// all possibilities.

package main

import (
	"fmt"
	"strconv"
)

func ValidIPAddresses(decimal string) (addresses []string) {
	length := len(decimal)

	if length < 4 {
		return
	}

	isValidOctet := func(str string) (octet int, valid bool) {
		// An octet starting with 0 is invalid.
		if str[0] == '0' && len(str) > 1 {
			return
		}

		number, err := strconv.Atoi(str)
		if err != nil || number < 0 || number > 255 {
			return
		}

		return number, true
	}

	getIPv4 := func(o1, o2, o3, o4 int) string {
		return strconv.Itoa(o1) + "." + strconv.Itoa(o2) + "." +
			strconv.Itoa(o3) + "." + strconv.Itoa(o4)
	}

	addresses = make([]string, 0)

	for i1 := 0; i1 < 3 && i1 < length-3; i1++ {
		octet1, ok1 := isValidOctet(decimal[0:i1+1])
		if !ok1 {
			// The first octet cannot be 0.
			continue
		}

		for i2 := i1+1; i2 < i1+4 && i2 < length-2; i2++ {
			octet2, ok2 := isValidOctet(decimal[i1+1:i2+1])
			if !ok2 {
				continue
			}

			for i3 := i2+1; i3 < i2+4 && i3 < length-1; i3++ {
				octet3, ok3 := isValidOctet(decimal[i2+1:i3+1])
				if !ok3 {
					continue
				}

				// If length of the last octet turns out to be
				// more than 3 digits, then it's invalid one.
				if length-i3-1 > 3 {
					continue
				}

				octet4, ok4 := isValidOctet(decimal[i3+1:])
				if !ok4 {
					continue
				}

				ipv4 := getIPv4(octet1, octet2, octet3, octet4)
				addresses = append(addresses, ipv4)
			}
		}
	}

	return
}

func main() {
	for _, str := range []string{"19216811"} {
		fmt.Println(str, "->", ValidIPAddresses(str))
	}
}
