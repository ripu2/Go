package main

import "fmt"

var revenue, expense, taxRate float64

func main() {
	revenue = inputTaker("Enter revenue: ")
	expense = inputTaker("Enter expense: ")
	taxRate = inputTaker("Enter tax rate: ")

	var earningsBeforeTax, profit, ratio = calulateProfit()

	fmt.Println("EBT is ", earningsBeforeTax)
	fmt.Println("Profit is ", profit)
	fmt.Printf("Ratio is %.2f ", ratio)

}

func calulateProfit() (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return
}

func inputTaker(label string) (input float64) {
	fmt.Print(label)
	fmt.Scan(&input)
	return
}
