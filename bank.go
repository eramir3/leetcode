package main

import (
	"fmt"

	"example.com/bank/fileops"
)

func main() {
	const inflationRate float64 = 2.5 // Assuming a constant inflation rate of 2% for the calculation
	var investmentAmount float64
	var years int
	var expectedReturnRate float64

	// fmt.Print("Enter the initial investment amount: ")
	outputText("Please enter the initial investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter the expected return rate: ")
	outputText("Please enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter the number of years: ")
	outputText("Please enter the number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := fileops.CalculateFutureValues(investmentAmount, expectedReturnRate, years)
	fmt.Println("Future Value of Investment: ", futureValue)
	fmt.Println("Future Real Value of Investment: ", futureRealValue)
}
