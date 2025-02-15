package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IoManager         iomanager.IoManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%2.f", price)] = fmt.Sprintf("%2.f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = result

	job.IoManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IoManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func NewTaxIncludedPriceJob(io iomanager.IoManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IoManager: io,
		TaxRate:   taxRate,
	}
}
