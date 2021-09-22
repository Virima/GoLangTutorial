package main

import "fmt"

func main() {

	// First type, this is go's while loop with For syntax
	a := 3
	for a <= 5 {
		fmt.Printf("Type 1 Loop, value: %d \n", a)
		a++
	}

	// Second type, standard For loop
	for b := 0; b <= 4; b += 2 {
		fmt.Printf("Type 2 Loop, value: %d \n", b)

	}

	// The use of break and continue
	for c := 0; c < 1000; c++ {
		if c != 0 && c%3 == 0 && c%7 == 0 && c%9 == 0 {
			fmt.Printf("%d \n", c)
			// continue // Skip anything below it and back to for
			break // exit the loop immediately
		}
	}

	//Switch case
	ans := 3

	switch ans {
	case 1, -1:
		fmt.Println("1. one")
	case 2:
		fmt.Println("2. two")
	default:
		fmt.Println("default, not a case")
	}

}
