package main

import "fmt"

func main() {
	fmt.Println("Welcome to Trivial Go Bank!")
	var accountBalance float64 = 1000
	for {
		showMenu()
		input := readInput("Enter your choice:")
		if input > 4 || input < 1 {
			fmt.Println("Invalid Choice!")
			continue
		}

		if input == 1 {
			fmt.Printf("You balance is: %.2f\n", accountBalance)

		} else if input == 2 {
			amount := readInput("Enter the amount you want to deposit:")
			if amount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}
			accountBalance += amount
		} else if input == 3 {
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
		} else {
			fmt.Println("Goodbye!")
			break
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
