FROM golang:1.20.3-alpine

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

# Make runserver.sh executable
RUN chmod +x runserver.sh

# Command to run the application
CMD ["sh", "runserver.sh"]