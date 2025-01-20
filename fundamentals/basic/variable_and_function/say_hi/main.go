package main

import "fmt"

func main() {
	name := "john doe"
	fmt.Print("Name :")
	fmt.Scan(&name)

	hello := "Hi, " + name
	fmt.Println(hello)
}
