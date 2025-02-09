package main

import (
	"fmt"
	"math"
)

func main2() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter initial investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter year: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected annual return rate (in percentage): ")
	fmt.Scan(&expectedReturnRate)

	var netReturn = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var inflationAdjustedReturn = netReturn / (math.Pow(1+inflationRate/100, years))

	fmt.Printf("Inflation-adjusted return: INR %.2f\n", inflationAdjustedReturn)
}
