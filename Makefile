include .env

migration-create:
	@echo "Creating migration..."
	goose -dir db/postgres create $(name) sql