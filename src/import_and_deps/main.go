package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func main() {
	e := echo.New()
	e.GET("/", handler)
	e.Start(":8080")
}
