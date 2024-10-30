# Use the official Golang image as the base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the service will run on
EXPOSE 8082

# Run the application
CMD ["./main"]

