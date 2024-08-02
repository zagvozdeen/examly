run:
	@go run cmd/api/main.go

up:
	@docker compose up -d

down:
	@docker compose down

restart r:
	@down
	@up

migration:
	@migrate create -ext sql -dir ./migrations -seq $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

