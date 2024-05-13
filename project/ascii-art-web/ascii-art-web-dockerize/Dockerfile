# Use an official Go runtime as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Change directory to the location of your main.go file
WORKDIR /app/cmd

# Build the Go app
RUN go build -o myapp .

# Command to run the executable
CMD ["./myapp"]
