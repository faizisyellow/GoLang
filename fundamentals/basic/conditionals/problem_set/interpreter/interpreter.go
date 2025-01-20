package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1. Input user.
	// 2. Split user's input into 3 variable (x,y,z)
	// 3. Do math operation.

	fmt.Print("Expression: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	x, y, z := extractString(input)

	switch y {
	case "+":
		fmt.Printf("%.1f", x+z)
	case "-":
		fmt.Printf("%.1f", x-z)
	case "*":
		fmt.Printf("%.1f", x*z)
	case "/":
		fmt.Printf("%.1f", x/z)

	}

}
func extractString(s string) (float64, string, float64) {
	splitString := strings.Split(s, " ")

	x, _ := strconv.ParseFloat(splitString[0], 64)
	y := splitString[1]
	z, _ := strconv.ParseFloat(splitString[2], 64)

	return x, y, z
}
