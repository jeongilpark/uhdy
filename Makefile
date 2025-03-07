.PHONY: all apply-openapi

# Default target that calls 'all'
all: tidy apply-openapi

.PHONY: tidy sqlc

# Helper function to execute a command in service directories
define SERVICE_COMMAND
	@if [ -z "$(service)" ]; then \
		for dir in services/*/ ; do \
			(if [ -d "$$dir" ]; then \
				(cd "$$dir" && $(1)); \
			fi); \
		done; \
	else \
		(cd "services/$(service)" && $(1)); \
	fi
endef

# The tidy target
tidy:
	$(call SERVICE_COMMAND, go mod tidy)

# The sqlc target
sqlc:
	$(call SERVICE_COMMAND, sqlc generate -f sqlc/sqlc.yaml)

#
# Docker compose
#
.PHONY: build up down
build:
	@if [ -z "$(service)" ]; then \
		for dir in services/* ; do \
			(if [ -d "$$dir" ] && [ "$$dir" != "services/utils" ]; then \
				docker build -t "docker.io/uhdy/$(notdir "$$dir")-service" -f "$$dir/Dockerfile" services; \
			fi); \
		done; \
	else \
		if [ -d "services/${service}" ]; then \
			docker build -t "docker.io/uhdy/${service}-service" -f "services/${service}/Dockerfile" services; \
		fi; \
	fi

up:
	@$(MAKE) build service=${service}
	@if [ -z "$(service)" ]; then \
		docker compose up -d; \
	else \
		docker compose up $(service) -d; \
	fi

down:
	docker compose down

#
# Scripts
#

# Apply Database DDL
apply_ddl: up_postgres
	./scripts/apply_ddl.sh ./services/*/sqlc/schema.sql

up_postgres:
	@$(MAKE) up service=postgres
