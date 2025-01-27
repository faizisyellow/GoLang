package main

import "fmt"

func main() {
	// cities := [8]string{"Los angels", "Seattle", "San dieago", "San francisco", "New york", "jersey city", "philadelphia", "chicago"}

	// fmt.Printf("Cities: %v\n", cities)
	/*
		Slicing array.
		[Start index :  Last index]
	*/
	// westArea := cities[:4]
	// fmt.Printf("Adress West Area before realocate: %p\n", westArea)

	// westArea = append(westArea, []string{"jakarta", "asd", "vdas", "sdah", "asjdha", "skdai", "asdjai"}...)
	// westArea[0] = "broklyn"
	// fmt.Printf("West Area: %v\n", westArea)
	// fmt.Printf("Salah Area: %v\n", westArea)
	// fmt.Printf("Cities: %v\n", cities)
	// cities[0] = "anjir"
	// fmt.Printf("West Area: %v\n", westArea)

	// fmt.Printf("Adress West Area: %p\n", westArea)
	// fmt.Printf("Address cities: %p\n", &cities[0])

	// var a [4]int64
	// fmt.Printf("Array Before: %v\n", a)
	// modifyArray(&a)
	// fmt.Printf("Array After: %v\n", a)

	s1 := make([]int, 5, 10)
	fmt.Printf("BEFORE ASIGNMENT: %v\n", s1)
	fmt.Printf("Length: %v, Cap: %v\n", len(s1), cap(s1))

	length := cap(s1)
	for i := 0; i < length; i++ {
		s1 = append(s1, i*2)
	}

	fmt.Printf("AFTER ASIGNMENT: %v\n", s1)
	fmt.Printf("Length: %v, Cap: %v\n", len(s1), cap(s1))
}

// func modifyArray(a *[4]int64) {
// 	(*a)[0] = 99 // dereference the array pointer first, then access index.

// 	// or, Go allows dereference automatically.
// 	a[0] = 99
// }
