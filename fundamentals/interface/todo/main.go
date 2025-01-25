package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	todoType "example.com/todo/todo_type"
)

// interface type
type saver interface {
	Save() error
}

// embeded existing interface
type outputtable interface {
	saver
	Display()
}

func main() {
	content := getTodoData()

	todo, _ := todoType.New(content)

	outputData(todo)

	fmt.Println("Saving the content succedeed!")

}

func saveData(value saver) {
	err := value.Save()
	if err != nil {
		panic(err)
	}
}

func outputData(value outputtable) {
	value.Display()

	err := value.Save()
	if err != nil {
		panic(err)
	}
}

func getTodoData() string {
	var content string

	input("Todo content: ", &content)

	return content

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
