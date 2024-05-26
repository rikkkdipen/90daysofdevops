package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Type casting
	// Input and output
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year when you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2024", 2024-input)

}
