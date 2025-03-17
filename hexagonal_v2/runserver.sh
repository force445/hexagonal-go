#!/bin/sh
#load config from .env
while ! nc -z ${POSTGRES_HOST} ${POSTGRES_PORT}; do
    echo "Waiting for the Postgres database to start...";
    sleep 1;
done;

# Run the Go FOR DEVELOPMENT ONLY
# cd cmd/http/
# dogo -c dogo.json

exec /app/bin/main
