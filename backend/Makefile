# Load environment variables from app.env
include app.env
export $(shell sed 's/=.*//' app.env)

# Run SQLC code generation
.PHONY: sqlc
sqlc:
	sqlc generate

# Make Postgresql container in Docker
.PHONY: postgres
postgres:
	docker run \
		--name $(DB_NAME) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(DB_PORT):5432 \
		postgres:latest

# Generate Protobuf files
.PHONY: proto
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=munshiji \
	--experimental_allow_proto3_optional \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc -f

# Run all migrations (up)
.PHONY: migrateup
migrateup:
	migrate -path ./db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" -verbose up

# Rollback all migrations (down)
.PHONY: migratedown
migratedown:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" -verbose down

# Run Evans gRPC client
.PHONY: evans
evans:
	evans --proto service_munshiji.proto --path proto --host localhost --port 9000

# Run the application
.PHONY: run
run:
	go run main.go