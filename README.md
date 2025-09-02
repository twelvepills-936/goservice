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


