#!/bin/bash

if [ ! -f "app.env" ]; then
    touch app.env
fi

export TERM=xterm

# INSTALL MIGRATE BIN
echo "Installing migrate..."
cd /go/bin
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
cd -
# EXECUTE MIGRATE BIN
echo "Executing migrate..."
migrate -path src/infrastructure/db/migrations -database "$DB_SOURCE" -verbose up

# INTALL SQLC
echo "Installing sqlc..."
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# EXECUTE SQLC
echo "Executing sqlc..."
sqlc generate

top -b