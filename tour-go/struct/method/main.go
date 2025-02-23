package main

import (
	"fmt"

	words "example.com/method/custom-string-method"
	"example.com/method/user"
)

func main() {
	myAccount := user.New("lizzy mcalpine", "lizzy@music.com", "hello-world")

	updateAccount := myAccount.ResetPassword()
	updateAccount.SetPassword("anjir")

	var hello words.Str = "anjir"

	hello.UpperOne()

	fmt.Println(*myAccount)
	fmt.Println(updateAccount)
	fmt.Println(hello)
}
