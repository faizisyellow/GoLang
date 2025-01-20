package main

import "fmt"

func main() {
	var x int
	var y int

	// input x
	fmt.Print("what's x: ")
	fmt.Scan(&x)

	// input y
	fmt.Print("what's y: ")
	fmt.Scan(&y)

	result := x + y

	fmt.Print(result)
}
