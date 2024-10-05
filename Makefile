include .env
export

run:
	@go run cmd/api/*.go

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
	@migrate -source file://migrations -database ${DB_ADDR} up

migrate-down:
	@migrate -source file://migrations -database ${DB_ADDR} down

bash:
	@docker compose exec node /bin/sh

build:
	@docker compose exec node npm run build

check:
	@docker compose exec node npm run check

prod-up:
	@docker compose --file compose.prod.yaml up --build -d
prod-down:
	@docker compose --file compose.prod.yaml down
prod-restart prod-r:
	@make prod-down
	@make prod-up
prod-build:
	@docker compose exec node npm run build
	@rm -rf ./public/dist || true
	@rm ./public/index.html || true
	@mv ./resources/dist ./public
	@mv ./public/dist/index.html ./public/index.html
prod-migrate-up:
	@docker compose exec migrate ./migrate up
prod-migrate-down:
	@docker compose exec migrate ./migrate down
