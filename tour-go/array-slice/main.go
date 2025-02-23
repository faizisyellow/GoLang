package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a, len(a), cap(a))
}
