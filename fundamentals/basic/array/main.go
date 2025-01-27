package main

import "fmt"

func main() {
	cities := [8]string{"Los angels", "Seattle", "San dieago", "San francisco", "New york", "jersey city", "philadelphia", "chicago"}

	fmt.Printf("Cities: %v\n", cities)
	/*
		Slicing array.
		[Start index :  Last index]
	*/
	westArea := cities[:4]
	fmt.Printf("Slice of cities, West Area: %v\n", westArea)
	fmt.Printf("Length of West Area: %v, Cap of West Area: %v\n", len(westArea), cap(westArea))

	var a [4]int64
	fmt.Printf("Array Before: %v\n", a)
	modifyArray(&a)
	fmt.Printf("Array After: %v\n", a)

}

func modifyArray(a *[4]int64) {
	(*a)[0] = 99 // dereference the array pointer first, then access index.

	// or, Go allows dereference automatically.
	a[0] = 99
}
