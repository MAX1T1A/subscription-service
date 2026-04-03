#!/bin/sh
set -e

echo "Running migrations..."
migrate -path /app/migrations -database "$DATABASE_URL" up
echo "Migrations done."

exec air
