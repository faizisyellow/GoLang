package main

import "fmt"

func main() {
	// CREATE SLICE ITERALS.
	wholeNumber := []int32{0, 1, 2, 3, 4, 5}
	fmt.Printf("Whole Number Slice: %v\n", wholeNumber)

	// CREATE SLICE WITH MAKE FUNCTION.
	integers := make([]int32, 5)
	fmt.Printf("Integers Slice: %v\n", integers)
	integers[0] = 100

	// ADDING NEW VALUES IN SLICE
	integers = append(integers, -1, 0, 1, 2, 3)
	fmt.Printf("Integers V1 Slice: %v\n", integers)

	fmt.Printf("Integers Length %v, Cap%v\n", len(integers), cap(integers))

	// fmt.Printf("My Scores: %v", ScoresOut())

	/*
		Slice literals.
		A slice literal is like an array literal without the length.
	*/
	ints := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(ints)
	trueIntegers := ints[1:]
	trueIntegers[0] = 12
	fmt.Println(trueIntegers)
	fmt.Println(ints)

	fmt.Println("=======CAR SLICE========")
	Car()
}
