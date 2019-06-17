// Write a program that tests if a string made up of the characters
// '(', ')', '[', and"}' is well-formed.

package main

import (
	"fmt"
	"./stack"
)

type Stack = stack.Stack
var InitStack = stack.InitStack

func ValidParentheses(str string) bool {
	stack := InitStack()
	for i := 0; i < len(str); i++ {
		if str[i] == '{' || str[i] == '(' || str[i] == '[' {
			stack.Push(str[i])
		} else if str[i] != '}' && str[i] != ')' && str[i] != ']' {
			continue
		} else {
			top := stack.Pop().(byte)

			// If closing one, then top of the stack must be
			// corresponding opening one to be valid string.
			if (str[i] == '}' && top != '{') ||
				(str[i] == ')' && top != '(') ||
				(str[i] == ']' && top != '[') {
					return false
			}
		}
	}

	// If there are elements left in the stack, then its invalid.
	return stack.Empty()
}

func main() {
	for _, expr := range []string{"{}()[]", "[()[]{()()}]", "{]",
		"[()[]{()()", "()(())" } {
		fmt.Println(expr, "=", ValidParentheses(expr))
	}
}
