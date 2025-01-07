# Команды по работе с локальной базой данных
# Переменные для подключения к базе данных
PG_USER=user
PG_PASSWORD=password
PG_DB=test
PG_PORT=5432

# Запуск локальной базы данных
.PHONY: docker-up
docker-up:
	docker run --name pg-17.2 -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=test -d postgres:17.2

# Остановка локальной базы данных
.PHONY: docker-down
docker-down:
	docker stop pg-17.2

# Запуск миграций
.PHONY: g-up
g-up:
	goose -dir db/migrations postgres "postgresql://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_DB)?sslmode=disable" up

# Откат миграций
.PHONY: g-down
g-down:
	goose -dir db/migrations postgres "postgresql://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_DB)?sslmode=disable" down

# Статус миграций
.PHONY: g-status
g-status:
	goose -dir db/migrations postgres "postgresql://$(PG_USER):$(PG_PASSWORD)@localhost:$(PG_PORT)/$(PG_DB)?sslmode=disable" status
	