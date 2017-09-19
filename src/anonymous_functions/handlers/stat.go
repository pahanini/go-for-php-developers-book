package handlers

import (
	ds "anonymous_functions/datastore"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

// Stat is a page handler for stat
func Stat(ds *ds.Datastore) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf(`{ "total": %d }`, ds.TotalHits()))
	}
}
