package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 5, 4, 9, 6, 4, 9, 3}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// As same as For loop but with Range func
	// Change to _ if not using i or element variable
	for index, element := range a {
		fmt.Printf("%d: %d \n", index, element)
	}

	// Print duplicated value in range a
	for i, element := range a {
		for j, element2 := range a {
			if element == element2 && j > i {
				fmt.Println(element)
			}
		}
	}

	// Print duplicated value in range a (effective)
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			temp := a[j]
			if temp == element {
				fmt.Println(element)
			}
		}
	}
}
