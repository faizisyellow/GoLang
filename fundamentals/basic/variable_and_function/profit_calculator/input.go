package main

import "fmt"

func input(text string) string {
	// print something
	fmt.Print(text)

	// get input from user
	var value string
	fmt.Scan(&value)

	return value
}
