package main

import (
	"fmt"
	"math"
)

func main() {
	var invAmount float64 = 1000
	var expectedReturn float64 = 5.5
	var years float64 = 10

	var futureValue = (invAmount) * math.Pow(1+expectedReturn/100, (years))

	fmt.Println(futureValue)
}
