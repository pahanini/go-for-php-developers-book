# Роутинг. Пакеты

**В этой главе мы добавим роутинг к приложению. Создадим собственный пакет и
познакомимся с экспортируемыми и неэкспортируемыми сущностями**

Устроим наше приложение следующим образом:

1. При обращении по адресу `/counter.js `
скрипт будет отдавать пустую страницу с 204 статусом и регистрировать
данное событие.

2. При обращении к по адресу `/stat.js` будет отдаваться JSON вида
`{ "total": N }`, где N - количество запросов к  `/counter.js`

Чтобы добавить оба наших маршрута изменим main.go следующим образом

```go

func stat(c echo.Context) error {
	return c.String(http.StatusOK, `{ "total": 0 }`)
}

func counter(c echo.Context) error {
	return c.String(http.StatusOK, "counter")
}

func main() {
	e := echo.New()
	e.GET("/stat.js", stat)
	e.GET("/counter.js", counter)
	e.Start(":8080")
}
```

В реальном приложении обработчиков может быть очень много, чтобы сохранить читаемость
кода давайте вынесем обработчики GET запросов в отдельный пакет. Создайте директорию
`src/app/handlers` и файлы `stat.go` и `counter.go` в ней и перенесите функции в эти файлы.

Файл counter.js будет выглядеть так:

```go
package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Counter handler
func Counter(c echo.Context) error {
	return c.String(http.StatusOK, "counter")
}

```

Первая стока задает имя пакета. Обратите также внимание то, что имя функции
теперь начинается с заглавной буквы. Аналогично полям стуктуры, сущности пакета
делятся на экспортируемые и не экспортируемые. Имена экспортируемых сущностей
должны начинаться с заглавной буквы, не экспортируемых со строчной.

`main.go` также нуждается в изменениях. Мы импортировали наш новый пакет
`app/handlers` и заменили обработчики запросов на обработчики из пакета.


```go
package main

import (
	"github.com/labstack/echo"
	"app/handlers"
)

func main() {
	e := echo.New()
	e.GET("/stat.js", handlers.Stat)
	e.GET("/counter.js", handlers.Counter)
	e.Start(":8080")
}
```
**Код находится в `src/routing`.**
