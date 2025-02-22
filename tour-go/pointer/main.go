package main

import "fmt"

func main() {
	var n int = 50
	fmt.Printf("VALUE OF N: %v\n", n)
	fmt.Printf("ADDRESS OF N: %v\n", &n)

	// pointers variable (a variable that store an address)
	var p *int = &n
	fmt.Printf("VALUE OF P: %v\n", p)
	fmt.Printf("ADDRESS OF P: %v\n", &p)

	var z **int = &p
	fmt.Printf("VALUE OF Z: %v\n", z)
	fmt.Printf("ADDRESS OF Z: %v\n", &z)

	fmt.Printf("If i go to the value pointers P i got: %v\n", *p)
	fmt.Printf("If i go to the value pointers Z i got: %v\n", *z)
}
