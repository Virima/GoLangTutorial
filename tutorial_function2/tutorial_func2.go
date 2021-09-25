package main

import (
	"fmt"
)

func printHello(x int) {
	fmt.Println("Hello!", x)
}

// Funception
func test2(myFunc func(int) int) {
	fmt.Println(myFunc(8))
}

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {

	// storing function inside of a variable
	test :=
		func(x int) int {
			fmt.Println("Hellowww", x)
			return x * -1
		}

	test3 := func(x int) int {
		return x * 7
	}

	fmt.Println(test(8))

	// x := printHello // assign type function
	// x(5)            // call function separately

	// Call funception
	test2(test3)

	/// Function closure ///
	returnFunc("bruh")()
	returnFunc("goodbye")()
}
