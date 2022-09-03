.PHONY: help dev proto

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Run development serverm
	docker compose up --build --force-recreate -d
	make proto
	./air-install.sh
	air --build.cmd "go build -o ./bin/rradar" --build.bin "./bin/rradar" --build.exclude_dir "rradar-ui"

proto: ## (Re)generate protobuf files
	protoc --proto_path=. --go_out=paths=source_relative:. pkg/proto/*.proto
	protoc-go-inject-tag -input="pkg/proto/*.pb.go"