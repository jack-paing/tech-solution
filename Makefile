DSN := mysql://root:yourpasswd@tcp(127.0.0.1:3306)/user_service?multiStatements=true
MIGRATE := migrate -database "$(DSN)" -path db/

run:
	go run ./

migrate:
	@echo "Running database migrations."
	@$(MIGRATE) up