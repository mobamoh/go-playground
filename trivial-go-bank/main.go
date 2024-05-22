package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/mobamoh/go-playground/trivial-go-bank/fileops"
)

func main() {

	accountBalance, err := fileops.ReadBalanceFromFile()
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to Trivial Go Bank!")
	fmt.Println("Contact us 24/7: ", randomdata.PhoneNumber())
	for {
		showMenu()
		input := readInput("Enter your choice:")
		if input > 4 || input < 1 {
			fmt.Println("Invalid Choice!")
			continue
		}

		switch input {

		case 1:
			fmt.Printf("You balance is: %.2f\n", accountBalance)

		case 2:
			amount := readInput("Enter the amount you want to deposit:")
			if amount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}
			accountBalance += amount
			fileops.WriteBalanceToFile(accountBalance)
		case 3:
			amount := readInput("Enter the amount you want to withdraw:")

			if amount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}
			if amount > accountBalance {
				fmt.Println("Not enough balance to perform this operation!")
				continue
			}
			accountBalance -= amount
			fileops.WriteBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}

func showMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func readInput(text string) float64 {
	fmt.Println(text)
	var input float64
	fmt.Scan(&input)
	return input
}
