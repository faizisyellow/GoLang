package main

import (
	"fmt"
	"os"

	ct "example.com/todo-list/complete-task"
	fo "example.com/todo-list/file-operation"
	so "example.com/todo-list/show-options"
	taskstruct "example.com/todo-list/task"
)

func main() {
	var options int
	var task string

	fmt.Println("Task Manager")

	so.ShowListOptions(&options)

	switch options {
	case 1:

		fmt.Print("New Task: ")
		fmt.Scan(&task)

		newTask := taskstruct.New(task)

		tasks, errLoadFile := fo.LoadFromFile("task.json")
		if errLoadFile != nil {
			fmt.Printf("Error while load task: %v", errLoadFile.Error())
		}

		tasks = append(tasks, newTask)

		errSaveFile := fo.SaveToFile("task.json", tasks)

		if errSaveFile != nil {
			fmt.Printf("Error while saving task: %v", errSaveFile.Error())
			break
		}

		for i, v := range tasks {
			fmt.Printf("%v. %v %v\n", i+1, v.Name, v.Status)
		}

	case 2:
		var task string
		fmt.Print("Task Name: ")
		fmt.Scan(&task)
		ct.CompleteTask(task)

	case 3:
		tasks, errLoadFile := fo.LoadFromFile("task.json")
		if errLoadFile != nil {
			fmt.Printf("Error while load task: %v", errLoadFile.Error())
		}

		for i, v := range tasks {
			fmt.Printf("%v. %v %v\n", i+1, v.Name, v.Status)
		}

	case 4:
		fmt.Println("exit...")
		os.Exit(1)

	default:
		fmt.Println("Invalid input. Try again")
	}

}
