package ct

import (
	"errors"
	"fmt"

	fo "example.com/todo-list/file-operation"
	taskstruct "example.com/todo-list/task"
)

func CompleteTask(task string) {
	data, errLoadFile := fo.LoadFromFile("task.json")
	if errLoadFile != nil {
		fmt.Printf("Error while loading tasks: %v\n", errLoadFile)
		return
	}

	if task == "" {
		fmt.Println(errors.New("task name is required"))
		return
	}

	var filteredTasks []taskstruct.Task
	var foundTask bool

	for _, v := range data {
		if v.Name != task {
			filteredTasks = append(filteredTasks, v)
		} else {
			foundTask = true
		}
	}

	if !foundTask {
		fmt.Println("Task not found")
		return
	}

	if err := fo.SaveToFile("task.json", filteredTasks); err != nil {
		fmt.Printf("Error while saving the task: %v\n", err)
		return
	}

	fmt.Println("Task completed successfully!")
}
