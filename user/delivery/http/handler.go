package http

import (
	"date_users_app/models"
	"date_users_app/user"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strconv"
)

type Handler struct {
	useCase user.UseCase
}

func NewHandler(useCase user.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetAllUsers(c *echo.Context) error {
	_, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	_, err := strconv.Atoi(c.QueryParam("skip"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	allUsers, err := h.useCase.GetAllUsers(c)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.String(http.StatusOK, allUsers)
}

func (h *Handler) CreateUser(c *echo.Context) error {
	user := &models.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

}

func (h *Handler) UpdateUser(c *echo.Context) error {

}
