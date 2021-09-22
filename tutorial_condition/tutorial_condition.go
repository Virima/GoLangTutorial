package main

import "fmt"

func main() {
	// to do comparison, need to be in the same type
	x := 5
	y := 6
	val := x < 5
	val2 := y > 3

	fmt.Printf("%t", val)
	fmt.Printf("%t", val2)

	s := "kelvin"
	t := "Kelvin"

	fmt.Printf("%t", s == t)

	// - Tutorial chained conditional - //
	chained := (true || false) && !false
	fmt.Printf("%t", chained)
}
