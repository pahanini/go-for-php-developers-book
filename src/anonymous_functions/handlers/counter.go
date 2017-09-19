package handlers

import (
	ds "anonymous_functions/datastore"
	"github.com/labstack/echo"
	"net/http"
)

// Counter handler
func Counter(ds *ds.Datastore) echo.HandlerFunc {
	return func(c echo.Context) error {
		ds.RegisterHit()
		return c.String(http.StatusOK, "ok")
	}
}
