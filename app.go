package main

import "fmt"

//"math"

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

	fmt.Println("\nWelcome to the bank of Hirunika! \nWhat do you want to do today?")
	fmt.Println("1. Check Balance \n2. Deposit Money \n3. Withdraw Money \n4. Exit")

	var custChoice int
	index := 1

	for {
		fmt.Print("Please enter your choice: ")
		fmt.Scan(&custChoice)
		calc(custChoice)
		index = custChoice
		if index >= 4 {
			break
		}

	}

}
