include .env
export

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
	@migrate -source file://migrations -database ${DB_ADDR} up

migrate-down:
	@migrate -source file://migrations -database ${DB_ADDR} down

bash:
	@docker compose exec node /bin/sh

build:
	@GOOS=linux GOARCH=amd64 go build -o examly ./cmd/api
	@cd resources && npm run build

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

deploy: build
deploy:
	ssh root@185.221.214.4 "cd /var/www/examly.ru && systemctl stop examly.service && rm examly || true && rm -rf public"
	scp examly root@185.221.214.4:/var/www/examly.ru
	scp .env root@185.221.214.4:/var/www/examly.ru
	scp Makefile root@185.221.214.4:/var/www/examly.ru
	scp -r resources/dist root@185.221.214.4:/var/www/examly.ru/public
	scp -r migrations root@185.221.214.4:/var/www/examly.ru/migrations
	scp .docker/services/deploy/nginx.conf root@185.221.214.4:/etc/nginx/sites-available/examly.ru
#	ssh root@185.221.214.4 "ln -s /etc/nginx/sites-available/examly.ru /etc/nginx/sites-enabled/examly.ru"
	scp .docker/services/deploy/examly.service root@185.221.214.4:/etc/systemd/system
	ssh root@185.221.214.4 "systemctl daemon-reload && systemctl restart examly.service && make migrate-up"

