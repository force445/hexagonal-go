FROM golang:1.20.3 AS builder

WORKDIR go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./bin