# Processing Bank Transfers

Учебный backend-проект на Go для обработки банковских переводов.

## Текущий этап

Реализован минимальный каркас проекта:

- точка входа API (`cmd/api/main.go`) с health-check `GET /health`;
- доменные модели `BankAccount` и `Transaction`;
- базовые интерфейсы `Repository` и `Service` для дальнейшей реализации бизнес-логики.

## Запуск

```bash
go run ./cmd/api
```

Проверка health-check:

```bash
curl http://localhost:8080/health
```
