package server

import (
	"First-Go-Project/internal/handler"
	"First-Go-Project/internal/repository"
	"First-Go-Project/internal/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	repo := repository.NewRepository()
	services := service.NewService(repo)
	h := handler.New(services)
	h.RegisterRoutes(e)

	return e
}

func Start(e *echo.Echo, port string) error {
	address := fmt.Sprintf(":%s", port)
	return e.Start(address)
}
