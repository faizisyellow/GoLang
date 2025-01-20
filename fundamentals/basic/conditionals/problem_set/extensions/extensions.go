package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Input user.
	// 2. If user's input valid media types return type of the file.
	// 3. Otherwise not valid or there's no suffix return application/octet-stream.

	fmt.Print("File Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputSanitize := strings.ToLower(input)

	fmt.Println(getExtension(inputSanitize))
}

func getExtension(s string) string {

	if strings.Contains(s, ".") {
		_, after, _ := strings.Cut(s, ".")

		result := strings.Split(after, ".")
		lastElement := result[len(result)-1]

		switch lastElement {
		case "jpg", "jpeg":
			return "image/jpeg"
		case "gif":
			return "image/gif"
		case "png":
			return "image/png"
		case "pdf":
			return "application/pdf"
		case "txt":
			return "text/plain"
		case "zip":
			return "application/zip"
		default:
			return "application/octet-stream"
		}

	} else {
		return "application/octet-stream"
	}

}
