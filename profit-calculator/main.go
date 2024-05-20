package main

import "fmt"

func main() {

	revenue := readInput("Enter the revnue amount: ")
	expenses := readInput("Enter the expenses amount: ")
	taxRate := readInput("Enter the tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("ETB: %.2f ---- Profit: %.2f ---- Ration: %.2f\n", ebt, profit, ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func readInput(text string) float64 {
	var in float64
	fmt.Print(text)
	fmt.Scan(&in)
	return in
}
