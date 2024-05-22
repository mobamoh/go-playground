package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	accountBalance, err := readBalanceFromFile()
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to Trivial Go Bank!")
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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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

func writeBalanceToFile(balance float64) {
	os.WriteFile("balance.txt", []byte(fmt.Sprint(balance)), 0644)
}

func readBalanceFromFile() (float64, error) {
	balancByte, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("no balance account found, please contact admin.")
	}
	balance, err := strconv.ParseFloat(string(balancByte), 64)
	if err != nil {
		return 0, errors.New("couldn't parse the balance amount, please contact admin.")
	}
	return balance, nil
}
