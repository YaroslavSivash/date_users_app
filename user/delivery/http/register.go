package http

import (
	"date_users_app/user"

	"github.com/labstack/echo"
)

func RegisterHTTPEndpoints(e *echo.Echo, uc user.UseCase) {
	h := NewHandler(uc)

	e.GET("/getAllUsers", h.GetAllUsers)
	e.POST("/createUser", h.CreateUser)
	e.POST("/updateUser", h.UpdateUser)

}
