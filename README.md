# Webinar #2

# Моки

## Написание моков вручную

См. `internal/core/application/application_mock.go`

## Моки без `go:generate`

Установите gomock локально:

```bash
go install go.uber.org/mock/mockgen@latest
```

Моки можно сгенерировать, передавая путь к файлу с интерфейсами:

```bash
mockgen -source=internal/core/services/repo.go \
    -destination=internal/core/services/internal/mocks/repo_mock.gen.go \
    -package=mocks
```

А можно сгенерировать, передавая import path пакета и список интерфейсов для мокирования:

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

## Паттерн tools

Иногда некоторые пакеты, как например `go.uber.org/mock/mockgen`, не требуются для запуска вашего приложения, а необходимы для выполнения каких-то действия при разработке проекта. Например, в нашем случае, это генерация моков для тестов.

NB: В некоторых пакетных менеджерах для других языков программирования (`npm`) существует возможность добавлять не только зависимости проекта, но и т.н. "devDependencies": это то, что описано выше.

В Go такой возможности пока (11.11.2024) нет, хотя ее добавление предполагается в будущем: https://go.googlesource.com/proposal/+/refs/changes/55/495555/5/design/48429-go-tool-modules.md

Пока для управления такими зависимостями у разработчиков есть несколько опций:

1. добавить в REAMDE.md инструкции по установке необходимых зависимостей локально + `go generate`
2. автоматизировать (хотя бы частично) использование зависимоcтей при помощи `docker`
3. использовать pattern `tools`

См. пример в директории `tools`

## Mockery

Другой популярный пакет для создания моков (см. сравнение их и других пакетов для мокирования здесь: https://gist.github.com/maratori/8772fe158ff705ca543a0620863977c2).

Установите `mockery` при помощи паттерна `tools` (см. `tools/tools.go`)

Конфигурация `mockery` находится в `.mockery.yaml`.

Сгенерируйте код:

```bash
go run github.com/vektra/mockery
```

Почему это не срабатывает сразу? Как это исправить?

Тесты с использованием `mockery` лежат в `internal/core/services/repo_mockery_test.go`.

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
