package boot

import (
	"go-jahitin/api-packages/handler"
	"go-jahitin/boot/setup"
	"go-jahitin/helper/config"

	"github.com/gin-gonic/gin"
)

var (
	sessionHandler = handler.NewSessionHandler()
	userHandler    = handler.NewUserHandler()
)

func InitializeRoutes(r *gin.Engine, cfg config.IConfig) {
	authMiddleware := setup.SetupAuthMiddleware(cfg)

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
