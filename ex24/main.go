package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// евклидово расстояние между точками
func (p1 *Point) Distance(p2 Point) float64 {

	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}

func main() {

	p1 := NewPoint(3.22, 1.45)
	p2 := NewPoint(-2.39, 5)

	d := p1.Distance(*p2)
	fmt.Printf("%.2f", d)
}
