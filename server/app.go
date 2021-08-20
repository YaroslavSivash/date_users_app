package server

import (
	"date_users_app/services"
	"date_users_app/user"
	"date_users_app/user/delivery/http"
	"date_users_app/user/repository/mongo"
	"date_users_app/user/usecase"
	"github.com/labstack/echo"

	"github.com/spf13/viper"
)

type App struct {
	httpServer *echo.Echo
	usersUC    user.UseCase
}

func NewApp() *App {
	db := services.InitDB()

	userRepo := mongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))
	return &App{
		usersUC: usecase.NewUserUseCase(userRepo),
	}
}

func (a *App) Run(port string) error {

	e := echo.New()
	http.RegisterHTTPEndpoints(e, a.usersUC)
	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
	return nil
}
