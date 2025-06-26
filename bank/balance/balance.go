package balance

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// User first letter as capital letter to export the functions
func WriteBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func ReadBalanceFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("testing own error")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("testing own error")
	}
	return balance, nil
}
