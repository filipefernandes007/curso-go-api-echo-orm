package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

type InternalHandler struct {
	echo *echo.Echo
}

func NewInternalHandler(echo *echo.Echo) *InternalHandler {
	return &InternalHandler{echo: echo}
}

func (h *InternalHandler) Healthy() {
	h.echo.GET("/api/healthy", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})
}