package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(5, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	result, message := multiple(5, 2)
	fmt.Print(message)
	fmt.Println(result)

	fmt.Println(split(18))
}
