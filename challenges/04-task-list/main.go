package main

import (
	"fmt"
)

func main() {
	learnGo := Task{
		name:        "Learn GO",
		description: "Learn the programming language Go",
		completed:   false,
	}

	implementService := Task{
		name:        "Implement service",
		description: "Implement a service with Go",
		completed:   false,
	}

	trivialTask := Task{
		name:        "Meh...",
		description: "meh...",
		completed:   false,
	}

	implementService.markCompleted()

	var tasks TaskList
	tasks.addTask(&learnGo)
	tasks.addTask(&implementService)
	tasks.addTask(&trivialTask)

	for i, task := range tasks.Tasks {
		fmt.Println(i, task)
	}

	tasks.deleteTask(2)

	for i := 0; i < len(tasks.Tasks); i++ {
		tasks.Tasks[i].printTask()
	}

}

// Task Class
type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) printTask() {
	// https://pkg.go.dev/fmt#hdr-Printing
	fmt.Printf("Name: %s\nDescription: %s\nCompleted: %t\n", t.name, t.description, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

// Task List
type TaskList struct {
	Tasks []*Task
}

func (tl *TaskList) addTask(t *Task) {
	fmt.Println("Adding task...")
	tl.Tasks = append(tl.Tasks, t)
}

func (tl *TaskList) deleteTask(index int) {
	fmt.Println("Deleting task...")
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
}
