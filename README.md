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

## Тестирование
1 (проверка данных миграции)
docker exec -it structure-api-db-1 psql -U postgres -d orgdb
SELECT * FROM departments;
SELECT * FROM employees;

<img width="1767" height="504" alt="image" src="https://github.com/user-attachments/assets/829ad417-5b45-40ad-a625-e9a4da4dff21" />

Примеры запросов для тестирования( не работает postman, сделал через powershell)

Примеры тестовых запросов для README.md
1️ Создать корневой департамент

Invoke-RestMethod -Uri "http://localhost:8080/departments" `
    -Method POST `
    -ContentType "application/json" `
    -Body '{"name":"IT_Main","parent_id":null}'

    
2️ Создать поддепартамент

Invoke-RestMethod -Uri "http://localhost:8080/departments" `
    -Method POST `
    -ContentType "application/json" `
    -Body '{"name":"Backend_Team","parent_id":1}'
    
3️ Создать сотрудника в корневом департаменте

Invoke-RestMethod -Uri "http://localhost:8080/departments/1/employees" `

    -Method POST `
    -ContentType "application/json" `
    -Body '{"full_name":"Ivan Ivanov","position":"Developer","hired_at":"2024-01-01"}'
4️ Создать сотрудника в поддепартаменте

Invoke-RestMethod -Uri "http://localhost:8080/departments/2/employees" `
    -Method POST `
    -ContentType "application/json" `
    -Body '{"full_name":"Olga Smirnova","position":"Backend Developer","hired_at":"2025-06-15"}'
5️ Получить департамент с поддеревом и сотрудниками

Invoke-RestMethod -Uri "http://localhost:8080/departments/1?depth=2&include_employees=true" `
    -Method GET
    
6️ Обновить департамент (изменить имя или родителя)

Invoke-RestMethod -Uri "http://localhost:8080/departments/2" `
    -Method PATCH `
    -ContentType "application/json" `
    -Body '{"name":"Backend_Updated","parent_id":1}'
    
7️ Удалить поддепартамент с каскадом

Invoke-RestMethod -Uri "http://localhost:8080/departments/2?mode=cascade" `
    -Method DELETE
    
8️ Удалить департамент с перераспределением сотрудников

Invoke-RestMethod -Uri "http://localhost:8080/departments/3?mode=reassign&reassign_to_department_id=1" `
    -Method DELETE
