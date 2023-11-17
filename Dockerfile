# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container
COPY . .

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 5000

# Command to run the executable
CMD ["./main"]
