package server

import (
	"context"
	"date_users_app/user"
	mongo2 "date_users_app/user/repository/mongo"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
	usersUC	user.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := mongo2.NewUserRepository(db, (viper.GetString("mongo.user_collection")))
	return &App{

	}
}

func (a *App) Run (port string) error{
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return a.httpServer.ListenAndServe()
}

func (a *App) Shutdown (ctx context.Context) error{
	return a.httpServer.Shutdown(ctx)
}
