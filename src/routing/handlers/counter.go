package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Counter handler
func Counter(c echo.Context) error {
	return c.String(http.StatusOK, "counter")
}
