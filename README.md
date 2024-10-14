# Ozinshe

Ozinshe — это веб-приложение для управления фильмами, созданное с использованием фреймворка Gin в Go. Это приложение позволяет пользователям находить, управлять и взаимодействовать с фильмами с помощью простого и интуитивно понятного API. Проект поддерживает управление пользователями, включая возможность отмечать фильмы как избранные или просмотренные.

## Функции
### Управление пользователями
* Регистрация пользователей: пользователи могут зарегистрироваться, чтобы создать учетную запись.
* Профиль пользователя: пользователи могут просматривать и обновлять информацию своего профиля.
* Управление избранным: пользователи могут добавлять и удалять фильмы из своего списка избранного.
* Отслеживание просмотренных фильмов: пользователи могут отслеживать просмотренные фильмы.

### Управление фильмами
* Поиск фильмов: пользователи могут искать фильмы по названию или жанру с поддержкой пагинации.
* Сведения о фильме: пользователи могут получать подробную информацию о конкретном фильме.
* Функции администратора: пользователи-администраторы могут добавлять, обновлять, удалять фильмы и  управлять снимками экрана.

## Документация API
Приложение использует Swagger для документации API. Каждая конечная точка хорошо документирована с параметрами, ожидаемыми ответами и обработкой ошибок.

### Полный список того, что было использовано:

* [gin](https://github.com/gin-gonic/gin) - Web framework
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [swag](https://github.com/swaggo/swag) - swag - Swagger
* [Docker](https://www.docker.com/) - Docker


### Getting Started

To run this application locally, follow these steps:

1. Clone the Repository:

```bash
git clone https://github.com/purkhanov/ozinshe

cd ozinshe
```

2. Run it on docker:

```bash
docker-compose up -d
```

### SWAGGER UI:
[http://localhost:8000/swagger/index.html#/](https://localhost:5000/swagger/index.html)
