include app.env

execute:
	docker compose up --build --force-recreate

create_migration:
	migrate create -ext sql -dir src/infrastructure/db/migrations -seq $(file_name)

migrate_up:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose up

migrate_up1:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose up 1

migrate_down:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose down

migrate_down1:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose down 1

query:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run src/main.go

mock:
	mockgen -destination=src/infrastructure/db/mock/store.go -package=mockdb github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/db/sqlc Store

.PHONY: execute create_migration migrate_up migrate_down query test server mock