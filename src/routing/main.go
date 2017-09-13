package main

import (
	"github.com/labstack/echo"
	"routing/handlers"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Main)
	e.GET("/counter.js", handlers.Counter)
	e.Start(":8080")
}
