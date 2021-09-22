package main

import "fmt"

func main() {
	// fmt.Printf("Hello %T %v", 10, 10)

	// fmt.Printf("The value of boolean is %t", true)

	// fmt.Printf("Number 1025 in binary: %b", 1025)
	// fmt.Printf("Number 3435 in hex: %X", 3435)

	// fmt.Printf("Scientific notation %e", 2.8490238492342)
	// fmt.Printf("Decimal %f", 2.8490238492342)
	// fmt.Printf("Large exponent decimal %g", 2.8490238492342)

	// fmt.Printf("Number: %s", "kelvin")
	// fmt.Printf("Number: %q", "kelvin")

	// String limitation and formatting //
	fmt.Printf("Hello %15q is cool \n", "kelvin")
	fmt.Printf("Hello %.2f is cool \n", 3.1422938)
	fmt.Printf("Hello %.f is cool \n", 3.1422938)
	fmt.Printf("Hello %.07d is cool \n", 35)

	// Storing fmt function in a variable
	var out string = fmt.Sprintf("Hello %.07d is cool \t hehe \n", 35)
	fmt.Println(out)
}
