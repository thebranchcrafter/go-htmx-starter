# Define variables
DOCKER_COMPOSE_FILE=docker-compose.yml
TEMPL_CMD=$(shell command -v templ 2>/dev/null)

# Define color codes
GREEN=\033[32m
BLUE=\033[34m
CYAN=\033[36m
YELLOW=\033[33m
RESET=\033[0m

# Default target
.PHONY: help
help:
	@echo -e "$(CYAN)Available commands:$(RESET)"
	@awk 'BEGIN {FS = ":.*?#"} /^[a-zA-Z_-]+:.*?#/ {printf "  $(GREEN)%-15s$(RESET) $(YELLOW)%s$(RESET)\n", $$1, $$2}' $(MAKEFILE_LIST)

# Docker Compose commands
.PHONY: up down start stop logs ps

up:                # Start Docker Compose services
	@echo -e "$(CYAN)Starting Docker Compose services...$(RESET)"
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

down:              # Stop and remove Docker Compose services
	@echo -e "$(CYAN)Stopping and removing Docker Compose services...$(RESET)"
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

start:             # Start existing Docker Compose services
	@echo -e "$(CYAN)Starting existing Docker Compose services...$(RESET)"
	docker-compose -f $(DOCKER_COMPOSE_FILE) start

stop:              # Stop Docker Compose services
	@echo -e "$(CYAN)Stopping Docker Compose services...$(RESET)"
	docker-compose -f $(DOCKER_COMPOSE_FILE) stop

logs:              # Tail logs for Docker Compose services
	@echo -e "$(CYAN)Tailing logs for Docker Compose services...$(RESET)"
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

ps:                # List Docker Compose services
	@echo -e "$(CYAN)Listing Docker Compose services...$(RESET)"
	docker-compose -f $(DOCKER_COMPOSE_FILE) ps

# Templ setup and server commands
.PHONY: templ-install templ-run

templ-install:     # Install templ if not already installed
	@if [ -z "$(TEMPL_CMD)" ]; then \
		echo -e "$(CYAN)Installing templ...$(RESET)"; \
		go install github.com/a-h/templ/cmd/templ@latest; \
	else \
		echo -e "$(GREEN)templ is already installed.$(RESET)"; \
	fi

templ-run: templ-install # Run templ generate with watch and proxy
	@echo -e "$(CYAN)Running templ generate with watch and proxy...$(RESET)"
	templ generate --watch --proxy="http://localhost:8081" --cmd="go run cmd/web/app/main.go"
