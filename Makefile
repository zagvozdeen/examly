include .env
export

run:
	@go run cmd/api/main.go

migration:
	@migrate create -ext sql -dir ./migrations -seq $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@migrate -source file://migrations -database ${DB_ADDR} up

migrate-down:
	@migrate -source file://migrations -database ${DB_ADDR} down

build:
	@GOOS=linux GOARCH=amd64 go build -o examly ./cmd/api
	@cd resources && npm run build

deploy: build
	ssh root@185.221.214.4 "cd /var/www/examly.ru && systemctl stop examly.service && rm examly || true && rm -rf public && rm -rf migrations"
	scp examly root@185.221.214.4:/var/www/examly.ru
	scp .env root@185.221.214.4:/var/www/examly.ru
	scp Makefile root@185.221.214.4:/var/www/examly.ru
	scp -r resources/dist root@185.221.214.4:/var/www/examly.ru/public
	scp -r migrations root@185.221.214.4:/var/www/examly.ru/migrations
	scp .docker/services/deploy/nginx.conf root@185.221.214.4:/etc/nginx/sites-available/examly.ru
#	ssh root@185.221.214.4 "ln -s /etc/nginx/sites-available/examly.ru /etc/nginx/sites-enabled/examly.ru"
	scp .docker/services/deploy/examly.service root@185.221.214.4:/etc/systemd/system
	ssh root@185.221.214.4 "systemctl daemon-reload && systemctl restart examly.service && make migrate-up"

