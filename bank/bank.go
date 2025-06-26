package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountBalance, _ = readBalanceFile()

func main() {

	// for i := 0; i < 2; i++ {
	for {
		choice := welcomeUser()
		if choice == 4 {
			break
		}
		bankChoiceActions(choice)
		writeBalanceToFile(accountBalance)
	}
	// fmt.Println("Final Balance:", accountBalance)
}

func welcomeUser() (choice int) {
	fmt.Println("Welcome to go bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposite Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Scan(&choice)
	return choice
}

func bankChoiceActions(choice int) {
	var deposite, withdraw float64
	// if else if
	if choice == 1 {
		fmt.Println("Your account balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Println("How much money do you want to deposite?")
		fmt.Scan(&deposite)
		if deposite <= 0 {
			fmt.Println("Invalid amount.")
		}
		// accountBalance = accountBalance + deposite
		accountBalance += deposite
		fmt.Println("Money deposited. Your account balance is: ", accountBalance)
	} else if choice == 3 {
		fmt.Println("How much money do you want to withdraw?")
		fmt.Scan(&withdraw)
		if withdraw > accountBalance {
			fmt.Println("Invalid amount.")
		}
		// accountBalance = accountBalance - withdraw
		accountBalance -= withdraw
		fmt.Println("Your account balance is: ", accountBalance)
	} else {
		fmt.Println("Thank you for chooseing our bank!")
	}
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func readBalanceFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("Testing own error")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Testing own error")
	}
	return balance, nil
}
