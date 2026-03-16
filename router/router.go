package router

import (
	"auth/config"
	"auth/handler"
	"auth/repository"
	"auth/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RouteController(dbPool *pgxpool.Pool, cfg *config.Config) *gin.Engine {

	authRepo := repository.NewAuthRepo(dbPool)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/users", authHandler.CreateUser)

	return r

}
