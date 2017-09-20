package handlers

import (
	"github.com/labstack/echo"
	ds "goroutines/datastore"
	"net/http"
)

// Counter handler
func Counter(ds *ds.Datastore) echo.HandlerFunc {
	return func(c echo.Context) error {
		ds.RegisterHit()
		return c.String(http.StatusOK, "ok")
	}
}
