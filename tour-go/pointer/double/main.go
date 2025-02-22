package main

import "fmt"

func main() {
	var x int = 60
	var y int = 30

	var p *int = &x

	var z **int = &p
	*z = &y
	**z = 50
	fmt.Println(**z)
	fmt.Println(y)
}
