package todoType

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"text"`
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("invalid data")
	}

	newNote := &Todo{
		Content: content,
	}

	return newNote, nil
}

func (content Todo) Display() {
	fmt.Println(content)
}

func (content Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(content)

	if err != nil {
		return err
	}

	// will return nil if no error
	return os.WriteFile(filename, json, 0644)
}
