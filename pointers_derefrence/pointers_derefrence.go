package main

import "fmt"

// * pass the pointer, not value
func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {

	// Operator & means get the pointer
	// Operator * means dereference

	// Look out the reference of a variable
	x := 7
	fmt.Println(&x)
	fmt.Println(x)

	y := &x
	fmt.Println(&y)
	// changing the value of x through the pointer y
	*y = 8
	fmt.Println(x, y)

	////////
	// to change a value inside of function, we'll be using
	// pointer and reference.
	fmt.Println()
	toChange := "Hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	// not changed because didn't pass the pointer
	fmt.Println()
	toChange2 := "Hello"
	fmt.Println(toChange2)
	changeValue2(toChange2)
	fmt.Println(toChange2)
}
