package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 100, errors.New("balance not found")
	}

	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 100, errors.New("failed to parse")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	// 0644 (permission to write file)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println(err)
		// to exit while error
		panic("Error")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check  balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is: %v\n", accountBalance)
		} else if choice == 2 {
			var depositAmount float64

			fmt.Printf("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Printf("Invalid amount. Must be greater than 0.")

				// To continue looping back from start.
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Balance updated! New amount: %v\n", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64

			fmt.Printf("How much you want to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Printf("Invalid amount. Must be greater than 0.")

				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Printf("Can not withdraw money more than you have ")

				continue
			}
			accountBalance -= withdrawAmount
			fmt.Printf("Balance updated! withdraw successfull: %v\n", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {

			// Exit in the terminal.
			os.Exit(1)

			// To exit the loops.
			break
		}
	}
	fmt.Println("Thanks for choosing our banks")
}
