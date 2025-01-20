package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// user input hello :) return 😊
	// user input goodbye :( return ☹️
	// user input hello :) goodbye :( return 😊 ☹️

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	if strings.Contains(input, ":)") && strings.Contains(input, ":(") {
		fmt.Println("hello 😊 goodbye ☹️")
	} else if strings.Contains(input, ":(") {
		fmt.Println("goodbye ☹️")
	} else if strings.Contains(input, ":)") {
		fmt.Println("hello 😊")
	} else {
		return
	}
}
