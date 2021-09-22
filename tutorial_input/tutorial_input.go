package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type something: ")
	// scanner.Scan()
	// input := scanner.Text()
	// fmt.Printf("You typed: %q", input)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()

	// The second parameter is for error handling
	// but for now put _ to ignore/not use it
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years gold at the end of 2021", 2021-input)
}
