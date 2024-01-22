#!/bin/sh

# We use set -e instruction to make sure that the script will exit immediately if a command returns a non-zero status.
set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

# After running migrate up, we will start the app.
echo "start the app"

# It basically means: takes all parameters  passed to the script and run it.
exec "$@"