package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var years float64 = 10
	expectedReturnRate := 5.5

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
