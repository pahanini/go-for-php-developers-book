package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Main page handler
func Main(c echo.Context) error {
	return c.String(http.StatusOK, "main page")
}
