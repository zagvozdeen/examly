run:
	@go run cmd/api/main.go

up:
	@docker compose up -d
	@docker compose exec -d node npm run dev

down:
	@docker compose down

restart r:
	@make down
	@make up

migration:
	@migrate create -ext sql -dir ./migrations -seq $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

bash:
	@docker compose exec node /bin/sh

build:
	@docker compose exec node npm run build

check:
	@docker compose exec node npm run check
