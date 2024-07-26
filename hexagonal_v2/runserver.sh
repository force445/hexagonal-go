#!/bin/sh
while ! nc -z ${DB_HOST} ${DB_PORT}; do
    echo "Waiting for the Postgres database to start...";
    sleep 1;
done;

# Run the Go application directly
cd cmd/http/
dogo -c dogo.json
