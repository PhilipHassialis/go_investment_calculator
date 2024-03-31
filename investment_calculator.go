package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years of invesment: ")
	fmt.Scan(&years)
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	futureRealValue := calculateFutureRealValue(futureValue, inflationRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adj. for inflation): %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) float64 {
	return investmentAmount * math.Pow((1+expectedReturnRate/100), years)
}

func calculateFutureRealValue(futureValue, inflationRate, years float64) float64 {
	return futureValue / math.Pow((1+inflationRate/100), years)
}
