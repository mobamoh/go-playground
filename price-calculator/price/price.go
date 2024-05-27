package price

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate     float64
	Prices      []float64
	TaxedPrices map[string]float64
}

func New(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	result := make(map[string]float64)
	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	job.TaxedPrices = result
	job.OutputJson()
}

func (job *TaxIncludedPriceJob) LoadPrices() {

	file, _ := os.Open("prices.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var priceLines []float64
	for scanner.Scan() {
		price, _ := strconv.ParseFloat(scanner.Text(), 64)
		priceLines = append(priceLines, price)
	}
	job.Prices = priceLines
}

func (job *TaxIncludedPriceJob) OutputJson() {
	fileName := fmt.Sprintf("taxed-prices-%.0f.json", job.TaxRate*100)
	file, _ := os.Create(fileName)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(job)
}
