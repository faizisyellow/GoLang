package notetype

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("invalid data")
	}

	newNote := &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	return newNote, nil
}

func (note Note) Display() {

	fmt.Printf("Your note titled %v has the following content:\n\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	filename := snakecase(note.Title)
	filename += ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	// will return nil if no error
	return os.WriteFile(filename, json, 0644)
}

func snakecase(str string) string {
	result := strings.ReplaceAll(str, " ", "_")
	result = strings.ToLower(result)

	return result
}
