package main

import (
	"fmt"
	"os"
)

func main() {
	accountBalance := 100.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check  balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Your balance is: %v\n", accountBalance)
	case 2:
		var depositAmount float64

		fmt.Printf("Your deposit: ")
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Printf("Invalid amount. Must be greater than 0.")
			return
		}
		accountBalance += depositAmount
		fmt.Printf("Balance updated! New amount: %v\n", accountBalance)
	case 3:

		var withdrawAmount float64

		fmt.Printf("How much you want to withdraw: ")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Printf("Invalid amount. Must be greater than 0.")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Printf("Can not withdraw money more than you have ")
			return
		}
		accountBalance -= withdrawAmount
		fmt.Printf("Balance updated! withdraw successfull: %v\n", accountBalance)
	default:
		// exit in the terminal.
		os.Exit(1)

		// Could use return.
		// return
	}

}
