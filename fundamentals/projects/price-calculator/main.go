package main

import (
	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result-%.0f.json", taxRate*100))
		cmd := cmdmanager.New()
		result := prices.NewTaxIncludedPriceJob(*cmd, taxRate)
		result.Process()
	}

}
