package main

import (
	"fmt"
	"strconv"
)

func main() {

	revenue, _ := strconv.ParseFloat(input("Revenue: "), 64)
	expense, _ := strconv.ParseFloat(input("Expense: "), 64)
	taxRate, _ := strconv.ParseFloat(input("Tax: "), 64)

	ebt, profit, ratio := calculateFinancials(revenue, expense, taxRate)

	// display ebt, profit and ration
	fmt.Printf("ebt: %0.f\n", ebt)
	fmt.Printf("profit: %0.f\n", profit)
	fmt.Printf("ratio: %0.f\n", ratio)
}
