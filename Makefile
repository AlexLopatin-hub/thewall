include .env
export

export PROJECT_ROOT=$(CURDIR)

env-up:
	@docker compose up -d thewall-postgres

env-down: env-port-close
	@docker compose down thewall-postgres

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@migrate create -ext sql -dir ./migrations -seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@migrate \
		-path ./migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

thewall-run:
	@docker compose up thewall-app

thewall-stop:
	@docker compose down thewall-app thewall-postgres
