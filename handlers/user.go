package handlers

import (
	"curso-go/api-echo-orm/entities"
	"curso-go/api-echo-orm/services"
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
)

// Construtor
type UserHandler struct {
	echo *echo.Echo
	db   *gorm.DB
}

func NewUserHandler(echo *echo.Echo, db *gorm.DB) *UserHandler {
	return &UserHandler{echo: echo, db: db}
}

func (h *UserHandler) GetUserByUsernameHandler() {
	h.echo.GET("/api/user/:username", func(c echo.Context) error {
		username := c.Param("username")

		service := services.NewUserService(h.db)
		response, err := service.GetByUsername(username)
		if err != nil {
			if errors.Is(err, entities.ErrUserDoesNotExist) {
				return c.JSON(http.StatusNotFound, entities.OutputError{
					Error: err.Error(),
				})
			}

			return c.JSON(http.StatusInternalServerError, nil)
		}

		return c.JSON(http.StatusOK, response)
	})
}
