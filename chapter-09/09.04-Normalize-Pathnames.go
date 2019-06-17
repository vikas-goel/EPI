// Write a program which takes a pathname, and returns the shortest equivalent
// pathname. Assume individual directories and files have names that use only
// alphanumeric characters. Subdirectory names may be combined using forward
// slashes (/), the current directory (.), and parent directory (..).

package main

import (
	"fmt"
	"strings"
	"./stack"
)

type Stack = stack.Stack
var InitStack = stack.InitStack

func Normalize(pathname string) string {
	stack := InitStack()

	if len(pathname) == 0 {
		return ""
	}

	startsWithRoot := false
	if pathname[0] == '/' {
		startsWithRoot = true
	}

	tokens := strings.Split(pathname, "/")
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == ".." {
			if !stack.Empty() {
				top := stack.Peek().(string)

				// If the top element is "..", then
				// add this again. Otherwise, pop.
				if top != ".." {
					stack.Pop()
				} else {
					stack.Push(tokens[i])
				}
			} else if !startsWithRoot {
				stack.Push(tokens[i])
			}
		} else if tokens[i] != "" && tokens[i] != "." {
			stack.Push(tokens[i])
		}
	}

	shortPath := stack.Pop().(string)
	for !stack.Empty() {
		shortPath = stack.Pop().(string) + "/" + shortPath
	}

	if startsWithRoot {
		shortPath = "/" + shortPath
	}

	return shortPath
}

func main() {
	for _, path := range []string{"/usr/bin/gcc", "scripts/awkscripts",
		"/usr/lib/../bin/gcc", "scripts//./../scripts/awkscripts/././",
		"sc//./../tc/awk/././", "../../../a/b/../c/./d/.."} {
		fmt.Println(path, "=", Normalize(path))
	}
}
