package server

import (

	"date_users_app/services"
	"date_users_app/user"
	"date_users_app/user/repository/mongo"
	"fmt"
	"date_users_app/user/delivery/http"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type App struct {
	httpServer *echo.Echo
	usersUC    user.UseCase
}

func NewApp() *App {
	db := services.InitDB()

	userRepo := mongo.NewUserRepository(db, (viper.GetString("mongo.user_collection")))
	return &App{}
}

func (a *App) Run(port string) error {

	e := echo.New()
	http.RegisterHTTPEndpoints(e, a.usersUC)
	fmt.Println("Starting server at " + (viper.GetString(port))
	e.Logger.Fatal(e.Start(port))
	return nil
}
