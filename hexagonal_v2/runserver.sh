#!/bin/sh
while ! nc -z database 5432; do 
    echo "Waiting for the Postgres database to start...";
    sleep 1;
done;

# Run the Go application directly
go run cmd/http/main.go