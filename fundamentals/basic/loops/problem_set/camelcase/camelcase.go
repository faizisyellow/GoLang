package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1. Input user.
	// 2. Loop through the strings.
	// 3. Find the upper one and add _ before lower it.

	fmt.Print("camelCase: ")
	var input string

	fmt.Scan(&input)

	fmt.Printf("snake_case: %v\n", toSnakeCase(input))

}

func toSnakeCase(s string) string {
	var result string

	for i, r := range s {
		if unicode.IsUpper(r) {
			// if not the first character in the strings
			if i > 0 {
				// add underscore
				result += "_"
			}
			// then, add the uppercase character
			result += strings.ToLower(string(r))
		} else {
			result += string(r)
		}
	}

	return result
}
