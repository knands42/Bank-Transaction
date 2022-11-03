#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "installing sqlc"
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
sqlc generate

echo "start the app"
exec "$@"