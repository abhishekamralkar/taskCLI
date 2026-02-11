package main

import (
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func addTask(list []Task, task string) []Task {
	newId := len(list) + 1
	newTask := Task{
		ID:   newId,
		Task: task,
		Done: false,
	}

	return append(list, newTask)
}

func listTasks(list []Task) {
	if len(list) == 0 {
		fmt.Println("No task in the list.. Yay!")
		return
	}
	fmt.Println("┌────┬──────────────────────────────────┬────────┐")
	fmt.Println("│ ID │ Task                             │ Status │")
	fmt.Println("├────┼──────────────────────────────────┼────────┤")
	for _, task := range list {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("│ %2d │ %-32s │ %s    │\n", task.ID, truncate(task.Task, 32), status)
	}
	fmt.Println("└────┴──────────────────────────────────┴────────┘")
}

func deleteTask(list []Task, id int) []Task {
	for i, t := range list {
		if t.ID == id {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func completeTask(list []Task, id int) bool {
	for i, t := range list {
		if t.ID == id {
			list[i].Done = true
			return true
		}
	}
	return false
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Todo Tool:\n")
		fmt.Fprintf(os.Stderr, "  -add string:  Add a new task\n")
		fmt.Fprintf(os.Stderr, "  -list:        List all tasks\n")
		fmt.Fprintf(os.Stderr, "  -delete int:  Remove a task by ID\n")
		fmt.Fprintf(os.Stderr, "  -done int:    Mark a task as done\n")
	}

	todoList := []Task{
		{ID: 1, Task: "Buy groceries"},
		{ID: 2, Task: "Wash car"},
		{ID: 3, Task: "Pay bills"},
	}

	addFlag := flag.String("add", "", "Task Description")
	listFlag := flag.Bool("list", false, "List all tasks")
	deleteFlag := flag.Int("delete", 0, "Task ID to delete")
	doneFlag := flag.Int("done", 0, "Mark task as done by ID")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Error: No command provided.")
		flag.Usage()
		return
	}

	if *addFlag != "" {
		todoList = addTask(todoList, *addFlag)
		fmt.Println("✓ Task added successfully.")
		listTasks(todoList)
	}

	if *deleteFlag != 0 {
		todoList = deleteTask(todoList, *deleteFlag)
		fmt.Println("✓ Task deleted.")
		listTasks(todoList)
	}

	if *doneFlag != 0 {
		if completeTask(todoList, *doneFlag) {
			fmt.Printf("✓ Task %d marked as done.\n", *doneFlag)
			listTasks(todoList)
		} else {
			fmt.Printf("✗ Task %d not found.\n", *doneFlag)
		}
	}

	if *listFlag {
		listTasks(todoList)
	}

}
