package main

import (
	"fmt"
)

func calc() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("\n\n\n---------------- calculater------------------- \n\n\n")
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expense)

	fmt.Print("Tax: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expense
	profit := EBT - (revenue * (taxRate / 100))

	fmt.Println("EBT: ", EBT)
	fmt.Println("Profit: ", profit)

}
