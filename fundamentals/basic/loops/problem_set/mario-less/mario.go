package main

import "fmt"

func main() {
	height := inputHeight()
	displayPyramid(height)
}

func displayPyramid(h int) {
	for i := 0; i < h; i++ {
		for j := h; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}
}

func inputHeight() int {
	var height int

	for {
		if height <= 0 {
			fmt.Print("Height: ")

			_, err := fmt.Scan(&height)
			if err != nil {
				// if not string, exit
				break
			}

		} else {
			break
		}
	}

	return height
}
