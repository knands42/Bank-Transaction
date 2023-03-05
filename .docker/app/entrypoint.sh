#!/bin/sh

touch app.env
export TERM=xterm

###################################### Install migrate bin ######################################
echo "Installing migrate..."
cd /go/bin
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
cd -
###################################### Execute migrate bin ######################################
echo "Executing migrate..."
migrate -path src/infrastructure/db/migrations -database "$DB_SOURCE" -verbose up

###################################### Install SQLC ######################################
echo "Installing sqlc..."
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

###################################### Running the application ######################################
/app/main