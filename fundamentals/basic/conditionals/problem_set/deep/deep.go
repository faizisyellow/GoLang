package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Input user. It should (42, forty-two, forty two).
	// 2. Outputting "YES" if the input is correct.
	// 3. Outputting "NO" if incorrect.

	fmt.Print("What is the Answer to Great Question of Life, the Universe, and Everything? ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputLower := strings.ToLower(input)

	// remove whitespace
	inputWithoutSpace := strings.Replace(inputLower, " ", "", -1)

	if inputWithoutSpace == "42" || inputWithoutSpace == "forty-two" || inputWithoutSpace == "fortytwo" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
