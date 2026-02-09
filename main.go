package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of Todo Tool:\n")
		fmt.Fprintf(os.Stderr, "  -add string:  Add a new task\n")
		fmt.Fprintf(os.Stderr, "  -list:        List all tasks\n")
		fmt.Fprintf(os.Stderr, "  -delete int:  Remove a task by ID\n")
		fmt.Fprintf(os.Stderr, "  -done int:    Mark a task as done\n")
		fmt.Fprintf(os.Stderr, "  -storage string: Storage backend (slice, file) [default: slice]\n")
	}

	addFlag := flag.String("add", "", "Task Description")
	listFlag := flag.Bool("list", false, "List all tasks")
	deleteFlag := flag.Int("delete", 0, "Task ID to delete")
	doneFlag := flag.Int("done", 0, "Mark task as done by ID")
	storageFlag := flag.String("storage", "slice", "Storage backend: slice or file")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Error: No command provided.")
		flag.Usage()
		return
	}

	// Initialize storage backend
	var storage Storage
	var err error

	switch *storageFlag {
	case "file":
		storage, err = NewFileStorage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to initialize file storage: %v\n", err)
			return
		}
	case "slice":
		storage = NewSliceStorage()
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown storage backend '%s'\n", *storageFlag)
		return
	}

	// Initialize display and service
	display := &TableDisplay{}
	service := NewTaskService(storage, display)

	// Execute commands
	if *addFlag != "" {
		if err := service.AddTask(*addFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}

	if *deleteFlag != 0 {
		if err := service.DeleteTask(*deleteFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}

	if *doneFlag != 0 {
		if err := service.CompleteTask(*doneFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}

	if *listFlag {
		if err := service.ListTasks(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
}
