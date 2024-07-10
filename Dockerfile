# Use a base image with Golang
FROM golang:1.22.5

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Set the working directory to the cmd directory for building the main.go file
WORKDIR /app/cmd

# Build the Go app
RUN go build -o /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Define the default command to run the application
CMD ["/app/main"]
