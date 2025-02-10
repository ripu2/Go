package banking

import (
	"fmt"

	"example.com/first-app/packages/communication"
)

func CheckBalance(balance float64) {
	fmt.Printf("Your current balance is: %.2f\n", balance)
}

func DepositMoney(balance float64, setBalance func(float64)) {
	var amount = communication.TakeInput("Enter the amount to deposit: ")
	if amount <= 0 {
		fmt.Printf("Please enter a valid amount!!!!")
		return
	}
	setBalance(balance + float64(amount))
	fmt.Printf("Deposit Successful. New balance is: %.2f\n", balance)

}

func WithdrawMoney(balance float64, setBalance func(float64)) {
	var amount = communication.TakeInput("Enter the amount to withdraw: ")
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
