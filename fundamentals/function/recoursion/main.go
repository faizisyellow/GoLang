package main

import "fmt"

func factor(n int) int {
	if n == 1 {
		return 1
	} else {
		fmt.Println("called")
		return n * factor(n-1)
	}
}

func main() {

	fmt.Println(factor(5))
}
