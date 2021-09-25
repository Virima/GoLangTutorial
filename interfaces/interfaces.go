package main

import (
	"fmt"
	"math"
)

// interface point that anything that has
// func area that return float64 is a type shape
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 7}

	fmt.Println(c1.area())
	fmt.Println(r1.area())

	// shapes := []shape{c1, r1}
	shapes := []shape{&c1, &r1}
	fmt.Println(shapes)

	for _, shape := range shapes {
		// fmt.Println(shape.area())
		fmt.Println(getArea(shape))
	}

}
