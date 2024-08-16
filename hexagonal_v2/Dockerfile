#FOR DEV
# FROM golang:1.20.3-alpine

# RUN go install github.com/liudng/dogo@latest

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go module files
# COPY go.mod .
# COPY go.sum .

# # Download dependencies
# RUN go mod download
# RUN go mod tidy

# # Copy the rest of the application code
# COPY . .

# # Copy the dogo.json file
# COPY cmd/http/dogo.json cmd/http/dogo.json

# # Make runserver.sh executable
# RUN chmod +x runserver.sh

# # Command to run the application
# CMD ["sh", "runserver.sh"]

#FOR DEPLOY
FROM golang:1.20.3-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download
RUN go mod tidy

# Copy the rest of the application code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/main cmd/http/main.go

FROM ubuntu:20.04

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/main /app/bin/main 
COPY --from=builder /app/runserver.sh /app

RUN apt-get update && apt-get install -y libc6

RUN apt-get install -y netcat

RUN chmod +x /app/runserver.sh

EXPOSE 8080

# Use the runserver.sh script as the entry point
CMD ["/app/runserver.sh"]