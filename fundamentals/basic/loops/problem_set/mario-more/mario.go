package main

import "fmt"

func main() {
	input := inputHeight()

	displayTwoPyramid(input)

}

func displayTwoPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Print("  ")

		for l := 0; l <= i; l++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}
}

func inputHeight() int {
	var input int

	for {
		if input <= 0 {
			fmt.Print("Height: ")

			_, err := fmt.Scan(&input)
			if err != nil {
				break
			}
		} else {
			break
		}
	}

	return input
}
