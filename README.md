# Сервис отправки уведомленй

Сервис отправки уведомлений, который может отправлять письма по SMTP и может быть расширен до отправки SMS и PUSH-уведомлений. Интерфейс подготовлен, но сам функционал опущен из-за отсутствия бесплатных решений.

## Запуск

Для запуска сервера использовать: **docker-compose up -d --build**

Запустить воркер: **go run ./cmd/worker/send_email_worker.go** *или* **go build -o worker ./cmd/worker/send_email_worker.go & ./worker**

## Конфигурация

Создаём **.env** файл из **.env.example** и меняем настройки на актуальные

Всё работает в докере, при необходимости редактируем Dockerfile и docker-compose.yaml

## Детали реализации

В сервисе используются следующие модули:

- github.com/emersion/go-sasl
- github.com/emersion/go-smtp
- github.com/gorilla/mux v1.8.0
- github.com/joho/godotenv v1.4.0
- github.com/rabbitmq/amqp091-go v1.5.0


