package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Get input from user.
	// 2. Split string input with space
	// 3. Join them with dots.

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	splitInput := strings.Split(input, " ")

	fmt.Print(strings.Join(splitInput, "..."))
}
