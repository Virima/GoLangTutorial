package main

import "fmt"

func main() {
	// a map named mp with string key
	// pointing on  an integer value
	// map doesn't keep the value order
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	// make an empty map
	// mp := make(map[string]int)

	fmt.Println(mp)

	fmt.Printf("The ammount of apple: %d\n", mp["apple"])

	mp["apple"] = 900
	fmt.Printf("The ammount of apple new: %d\n", mp["apple"])

	mp["grape"] = 400
	fmt.Printf("New map key added:")
	fmt.Println(mp)

	// Delete from map
	delete(mp, "pear")
	fmt.Println("After pear deleted:")
	fmt.Println(mp)

	// Check if key exists or not in a map
	// val = value of the key, if available
	// ok = boolean represents if the key available or not
	val, ok := mp["apple"]
	val2, ok2 := mp["bandeng"]
	fmt.Println(val, ok)
	fmt.Println(val2, ok2)

	//Check length of the map
	fmt.Printf("length of the map: %d \n", len(mp))

	// Note:
	// Use array if you care about the data index
	// Use map if you don't care about index (faster)
}
