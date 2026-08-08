package main

import (
	"fmt"

	"example.com/pilot-app/Utility"
)

const fileName = "Balance.txt"

func calc() {
	//var revenue float64
	//var expense float64
	//var taxRate float64

	/*fmt.Print("\n\n\n---------------- calculater------------------- \n\n\n")
	EBT := revenue - expense
	profit := EBT - (revenue * (taxRate / 100))

	return EBT, profit */
	fmt.Println("\nWelcome to the bank of Hirunika! \nWhat do you want to do today?")
	fmt.Println("1. Check Balance \n2. Deposit Money \n3. Withdraw Money \n4. Exit")

	var choice int
	index := 1

	for {
		fmt.Print("Please enter your choice: ")
		fmt.Scan(&choice)
		//calc(custChoice)
		switch choice {
		case 1:
			value := userInfo()
			fmt.Println("Your balance is: ", value)
		case 2:
			depositMoney()
		case 3:
			withdrawMoney()
		default:
			fmt.Println("Pleasure doing business with you")
		}
		index = choice
		if index >= 4 {
			break
		}

	}

}

func userInfo() float64 {
	info := Utility.GetValueFromFile(fileName)
	return info

}

func depositMoney() {
	var depAmmount float64
	fmt.Print("Please enter your ammount to deposit: ")
	fmt.Scan(&depAmmount)
	info := userInfo()
	newBal := info + depAmmount
	fmt.Println("Your ammount has been deposit \nYour new ammount is: ", newBal)
	Utility.WriteBalanceToFile(newBal)
}

func withdrawMoney() {
	var withAmmount float64
	fmt.Print("Please enter your ammount to withdraw: ")
	fmt.Scan(&withAmmount)
	info := userInfo()
	if info > withAmmount {
		newBal := info - withAmmount
		fmt.Println("Your ammount has been withdrown \nYour new ammount is: ", newBal)
		Utility.WriteBalanceToFile(newBal)
	} else {
		fmt.Println("You cannot withdraw this much......")
	}

}
