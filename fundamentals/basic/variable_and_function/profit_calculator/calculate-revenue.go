package main

func calculateFinancials(revenue, expense, tax float64) (float64, float64, float64) {
	// calculate revenue before tax.
	ebt := revenue - expense

	// calculate revenue after tax.
	profit := ebt * (1 - tax/100)

	ratio := ebt / profit

	return ebt, profit, ratio
}
