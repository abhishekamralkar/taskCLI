# Task CLI

A simple command-line task manager that stores tasks in a JSON file.

## Features

- **Add tasks**: Create new tasks and save them to the database
- **List tasks**: View all tasks with their completion status
- **Mark complete**: Mark tasks as done
- **Delete tasks**: Remove tasks from the list

## Installation

Ensure you have Go installed, then clone or download this project.

```bash
cd taskStorage
go run main.go [flags]
```

## Usage

### List all tasks
```bash
go run main.go -list
```

Output:
```
ID  | Done | Task
----|------|-------------------
1   | [ ]  | Buy groceries
2   | [X]  | Complete project
```

### Add a new task
```bash
go run main.go -add "Your task description"
```

Example:
```bash
go run main.go -add "Write documentation"
```

### Mark a task as complete
```bash
go run main.go -done <task_id>
```

Example:
```bash
go run main.go -done 1
```

### Delete a task
```bash
go run main.go -del <task_id>
```

Example:
```bash
go run main.go -del 2
```

## Data Storage

Tasks are stored in `db.json` in JSON format. The file is automatically created and updated when you add, complete, or delete tasks.

Example `db.json`:
```json
[
 {
  "id": 1,
  "task": "Buy groceries",
  "done": false
 },
 {
  "id": 2,
  "task": "Complete project",
  "done": true
 }
]
```

## Build

To compile the application into an executable:

```bash
go build -o taskcli
./taskcli -list
```
