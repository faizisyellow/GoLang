package taskstruct

type Task struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func New(taskName string) Task {
	newTask := Task{Name: taskName, Status: "progress"}
	return newTask
}
