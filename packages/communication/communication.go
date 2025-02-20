package communication

import "fmt"

func PrintMenu() (choice string) {
	fmt.Println("What do you want to do")
	fmt.Println("1: Check Balance")
	fmt.Println("2: Deposit Money")
	fmt.Println("3: Withdraw Money")
	fmt.Println("4: Exit")

	choice = TakeInput("Enter Your Choice: ")
	return
}

func TakeInput(label string) (input string) {
	fmt.Print(label)
	fmt.Scanln(&input)
	return
}

func PrintValues(label string, value string) {
	fmt.Printf("%s: %s\n", label, value)
}
