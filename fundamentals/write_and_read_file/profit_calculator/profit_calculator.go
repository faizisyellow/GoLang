package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

/**
 * 1. Validate user input
 *  show error & exit if invalid input is provided
 *  - no negative numbers
 *  - not 0
 * 2. Store calculate results into file
 **/

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		log.Fatal(err)
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		log.Fatal(err)

	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		log.Fatal(err)

	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	storeResultToFile(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)

	if err != nil {
		panic(err)
	}

	if userInput <= 0 {
		return 0, errors.New("only postive number allowed")
	}

	return userInput, nil
}

func storeResultToFile(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT: %1.f\nPROFIT: %1.f\nRATIO: %1.f\n", ebt, profit, ratio)

	os.WriteFile("result.txt", []byte(result), 0644)

}
