// Implement run-length encoding and decoding functions. Assume the string to
// be encoded consists of letters of the alphabet, with no digits, and the
// string to be decoded is a valid encoding.

package main

import "fmt"

func Encode(str string) string {
	length := len(str)
	strByte := make([]byte, 0)

	if length < 1 {
		return str
	}

	for i, freq := 1, 1; i < length; i++ {
		if str[i] != str[i-1] {
			strByte = append(strByte, byte('0'+freq), str[i-1])
			freq = 1
		} else {
			freq++
			if i != length-1 {
				continue
			}
			strByte = append(strByte, byte('0'+freq), str[i])
		}
	}

	return string(strByte[0:len(strByte)])
}

func Decode(str string) string {
	strByte := make([]byte, 0)

	if len(str) < 1 {
		return str
	}

	for i := 0; i < len(str); i += 2 {
		for count := str[i] - '0'; count > 0; count-- {
			strByte = append(strByte, str[i+1])
		}
	}

	return string(strByte[0:len(strByte)])
}

func main() {
	for _, str := range []string{"aaaabcccaa", "eeeffffee"} {
		encoded := Encode(str)
		decoded := Decode(encoded)
		fmt.Printf("{%v, %v, %v}\n", str, encoded, decoded)
	}
}
