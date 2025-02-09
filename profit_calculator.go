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

	outputGenerator("EBT is ", earningsBeforeTax)
	outputGenerator("Profit is ", profit)
	outputGenerator("Ratio is ", ratio)

	outputGenerator("Type check %T", earningsBeforeTax)
}

func outputGenerator(label string, value float64) {
	fmt.Printf("%s: %.2f\n", label, value)
}
