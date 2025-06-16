package server

import (
	"First-Go-Project/internal/handler"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	h := handler.New()
	h.RegisterRoutes(e)

	return e
}

func Start(e *echo.Echo, port string) error {
	address := fmt.Sprintf(":%s", port)
	return e.Start(address)
}
