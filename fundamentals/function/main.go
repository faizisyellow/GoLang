/*
Deep dive into function in golang
*/
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// pass function as argument without param and return.
	double := transformNumbers(&numbers, doubleing)
	fmt.Println(double)

	// pass function as argument without param and return.
	tripples := transformNumbers(&numbers, tripple)

	fmt.Println(tripples)
}

func transformNumbers(x *[]int, transform func(int) int) []int {
	var dbnumber []int
	for _, v := range *x {
		// calling transform function with one argument
		dbnumber = append(dbnumber, transform(v))
	}

	return dbnumber
}

func doubleing(x int) int {
	return x * 2
}

func tripple(i int) int {
	return i * 3
}
