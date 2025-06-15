package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	return e
}

func Start(e *echo.Echo, port string) error {
	address := fmt.Sprintf(":%s", port)
	return e.Start(address)
}
