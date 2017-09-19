package main

import (
	"anonymous_functions/datastore"
	"anonymous_functions/handlers"
	"github.com/labstack/echo"
)

func main() {
	ds := datastore.Datastore{}

	e := echo.New()
	e.GET("/stat.js", handlers.Stat(&ds))
	e.GET("/counter.js", handlers.Counter(&ds))
	e.Start(":8080")
}
