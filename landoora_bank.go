package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/first-app/packages/banking"
	"example.com/first-app/packages/communication"
	"example.com/first-app/packages/fileops"
	"github.com/joho/godotenv"
)

var balance float64
var balanceFileNmae string

func landoora_bank() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	balanceFileNmae = os.Getenv("FILE_NAME")
	balanceText, _ := fileops.ReadFromFile(balanceFileNmae)
	balanceValue, _ := strconv.ParseFloat(balanceText, 64)
	balance = balanceValue

	for {
		if !performBanking() {
			break
		}
	}
	fmt.Println("Goodbye !!, Thanks for choosing our bank, Have a good day")

}

func updateBalance(newBalance float64) {
	balance = newBalance
}

func performBanking() bool {
	var choice = communication.PrintMenu()

	switch choice {
	case 1:
		banking.CheckBalance(balance)
		return true
	case 2:
		banking.DepositMoney(balance, updateBalance)
		return true
	case 3:
		banking.WithdrawMoney(balance, updateBalance)
		return true

	default:
		fileops.WriteToFile(balanceFileNmae, balance)
		return false
	}
}
