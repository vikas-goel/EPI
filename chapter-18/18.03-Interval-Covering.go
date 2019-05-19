// You are given a set of closed intervals. Design an efficient algorithm for
// finding a minimum sized set of numbers that covers all the intervals.

package main

import (
	"fmt"
	"sort"
)

type Intervals [][]int

func IntervalCovering(times Intervals) (visitTimes []int) {
	if len(times) < 1 {
		return
	}

	// Sort the takss in the order of their finish times.
	sort.Sort(times)

	// The optimized visiting time would be based on end time of the
	// intervals.
	visitTimes = append(visitTimes, times[0][1])

	for i := 1; i < len(times); i++ {
		// If the current start time occurs after the last visit time,
		// then add the end time of the interval to the list.
		if times[i][0] > visitTimes[len(visitTimes)-1] {
			visitTimes = append(visitTimes, times[i][1])
		}
	}

	return
}

func (this Intervals) Len() int {
	return len(this)
}

func (this Intervals) Less(i, j int) bool {
	if this[i][1] < this[j][1] {
		return true
	} else if this[i][1] > this[j][1] {
		return false
	} else if this[i][0] < this[j][0] {
		return true
	} else {
		return false
	}
}

func (this Intervals) Swap(i, j int) {
	this[i][0], this[j][0] = this[j][0], this[i][0]
	this[i][1], this[j][1] = this[j][1], this[i][1]
}

func main() {
	times1 := Intervals{{0, 3}, {2, 6}, {3, 4}, {6, 9}}
	times2 := Intervals{{1, 2}, {2, 3}, {3, 4}, {2, 3}, {3, 4}, {4, 5}}
	for _, t := range []Intervals{times1, times2} {
		fmt.Print("Task intervals ", t, " = Visit time ")
		fmt.Println(IntervalCovering(t))
	}
}
