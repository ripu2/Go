package main

import "fmt"

func main() {
	var revenue, expense, taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expense: ")
	fmt.Scan(&expense)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expense
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Println("EBT is ", earningsBeforeTax)
	fmt.Println("Profit is ", profit)
	fmt.Println("Ratio is ", ratio)
}
