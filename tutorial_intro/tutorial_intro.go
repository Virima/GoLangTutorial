package main

import "fmt"

func main() {
	var name string
	name = "Tim"
	name = "Kelvin"

	// Explicit defining variable
	var number uint16 = 260
	number = number + 5

	// Implicit defining variable
	// GoLang guesses the data type
	var whatType = 555

	anotherWayToDefine := 6

	fmt.Println(name)
	fmt.Println("Hello Number", number)

	fmt.Printf("%T", whatType)
	fmt.Println("Implicit variable", whatType)

	fmt.Println(anotherWayToDefine)

	// Assigning with emptydefault value
	var num_default uint64
	var bool_default bool
	println(num_default)
	println(bool_default)

	bool_default = true
	println(bool_default)
}
