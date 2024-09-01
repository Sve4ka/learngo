package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"learngo/docs"
	"learngo/internal/delivery/handlers"
	"learngo/internal/repository/user"
	userserv "learngo/internal/service/user"
	"learngo/pkg/log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, log *log.Logs) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	//r.SetTrustedProxies([]string{"127.0.0.1"})
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := r.Group("/user")

	userRepo := user.InitUserRepository(db)
	userService := userserv.InitUserService(userRepo, log)
	userHandler := handlers.InitUserHandler(userService)

	userRouter.POST("/create", userHandler.CreateUser)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
