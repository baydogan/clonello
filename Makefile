DB_URL=postgresql://root:secret@localhost:5432/users?sslmode=disable
POSTGRES_CONTAINER_NAME=postgres
POSTGRES_IMAGE=postgres:14-alpine
POSTGRES_NETWORK=clonello-network
POSTGRES_PORT=5432
POSTGRES_USER=root
POSTGRES_PASSWORD=secret
POSTGRES_DB=users

# Migration Path
MIGRATION_PATH=services/user-service/migrations

.PHONY: all help network start-postgres stop-postgres clean-postgres create-user-db dropdb migrateup-users migratedown-users

all: help

help:
	@echo "Available Commands:"
	@echo "  make network              -> Create Docker network"
	@echo "  make start-postgres       -> Start PostgreSQL container"
	@echo "  make stop-postgres        -> Stop PostgreSQL container"
	@echo "  make clean-postgres       -> Stop and remove PostgreSQL container"
	@echo "  make create-user-db       -> Create 'users' database"
	@echo "  make dropdb               -> Drop 'users' database"
	@echo "  make migrateup-users      -> Apply all migrations"
	@echo "  make migratedown-users    -> Roll back last migration"

network:
	@echo "Creating Docker network ($(POSTGRES_NETWORK)) if it doesn't exist..."
	docker network create $(POSTGRES_NETWORK) || true

start-postgres: network
	@echo "Starting PostgreSQL container ($(POSTGRES_CONTAINER_NAME))..."
	docker run -d \
		--name $(POSTGRES_CONTAINER_NAME) \
		--network $(POSTGRES_NETWORK) \
		-p $(POSTGRES_PORT):5432 \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		$(POSTGRES_IMAGE)

stop-postgres:
	@echo "Stopping PostgreSQL container ($(POSTGRES_CONTAINER_NAME))..."
	docker stop $(POSTGRES_CONTAINER_NAME) || true

clean-postgres: stop-postgres
	@echo "Removing PostgreSQL container ($(POSTGRES_CONTAINER_NAME))..."
	docker rm $(POSTGRES_CONTAINER_NAME) || true

create-user-db:
	@echo "Creating '$(POSTGRES_DB)' database..."
	docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)

dropdb:
	@echo "Dropping '$(POSTGRES_DB)' database..."
	docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb $(POSTGRES_DB) || true

migrateup-users:
	@echo "Applying all migrations from $(MIGRATION_PATH)..."
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

migratedown-users:
	@echo "Rolling back the last migration..."
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down
