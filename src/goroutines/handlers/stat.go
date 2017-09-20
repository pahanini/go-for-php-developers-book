package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	ds "goroutines/datastore"
	"net/http"
)

// Stat is a page handler for stat
func Stat(ds *ds.Datastore) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf(`{ "total": %d }`, ds.TotalHits()))
	}
}
