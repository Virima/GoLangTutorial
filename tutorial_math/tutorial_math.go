package main

import "fmt"

func main() {
	// math operation needs to be in same type
	var num1 int = 8
	var num2 int = 4

	answer := num1 + num2
	fmt.Printf("%d", answer)

	find_mod := num1 % num2
	fmt.Printf("%d", find_mod)

	var num3 float64 = 8
	var num4 int = 4

	// basic convertion example
	answer2 := num3 / float64(num4)
	fmt.Printf("%f", answer2)
}
