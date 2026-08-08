package Utility

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "Balance.txt"

func WriteBalanceToFile(balanceWrite float64) {
	balanceText := fmt.Sprint(balanceWrite)
	os.WriteFile(fileName, []byte(balanceText), 0644)

}
func GetValueFromFile(fileN string) float64 {
	data, _ := os.ReadFile(fileN)
	balanceData := string(data)
	balance, _ := strconv.ParseFloat(balanceData, 64)
	return balance
}
