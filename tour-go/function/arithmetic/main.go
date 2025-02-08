package main

import "fmt"

type Arithmetic struct {
	addition       int
	substraction   int
	multiplication int
	division       int
}

func main() {
	sum := addition(3, 5)
	subs := subtraction(10, 5)
	multiply := multiplication(5, 2)
	div := division(10, 2)

	arithmetic := Arithmetic{addition: sum, substraction: subs, multiplication: multiply, division: div}

	fmt.Printf("ADDITION: %v\nSUBSTRACTION: %v\nMULTIPLICATION: %v\nDIVISION: %v\n", arithmetic.addition, arithmetic.substraction, arithmetic.multiplication, arithmetic.division)
}

func addition(x, y int) int {
	sum := x + y

	return sum
}
func subtraction(x, y int) int {
	difference := x - y

	return difference
}
func multiplication(x, y int) int {
	result := x * y

	return result
}

func division(x, y int) int {
	result := x / y

	return result
}
