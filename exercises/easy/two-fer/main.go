package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)

	shareName := shareWith(name)
	fmt.Println(shareName)
}

func shareWith(name string) string {
	if name == "" {
		name = "you"
	}

	getName := fmt.Sprintf("One for %v, one for me", name)

	return getName
}
