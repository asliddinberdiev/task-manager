# Build stage
FROM golang:1.22.2-alpine AS builder
# Install necessary dependencies
RUN apk update && \
    apk upgrade --update-cache --available && \
    apk add git make curl bash
    
# Set the working directory
WORKDIR /app
# Copy Go module files 
COPY go.mod go.sum ./
# Download packages dependencies
RUN go mod download
# Copy the entire application code
COPY . .
# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
# Migration cli 
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz


# Final stage
FROM alpine:3.14
# Set the working directory
WORKDIR /app
# Copy necessary files from the build stage
COPY --from=builder /app/migrations/ ./migrations/
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
