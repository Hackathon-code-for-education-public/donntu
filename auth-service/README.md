# AUTH SERVICE

Сервис авторизации

## Стек

- PostgreSQL - хранение данных пользователей
- Redis - хранение Refresh токенов пользователей
- Wire - Dependency Injections
- gRPC - транспорт

## Среда окружения

### Пример ENV-файлов
#### ENV-Переменные для PostgreSQL `.env.redis`
```
REDIS_ARGS="--requirepass root"
```

#### ENV-Переменные для PostgreSQL `.env.pg`
```
POSTGRES_USER - имя пользователя Postgre
POSTGRES_PASSWORD - пароль к пользователю
POSTGRES_DB - название бд
```

#### ENV-Переменные для работы приложения - `.env.prod`
```
AUTH_SERVICE_HOST=0.0.0.0
AUTH_SERVICE_PORT=5252

DB_USER=postgres
DB_PASS=postgres

DB_HOST=postgres
DB_PORT=5432
DB_NAME=credentials

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_DB=0
REDIS_PASS=root

JWT_ACCESS_SECRET=access
JWT_REFRESH_SECRET=refresh

JWT_ACCESS_TTL=10
JWT_REFRESH_TTL=21600

LOGGER_LEVEL=debug
```