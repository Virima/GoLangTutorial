package main

import "fmt"

func calculate(x, y int) (r1, r2 int) {
	defer fmt.Println("Execute after the end of function")
	r1 = x + y
	r2 = x - y
	fmt.Println("Before defer")
	return
}

func main() {
	// ans1, ans2 := calculate(4,7)

	fmt.Println(calculate(5, 6))
}
