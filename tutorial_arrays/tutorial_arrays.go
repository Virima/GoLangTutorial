package main

import "fmt"

func main() {
	// Go array needs in the same type
	// Array is fixed length, init the max length
	var empty_arr [5]int
	var arr [5]int

	fmt.Println(empty_arr)

	arr[0] = 100
	arr[4] = 900
	fmt.Println(arr)
	fmt.Println(arr[4])

	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)
	fmt.Printf("Arr2 has %d elements \n", len(arr2))

	// Printing Array with Loop
	arr3 := [3]int{4, 5, 6}
	sum := 0

	for i := 0; i < len(arr3); i++ {
		sum += arr3[i]
	}

	fmt.Printf("The sum of arr3 = %d \n", sum)

	// 2 Dimensional array example
	arr2D := [3][3]int{{1, 2, 3 * 5}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[0][2])
}
