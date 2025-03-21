package prices

import (
	"Go-Basics/price-calculator/util"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	IoManager         util.IoManager    `json:"-"`
}

func New(io util.IoManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
		IoManager:   io,
	}
}

func (job TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	return job.IoManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IoManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := util.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}
