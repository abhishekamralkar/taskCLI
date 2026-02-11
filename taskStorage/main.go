package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

const dbFile = "db.json"

func loadDB() ([]Task, error) {

	file, err := os.ReadFile(dbFile)

	if err != nil {
		return []Task{}, nil
	}

	var tasks []Task

	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func saveDB(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile(dbFile, data, 0644)
}

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "Update task")
	del := flag.Int("del", 0, "Delete task")
	flag.Parse()

	tasks, err := loadDB()
	if err != nil {
		log.Fatal(err)
	}

	if *add != "" {
		newID := 1
		if len(tasks) > 0 {
			newID = tasks[len(tasks)-1].ID + 1
		}
		tasks = append(tasks, Task{ID: newID, Task: *add, Done: false})
		saveDB(tasks)
		fmt.Printf("✔ Added: %s\n", *add)
		return
	}

	if *list {
		if len(tasks) == 0 {
			fmt.Println("Your todo list is empty.")
			return
		}
		fmt.Println("ID  | Done | Task")
		fmt.Println("----|------|-------------------")
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "X"
			}
			fmt.Printf("%-3d | [%s]  | %s\n", t.ID, status, t.Task)
		}
		return
	}

	if *done > 0 {
		success := false
		for i := range tasks {
			if tasks[i].ID == *done {
				tasks[i].Done = true
				success = true
			}
		}
		if success {
			saveDB(tasks)
			fmt.Printf("✔ Task %d marked as complete.\n", *done)
		} else {
			fmt.Printf("✘ Task %d not found.\n", *done)
		}
		return
	}

	if *del > 0 {
		newTasks := []Task{}
		found := false
		for _, t := range tasks {
			if t.ID == *del {
				found = true
				continue
			}
			newTasks = append(newTasks, t)
		}
		if found {
			saveDB(newTasks)
			fmt.Printf("✔ Task %d deleted.\n", *del)
		} else {
			fmt.Printf("✘ Task %d not found.\n", *del)
		}
		return
	}
	flag.Usage()
}
