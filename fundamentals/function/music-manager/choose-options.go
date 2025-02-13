package main

import "fmt"

func chooseOptions() (option int) {
	fmt.Println("Options available: ")
	fmt.Println("1. Add new song")
	fmt.Println("2. Show all song")
	fmt.Println("3. Remove song")
	fmt.Println("4. Filter song")
	fmt.Println("5. Sort song")

	fmt.Print("Your choose: ")
	fmt.Scan(&option)

	return
}
