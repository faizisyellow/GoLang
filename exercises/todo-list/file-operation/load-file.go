package fo

import (
	"encoding/json"
	"os"

	taskstruct "example.com/todo-list/task"
)

func LoadFromFile(filename string) ([]taskstruct.Task, error) {
	var tasks []taskstruct.Task

	file, err := os.ReadFile(filename)
	if err != nil {
		// If file does not exist, return an empty task list
		if os.IsNotExist(err) {
			return []taskstruct.Task{}, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
