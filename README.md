# Webinar #2

# Моки

## Написание моков вручную

См. `internal/core/application/application_mock.go`

## Моки без `go:generate`

Установите gomock локально:

```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

Моки можно сгенерировать, передавая путь к файлу с интерфейсами:

```bash
mockgen -source=internal/core/services/repo.go \
    -destination=internal/core/services/internal/mocks/repo_mock.gen.go \
    -package=mocks
```

А можно сгенерировать, передавая import path пакета и список интерфейсов для мокирования. Для этого сначал необходимо добавить зависимость:

```bash
go get github.com/golang/mock/mockgen/model@v1.6.0
```

А после выполнить эту комманду:

```bash
mockgen -destination=internal/core/services/internal/mocks-reflect/repo_mock.gen.go \
    -package=mocks \
    mocks/internal/core/services Store,Foobar
```

Разница между source-mode и reflect-mode:

1. source-mode работает с неэкспортированными интерфейсами
2. reflect-mode позволяет выбирать интерфейсы для мокирования
3. reflect-mode работает в случае сложной организации репозитория (например, когда модули не используются)

См. обсуждение: https://github.com/golang/mock/issues/406

См. пример использования моков в тестах в `internal/core/services/repo_test.go`

## Моки с `go:generate`

Для генерации моков можно использовать `go:generate` (см. https://go.dev/blog/generate).

См. пример в `internal/core/services/repo_test.go`.

Для генерации кода выполните:

```bash
go generate ./...
```

в корне проекта.

# Makefile

Дока по Makefile лежит здесь: https://www.gnu.org/software/make/manual/make.html

См. Makefile в корне проекта

# Линтеры

## `golangci-lint`

См. https://golangci-lint.run/

Для запуска линтеров выполните:

```bash
make lint
```

Для того, чтобы запустить `golangci-lint` с файлом конфигурации выполните:

```bash
make lint GOLANGCI_LINT_CONFIG_FILE=lint.yml 
```
