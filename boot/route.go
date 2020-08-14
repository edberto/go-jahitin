package boot

import (
	"go-jahitin/api-packages/handler"

	"github.com/gin-gonic/gin"
)

var (
	sessionHandler = handler.NewSessionHandler()
	userHandler    = handler.NewUserHandler()
)

func InitializeRoutes(r *gin.Engine) {
	api := r.Group("")
	{
		session := api.Group("/session")
		{
			session.POST("/login", sessionHandler.Login)
			session.POST("/refresh", sessionHandler.Refresh)
			session.DELETE("/logout", sessionHandler.Logout)
		}

		user := api.Group("/user")
		{
			user.GET("/:id", userHandler.GetOne)
			user.POST("/register", userHandler.Register)
		}
	}
}
