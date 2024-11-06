# Todo-Rest API


**Todo-Rest API** — это серверное приложение для управления задачами с использованием Go. Оно предоставляет REST API для создания, редактирования и удаления задач. Приложение поддерживает аутентификацию с использованием JWT-токенов, а также хранит данные задач в SQLite базе данных.

## Описание
Данное приложение представляет собой API для системы управления задачами, которое позволяет пользователям создавать задачи с указанием сроков, описания и повторяющихся событий. Также реализован механизм аутентификации и авторизации на основе токенов.

## Функциональные возможности:
 - Создание, редактирование и удаление задач
 - Поддержка повторяющихся задач
 - Аутентификация с использованием JWT
 - Хранение данных в базе данных SQLite
 - Валидация данных на уровне API
 - Поддержка переменных окружения для конфигурации
## Технологии
 - Go 1.23.2
 - SQLite для хранения данных
 - Chi для роутинга
 - Docker для контейнеризации
 - Buildx и GitHub Actions для CI/CD
 - JWT для аутентификации
 - Окружение для разработки: WSL/Ubuntu

## Установка и запуск

### Локальная установка