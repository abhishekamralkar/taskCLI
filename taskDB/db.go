package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// initDB opens (or creates) the SQLite database file and ensures the schema exists.
func initDB(filepath string) {
	var err error
	db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	// Verify the connection is alive.
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	createTable()
}

// createTable creates the tasks table if it does not already exist.
func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id    INTEGER PRIMARY KEY AUTOINCREMENT,
		done  BOOLEAN NOT NULL DEFAULT 0,
		title TEXT    NOT NULL
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
}

// addTask inserts a new task into the database.
func addTask(text string) {
	_, err := db.Exec("INSERT INTO tasks (title) VALUES (?)", text)
	if err != nil {
		log.Fatalf("failed to add task: %v", err)
	}
	fmt.Printf("✔ Added: %s\n", text)
}

// listTasks prints all tasks ordered by ID.
func listTasks() {
	rows, err := db.Query("SELECT id, done, title FROM tasks ORDER BY id")
	if err != nil {
		log.Fatalf("failed to query tasks: %v", err)
	}
	defer rows.Close()

	fmt.Printf("\n%-4s | %-4s | %s\n", "ID", "Done", "Task")
	fmt.Println("-----|------|----------------------------------")

	count := 0
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Done, &t.Title); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		done := "[ ]"
		if t.Done {
			done = "[✔]"
		}
		fmt.Printf("%-4d | %-4s | %s\n", t.ID, done, t.Title)
		count++
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("row iteration error: %v", err)
	}

	if count == 0 {
		fmt.Println("No tasks found. Add one with: todo -add \"your task\"")
	}
	fmt.Println()
}

// markDone sets a task's done flag to true by ID.
func markDone(id int) {
	res, err := db.Exec("UPDATE tasks SET done = 1 WHERE id = ?", id)
	if err != nil {
		log.Fatalf("failed to update task: %v", err)
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		fmt.Printf("no task found with ID %d\n", id)
		return
	}
	fmt.Printf("✔ Marked as done: task %d\n", id)
}

// deleteTask removes a task from the database by ID.
func deleteTask(id int) {
	res, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatalf("failed to delete task: %v", err)
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		fmt.Printf("no task found with ID %d\n", id)
		return
	}
	fmt.Printf("✔ Deleted task %d\n", id)
}
