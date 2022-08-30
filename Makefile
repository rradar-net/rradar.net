.PHONY: help dev

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Run the development server
	make proto
	go run main.go

proto: ## (Re)generate protobuf files
	protoc --proto_path=. --go_out=paths=source_relative:. pkg/proto/*.proto
	protoc-go-inject-tag -input="pkg/proto/*.pb.go"