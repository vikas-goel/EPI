// Write a program that takes as input a number and returns all the strings
// with that number of matched pairs of parens.

package main

import "fmt"

func GenerateParens(count int) []string {
	pstr := make([]byte, count*2)
	list := make([]string, 0, 10)

	addParens(&list, count, count, pstr, 0)
	return list
}

func addParens(list *[]string, leftRem, rightRem int, str []byte, index int) {
	// Make sure left parentheses is not exhausted and the syntax is valid.
	if leftRem < 0 || rightRem < leftRem {
		return
	}

	if leftRem == 0 && rightRem == 0 {
		*list = append(*list, string(str))
	} else {
		str[index] = '('
		addParens(list, leftRem-1, rightRem, str, index+1)

		str[index] = ')'
		addParens(list, leftRem, rightRem-1, str, index+1)
	}
}

func main() {
	fmt.Println(GenerateParens(6))
}
