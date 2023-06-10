include app.local.env

############################### Docker ###############################
docker_down:
	docker-compose down

docker_up: docker_down
	docker-compose up --build --force-recreate

docker_up_deps_app: docker_down
	docker-compose up postgres

docker_up_deps_test: docker_down
	docker-compose up postgres-test

############################### Migrate ###############################
migration_create:
	migrate create -ext sql -dir src/infrastructure/db/migrations -seq $(NAME)

migration_up:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" up

migrate_up1:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose up 1

migrate_down:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose down

migrate_down1:
	migrate -path src/infrastructure/db/migrations -database "$(DB_SOURCE)" -verbose down 1


############################### Sqlc ###############################
sqlc_query:
	sqlc generate

############################### Mockgen ###############################
DB_PATH = "src/infrastructure/db/sqlc/store.go"
DB_PATH_MOCKGEN = src/infrastructure/db/mock

mockgen:
	mockgen -source=$(DB_PATH) -destination=$(DB_PATH_MOCKGEN)/`basename $(DB_PATH)`

############################### App ###############################
app_run:
	go run src/main.go

app_build:
	go build -o main src/main.go

app_test:
	go test -v ./...

############################### Phony ###############################
.PHONY:
	docker_down docker_up docker_up_deps migration_create migration_up migrate_up1 migrate_down migrate_down1 sqlc_query mockgen app_run app_build app_tests