package main

import "fmt"

func main() {

	var revenue, expenses, taxRate float64

	fmt.Print("Enter the revnue amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses amount: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("ETB: %.2f ---- Profit: %.2f ---- Ration: %.2f\n", ebt, profit, ratio)

}
