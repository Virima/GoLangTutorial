package main

import "fmt"

// Make a struct with capital first letter

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	center *Point
}

func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	var p1 Point = Point{1, 2}
	var p2 Point = Point{-5, 7}

	fmt.Println(p1.x, p2.x)
	fmt.Println(p2)

	p1.x = 7

	fmt.Println(p1)
	//////////////////////////////////

	// assign Point without y value
	// y will be 0
	p3 := Point{x: 3}
	fmt.Println(p3)

	// Point with pointer
	p4 := &Point{y: 3}
	fmt.Println(p4)
	changeX(p4)
	fmt.Println(p4)

	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1.center.x)

}
