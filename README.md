# Subscription Service

REST API для управления онлайн-подписками пользователей. Поддерживает CRUD операции, историю платежей и автоматическое продление подписок.

## Запуск

```bash
cp .env.example .env
task up
```

Swagger UI: `http://localhost:8080/swagger/index.html`

## Стек

Go, Gin, PostgreSQL, Docker

## Структура

```
cmd/app/          — точка входа
config/           — конфиг из env
internal/
  api/            — роутер, хендлеры, middleware
  service/        — бизнес-логика
  repository/     — работа с БД
  model/          — структуры данных
  worker/         — воркер автопродления
migrations/       — SQL-миграции
```

## Воркер

При старте запускается горутина, которая периодически проверяет подписки с истекающим сроком. Если `auto_renew = true` — продлевает подписку и создаёт запись в `payments`. Если нет — переводит в `expired`.

Интервал и порог настраиваются через `WORKER_INTERVAL_SECONDS` и `WORKER_RENEWAL_THRESHOLD_SECONDS` в `.env`.
