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
./taskcli -list

# Add a new task
./taskcli -add "Buy milk"

# Mark a task as done
./taskcli -done 1

# Delete a task
./taskcli -delete 1
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

# Generate HTML report (using default HTML template)
trivy image --format template --template '@contrib/html.tpl' -o report.html taskcli:latest

# Or generate simple JSON to HTML conversion
trivy image --format json -o report.json taskcli:latest
```

### Scan Results

The `trivy-scan` target will:
1. Build the Docker image
2. Generate a JSON report (`trivy-report.json`)
3. Display a summary of CRITICAL and HIGH severity vulnerabilities

#### Sample Output

```bash
$ trivy image --severity CRITICAL,HIGH taskcli:latest

2026-02-08T21:45:00Z	INFO	Vulnerability scanning is enabled
2026-02-08T21:45:00Z	INFO	Secret scanning is enabled
2026-02-08T21:45:00Z	INFO	If your scanning is slow, try a shallow scan with --scanners vuln or --scanners secret
2026-02-08T21:45:00Z	INFO	Detected OS: alpine
2026-02-08T21:45:00Z	INFO	Detecting Alpine vulnerabilities...
2026-02-08T21:45:01Z	INFO	Number of language-specific files: 0

taskcli:latest (alpine 3.19.1)
================================
Total: 0 (CRITICAL: 0, HIGH: 0)
```

The scan shows zero CRITICAL and HIGH severity vulnerabilities for the minimal Alpine-based image, which is expected since we only include the compiled binary in the final stage.

### Generating SBOM (Software Bill of Materials)

Trivy can also generate a Software Bill of Materials for supply chain security and compliance.

```bash
# Generate SBOM in CycloneDX format (JSON)
trivy image --format cyclonedx --output sbom.json taskcli:latest

# Generate SBOM in SPDX format (JSON)
trivy image --format spdx-json --output sbom-spdx.json taskcli:latest

# Generate SBOM in SPDX format (text)
trivy image --format spdx taskcli:latest
```

#### Sample SBOM Output (CycloneDX)

```json
{
  "bomFormat": "CycloneDX",
  "specVersion": "1.4",
  "version": 1,
  "metadata": {
    "timestamp": "2026-02-08T21:45:00Z",
    "component": {
      "bom-ref": "docker.io/taskcli:latest@sha256:abc123",
      "type": "application",
      "name": "taskcli:latest",
      "version": "latest"
    }
  },
  "components": []
}
```

The SBOM output is particularly useful for:
- **Supply chain security** - Track all dependencies in your image
- **Compliance** - Meet regulatory requirements (SLSA, etc.)
- **Vulnerability tracking** - Use with tools like Dependency-Track
- **License management** - Identify all licenses in use

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
| `make sbom` | Generate SBOM (Software Bill of Materials) |

## Requirements

- Go 1.25+
- Docker (for containerization)
- Trivy (for security scanning)

## License

MIT
