include app.env

execute:
	docker compose up --build --force-recreate

create_migration:
	migrate create -ext sql -dir src/infrastructure/db/migrations -seq example_schema

migrate_up:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose up

migrate_down:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose down

generate_query:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run src/main.go

.PHONY: execute create_migration migrate_up migrate_down generate_query test server