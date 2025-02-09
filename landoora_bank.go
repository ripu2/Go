package main

import "fmt"

var balance float64

func main() {
	fmt.Println("Welcome to landoora Bank")

	for {
		if !performBanking() {
			break
		}
	}
	fmt.Println("Goodbye !!, Thanks for choosing our bank, Have a good day")

}

func performBanking() bool {
	var choice = printMenu()

	switch choice {
	case 1:
		checkBalance()
		return true
	case 2:
		depositMoney()
		return true
	case 3:
		withdrawMoney()
		return true

	default:
		return false
	}
}

func checkBalance() {
	fmt.Printf("Your current balance is: %.2f\n", balance)
}

func depositMoney() {
	var amount = takeInput("Enter the amount to deposit: ")
	if amount <= 0 {
		fmt.Printf("Please enter a valid amount!!!!")
		return
	}
	balance += float64(amount)
	fmt.Printf("Deposit Successful. New balance is: %.2f\n", balance)

}

func withdrawMoney() {
	var amount = takeInput("Enter the amount to withdraw: ")
	if amount <= 0 {
		fmt.Printf("Please enter a valid amount!!!!")
		return
	}
	if amount > balance {
		fmt.Println("Insufficient balance")
	} else {

		balance -= float64(amount)
		fmt.Printf("Withdrawal Successful. New balance is: %.2f\n", balance)
	}
}

func printMenu() (choice float64) {
	fmt.Println("What do you want to do")
	fmt.Println("1: Check Balance")
	fmt.Println("2: Deposit Money")
	fmt.Println("3: Withdraw Money")
	fmt.Println("4: Exit")

	choice = takeInput("Enter Your Choice: ")
	return
}

func takeInput(label string) (input float64) {
	fmt.Print(label)
	fmt.Scan(&input)
	return
}
