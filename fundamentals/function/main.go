package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	double := tranformsNumbers(&numbers, double)
	fmt.Println(double)
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
