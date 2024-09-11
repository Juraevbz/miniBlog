### ТЗ к платформе Mini_Blog

# Стек технологии
* Go
* Фреймворк Gin-Gonic
* GORM драйвер posgresql
* Конфиг env
* Логгер Slog
* Unit-tests библиотека Mockery
* Документация API Swagger-api v.3.0

# Сушности
1.Пользователь (user)
- id int
- username string
- passwort_hash string


Запус приложения - для запуска введите следуюшие команды:

1. export DSN='host=localhost user=postgres password=postgres dbname=miniblog port=5433 sslmode=disable'