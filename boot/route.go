package boot

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/handler"
	"go-jahitin/boot/setup"
	"go-jahitin/helper/config"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, cfg config.IConfig, toolkit *apipackages.Toolkit) {
	//Handlers
	sessionHandler := handler.NewSessionHandler(toolkit)
	userHandler := handler.NewUserHandler(toolkit)

	//Middlewares
	authMiddleware := setup.SetupAuthMiddleware(cfg, toolkit)

	//Routes
	api := r.Group("")
	{

		session := api.Group("/session")
		{
			session.POST("/login", sessionHandler.Login)
			session.Use(authMiddleware)
			session.POST("/refresh", sessionHandler.Refresh)
			session.DELETE("/logout", sessionHandler.Logout)
		}

		user := api.Group("/user")
		{
			user.POST("/register", userHandler.Register)
			user.Use(authMiddleware)
			user.GET("/:id", userHandler.GetOne)

		}
	}
}
