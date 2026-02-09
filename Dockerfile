# Multi-stage build
# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o taskCli .

# Stage 2: Runtime
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/taskCli .

ENTRYPOINT ["./taskCli"]
CMD ["-list"]
