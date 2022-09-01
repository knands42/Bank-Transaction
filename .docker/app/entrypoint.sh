#!/bin/bash

if [ ! -f ".env" ]; then
    cp .env.example .env
fi

export TERM=xterm

# INSTALL MIGRATE BIN
echo "Installing migrate..."
mkdir /gobin
cd /go/bin
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
cd -
# EXECUTE MIGRATE BIN
echo "Executing migrate..."
migrate -path src/infrastructure/db/migrations -database "postgresql://postgres:root@bank-transactions-simulations-db:5432/simple_bank?sslmode=disable" -verbose up

# INTALL SQLC
echo "Installing sqlc..."
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# EXECUTE SQLC
echo "Executing sqlc..."
sqlc generate

top