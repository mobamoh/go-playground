package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceTextFile = "balance.txt"

func WriteBalanceToFile(balance float64) {
	os.WriteFile(balanceTextFile, []byte(fmt.Sprint(balance)), 0644)
}

func ReadBalanceFromFile() (float64, error) {
	balancByte, err := os.ReadFile(balanceTextFile)
	if err != nil {
		return 0, errors.New("no balance account found, please contact admin")
	}
	balance, err := strconv.ParseFloat(string(balancByte), 64)
	if err != nil {
		return 0, errors.New("couldn't parse the balance amount, please contact admin")
	}
	return balance, nil
}
