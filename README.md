# Tradesoft API Client for Go

Это клиент для работы с API сервиса Tradesoft, реализованный на языке Go.

## Описание

Библиотека предоставляет удобный интерфейс для взаимодействия с API Tradesoft, включая методы для работы с поставщиками, информацией о деталях, аналогами и SMS-уведомлениями.

## Структура пакетов

- `provider` - работа с поставщиками
- `info` - информация о деталях
- `analog` - поиск аналогов
- `messenger` - отправка SMS-сообщений
- `common` - общие типы данных

## Установка

```bash
go get github.com/WorldException/go_tradesoft_api
```

## Использование

```go
import "github.com/WorldException/go_tradesoft_api"

client := tradesoft.NewClient("https://api.tradesoft.ru/3", "username", "password")

// Пример вызова метода получения списка поставщиков
providers, err := client.Provider().GetProviderList()
if err != nil {
    // обработка ошибки
}
```

### Параметры командной строки

Команда `go run cmd/check/main.go` поддерживает следующие параметры:

- `-u` - имя пользователя (обязательный параметр)
- `-p` - пароль (обязательный параметр)
- `-a` - артикул (опциональный параметр)
- `-b` - бренд (опциональный параметр)

Пример использования:

```bash
go run cmd/check/main.go -u username -p password -a 0356150003 -b BOSCH
```

## Документация

Документация по API доступна на сайте: https://service.tradesoft.ru/3/docs

## Лицензия

MIT
