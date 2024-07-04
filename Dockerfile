# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o cmd/main

# Expose the port that the application will listen on
EXPOSE 8080

# Run the application
CMD ["./cmd/main"]