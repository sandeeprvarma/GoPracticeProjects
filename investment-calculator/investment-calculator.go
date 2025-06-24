package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentValue float64 = 1000
	// var expectedReturnRate = 5.5
	// var years float64 = 10
	// investmentValue := 1000.0
	// expectedReturnRate := 5.5
	// years := 10.0
	var investmentValue float64
	const inflationRate = 6.5
	years, expectedReturnRate := 10.0, 5.5
	fmt.Println("Investment Amount")
	fmt.Scan(&investmentValue)

	fmt.Println("Expected return rate")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Investment years")
	fmt.Scan(&years)
	futureValue := investmentValue * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue * math.Pow((1+inflationRate/100), years)
	fmt.Println(futureValue)
	fmt.Print(futureRealValue)
}
