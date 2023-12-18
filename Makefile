SHELL := /bin/bash
BINARY_NAME = dancing-pony
DC = docker-compose
MIGRATION_DIR = internal/platform/migration/files
MIGRATE = migrate -database ${DB_SOURCE} -path ${MIGRATION_DIR}
MODULE = $(shell go list -m)
PACKAGES := $(shell go list ./... | grep -v /vendor/)

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

build: ## Build the binary
	@echo "Building Dancing Pony..."
	@go build -o ./${BINARY_NAME} .
	@echo "Dancing Pony built!"

test: ## Test the project
	@echo "Testing..."
	@go test -v -cover -short ./...
	@echo "Done!"

start: build ## Build binary & start Api server
	@echo "Starting Dancing Pony..."
	@./${BINARY_NAME}

start-dev: ## Start application with Hot reload
	@CompileDaemon --build="make build" --command="./${BINARY_NAME}"

run: ## Start Api server
	@go run ./main.go

lint: ## Run the linter (perform static analysis)
	golangci-lint run ./...

fmt: ## Apply code formatting
	@go fmt $(PACKAGES)

##@ Migration

migrate-new: ## Creating migration files
	@read -p "Enter migration name: " Mname; \
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq "$$Mname"

migrate: ## Apply all up migrations
	$(MIGRATE) up

migrate-down: ## Apply all down migrations
	$(MIGRATE) down

migrate-drop: ## Drop everything inside database
	$(MIGRATE) drop

migrate-force: ## Set version but don't run migration (ignores dirty state)
	@read -p "Specify version: " Mversion; \
	$(MIGRATE) force "$$Mversion"

migrate-rollback: ## Migration rollback to version V
	@read -p "Specify version: " Mversion; \
	$(MIGRATE) goto "$$Mversion"

migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	$(MIGRATE) drop
	@echo "Running all database migrations..."
	$(MIGRATE) up

migrate-v: ## Print current migration version
	$(MIGRATE) version

##@ Docker

dc-build: ## Rebuild docker images for container
	$(DC) build

dc-up-build: ## Rebuild docker images for container & start application
	$(DC) up -d --build

dc-up: ## Start containers applications
	$(DC) up -d && $(DC) logs -f app

dc-logs: ## Show all container logs
	$(DC) logs -f

dc-app-logs: ## Show container app logs
	$(DC) logs -f app

dc-stop: ## Stop container services
	$(DC) stop

dc-down: ## Stop and remove containers
	$(DC) down
