// Write a program which tests if two rectangles have a nonempty intersection.
// If the intersection is nonempty, return the rectangle formed by their
// intersection.

package main

import "fmt"

type Point struct {
	x, y int
}

type Rectangle struct {
	LB, RT Point
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func Intersection(r1, r2 Rectangle) (Rectangle, bool) {
	xlb := max(r1.LB.x, r2.LB.x)
	xrt := min(r1.RT.x, r2.RT.x)
	ylb := max(r1.LB.y, r2.LB.y)
	yrt := min(r1.RT.y, r2.RT.y)

	if xrt - xlb < 0 || yrt - ylb < 0 {
		return Rectangle{Point{0, 0}, Point{0, 0}}, false
	}

	return Rectangle{Point{xlb, ylb}, Point{xrt, yrt}}, true
}

func main() {
	r1 := Rectangle{Point{6, 4}, Point{15, 8}}
	r2 := Rectangle{Point{3, 7}, Point{11, 12}}

	r3, ok := Intersection(r1, r2)
	fmt.Printf("Intersection(%v, %v) = %v, %v\n", r1, r2, r3, ok)

	r4 := Rectangle{Point{6, 4}, Point{15, 8}}
	r5 := Rectangle{Point{3, 1}, Point{5, 3}}

	r6, ok := Intersection(r4, r5)
	fmt.Printf("Intersection(%v, %v) = %v, %v\n", r4, r5, r6, ok)
}
