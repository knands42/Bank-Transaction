#!/bin/bash

if [ ! -f "app.env" ]; then
    touch app.env
fi

export TERM=xterm

# INSTALL DEPENDENCIES
apt update && apt upgrade -y
apt install make gcc g++ bash curl -y
ENV PATH="$PATH:/bin/bash"

# INSTALL MIGRATE BIN
echo "Installing migrate..."
cd /go/bin
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
cd -
# EXECUTE MIGRATE BIN
echo "Executing migrate..."
migrate -path src/infrastructure/db/migrations -database "postgresql://postgres:root@Database-Transactions-Simulation-db:5432/simple_bank?sslmode=disable" -verbose up

# INTALL SQLC
echo "Installing sqlc..."
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# EXECUTE SQLC
echo "Executing sqlc..."
sqlc generate
