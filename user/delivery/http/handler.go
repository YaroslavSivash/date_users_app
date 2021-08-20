package http

import (
	"date_users_app/models"
	"date_users_app/user"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
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

func (h *Handler) GetAllUsers(c echo.Context) error {

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	skip, err := strconv.Atoi(c.QueryParam("skip"))
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	allUsers, err := h.useCase.GetAllUsers(c, skip, limit)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, allUsers)
}

func (h *Handler) CreateUser(c echo.Context) error {

	userBody := &models.User{}

	err := json.NewDecoder(c.Request().Body).Decode(&userBody)
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

	userNew := h.useCase.CreateUser(c, userBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, userNew)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	userBody := &models.User{}

	err := json.NewDecoder(c.Request().Body).Decode(&userBody)
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

	userUp := h.useCase.UpdateUser(c, userBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, userUp)
}
