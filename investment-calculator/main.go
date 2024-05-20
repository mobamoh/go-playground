package main

import (
	"fmt"
	"math"
)

const InflationRate = 2.5

func main() {

	var investmentValue, expectedReturnRate, years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentValue)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentValue * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+InflationRate/100, years)

	fmt.Printf("The future investment is %.2f\n", futureRealValue)
}
