# go

Шаблон golang сервиса

## Структура проекта

cmd/service/main.go - точка входа в приложение

internal/usecase - бизнес логика приложения
internal/repository - методы для работы с БД
internal/service - транспортный слой grpc сервера

internal/migrations - sql файлы миграции БД

api/ - proto и swagger файлы

pkg/api/ - сгенерированные файлы для работы с grpc
pkg/config/ - методы для конфигурации приложения


## Как запустить проект
0. Установка зависимостей
```go
go install \                                       
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
1. Необходмио описать методы в api/service.proto
2. Сгенерировать необходимые для работы
```
./scripts/generate.sh
```

## Конфигурация PostgreSQL через ENV

Сервис поддерживает настройку подключения к базе данных PostgreSQL через переменные окружения.  
Все параметры имеют значения по умолчанию, которые можно переопределить.
Ниже приведены переменные с их типами, значениями по умолчанию и описанием.

| Переменная                              | Тип      | Значение по умолчанию                          | Описание                                                                         |
|-----------------------------------------|----------|------------------------------------------------|----------------------------------------------------------------------------------|
| `APP_NAMESPACE`                         | string   | `ecom`                                         | Пространство имён приложения (namespace)                                         |
| `APP_NAME`                              | string   | `go-template`                                  | Имя приложения                                                                   |
| `APP_ENVIRONMENT`                       | string   | `dev`                                          | Окружение (`dev`, `stage`, `prod`)                                               |
| `APP_IMAGE`                             | string   | `registry-gitlab16.skiftrade.kz/templates1/go` | ---                                                                              |
| `IMAGE_TAG`                             | string   | `dev`                                          | Тег образа (уточнить акутальность)                                               |
| `APP_SWAGGER_FILE`                      | string   | `/api/service.swagger.json`                    | Путь до Swagger файла                                                            |
| `APP_HTTP_PORT`                         | int      | `8090`                                         | Порт HTTP API                                                                    |
| `APP_GRPC_PORT`                         | int      | `8091`                                         | Порт gRPC API                                                                    |
| `APP_LOG_LEVEL`                         | string   | `debug`                                        | Уровень логирования приложения (`trace`, `debug`, `info`, `warn`, `error`)       |
| `DOC_PORT`                              | int      | `9081`                                         | Порт для документации                                                            |
| `JOB_PORT`                              | int      | `8410`                                         | Порт для фоновых job                                                             |
| `GOLANG_PROTOBUF_REGISTRATION_CONFLICT` | string   | `warn`                                         | Поведение при конфликте регистрации protobuf (`warn`, `error`, `ignore`)         |
| `TELEMETRY_LOGGER_DISABLE`              | bool     | `true`                                         | Отключение логгера в telemetry                                                   |
| `TELEMETRY_METRIC_DISABLE`              | bool     | `true`                                         | Отключение метрик в telemetry                                                    |
| `TELEMETRY_TRACING_DISABLE`             | bool     | `true`                                         | Отключение трассировок в telemetry                                               |
| `PG_HOST`                               | string   | `localhost`                                    | Хост базы данных                                                                 |
| `PG_PORT`                               | string   | `5432`                                         | Порт базы данных                                                                 |
| `PG_USER`                               | string   | `postgres`                                     | Имя пользователя                                                                 |
| `PG_PASS` / `PG_PASSWORD`               | string   | `postgres`                                     | Пароль пользователя                                                              |
| `PG_DBNAME` / `PG_NAME`                 | string   | `template`                                     | Имя базы данных                                                                  |
| `PG_SSLMODE`                            | string   | `disable`                                      | Режим SSL (`disable`, `require`, `verify-ca`, `verify-full`)                     |
| `PG_SSLROOTCERT`                        | string   | `` (пусто)                                     | Путь к CA-сертификату                                                            |
| `PG_DEBUG`                              | bool     | `false`                                        | Включить отладочный режим                                                        |
| `PG_DRIVER_LOG_LEVEL`                   | string   | `info`                                         | Уровень логирования драйвера (`trace`, `debug`, `info`, `warn`, `error`, `none`) |
| `PG_SIMPLE_PROTOCOL`                    | bool     | `false`                                        | Использовать простой протокол без prepared statements                            |
| `PG_POOL_STAT_PERIOD`                   | duration | `30s`                                          | Период публикации статистики пула                                                |
| `PG_POOL_MAX_CONNS`                     | int64    | `10`                                           | Максимальное количество соединений                                               |
| `PG_POOL_MIN_CONNS`                     | int64    | `0`                                            | Минимальное количество соединений                                                |
| `PG_POOL_MAX_CONN_LIFETIME`             | duration | `1h`                                           | Максимальное время жизни соединения                                              |
| `PG_POOL_MAX_CONN_IDLE_TIME`            | duration | `30m`                                          | Максимальное время простоя соединения                                            |
| `PG_POOL_HEALTH_CHECK_PERIOD`           | duration | `1m`                                           | Период проверки состояния соединений                                             |

---
