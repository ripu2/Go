package banking

import (
	"fmt"
	"strconv"

	"example.com/first-app/packages/communication"
)

func CheckBalance(balance float64) {
	fmt.Printf("Your current balance is: %.2f\n", balance)
}

func DepositMoney(balance float64, setBalance func(float64)) {
	var amount, _ = strconv.ParseFloat(communication.TakeInput("Enter the amount to deposit: "), 64)
	if amount <= 0 {
		fmt.Printf("Please enter a valid amount!!!!")
		return
	}
	setBalance(balance + float64(amount))
	fmt.Printf("Deposit Successful. New balance is: %.2f\n", balance)

}

func WithdrawMoney(balance float64, setBalance func(float64)) {
	var amount, _ = strconv.ParseFloat(communication.TakeInput("Enter the amount to withdraw: "), 64)
	if amount <= 0 {
		fmt.Printf("Please enter a valid amount!!!!")
		return
	}
	if amount > balance {
		fmt.Println("Insufficient balance")
	} else {

		setBalance(balance - float64(amount))
		fmt.Printf("Withdrawal Successful. New balance is: %.2f\n", balance)
	}
}
