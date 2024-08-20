package main

import "fmt"

type shape interface {
	getArea()
}
type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	square := square{
		sideLength: 2,
	}
	triangle := triangle{
		height: 6,
		base: 2,
	}
	printArea(square)
	printArea(triangle)
}

func printArea(s shape) {
	s.getArea()
}
func (s square) getArea() {
	fmt.Println("Area of square =", s.sideLength * s.sideLength)
}

func (t triangle) getArea() {
	fmt.Println("Area of triangle =", 0.5 * t.base * t.height)
}