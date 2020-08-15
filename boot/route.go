package boot

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/handler"
	"go-jahitin/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, toolkit *apipackages.Toolkit) {
	//Handlers
	sessionHandler := handler.NewSessionHandler(toolkit)
	userHandler := handler.NewUserHandler(toolkit)

	//Middlewares
	authMiddleware := middleware.SetupAuthMiddleware(toolkit)

	//Routes
	api := r.Group("")
	{

		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "go-jahitin",
			})
		})

		session := api.Group("/session")
		{
			session.POST("/login", sessionHandler.Login)
			session.POST("/refresh", sessionHandler.Refresh)
			session.Use(authMiddleware)
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
