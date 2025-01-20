package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1. Get strings(input) from user.
// 2. Print the input on lowercase.

func main() {

	// Use bufio to input text with space.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Get string input
	input := scanner.Text()

	// print in lowercase.
	fmt.Printf("%s", strings.ToLower(input))

}
