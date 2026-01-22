# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED=0 is required for glebarez/sqlite to work purely in Go without C dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Seed the database during build
# This creates tugas1.db with initial data
RUN SEED_DATA=true SEED_EXIT=true ./main

# Runtime Stage
FROM alpine:latest  

# Install tzdata for correct time handling
RUN apk add --no-cache tzdata

WORKDIR /app

# Copy the Pre-built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/web ./web
# Copy the seeded database
COPY --from=builder /app/tugas1.db .

# Ensure database permissions are correct (readable and writable)
# SQLite also needs write permission for the directory to create journal files
RUN chmod 666 tugas1.db && chmod 777 .

# Expose port 8080 to the outside world
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Command to run the executable
CMD ["./main"]
