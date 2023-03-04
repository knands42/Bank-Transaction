include app.env

############################### DOCKER ###############################
docker_down:
	docker-compose down

docker_up: docker_down
	docker-compose up --build --force-recreate

############################### MIGRATE ###############################
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


############################### SQLC ###############################
sqlc_query:
	sqlc generate

############################### MOCKGEN ###############################
DB_PATH = "src/infrastructure/db/sqlc/store.go"
DB_PATH_MOCKGEN = src/infrastructure/db/mock

mockgen:
	@for file in $(DB_PATH); do \
		mockgen -source=$$file -destination=$(DB_PATH_MOCKGEN)/`basename $$file` ; \
	done

############################### APP ###############################
app_run:
	go run src/main.go

app_build:
	go build -o main src/main.go

app_test:
	go test -v ./...

.PHONY: 