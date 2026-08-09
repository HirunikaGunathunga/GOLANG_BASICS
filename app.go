package main

import (
	"fmt"

	note "example.com/pilot-app/Note"
)

// "math"

func main() {
	//invAmount, years, expectedReturn := 1000.0, 10.0, 5.5
	/*v ar invAmount float64
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
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expense)

	fmt.Print("Tax: ")
	fmt.Scan(&taxRate)

	EBT, Profit := calc(revenue, expense, taxRate)
	fmt.Println("EBT: ", EBT, "\nProfit: ", Profit) */
	var func1 int
	fmt.Println("What func you need? ")
	fmt.Scan(&func1)

	switch func1 {
	case 1:
		calc()
	case 2:
		pointers()
	case 3:
		structFunc()
	default:
		note.NoteFunc()
	}

}
