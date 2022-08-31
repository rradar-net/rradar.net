.PHONY: help dev proto

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Run development serverm
	docker compose up --build --force-recreate -d
	make proto
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
	air --build.cmd "go build -o tmp/bin/rradar" --build.bin "./tmp/bin/rradar"

proto: ## (Re)generate protobuf files
	protoc --proto_path=. --go_out=paths=source_relative:. pkg/proto/*.proto
	protoc-go-inject-tag -input="pkg/proto/*.pb.go"