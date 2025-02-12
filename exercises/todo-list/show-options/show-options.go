package so

import "fmt"

func ShowListOptions(options *int) {

	fmt.Println("Options: ")
	fmt.Println("1. Add new task")
	fmt.Println("2. Complete task")
	fmt.Println("3. Show list task")
	fmt.Println("4. Exit program")
	fmt.Scan(options)
}
