package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	// to init a slice, empty the bracket
	// grab value index 1-3 but do not include 3
	var slice_example []int = arr[1:3]

	fmt.Println(arr)
	fmt.Println(slice_example)

	fmt.Println(len(slice_example))
	// capacity == length in arrays
	// capacity != length in slices
	fmt.Println(cap(slice_example))

	fmt.Println(slice_example[:cap(slice_example)])

	// Make a slice not from array
	var a []int = []int{6, 7, 8, 9, 10}

	// Add element to a slice
	a = append(a, 11, 12)  // same variable
	b := append(a, 11, 12) // create new slice
	fmt.Println(b)

	// Another way to make a slice with function 'make'
	c := make([]int, 5)
	fmt.Println(c) // empty slice
}
