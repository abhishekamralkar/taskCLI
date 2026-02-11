# TaskCLI - Go Task Management Suite

A collection of command-line task management applications written in Go, demonstrating different approaches to building CLI tools and data persistence.

## Project Structure

This monorepo contains two implementations of task management tools:

### 1. **taskSimple** - In-Memory Task Manager
A lightweight CLI application with Docker support and security scanning capabilities.

**Features:**
- In-memory task storage
- Quick and responsive
- Docker containerization with multi-stage builds
- Security scanning with Trivy
- Make-based build system

**Use Case:** Ideal for temporary task tracking or learning purposes.

**Quick Start:**
```bash
cd taskSimple
make build
./taskcli -list
```

**Learn More:** See [taskSimple/README.md](taskSimple/README.md)

---

### 2. **taskStorage** - Persistent File-Based Task Manager
A file-based task manager that persists data to JSON, ensuring tasks survive application restarts.

**Features:**
- JSON file-based persistence
- Automatic database creation and updates
- Durable task storage
- Simple and lightweight

**Use Case:** Perfect for maintaining task lists across sessions.

**Quick Start:**
```bash
cd taskStorage
go run main.go -list
```

**Learn More:** See [taskStorage/README.md](taskStorage/README.md)

---

## Common Commands

Both implementations support the following operations:

### List all tasks
```bash
go run main.go -list
# or (taskSimple)
./taskcli -list
```

### Add a new task
```bash
go run main.go -add "Your task description"
```

### Mark a task as complete
```bash
go run main.go -done <task_id>
```

### Delete a task
```bash
go run main.go -del <task_id>
# or (taskSimple)
go run main.go -delete <task_id>
```

---

## Installation

### Prerequisites
- Go 1.25.4 or higher
- (Optional) Docker for containerized execution
- (Optional) Make for build automation

### Clone the Repository
```bash
git clone https://github.com/abhishekamralkar/taskcli.git
cd taskcli
```

### Choose Your Implementation

**For In-Memory (taskSimple):**
```bash
cd taskSimple
make build
./taskcli --help
```

**For Persistent Storage (taskStorage):**
```bash
cd taskStorage
go build -o taskcli
./taskcli -help
```

---

## Comparison Matrix

| Feature | taskSimple | taskStorage |
|---------|-----------|------------|
| **Persistence** | In-memory | JSON file |
| **Data Survives Restart** | ❌ No | ✅ Yes |
| **Docker Support** | ✅ Yes | ✅ Yes |
| **Security Scanning** | ✅ Trivy | ✅ Yes |
| **Use Case** | Learning/Testing | Production Tasks |
| **Build Tool** | Make | Go build |

---

## Development

### Project Layout
```
taskCli/
├── taskSimple/           # In-memory CLI with Docker
│   ├── main.go
│   ├── Dockerfile
│   ├── Makefile
│   └── README.md
├── taskStorage/          # File-based persistent CLI
│   ├── main.go
│   ├── db.json
│   └── README.md
├── go.mod
├── LICENSE
└── README.md
```

### Running Tests
```bash
cd taskSimple
make test

cd ../taskStorage
go test ./...
```

### Building Docker Images
```bash
cd taskSimple
make docker-build
make docker-run
```

---

## License
See [LICENSE](LICENSE) file for details.

---

## Contributing
Feel free to fork, modify, and improve these implementations!

## Author
Abhishek Amralkar
