package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	moreNumbers := []int{2, 4, 6}

	double := tranformsNumbers(&numbers, double)
	fmt.Println(double)

	// return a function
	tranformNumberFn1 := transformNumbers(&moreNumbers)

	// execute the function
	anotherNumbers := tranformNumberFn1(2)
	fmt.Print(anotherNumbers)
}

func transformNumbers(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		// anonymous function
		return func(n int) int {
			return n * 2
		}
	} else {
		return tripple
	}
}

func tranformsNumbers(n *[]int, tranform func(int) int) []int {
	dnumbers := []int{}

	for _, v := range *n {
		dnumbers = append(dnumbers, tranform(v))
	}

	return dnumbers
}

func double(n int) int {
	return n * 2
}
func tripple(n int) int {

	return n * 3
}
