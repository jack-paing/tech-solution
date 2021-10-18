CONFIG_FILE ?= ./configs/config.yaml
DSN ?= $(shell sed -n 's/^dsn:[[:space:]]\(.*\)/\1/p' $(CONFIG_FILE))
MIGRATE := migrate -database mysql://"$(DSN)" -path db/
TEST_DSN ?= $(shell sed -n 's/^test_dsn:[[:space:]]\(.*\)/\1/p' $(CONFIG_FILE))
TEST_MIGRATE := migrate -database mysql://"$(TEST_DSN)" -path db/
UNIT_TEST_DIR=$(shell go list ./... | grep -v /app | grep -v internal/config | grep -v internal/health)

run:
	go run ./app/api

migrate:
	@echo "Running database migrations."
	@$(MIGRATE) up

migrate-down:
	@$(MIGRATE) down

mock:
	go get github.com/vektra/mockery/v2/.../
	cd internal/repo/;\
		mockery --name=CardRepo --filename=card.go;
	cd internal/card/; \
		mockery --name=Service --filename=service_mock.go --structname=ServiceMock --inpackage

setup-test-db:
	@$(TEST_MIGRATE) up

unit-test:
	go get gotest.tools/gotestsum
	gotestsum --jsonfile report.json fmt -race -v $(UNIT_TEST_DIR)