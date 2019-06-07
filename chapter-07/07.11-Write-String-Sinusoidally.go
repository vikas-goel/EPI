// Write a program which takes as input a string s and returns the snakestring
// of s.

package main

import "fmt"

func PrintSinusoidally(str string) string {
	sinuosdial, pos := make([]byte, len(str)), 0

	fmt.Println()

	for i := 1; i < len(str); i += 4 {
		fmt.Printf("  %c     ", str[i])
		sinuosdial[pos] = str[i]
		pos++
	}

	fmt.Println()

	for i := 0; i < len(str); i += 2 {
		fmt.Printf("%c   ", str[i])
		sinuosdial[pos] = str[i]
		pos++
	}

	fmt.Println()

	for i := 3; i < len(str); i += 4 {
		fmt.Printf("      %c ", str[i])
		sinuosdial[pos] = str[i]
		pos++
	}

	fmt.Println()
	fmt.Println()

	return string(sinuosdial)
}

func main() {
	for _, str := range []string{"Hello_World!", "Welcome"} {
		fmt.Println(str, "->", PrintSinusoidally(str))
	}
}
