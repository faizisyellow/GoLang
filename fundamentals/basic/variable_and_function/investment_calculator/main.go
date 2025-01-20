package main

import (
	"fmt"
	"math"
)

// Constant variable.
const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Invesment Amount: ")

	// Get input and store in variable.
	fmt.Scan(&investmentAmount)

	outputText("Years :")
	fmt.Scan(&years)

	outputText("Expacted Rate :")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Future value (adjusted for Inflation ): %.2f\n", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
