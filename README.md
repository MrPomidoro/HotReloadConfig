# Пример использования паттерна "наблюдатель" в Golang

Этот репозиторий содержит пример использования паттерна "наблюдатель" в языке программирования Golang. Пример демонстрирует, как можно создать HTTP-сервер, который перезагружается при изменении конфигурационного файла.

## Установка

Для установки необходимо выполнить команду:

```azure
go get github.com/spf13/viper
```


## Использование

Для запуска HTTP-сервера с поддержкой перезагрузки конфигурационного файла, необходимо выполнить команду:

```azure
go run main.go
```


HTTP-сервер будет запущен на порту 8080. При изменении конфигурационного файла (файл "config.yaml"), сервер будет перезапущен с новыми настройками.

## Конфигурация

Файл конфигурации "config.yaml" содержит следующие настройки:

* `port` - порт, на котором будет запущен HTTP-сервер;
* `read_timeout` - время ожидания чтения запроса от клиента (в миллисекундах);
* `write_timeout` - время ожидания отправки ответа клиенту (в миллисекундах).

Пример конфигурационного файла:

```yaml
port: 8080
read_timeout: 10000
write_timeout: 10000
