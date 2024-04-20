# File Service

Сервис авторизации

## Стек

- MinIO - объектное хранилище
- Wire - Dependency Injections
- gRPC - транспорт

## Запуск

```
docker compose up -d --build
```

При первом запуске необходимо создать Access Key в панели MinIO.

Переходим на http://localhost:9001. И вводим туда значения переменных (файл .env.minio) `MINIO_ROOT_USER` и `MINIO_ROOT_PASSWORD`.
При успешном входе переходим в Access Keys (в меню слева). Создаём пару ключей и записываем в переменные приложения (файл .env.prod)
`MINIO_ACCESS_KEY` для access key и `MINIO_SECRET_KEY` для secret key соответственно

Перезапускаем приложение 
```
docker compose down
docker compose up -d
```

## Среда окружения

### Пример ENV-файлов
#### ENV-Переменные для MinIO `.env.minio`
```
MINIO_ROOT_USER=root
MINIO_ROOT_PASSWORD=rootroot
```

#### ENV-Переменные для работы приложения - `.env.prod`
```
APP_HOST=0.0.0.0
APP_PORT=5253
LOGGER_LEVEL=debug
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=nX9NOGhFtNR9O0IAb5Wa
MINIO_SECRET_KEY=v8G2IBHxJJnhzOv53LbUV9v69rJcxk6yzLAEBztJ
MINIO_DB=0
MINIO_BUCKET=local
MINIO_SECURE=false
```