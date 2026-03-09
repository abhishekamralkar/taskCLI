package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addFlag := flag.String("add", "", "Add a new task")
	listFlag := flag.Bool("list", false, "List all tasks")
	doneFlag := flag.Int("done", 0, "Mark a task as done by ID")
	deleteFlag := flag.Int("delete", 0, "Delete a task by ID")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of taskcli:\n")
		fmt.Fprintf(os.Stderr, "  taskcli -add \"task description\"   Add a new task\n")
		fmt.Fprintf(os.Stderr, "  taskcli -list                     List all tasks\n")
		fmt.Fprintf(os.Stderr, "  taskcli -done <id>                Mark a task as done\n")
		fmt.Fprintf(os.Stderr, "  taskcli -delete <id>              Delete a task\n")
	}

	flag.Parse()

	// Open (or create) the SQLite database.
	initDB("todo.db")

	switch {
	case *addFlag != "":
		addTask(*addFlag)
	case *listFlag:
		listTasks()
	case *doneFlag != 0:
		markDone(*doneFlag)
	case *deleteFlag != 0:
		deleteTask(*deleteFlag)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
