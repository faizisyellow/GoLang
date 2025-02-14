package main

import "fmt"

func main() {
	double := createTransform(2)
	tripple := createTransform(3)

	fmt.Println("Factor 2: ", double(4))
	fmt.Println("Factor 3: ", tripple(6))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
