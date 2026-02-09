# Task CLI - Todo Task Manager

A simple command-line todo task manager written in Go with Docker support and security scanning.

## Features

- ✅ Add tasks
- ✅ List tasks with status indicators
- ✅ Mark tasks as done
- ✅ Delete tasks
- 🐳 Containerized with multi-stage Docker build
- 🔒 Security scanning with Trivy

## Usage

### Local Usage

```bash
# List all tasks
./taskCli -list

# Add a new task
./taskCli -add "Buy milk"

# Mark a task as done
./taskCli -done 1

# Delete a task
./taskCli -delete 1
```

### Using Make

```bash
# Build the binary
make build

# Run locally
make run

# Build Docker image
make docker-build

# Run in Docker
make docker-run

# Scan for vulnerabilities
make trivy-scan
```

## Docker

### Multi-stage Build

The Dockerfile uses a two-stage build process:

1. **Builder Stage**: Compiles the Go binary using `golang:1.25-alpine`
2. **Runtime Stage**: Minimal `alpine:latest` image with only the compiled binary

This approach keeps the final image size small (~15MB).

### Building and Running

```bash
# Build the image
docker build -t taskcli:latest .

# Run the container
docker run --rm taskcli:latest -list

# Run with specific command
docker run --rm taskcli:latest -add "New task"
```

## Security Scanning with Trivy

[Trivy](https://github.com/aquasecurity/trivy) is used to scan the Docker image for vulnerabilities.

### Install Trivy

```bash
# On macOS
brew install trivy

# On Linux (Ubuntu/Debian)
sudo apt-get install trivy

# Using Docker
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock aquasecurity/trivy image taskcli:latest
```

### Running Scans

```bash
# Full scan with JSON report
make trivy-scan

# Manual scan with terminal output
trivy image taskcli:latest

# Scan specific severity levels
trivy image --severity CRITICAL,HIGH taskcli:latest

# Generate HTML report
trivy image --format template --template '@/contrib/html.tpl' -o report.html taskcli:latest
```

### Scan Results

The `trivy-scan` target will:
1. Build the Docker image
2. Generate a JSON report (`trivy-report.json`)
3. Display a summary of CRITICAL and HIGH severity vulnerabilities

## Project Structure

```
taskCli/
├── main.go          # Application source code
├── go.mod           # Go module definition
├── Dockerfile       # Multi-stage Docker build
├── Makefile         # Build and utility targets
└── README.md        # This file
```

## Commands Reference

| Command | Description |
|---------|-------------|
| `make help` | Show all available targets |
| `make build` | Build the binary |
| `make run` | Build and run locally |
| `make clean` | Clean build artifacts |
| `make fmt` | Format Go code |
| `make docker-build` | Build Docker image |
| `make docker-run` | Run in Docker container |
| `make docker-clean` | Remove Docker image/container |
| `make trivy-scan` | Scan image for vulnerabilities |

## Requirements

- Go 1.25+
- Docker (for containerization)
- Trivy (for security scanning)

## License

MIT
