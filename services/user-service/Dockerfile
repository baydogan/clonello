# Base image
FROM golang:1.23.3-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker's caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy && go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o user-service cmd/main.go

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./user-service"]
