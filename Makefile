.PHONY: help dev proto ent test-integration

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Run development server
	docker compose -p rradar-dev -f docker/dev/docker-compose.yml up --force-recreate --build

proto: ## (Re)generate protobuf files
	rm -rf pkg/proto/*.pb.go
	protoc --proto_path=./pkg/proto --go_out=paths=source_relative:./pkg/proto $(shell cd pkg/proto && ls -d *.proto)
	protoc-go-inject-tag -input="pkg/proto/*.pb.go"

ent: ## (Re)generate ent assets
	go generate ./ent
