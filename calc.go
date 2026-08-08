package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToFile(balanceWrite float64) {
	balanceText := fmt.Sprint(balanceWrite)
	os.WriteFile("Balance.txt", []byte(balanceText), 0644)

}
func getValueFromFile() float64 {
	data, _ := os.ReadFile("Balance.txt")
	balanceData := string(data)
	balance, _ := strconv.ParseFloat(balanceData, 64)
	return balance
}

func calc(choise int) {
	//var revenue float64
	//var expense float64
	//var taxRate float64

	/*fmt.Print("\n\n\n---------------- calculater------------------- \n\n\n")
	EBT := revenue - expense
	profit := EBT - (revenue * (taxRate / 100))

	return EBT, profit */

	switch choise {
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

}

func userInfo() float64 {
	info := getValueFromFile()
	return info

}

func depositMoney() {
	var depAmmount float64
	fmt.Print("Please enter your ammount to deposit: ")
	fmt.Scan(&depAmmount)
	info := userInfo()
	newBal := info + depAmmount
	fmt.Println("Your ammount has been deposit \nYour new ammount is: ", newBal)
	writeBalanceToFile(newBal)
}

func withdrawMoney() {
	var withAmmount float64
	fmt.Print("Please enter your ammount to withdraw: ")
	fmt.Scan(&withAmmount)
	info := userInfo()
	if info > withAmmount {
		newBal := info - withAmmount
		fmt.Println("Your ammount has been withdrown \nYour new ammount is: ", newBal)
		writeBalanceToFile(newBal)
	} else {
		fmt.Println("You cannot withdraw this much......")
	}

}
