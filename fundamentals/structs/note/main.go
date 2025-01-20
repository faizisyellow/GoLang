package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	notetype "example.com/note/note_type"
)

func main() {
	title, content := getNoteData()

	note, err := notetype.New(title, content)

	if err != nil {
		panic(err)
	}
	note.Display()

	err = note.Save()
	if err != nil {
		panic(err)
	}

	fmt.Println("Saving the note succedeed!")

}

func getNoteData() (string, string) {
	var title string
	var content string

	input("Note Title: ", &title)

	input("Note Content: ", &content)

	return title, content

}

func input(prompt string, value *string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	*value = text
}
