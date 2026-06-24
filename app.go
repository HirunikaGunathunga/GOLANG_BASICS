package main

import (
	"fmt"
	"math"
)

func main() {
	//invAmount, years, expectedReturn := 1000.0, 10.0, 5.5
	var invAmount float64
	var expectedReturn float64
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&invAmount)

	fmt.Print("Expected Return: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := (invAmount) * math.Pow(1+expectedReturn/100, (years))

	fmt.Println(futureValue)
}
