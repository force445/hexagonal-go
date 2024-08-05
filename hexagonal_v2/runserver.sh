#!/bin/sh
#load config from .env
export $(cat .env | xargs)

echo "DB_HOST: ${DB_HOST}"
echo "DB_PORT: ${DB_PORT}"

while ! nc -z ${DB_HOST} ${DB_PORT}; do
    echo "Waiting for the Postgres database to start...";
    sleep 1;
done;

# Run the Go FOR DEVELOPMENT ONLY
# cd cmd/http/
# dogo -c dogo.json

exec /app/bin/main
