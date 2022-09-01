execute:
	docker compose up --build --force-recreate

create_migration:
	migrate create -ext sql -dir src/infrastructure/db/migrations -seq example_schema

migrate_up:
	migrate -path src/infrastructure/db/migrations -database "$(POSTGRES_URL)" -verbose up

migrate_down:
	migrate -path src/infrastructure/db/migrations -database "$(POSTGRES_URL)" -verbose down

generate_query:
	docker run --rm -v "${pwd}:/src" -w /src kjconroy/sqlc generate

test:
	go test -v ./...

.PHONY: execute create_migration migrate_up migrate_down generate_query test