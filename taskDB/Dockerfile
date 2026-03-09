# Multi-stage build
# Stage 1: Build
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum* ./
COPY main.go .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o taskCli .

# Stage 2: Runtime
FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/taskCli .

ENTRYPOINT ["./taskCli"]
CMD ["-list"]
