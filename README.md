# Subscription Service

REST API для агрегации данных об онлайн-подписках пользователей с автоматическим продлением.

## Стек

- **Go** + **Gin** — HTTP-сервер
- **PostgreSQL** — база данных
- **golang-migrate** — миграции (запускаются через `entrypoint.sh` перед стартом приложения)
- **swaggo/swag** — генерация Swagger-документации из аннотаций
- **Air** — hot-reload в dev-режиме
- **Docker Compose** + **Taskfile** — оркестрация

---

## Запуск

### Требования

- [Docker](https://docs.docker.com/get-docker/)
- [Task](https://taskfile.dev/installation/)

### 1. Создать `.env`

```bash
cp .env.example .env
```

### 2. Запустить

**Dev-режим** (hot-reload при изменении файлов, логи в терминале):
```bash
task dev
```

**Production:**
```bash
task up
```

---

## Команды

| Команда | Описание |
|---|---|
| `task dev` | Dev-режим с hot-reload |
| `task dev-down` | Остановить dev |
| `task up` | Запустить в фоне |
| `task down` | Остановить |
| `task migrate-up` | Применить миграции вручную |
| `task migrate-down` | Откатить последнюю миграцию |
| `task migrate-create -- name` | Создать новую миграцию |

---

## API

Базовый URL: `http://localhost:8080/api/v1`

Swagger UI: `http://localhost:8080/swagger/index.html`

### Эндпоинты

| Метод | Путь | Описание |
|---|---|---|
| `POST` | `/subscriptions` | Создать подписку |
| `GET` | `/subscriptions` | Список подписок |
| `GET` | `/subscriptions/:id` | Получить подписку по ID |
| `PUT` | `/subscriptions/:id` | Обновить подписку |
| `DELETE` | `/subscriptions/:id` | Удалить подписку |
| `GET` | `/subscriptions/cost` | Суммарная стоимость за период |
| `GET` | `/payments?subscription_id=` | История платежей по подписке |

### Пример создания подписки

```bash
curl -X POST http://localhost:8080/api/v1/subscriptions \
  -H "Content-Type: application/json" \
  -d '{
    "service_name": "Yandex Plus",
    "price": 400,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "start_date": "07-2025",
    "end_date": "08-2025",
    "auto_renew": true
  }'
```

### Query-параметры

**GET /subscriptions**

| Параметр | Тип | Описание |
|---|---|---|
| `user_id` | UUID | Фильтр по пользователю |
| `service_name` | string | Фильтр по названию сервиса |
| `status` | `active` / `expired` | Фильтр по статусу |

**GET /subscriptions/cost**

| Параметр | Тип | Описание |
|---|---|---|
| `user_id` | UUID | Фильтр по пользователю |
| `service_name` | string | Фильтр по названию сервиса |
| `start_period` | MM-YYYY | Начало периода |
| `end_period` | MM-YYYY | Конец периода |

---

## Воркер автопродления

При старте запускается горутина, которая с интервалом `WORKER_INTERVAL_SECONDS` проверяет активные подписки. Если до `end_date` осталось меньше `WORKER_RENEWAL_THRESHOLD_SECONDS` секунд:

- `auto_renew = true` — создаётся запись в `payments`, подписка продлевается на исходный срок
- `auto_renew = false` — статус меняется на `expired`

Для демонстрации рекомендуется установить в `.env`:
```
WORKER_INTERVAL_SECONDS=10
WORKER_RENEWAL_THRESHOLD_SECONDS=60
```

---
