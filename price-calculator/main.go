package main

import "github.com/mobamoh/go-playground/price-calculator/price"

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		job := price.New(taxRate)
		job.Process()
	}
}
