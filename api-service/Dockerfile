# Dockerfile for Go service
FROM golang:1.23-alpine

WORKDIR /app

# Copy all of the Go service source code
COPY . .

# Install the Go dependencies
RUN go mod tidy

# Build the Go service
RUN go build -o main

# Expose port 4444
EXPOSE 4444

# Command to run the Go service
CMD ["./main"]
