# Todo App

Backend-сервис для управления задачами, написанный на Go.

Проект реализует REST API для работы с пользователями, задачами и статистикой. В качестве хранилища используется PostgreSQL. Приложение разделено на функциональные модули и слои, а запуск инфраструктуры и самого сервиса автоматизирован с помощью Docker Compose и Makefile.

## Стек

* **Go 1.25.4**
* **PostgreSQL 18**
* **pgx/v5** — работа с PostgreSQL
* **net/http** — HTTP API
* **Swagger** — документация API
* **go-playground/validator** — валидация входных данных
* **Google UUID** — идентификаторы сущностей
* **Uber Zap** — структурированное логирование
* **Docker / Docker Compose**
* **golang-migrate** — миграции базы данных
* **Makefile** — команды для разработки и запуска

## Возможности

### Tasks

* создание задач;
* получение списка задач;
* получение отдельной задачи;
* изменение задач;
* удаление задач;
* изменение статуса задач.

### Users

* работа с пользователями;
* получение пользовательских данных;
* привязка задач к пользователям.

### Statistics

* получение статистики по задачам.

API работает с префиксом `/api/v1`.

Полное описание HTTP API доступно в [`docs/swagger.yaml`](docs/swagger.yaml).

## Архитектура

Проект имеет слоистую архитектуру. Feature-модули изолированы друг от друга и имеют одинаковую структуру:

```text
HTTP transport
      ↓
   Service
      ↓
 Repository
      ↓
 PostgreSQL
```

Общие инфраструктурные компоненты вынесены в `internal/core`.

```text
.
├── cmd/
│   └── todoapp/
│       ├── main.go
│       └── Dockerfile
│
├── internal/
│   ├── core/
│   │   ├── config/
│   │   ├── domain/
│   │   ├── errors/
│   │   ├── logger/
│   │   ├── repository/
│   │   │   └── postgres/
│   │   └── transport/
│   │       └── http/
│   │
│   └── features/
│       ├── tasks/
│       │   ├── repository/
│       │   ├── service/
│       │   └── transport/
│       │
│       ├── users/
│       │   ├── repository/
│       │   ├── service/
│       │   └── transport/
│       │
│       ├── statistics/
│       │   ├── repository/
│       │   ├── service/
│       │   └── transport/
│       │
│       └── web/
│
├── migrations/
├── docs/
├── public/
├── Makefile
├── docker-compose.yaml
├── go.mod
└── .env.example
```

## Технические особенности

* PostgreSQL используется как основное хранилище;
* работа с БД реализована через `pgx/v5` и connection pool;
* схема базы данных управляется через `golang-migrate`;
* входные данные валидируются через `go-playground/validator`;
* для идентификаторов используются UUID;
* конфигурация приложения задаётся через переменные окружения;
* логирование реализовано через Uber Zap;
* REST API версионируется через `/api/v1`;
* API документируется через Swagger;
* приложение поддерживает graceful shutdown.

Graceful shutdown обрабатывает системные сигналы `SIGINT` и `SIGTERM` через `signal.NotifyContext`. При завершении используется контекст приложения и настраиваемый timeout для остановки HTTP-сервера.

## Конфигурация

Конфигурация приложения задаётся через переменные окружения.

Пример:

```env
HTTP_ADDR=:5050
HTTP_SHUTDOWN_TIMEOUT=30s
ALLOWED_ORIGINS=http://localhost:5050/,null

POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
POSTGRES_TIMEOUT=10s

LOGGER_LEVEL=DEBUG
TIME_ZONE=UTC
```

Готовый шаблон переменных находится в [`.env.example`](.env.example).

## Запуск

Для запуска проекта необходимы Docker, Docker Compose и Make.

Создать `.env` на основе `.env.example`:

```bash
cp .env.example .env
```

Запустить PostgreSQL:

```bash
make env-up
```

Применить миграции:

```bash
make migrate-up
```

Запустить приложение в Docker:

```bash
make todoapp-deploy
```

Проверить состояние контейнеров:

```bash
make ps
```

После запуска приложение доступно на:

```text
http://localhost:5050
```

### Локальный запуск

PostgreSQL можно запустить в Docker, а само Go-приложение — локально:

```bash
make env-up
make migrate-up
make todoapp-run
```

## Миграции

Для работы со схемой базы данных используется `golang-migrate`.

Создать новую миграцию:

```bash
make migrate-create seq=название_миграции
```

Применить миграции:

```bash
make migrate-up
```

Откатить миграции:

```bash
make migrate-down
```

## Swagger

Swagger-документация генерируется из аннотаций Go-кода с использованием `swaggo`.

Для генерации:

```bash
make swagger-gen
```

Документация генерируется в директорию `docs`.

## Docker Compose

Docker Compose используется для запуска приложения и необходимой инфраструктуры:

* Go-приложение;
* PostgreSQL;
* отдельный контейнер для выполнения миграций;
* port-forwarder для PostgreSQL;
* контейнер для генерации Swagger-документации.

## Make-команды

| Команда                        | Назначение                              |
| ------------------------------ | --------------------------------------- |
| `make env-up`                  | Запустить PostgreSQL                    |
| `make env-down`                | Остановить PostgreSQL                   |
| `make env-cleanup`             | Удалить данные PostgreSQL               |
| `make env-port-forward`        | Запустить port forwarding PostgreSQL    |
| `make env-port-close`          | Остановить port forwarding              |
| `make migrate-create seq=name` | Создать миграцию                        |
| `make migrate-up`              | Применить миграции                      |
| `make migrate-down`            | Откатить миграции                       |
| `make todoapp-run`             | Запустить приложение локально           |
| `make todoapp-deploy`          | Собрать и запустить приложение в Docker |
| `make todoapp-undeploy`        | Остановить приложение                   |
| `make swagger-gen`             | Сгенерировать Swagger                   |
| `make ps`                      | Показать состояние Docker-контейнеров   |

## Проект

Проект создан как практическая работа для изучения backend-разработки на Go и построения полноценного REST-сервиса с PostgreSQL.

В процессе разработки реализованы работа с PostgreSQL через `pgx`, миграции базы данных, разделение transport/service/repository, feature-модули, контейнеризация, валидация входных данных, структурированное логирование, Swagger-документация и graceful shutdown.

