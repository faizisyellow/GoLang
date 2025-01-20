package main

import "fmt"

func main() {
	age := 20

	fmt.Printf("Before age: %v\n", age)

	getAdultsYear(&age)
	fmt.Printf("After age: %v\n", age)

	x := 1
	y := 2

	fmt.Printf("x is %v, y is %v\n", x, y)
	// pass a pointer
	swap(&x, &y)
	fmt.Printf("x is %v, y is %v", x, y)

}

func getAdultsYear(n *int) {

	// go to that address and change the value.
	*n -= 18

}

func swap(a, b *int) {
	temp := *a

	*a = *b
	*b = temp
}
