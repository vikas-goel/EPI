// Given service times for a set of queries, compute a schedule for processing
// the queries that minimizes the total waiting time. Return the minimum
// waiting time.

package main

import (
	"fmt"
	"sort"
)

func WaitingTime (serviceTimes []int) (waitingTime int) {
	// Sort the servicing times to minimize the waiting time.
	// The higher servicing times can go later for optimization.
	sort.Ints(serviceTimes)

	// A service time adds up as waiting time for each following service.
	for i, numServices := 1, len(serviceTimes); i < numServices; i++ {
		waitingTime += (serviceTimes[i-1]*(numServices-i))
	}

	return
}

func main() {
	servicing := []int{2, 5, 1, 3}
	fmt.Print("Service Times ", servicing, " = Waiting time ")
	fmt.Println(WaitingTime(servicing))
}
