.PHONY: help dev proto ent integration integration-down

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Run development server
	docker compose -p rradar-dev -f docker/dev/docker-compose.yml up --force-recreate --build

integration: ## Run integration tests
	docker container rm -f rradar-integration-postgres
	docker volume rm -f rradar-integration_pgdata
	docker compose -p rradar-integration -f docker/integration/docker-compose.yml up --attach runner --force-recreate --build --abort-on-container-exit

integration-down: ## Remove integration tests containers and volumes
	docker compose -p rradar-integration -f docker/integration/docker-compose.yml down -v

proto: ## (Re)generate protobuf files
	rm -rf pkg/proto/*.pb.go
	protoc --proto_path=./pkg/proto --go_out=paths=source_relative:./pkg/proto $(shell cd pkg/proto && ls -d *.proto)
	protoc-go-inject-tag -input="pkg/proto/*.pb.go"

ent: ## (Re)generate ent assets
	go generate ./ent
