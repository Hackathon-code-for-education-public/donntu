# AUTH SERVICE

Сервис авторизации

## Среда окружения

### Example of env files
#### .env.redis
```
REDIS_ARGS="--requirepass root"
```

#### .env.pg
```
POSTGRES_USER - имя пользователя Postgre
POSTGRES_PASSWORD - пароль к пользователю
POSTGRES_DB - название бд
```

#### .env.prod
```
DB_USER=postgres
DB_PASS=postgres

DB_HOST=postgres
DB_PORT=5432
DB_NAME=users

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_DB=0
REDIS_PASS=root

JWT_ACCESS_SECRET=access
JWT_REFRESH_SECRET=refresh

JWT_ACCESS_TTL=10
JWT_REFRESH_TTL=1440

LOGGER_LEVEL=info
```