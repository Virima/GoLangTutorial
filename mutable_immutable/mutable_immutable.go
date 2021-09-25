package main

import "fmt"

func main() {
	// var x []int = []int{3, 4, 5}
	// y := x
	// y[0] = 100

	// fmt.Println(x, y) // change both x and y because of slice behavior

	var x map[string]int = map[string]int{"hello": 3}
	y := x
	y["y"] = 100 // modify y and also x, also from map behavior
	fmt.Println(x, y)

	// For array it's a bit different //
	var a [2]int = [2]int{3, 4}
	b := a
	b[0] = 100

	fmt.Println(a, b)

	// Note: Map, Slice, and Array are mutable d.t.

	// the difference between mutable and immutable is
	// how the data stored in the RAM

	// to modify value of a slice, pass it through a function
	var c []int = []int{3, 4, 5}
	fmt.Println(c)
	modifyFirst(c, 100)
	fmt.Println(c)
}

// func to change a value of a slice
func modifyFirst(slice []int, new_val int) {
	slice[0] = new_val
}
