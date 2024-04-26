# Use an official Golang runtime as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go get -u ./...

# Build the Go application
RUN go build -o app ./cmd

# Expose the port the application runs on
EXPOSE 8080

# Define the command to run the application
CMD ["./app"]
