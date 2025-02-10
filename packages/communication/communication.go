package communication

import "fmt"

func PrintMenu() (choice float64) {
	fmt.Println("What do you want to do")
	fmt.Println("1: Check Balance")
	fmt.Println("2: Deposit Money")
	fmt.Println("3: Withdraw Money")
	fmt.Println("4: Exit")

	choice = TakeInput("Enter Your Choice: ")
	return
}

func TakeInput(label string) (input float64) {
	fmt.Print(label)
	fmt.Scan(&input)
	return
}
