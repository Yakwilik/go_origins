.PHONY: compose-build compose-up compose-log build-up compose-down migrate

sess-db=sessions-db

all:  build-up migrate runserver

runserver:
	go run cmd/redditclone/main.go

compose-build:
	docker-compose build

restart:
	docker compose restart --timeout=1

compose-up:
	docker-compose up -d

compose-log:
	docker-compose logs -f
build-up:
	docker-compose up -d --build

compose-down:
	docker-compose down -v

#"root:love@tcp(localhost:3306)/golang?"
#mysql://user:password@tcp(host:port)/dbname?query
migrate:
	make migrate-sess up
	make migrate-users up

migrate-sess:
	migrate -path ./schema/sessions -database 'mysql://service_admin:qwerty@tcp(localhost:3301)/redditclone'  $(filter-out $@,$(MAKECMDGOALS))

migrate-users:
	migrate -path ./schema/users -database 'mysql://service_admin:qwerty@tcp(localhost:3302)/redditclone' $(filter-out $@,$(MAKECMDGOALS))

connect-sess-db:
	docker-compose exec $(sess-db) mysql -u service_admin -p redditclone

get:
	go get "github.com/$(filter-out $@,$(MAKECMDGOALS))"


%:
	@: