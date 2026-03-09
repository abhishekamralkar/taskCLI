package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID          int
	Title       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

const createTableSQL = `
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL,
    completed_at DATETIME
);
CREATE INDEX IF NOT EXISTS idx_done ON tasks(done);
CREATE INDEX IF NOT EXISTS idx_created_at ON tasks(created_at);
`

type TodoDB struct {
	db *sql.DB
}

func NewTodoDB(dbPath string) (*TodoDB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Enable foreign keys and WAL mode for better concurrency
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}
	if _, err := db.Exec("PRAGMA journal_mode = WAL"); err != nil {
		return nil, err
	}

	// Create tables
	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &TodoDB{db: db}, nil
}

func (t *TodoDB) Close() error {
	return t.db.Close()
}

// Add a new task
func (t *TodoDB) AddTask(title string) (*Task, error) {
	result, err := t.db.Exec(
		"INSERT INTO tasks (title, done, created_at) VALUES (?, ?, ?)",
		title, false, time.Now(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add task: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return t.GetTask(int(id))
}

// Get a single task
func (t *TodoDB) GetTask(id int) (*Task, error) {
	task := &Task{}
	var completedAt sql.NullTime

	err := t.db.QueryRow(
		"SELECT id, title, done, created_at, completed_at FROM tasks WHERE id = ?",
		id,
	).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt, &completedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task not found")
	}
	if err != nil {
		return nil, err
	}

	if completedAt.Valid {
		task.CompletedAt = &completedAt.Time
	}

	return task, nil
}

// List all tasks
func (t *TodoDB) ListTasks(showDone bool) ([]*Task, error) {
	query := "SELECT id, title, done, created_at, completed_at FROM tasks"
	if !showDone {
		query += " WHERE done = 0"
	}
	query += " ORDER BY created_at DESC"

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		task := &Task{}
		var completedAt sql.NullTime

		err := rows.Scan(
			&task.ID, &task.Title, &task.Done,
			&task.CreatedAt, &completedAt,
		)
		if err != nil {
			return nil, err
		}

		if completedAt.Valid {
			task.CompletedAt = &completedAt.Time
		}

		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

// Mark task as done
func (t *TodoDB) CompleteTask(id int) error {
	result, err := t.db.Exec(
		"UPDATE tasks SET done = 1, completed_at = ? WHERE id = ?",
		time.Now(), id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}

// Delete a task
func (t *TodoDB) DeleteTask(id int) error {
	result, err := t.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}

func main() {
	// Command-line flags
	addTask := flag.String("add", "", "Add a new task")
	listTasks := flag.Bool("list", false, "List all tasks")
	showDone := flag.Bool("done", false, "Show completed tasks (use with -list)")
	completeID := flag.Int("complete", 0, "Mark task as complete")
	deleteID := flag.Int("delete", 0, "Delete a task")

	flag.Parse()

	// Initialize database
	db, err := NewTodoDB("todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Handle -add flag
	if *addTask != "" {
		handleAdd(db, *addTask)
		return
	}

	// Handle -list flag
	if *listTasks {
		handleList(db, *showDone)
		return
	}

	// Handle -complete flag
	if *completeID > 0 {
		handleComplete(db, *completeID)
		return
	}

	// Handle -delete flag
	if *deleteID > 0 {
		handleDelete(db, *deleteID)
		return
	}

	// No flags provided
	flag.Usage()
}

func handleAdd(db *TodoDB, title string) {
	task, err := db.AddTask(title)
	if err != nil {
		log.Fatalf("Error adding task: %v", err)
	}
	fmt.Printf("✔ Added task #%d: %s\n", task.ID, task.Title)
}

func handleList(db *TodoDB, showDone bool) {
	tasks, err := db.ListTasks(showDone)
	if err != nil {
		log.Fatalf("Error listing tasks: %v", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}

	fmt.Printf("%-4s | %-6s | %s\n", "ID", "Done", "Task")
	fmt.Println("-----|--------|-----------------------------")

	for _, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%-4d | %-6s | %s\n", task.ID, status, task.Title)
	}
}

func handleComplete(db *TodoDB, id int) {
	if err := db.CompleteTask(id); err != nil {
		log.Fatalf("Error completing task: %v", err)
	}
	fmt.Printf("✔ Completed task #%d\n", id)
}

func handleDelete(db *TodoDB, id int) {
	if err := db.DeleteTask(id); err != nil {
		log.Fatalf("Error deleting task: %v", err)
	}
	fmt.Printf("✔ Deleted task #%d\n", id)
}
