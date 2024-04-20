# File Service

Сервис авторизации

## Стек

- MinIO - объектное хранилище
- Wire - Dependency Injections
- gRPC - транспорт

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