CONFIG_FILE ?= ./configs/config.yaml
DSN ?= $(shell sed -n 's/^dsn:[[:space:]]\(.*\)/\1/p' $(CONFIG_FILE))
MIGRATE := migrate -database "$(DSN)" -path db/

run:
	go run ./app/api

migrate:
	@echo "Running database migrations."
	@$(MIGRATE) up