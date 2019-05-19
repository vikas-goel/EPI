// In the gasup problem, a number of cities are arranged on a circular road.
// You need to visit all the cities and come back to the starting city. A
// certain amount of gas is available at each city. The amount of gas summed
// up over all cities is equal to the amount of gas required to go around the
// road once. Your gas tank has unlimited capacity. Call a city ample if you
// can begin at that city with an empty tank, refill at it, then travel through
// all the remaining cities, refilling at each, and return to the ample city,
// without running out of gas at any point.

package main

import "fmt"

type City struct {
	name string
	distance, gas int
}

func AmpleCity(cities []City, mileage int) (cityIndex int) {
	if len(cities) == 0 {
		return -1
	}

	// Keep traversing and tracking the gas usage. The city where the
	// tank goes minimum is the ample city. Starting from such ample city
	// will never see the gas tank getting empty on the way.
	availableTank, minimumTank := cities[0].gas, cities[0].gas
	for i := 1; i < len(cities); i++ {
		availableTank -= cities[i-1].distance/mileage
		if availableTank < minimumTank {
			minimumTank = availableTank
			cityIndex = i
		}

		availableTank += cities[i].gas
	}

	return
}

func main() {
	cities := []City{{"A",900,50}, {"B",600,20}, {"C",200,5}, {"D",400,30}, {"E",600,25}, {"F",200,10}, {"G",100,10}}

	fmt.Printf("Ample city for %v is %v.\n",
		cities, cities[AmpleCity(cities, 20)].name)
}
