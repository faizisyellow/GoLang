package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Print user.
	// 2. If user print hello return $0.
	// 3. If user print start with "h" but not hello return $20.
	// 4. Otherwise return $100.

	fmt.Print("Greeting: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputSanitize := strings.ToLower(input)

	if inputSanitize == "hello" {
		fmt.Println("$0")
	} else if strings.Index(inputSanitize, "h") == 0 && !strings.Contains(inputSanitize, "hello") {
		fmt.Println("$20")
	} else {
		fmt.Println("$100")
	}

}
