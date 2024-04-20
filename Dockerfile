# Use the official Go image as a base image for the build stage
FROM golang:1.21 AS builder

WORKDIR /app

# Copy the Go mod and sum files, and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Change to the directory containing main.go
WORKDIR /app/cmd

# Build the application
RUN go build -o ../main .

# Start a new stage from scratch
FROM debian:bookworm-slim

WORKDIR /app

# Use non-root user
RUN useradd -m dietuser
USER dietuser

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

ENTRYPOINT ["/app/main"]
