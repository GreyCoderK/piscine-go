package main

import "fmt"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points:= new(point)

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
