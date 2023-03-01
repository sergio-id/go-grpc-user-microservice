# ==============================================================================
# Main

run:
	go run cmd/main.go -config=./config/config.yaml

test:
	go test -cover ./...


# ==============================================================================
# Modules support

tidy:
	go mod tidy

# ==============================================================================
# Linters https://golangci-lint.run/usage/install/

run-linter:
	@echo Starting linters
	golangci-lint run ./...


# ==============================================================================
# Go migrate postgresql

force:
	migrate -database postgres://postgres:postgres@localhost:5432/user?sslmode=disable -path migrations force 1

version:
	migrate -database postgres://postgres:postgres@localhost:5432/user?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5432/user?sslmode=disable -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5432/user?sslmode=disable -path migrations down 1


# ==============================================================================
# Docker compose commands

local_up:
	@echo Starting local docker compose without user service container
	docker-compose -f docker-compose-local.yml up --build

local_down:
	docker-compose -f docker-compose-local.yml down

docker_up:
	@echo Starting local docker compose with .env all containers
	docker-compose -f docker-compose.yml --env-file=./.env up --build

docker_down:
	docker-compose -f docker-compose.yml down


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)


# ==============================================================================
# Proto

proto_user:
	@echo Generating gRPC proto user_service
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto
