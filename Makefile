include app.local.env

############################### Docker ###############################
docker_down:
	docker-compose down

docker_up: docker_down
	docker-compose up app --build --force-recreate

docker_up_deps_app: docker_down
	docker-compose up postgres

docker_up_deps_test: docker_down
	docker-compose up postgres-test

############################### Migrate ###############################
migration_create:
	migrate create -ext sql -dir app/internal/db/migrations -seq $(NAME)

migration_up:
	migrate -path app/internal/db/migrations -database "$(DB_SOURCE)" up

migrate_up1:
	migrate -path app/internal/db/migrations -database "$(DB_SOURCE)" -verbose up 1

migrate_down:
	migrate -path app/internal/db/migrations -database "$(DB_SOURCE)" -verbose down

migrate_down1:
	migrate -path app/internal/db/migrations -database "$(DB_SOURCE)" -verbose down 1


############################### Sqlc ###############################
sqlc_query:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

############################### Mockgen ###############################
DB_PATH = "app/internal/db/sqlc/store.go"
DB_AUX_PATH = "github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/db/sqlc=querier.go"
DB_PATH_MOCKGEN = app/internal/db/mock

mockgen:
	mockgen -source app/internal/db/sqlc/store.go -destination app/internal/db/mock/store.go -package mock_sqlc Store

############################### App ###############################
app_run:
	go run app/cmd/main.go

app_build:
	go build -o main app/cmd/main.go

app_test:
	go test -v ./...

############################### Phony ###############################
.PHONY:
	docker_down docker_up docker_up_deps migration_create migration_up migrate_up1 migrate_down migrate_down1 sqlc_query mockgen app_run app_build app_tests