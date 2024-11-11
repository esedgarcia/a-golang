# Use an official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR C:\Users\esedg\OneDrive\Documents\Code\containers-assignment\a-golang

# Copy the Go file into the container
COPY greetings.go .

# Build the Go program
RUN go build -o greetings greetings.go

# Set the command to run the compiled executable
CMD ["./greetings"]
