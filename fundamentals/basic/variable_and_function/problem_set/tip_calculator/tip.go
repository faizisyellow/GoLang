package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1. Input dollars as string return float
	// 2. Input percent as string return float
	// 3. Calculate meal * percent

	var dollars string
	var percent string

	fmt.Print("How much was the meal? ")
	fmt.Scan(&dollars)

	fmt.Print("What percentage would you like to tip? ")
	fmt.Scan(&percent)

	tip := dollarsToFloat(dollars) * percentToFloat(percent)
	fmt.Printf("Leave: $%.2f", tip)
}

func dollarsToFloat(dollars string) float64 {

	dollarWithoutSign := strings.Trim(dollars, "$")

	dollarRealValue, _ := strconv.ParseFloat(dollarWithoutSign, 64)

	return dollarRealValue
}

func percentToFloat(percent string) float64 {
	percentWithoutSign := strings.Trim(percent, "%")

	percentRealValue, _ := strconv.ParseFloat(percentWithoutSign, 64)

	// Divided by 100 to get percentage.
	percentRealValue /= 100

	return percentRealValue

}
