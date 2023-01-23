package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func main() {
	points := point{}
	points.x = 42
	points.y = 21
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
