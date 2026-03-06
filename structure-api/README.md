# Org Structure API

API для управления организационной структурой компании.

## Технологии

- Go
- net/http
- GORM
- PostgreSQL
- goose migrations
- Docker / docker-compose

## Запуск

1. Клонировать репозиторий

2. Запустить

docker-compose up --build

3. Применить миграции

goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/orgdb?sslmode=disable" up

API будет доступен:

http://localhost:8080

## Основные endpoints

POST /departments  
POST /departments/{id}/employees  
GET /departments/{id}