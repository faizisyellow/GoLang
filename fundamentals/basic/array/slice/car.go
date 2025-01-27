package main

import "fmt"

func Car() {
	cars := make([]string, 2)
	fmt.Printf("len=%d cap=%d %v\n", len(cars), cap(cars), cars)
	cars[0] = "Lambo"
	cars[1] = "Lexus"

	// create new array
	cars = append(cars, "Honda")
	fmt.Printf("FIRST APPEND: len=%d cap=%d %v\n", len(cars), cap(cars), cars)

	cars = append(cars, "Mercy")
	fmt.Printf("SECOND APPEND: len=%d cap=%d %v\n", len(cars), cap(cars), cars)

	// create new array
	cars = append(cars, "Audi")
	fmt.Printf("THIRD APPEND: len=%d cap=%d %v\n", len(cars), cap(cars), cars)

}
