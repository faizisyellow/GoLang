package main

import "fmt"

func main() {
	xInt := 1
	yInt := 5
	z := add(xInt, yInt)
	fmt.Println(z)

	xFloat := 2.5
	yFloat := 5.5

	zFloat := add(xFloat, yFloat)
	fmt.Println(zFloat)

	greet := "hello "
	who := "lizzy mcalpine"

	greetings := add(greet, who)
	fmt.Println(greetings)
}

// generic function (dynamic type)
func add[T int | float64 | string](x, y T) T {
	fmt.Printf("Passing argument x as: %T\n", x)
	fmt.Printf("Passing argument y as: %T\n", y)
	return x + y
}
