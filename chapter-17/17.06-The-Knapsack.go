// Write a program for the knapsack problem that selects a subset of items
// that has maximum value and satisfies the weight constraint. All items have
// integer weights and values. Return the value of the subset.

package main

import "fmt"

type Object struct {
	name string
	weight, value int
}

func Knapsack(items []Object, weightTolerance int) int {
	if weightTolerance <= 0 || len(items) == 0 {
		return 0
	}

	memo := make(map[int]map[int]int)

	var compute func(int, int) int
	compute = func(index, weight int) (value int) {
		if weight <= 0 || index >= len(items) {
			return 0
		}

		// If already evaluated for this index and weight pait, return
		// the result from the cache.
		if thisValue, ok := memo[index][weight]; ok {
			return thisValue
		} else {
			memo[index] = make(map[int]int)
		}

		withIndex, withoutIndex := 0, compute(index+1, weight)

		// Include this item.
		remainingWeight := weight - items[index].weight
		if remainingWeight > 0 {
			withIndex = compute(index+1, remainingWeight)
		}

		if remainingWeight >= 0 {
			withIndex += items[index].value
		}

		value = withIndex
		if withoutIndex > value {
			value = withoutIndex
		}

		memo[index][weight] = value

		return
	}

	return compute(0, weightTolerance)
}

func main() {
	items := []Object{{"A",5,60}, {"B",3,50}, {"C",4,70}, {"D",2,30}}

	fmt.Println("Items = ", items)
	for w := 0; w < 6; w++ {
		fmt.Printf("Weight = %v, Value = %v.\n", w, Knapsack(items, w))
	}
}
