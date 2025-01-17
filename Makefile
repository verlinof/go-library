# Load .env file
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# Variables
REPO_URL = https://github.com/verlinof/softlancer-go.git
BUILD_DIR = ./build
CMD_DIR = ./cmd
ENV_FILE = .env

# dsn
DSN = ${DB_DRIVER}://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

# Environment Configurations
ENV = development
GO_ENV = $(ENV)

# Default target
.PHONY: all
all: run

.PHONY: migrate-up
migrate-up:
	@GOOSE_DRIVER=${DB_DRIVER} GOOSE_DBSTRING=${DSN} goose --dir=./db/migration up

# Run Seeders
.PHONY: seed
seed:
	go run $(CMD_DIR)/seeders/main.go

# Run the Application in development mode
.PHONY: run
run:
	go run $(CMD_DIR)/api/main.go

# Run with Air for hot reloading (development only)
.PHONY: air
air:
	air

# Build the Application
.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	GO_ENV=production go build -o $(BUILD_DIR)/app $(CMD_DIR)/api/main.go

# Set environment to release and run in production mode
.PHONY: release
release:
	@echo "Setting environment to release mode"
	$(eval GO_ENV := release)
	$(eval ENV := release)
	@echo "Environment set to: $(ENV)"
	go run $(CMD_DIR)/api/main.go
