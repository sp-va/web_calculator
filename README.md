# Краткое описание

Веб-калькулятор для вычисления арифметических выражений с операторами сложения, вычитания, умножения, деления и скобками.

# Структура проекта
```
web_calculator/
│
├── cmd/
│   └── calculator/
│       └── main.go                 // Точка входа приложения
├── internal/
│   ├── middleware/
│   │   └── middleware.go           // Проверка корректности запросов
│   ├── routes/
│   │   └── routes.go               // Обработчики
│   ├── service/
│   │   └── calculator.go           // Калькулятор
│   ├── types/
│   │   └── types.go                // Типы, которые используются в нескольких файлах
│
├── go.mod                          // Модуль Go
└── README.md                       // Информация о проекте
```

# Использование проекта

## Запуск

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/sp-va/web_calculator.git
   cd web_calculator

2. Запустить сервер
    ```bash
    go run ./cmd/calculator/main.go

3. Отправить запросы на `http://localhost:8080/api/v1/calculate`

## Примеры запросов
На Linux:
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "-1+(2*3/4)*5"
}'
```
На Windows:
```bash
curl -Uri 'http://localhost:8080/api/v1/calculate' `
     -Method POST `
     -ContentType 'application/json' `
     -Body '{"expression": "-1+(2*3/4)*5"}'
```

## Примеры ответа:

### Успех 200
```bash
{"result":"6.5"}
```

### Ошибка 400
```bash
{"error": "wrong request formatting"}
{"error": "cant read request body"}
{"error": "json invalid"}
```

### Ошибка 405
```bash
{"error": "method not allowed"}
```

### Ошибка 422
```bash
{"error": "expression wrong"}
{"error": "division by zero"}
```

### Ошибка 500
```bash
{"error": "server error"}
```
