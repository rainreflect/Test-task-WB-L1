package tasks

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func Task24() {
	p1 := newPoint(1, 2)
	p2 := newPoint(5, 4)
	fmt.Println(math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2)))
}
