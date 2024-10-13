#!/usr/bin/env sh

set -e

until psql postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB} -c '\q'; do
  echo "Postgres is unavailable - waiting..."
  sleep 1
done

echo "Postgres is up - begin migrating"
sql-migrate up -env="production"
