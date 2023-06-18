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
mockgen:
	mockgen -source app/internal/db/sqlc/store.go -destination app/internal/db/mock/store.go

############################### App ###############################
app_run:
	go run app/cmd/main.go

app_build:
	go build -o main app/cmd/main.go

app_test:
	go test -v ./...

############################### Proto ###############################
proto:
	protoc --proto_path=app/pkg/proto --go_out=app/pkg/proto/pb --go_opt=paths=source_relative \
 	--go-grpc_out=app/pkg/proto/pb --go-grpc_opt=paths=source_relative \
 	app/pkg/proto/*.proto

############################### Phony ###############################
.PHONY:
	docker_down docker_up docker_up_deps migration_create migration_up migrate_up1 migrate_down migrate_down1 sqlc_query mockgen app_run app_build app_tests proto