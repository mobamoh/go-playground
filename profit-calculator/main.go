package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := readInput("Enter the revnue amount: ")
	if err != nil {
		panic(err)
	}
	expenses, err := readInput("Enter the expenses amount: ")
	if err != nil {
		panic(err)
	}
	taxRate, err := readInput("Enter the tax rate: ")
	if err != nil {
		panic(err)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	if saveFinancialstoFile(ebt, profit, ratio) != nil {
		panic("couldn't save financials to file")
	}
	fmt.Printf("ETB: %.2f ---- Profit: %.2f ---- Ration: %.2f\n", ebt, profit, ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func readInput(text string) (float64, error) {
	var in float64
	fmt.Print(text)
	fmt.Scan(&in)
	if in <= 0 {
		return 0, errors.New("invalid input. value should be >=0")
	}
	return in, nil
}

func saveFinancialstoFile(ebt, profit, ratio float64) error {
	return os.WriteFile("financials.txt", []byte(fmt.Sprintf("%f,%f,%f", ebt, profit, ratio)), 0644)
}
