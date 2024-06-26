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

	futureValue, futureRealValue := calculateValues(investmentAmount, expectedReturnRate, inflationRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adj. for inflation): %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func calculateValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv = fv / math.Pow((1+inflationRate/100), years)
	return
}
