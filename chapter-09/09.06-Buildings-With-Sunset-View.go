// Design an algorithm that processes buildings in east-to-west order and
// returns the set of buildings which view the sunset. Each building is
// specified by its height.

package main

import (
	"fmt"
	"./stack"
)

type Stack = stack.Stack
var InitStack = stack.InitStack

type Building struct {
	Id, Height int
}

// The specified buildings are assumed to be in east-to-west order.
func SunsetView(buildings []Building) []Building {
	if len(buildings) == 0 {
		return nil
	}

	stack := InitStack()
	stack.Push(buildings[0])

	// Traverse from east to west. Remove buildings east to the current one
	// that are smaller in height. Push the current building in the stack
	// afterward.
	for i := 1; i < len(buildings); i++ {
		for !stack.Empty() &&
			stack.Peek().(Building).Height <= buildings[i].Height {
			stack.Pop()
		}

		stack.Push(buildings[i])
	}

	// The stack maintains the list of buildings that can view sunset.
	view := make([]Building, 0)
	for !stack.Empty() {
		view = append(view, stack.Pop().(Building))
	}

	// Return the buildings in west-to-east order. The height of the
	// buildings in west-to-east order must be increasing.
	return view
}

func main() {
	buildings := []Building{{1, 9}, {2, 5}, {3, 7}, {4, 2}, {5, 3}, {6, 1}}
	fmt.Printf("%+v -> %+v\n", buildings, SunsetView(buildings))
}
