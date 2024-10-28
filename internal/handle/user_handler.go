package handle

import (
	"net/http"
	"strconv"

	db "github.com/TiagoAmaralFerreira/go-echo/db/sqlc"
	"github.com/TiagoAmaralFerreira/go-echo/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user db.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	createdUser, err := h.service.CreateUser(c.Request().Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetUser(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
