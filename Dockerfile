# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED=0 is required for glebarez/sqlite to work purely in Go without C dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Runtime Stage
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
