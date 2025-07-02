#!/bin/sh
set -e

host="$1"
port="$2"
shift 2
cmd="$@"

echo "Waiting for PostgreSQL at $host:$port..."
until PGPASSWORD="$DB_PASSWORD" psql -h "$host" -p "$port" -U "$DB_USER" -d "$DB_NAME" -c '\\q' > /dev/null 2>&1; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "Postgres is up - executing command"
exec $cmd
