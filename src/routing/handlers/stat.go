package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Stat is a page handler for stat
func Stat(c echo.Context) error {
	return c.String(http.StatusOK, `{ "total": 0 }`)
}
